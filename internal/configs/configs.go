package configs

import (
	"fmt"
	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func GetConfig() *Configuration {
	return getFromEnv()
}

func getFromEnv() *Configuration {
	cfg := &Configuration{}

	// TODO: remove app config if not needed
	cfg.App = getAppConfig()
	cfg.DB = getDBConfig()
	cfg.HTTPServer = getHTTPServerConfig()
	cfg.APIClient = getAPIClientConfig()

	return cfg
}

func getDBConfig() *DatabaseConfig {
	dbCfg := &DatabaseConfig{}

	dbCfg.Host = envy.Get("DB_HOST", "localhost")
	dbCfg.User = envy.Get("DB_USER", "postgres")
	dbCfg.Password = envy.Get("DB_PASSWORD", "postgres")
	dbCfg.DBName = envy.Get("DB_NAME", "cars")
	dbCfg.SSLMode = envy.Get("DB_SSLMODE", "disable")
	port, err := strconv.Atoi(envy.Get("DB_PORT", "5432"))
	if err != nil {
		log.Panic().Err(err).Msg("cannot convert DB_PORT")
	}
	dbCfg.Port = port

	dbCfg.DSN = getDbDSN(dbCfg)

	return dbCfg
}

func getDbDSN(dbConfig *DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password,
		dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode,
	)
}

func getAPIClientConfig() *APIClientConfiguration {
	apiClientCfg := &APIClientConfiguration{}

	apiClientCfg.Host = envy.Get("API_HOST", "localhost")
	apiClientCfg.Scheme = envy.Get("API_SCHEME", "http")
	apiClientCfg.DefaultHeader = getDefaultHeaders()
	apiClientCfg.ServerAddress = envy.Get("API_SERVER_ADDRESS", "localhost")
	apiClientCfg.ServerPort = envy.Get("API_SERVER_PORT", "8081")
	duration, err := time.ParseDuration(envy.Get("API_TIMEOUT_RESPONSE", "10s"))
	if err != nil {
		log.Panic().Err(err).Msg("cannot parse API_TIMEOUT_RESPONSE")
	}
	apiClientCfg.TimeoutResponse = duration
	//duration, err = time.ParseDuration(envy.Get("API_TIME_TO_RETRY", "5s"))
	//if err != nil {
	//	log.Panic().Err(err).Msg("cannot parse API_TIME_TO_RETRY")
	//}
	//apiClientCfg.TimeToRetry = duration
	//apiClientCfg.RetryEnabled = envy.Get("API_RETRY_ENABLED", "false") == "true"

	return apiClientCfg
}

func getAppConfig() *App {
	appCfg := &App{}

	return appCfg
}

func getHTTPServerConfig() *HTTPServerConfig {
	httpServerCfg := &HTTPServerConfig{}

	httpServerCfg.Host = envy.Get("HTTP_SERVER_HOST", "localhost")
	httpServerCfg.Port = envy.Get("HTTP_SERVER_PORT", "8080")

	return httpServerCfg
}

func getDefaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   envy.Get("API_USER_AGENT", "cars-api-client"),
	}
}
