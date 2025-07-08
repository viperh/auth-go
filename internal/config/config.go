package config

import (
	"fmt"
	"os"
)

type Config struct {
	DbHost string `env:"DB_HOST"`
	DbPort string `env:"DB_PORT"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbName string `env:"DB_NAME"`
	DbSSL  string `env:"DB_SSL"`

	Port   string `env:"PORT"`
	Secret string `env:"SECRET"`
}

func New() *Config {
	return &Config{
		DbHost: mustGet("DB_HOST"),
		DbPort: mustGet("DB_PORT"),
		DbUser: mustGet("DB_USER"),
		DbPass: mustGet("DB_PASS"),
		DbName: mustGet("DB_NAME"),
		DbSSL:  mustGet("DB_SSL"),
		Port:   mustGet("PORT"),
		Secret: mustGet("SECRET"),
	}
}

func mustGet(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("Environment variable %s is not set", key))
}
