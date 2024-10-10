package configs

import (
	"time"

	"github.com/caarlos0/env"
)

var (
	Env          Environment
	LocalAddress string
)

type Environment struct {
	PodNamespace        string        `env:"MY_POD_NAMESPACE" envDefault:"192.168.115.148"`
	PodName             string        `env:"MY_POD_NAME" envDefault:"192.168.115.148"`
	PodIP               string        `env:"MY_POD_IP" envDefault:"192.168.115.148"`
	GrpcPort            string        `env:"GRPC_PORT" envDefault:"8090"`
	GatewayPort         string        `env:"GRPC_GATEWAY_PORT" envDefault:"80"`
	PlayMentTimeout     time.Duration `env:"PLAY_MENT_TIMEOUT" envDefault:"60s"`
	RingWaitTimeout     time.Duration `env:"RING_WAIN_TIMEOUT" envDefault:"180s"`
	TransactionTimeout  time.Duration `env:"TRANSACTION_TIMEOUT" envDefault:"180s"`
	JaegerControllerDns string        `env:"BCLOUD_TRACES_URI" envDefault:"http://jaeger-collector.istio-system:14268/api/traces"`
	KafkaBroker         []string      `env:"BCLOUD_KAFKA_BROKER_URI" envSeparator:"," envDefault:"100.100.103.152:9092"`
	QueueSvcDns         string        `env:"BCLOUD_QFLOW_SVC_URI" envDefault:"100.100.103.163:50052"`
	CallFlowSvcDns      string        `env:"BCLOUD_CALLFLOW_SVC_URI" envDefault:"100.100.103.163:50052"`
	SCEFlowSvcDns       string        `env:"BCLOUD_SCE_FLOW_SVC_URI" envDefault:"sceflow-svc:8080"`
	ArchDataSvcDns      string        `env:"BCLOUD_ARCHDATA_SVC_URI" envDefault:"100.100.103.163:6001"`
	PresenceSvcDns      string        `env:"BCLOUD_PRESENCE_SVC_URI" envDefault:"100.100.103.163:6004"`
	RealtimeSvcDns      string        `env:"BCLOUD_REALTIMEDATA_SVC_URI" envDefault:"100.100.103.163:6003"`
	RecordGatewaySvcDns string        `env:"BCLOUD_RECORD_GATEWAY_SVC_URI" envDefault:"recgw-svc.bcloud-core.svc.cluster.local:80"`
	TrunkSvcDns         string        `env:"BCLOUD_TRUNK_SVC_URI" envDefault:"trunksip.bcloud-core.svc.cluster.local:80"`
	BotFlowSvcDns       string        `env:"BCLOUD_BOT_FLOW_SVC_URI" envDefault:"botflow-svc.bcloud-bot.svc.cluster.local:8080"`
	CryptoApiDns        string        `env:"BCLOUD_CRYPTO_API_SVC_URI" envDefault:"http://api-crypto-svc.bcloud-backend.svc.cluster.local:80"`
}

func init() {
	if err := env.Parse(&Env); err != nil {
	}
	LocalAddress = Env.PodIP + ":" + Env.GrpcPort
}

func Print() {
}
