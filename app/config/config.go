package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Load() *Config {
	cfg := &Config{}

	// Service configuration
	flag.StringVar(&cfg.Port, "PORT", getEnv("PORT", "8080"), "HTTP Server address")
	flag.StringVar(&cfg.DBHost, "DB_HOST", getEnv("DB_HOST", "localhost"), "Postgres host")
	flag.StringVar(&cfg.DBPort, "DB_PORT", getEnv("DB_PORT", "5432"), "Postgres port")
	flag.StringVar(&cfg.DBUser, "DB_USER", getEnv("DB_USER", "postgres"), "Postgres user")
	flag.StringVar(&cfg.DBPass, "DB_PASSWORD", getEnv("DB_PASSWORD", "postgres"), "Postgres password")
	flag.StringVar(&cfg.DBName, "DB_NAME", getEnv("DB_NAME", "football_db"), "Postgres database name")

	flag.Parse()

	return cfg
}

func (c *Config) BuildDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
