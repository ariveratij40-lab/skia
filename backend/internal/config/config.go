package config

import (
	"os"
	"time"
)

// Config contiene toda la configuración de la aplicación
type Config struct {
	DatabaseURL         string
	JWTSecret           string
	JWTExpiration       time.Duration
	JWTRefreshExpiration time.Duration
	AppEnv              string
	AppPort             string
	LogLevel            string
}

// Load carga la configuración desde variables de entorno
func Load() *Config {
	return &Config{
		DatabaseURL:          getEnv("DATABASE_URL", ""),
		JWTSecret:            getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiration:        parseDuration(getEnv("JWT_EXPIRATION", "15m")),
		JWTRefreshExpiration: parseDuration(getEnv("JWT_REFRESH_EXPIRATION", "7d")),
		AppEnv:               getEnv("APP_ENV", "development"),
		AppPort:              getEnv("APP_PORT", "8080"),
		LogLevel:             getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return 15 * time.Minute
	}
	return d
}
