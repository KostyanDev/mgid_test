package configuration

import "github.com/caarlos0/env/v6"

// Config for service
type Config struct {
	UserGRPCConfig UserGRPCConfig
}

// UserGRPCConfig for grpc.
type UserGRPCConfig struct {
	Port string `env:"DEVICES_INTERNAL_GRPC_PORT,required"`
	Host string `env:"DEVICES_INTERNAL_GRPC_HOST,required"`
}

// Parse and validate configuration.
func Parse() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
