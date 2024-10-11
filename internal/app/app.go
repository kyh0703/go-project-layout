package app

import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kyh0703/go-project-layout/configs"
	"github.com/kyh0703/go-project-layoutinternal/adaptor/rpc"
	"github.com/kyh0703/go-project-layoutinternal/app/common"
	"github.com/kyh0703/go-project-layoutpkg/cache"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/ipron-ne/iCore/ilog"
	"gitlab.com/ipron-ne/iCore/otrace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	callpb "gitlab.com/ipron-ne/grpc-idl/gen/go/call"
	probe "gitlab.com/ipron-ne/iCore/k8s-probe"
	platform_server "gitlab.com/ipron-ne/iproncloud-platform-server/v2"
	lgrpc "module github.com/kyh0703/go-project-layoutpkg/grpc"
)

type App struct {
	rpc          rpc.Rpc
	cacheManager cache.HookHandler
	commonApi    *platform_server.PlatformAPI
	listener     net.Listener
	grpcServer   *grpc.Server
	httpServer   *http.Server
	perfServer   *http.Server
	tracer       *otrace.Provider
	wg           sync.WaitGroup
}

func ProvideApp(
	ctx context.Context,
	rpc rpc.Rpc,
	commonApi *platform_server.PlatformAPI,
	cacheManager cache.HookHandler,
	tracer *otrace.Provider,
	apiCtrl common.ApiController,
	devCtrl common.DeviceController,
) (*App, error) {
	app := &App{
		rpc:          rpc,
		cacheManager: cacheManager,
		commonApi:    commonApi,
		tracer:       tracer,
	}
	app.initializeLogger()
	configs.Print()
	app.cacheManager.Initialize()

	if err := app.rpc.Connect(ctx); err != nil {
		return nil, err
	}

	if err := app.cacheManager.Start(app.commonApi); err != nil {
		return nil, err
	}

	app.tracer = tracer

	probe.HookGracefulShutdown(app.onShutdown)

	if err := app.listenServer(ctx, apiCtrl, devCtrl, intercept); err != nil {
		return nil, err
	}

	probe.ServiceReadyProbe()

	return app, nil
}

func (app *App) Close() {
	time.Sleep(configs.Env.TransactionTimeout)
	app.tracer.Close()
	app.grpcServer.GracefulStop()
	app.httpServer.Close()
	app.perfServer.Close()
	app.rpc.Close()
	app.commonApi.Close()
	app.wg.Wait()
}

func (app *App) initializeLogger() {
	ilog.InitLog("/app/log", "call-%Y%M%D-%S.log")
	opt := ilog.DefOption()
	opt.AnsiColor = true
	ilog.SetOption(&opt)
}

func (app *App) ListenPerfServer() {
	app.perfServer = &http.Server{
		Addr:         ":6060",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.wg.Add(1)
	go func() {
		if err := app.perfServer.ListenAndServe(); err != nil {
			ilog.Fatal(err)
		}
		ilog.Info("start listen perf server")
	}()
}

func (app *App) ListenGatewayServer(mux *http.ServeMux) {
	app.httpServer = &http.Server{
		Addr:         ":" + configs.Env.GatewayPort,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		if err := app.httpServer.ListenAndServe(); err != nil {
			ilog.Fatal("listen gateway fail: %w", err)
		}
		ilog.Info("start listen http server")
	}()
}

func (app *App) listenServer(
	ctx context.Context,
	apiCtrl common.ApiController,
	devCtrl common.DeviceController,
) error {
	listener, err := net.Listen("tcp", ":"+configs.Env.GrpcPort)
	if err != nil {
		return err
	}
	app.listener = listener

	// create a grpc server object
	app.grpcServer = lgrpc.Server()

	// attach the service to the server
	callpb.RegisterApiServer(app.grpcServer, apiCtrl)
	callpb.RegisterDeviceServer(app.grpcServer, devCtrl)

	// start server
	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		if err := app.grpcServer.Serve(app.listener); err != nil {
			ilog.Fatal("listen server fail: %w", err)
		}

		ilog.Info("start listen grpc server")
	}()

	// dial context
	grpcPort := "0.0.0.0:" + configs.Env.GrpcPort
	conn, err := grpc.DialContext(
		ctx,
		grpcPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	// register greeter
	gwmux := runtime.NewServeMux(
		runtime.WithErrorHandler(intercept.CustomHttpError),
	)
	if err = callpb.RegisterApiHandler(ctx, gwmux, conn); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.Handle("/metrics", promhttp.Handler())

	app.ListenPerfServer()
	app.ListenGatewayServer(mux)
	return nil
}

func (app *App) onShutdown() {
	app.Close()
}
