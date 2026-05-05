package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string
	Port string
	JWTSecret string
}

func Load() *Config {
	_ = godotenv.Load() // Ignore error

	cfg := &Config {
		AppEnv: getEnv("APP_ENV", "development"),
		Port: getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "")
	}

	validate(cfg)

	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func validate(cfg *Config) {
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}
}
