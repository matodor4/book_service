package main

import "time"

type Config struct {
	Address                 string        `env:"TRANSPORT_ADDRESS"`
	Port                    string        `env:"TRANSPORT_PORT"`
	LogLevel                string        `env:"LOGGER_LEVEL"`
	GracefulShutdownTimeout time.Duration `env:"GRACEFUL_SHUTDOWN"`
}

func LoadConfig() (*Config, error) {
	cfg := Default()

	return &cfg, nil
}

func Default() Config {
	return Config{
		Address:                 "127.0.0.1",
		Port:                    "8080",
		LogLevel:                "debug",
		GracefulShutdownTimeout: time.Minute * 5,
	}
}
