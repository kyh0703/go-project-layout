package main

import "context"

type app struct {
	perfSvr *PerfServer
}

func NewApp(perfSvr *PerfServer) *app {
	return &app{
		perfSvr: perfSvr,
	}
}

func (a *app) Run(ctx context.Context) error {
	// Run the application
	return nil
}

func (a *app) Stop(ctx context.Context) error {
	// Stop the application
	a.perfSvr.Shutdown(ctx)
	return nil
}
