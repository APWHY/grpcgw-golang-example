package app

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

// Config -- Struct that holds all configuration options for this service
type Config struct {
	Environment                  string `env:"ENVIRONMENT,required"`
	IsDebuggingEnabled           bool   `env:"ENABLE_DEBUGGING"`
	EcommCloudDbConnectionString string `env:"ECOMM_CLOUD_DB_CONNECTION_STRING,required"`
	APIPort                      string `env:"API_PORT" envDefault:":8080"`
	GRPCPort                     string `env:"GRPC_PORT" envDefault:":8081"`
	GRPCHost                     string `env:"GRPC_HOST" envDefault:"localhost"`
}

// LoadConfig -- Checks the env for the ENVIRONMENT variable to figure out which config/<ENVIRONMENT>.toml file to load
func LoadConfig() *Config {

	initialEnv := os.Getenv("ENVIRONMENT")
	if len(strings.TrimSpace(initialEnv)) == 0 {
		log.Fatal("'ENVIRONMENT' variable not set, exiting.")
	}
	cfg := Config{}

	if _, err := toml.DecodeFile("configs/"+initialEnv+".toml", &cfg); err != nil {
		log.Fatalf("Could not load %s config with error: %s", initialEnv, err.Error())
	}
	err := env.Parse(&cfg)

	if err != nil {
		log.Fatalf("Failed to load env variables. %+v\n", err)
	}

	return &cfg
}

// TODO remove this function and it's related test

// Sum -- is only temporary bruh wwtf you want
func Sum(x int, y int) int {
	return x + y
}
