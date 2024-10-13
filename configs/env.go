package configs

import (
	"log"
	"time"

	"github.com/caarlos0/env"
)

var (
	Env          Environment
	LocalAddress string
)

type Environment struct {
	AppPort            string        `env:"APP_PORT" envDefault:"8000"`
	ServerReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"5s"`
	TransactionTimeout time.Duration `env:"TRANSACTION_TIMEOUT" envDefault:"180s"`
	// Secret
	AccessSecretKey  string `env:"ACCESS_SECRET_KEY" envDefault:"secret"`
	RefreshSecretKey string `env:"REFRESH_SECRET_KEY" envDefault:"refresh"`
	CookieSecretKey  string `env:"COOKIE_SECRET_KEY"`
	// Database
	DBType string `env:"DB_TYPE" envDefault:"mysql"`
	DBName string `env:"DB_NAME" envDefault:"go_project_layout"`
	DBUrl  string `env:"DB_URL" envDefault:""`
	// Cache
	CacheUrl string `env:"CACHE_URL" envDefault:"localhost:6379"`
}

func init() {
	if err := env.Parse(&Env); err != nil {
		log.Fatal(err)
	}
}

func Print() {
	log.Println("APP_HOST:", Env.AppPort)
	log.Println("TRANSACTION_TIMEOUT:", Env.TransactionTimeout)
}
