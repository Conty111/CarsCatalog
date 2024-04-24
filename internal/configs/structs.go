package configs

import (
	"time"
)

// APIClientConfiguration stores the configuration of the external API client
type APIClientConfiguration struct {
	Host            string
	Scheme          string
	DefaultHeader   map[string]string
	UserAgent       string
	ServerAddress   string
	ServerPort      string
	TimeoutResponse time.Duration
	TimeToRetry     time.Duration
	RetryEnabled    bool
}

type Configuration struct {
	App        *App
	DB         *DatabaseConfig
	HTTPServer *HTTPServerConfig
	APIClient  *APIClientConfiguration
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
	DSN      string
}

type HTTPServerConfig struct {
	Host string
	Port string
}

type App struct{}
