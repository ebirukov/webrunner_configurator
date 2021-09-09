package internal

import (
	"errors"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"strconv"
)

var (
	InvalidKeyError = errors.New("unsecure key in the request header. Setup more secure key")
)

type ServerConfig struct {
	Port int
}

type Config struct {
	Server   ServerConfig
	DBConfig DBConfig
	Security SecurityConfig
}

type SecurityConfig struct {
	AuthRole      string
	KeyAuthConfig middleware.KeyAuthConfig
}

type DBConfig struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

func NewCommandlineConfig() *Config {
	var port = flag.Int("port", LookupEnvOrInt("WRC_PORT", 8080), "Port for HTTP server")
	var dbName = flag.String("db.name", LookupEnvOrString("WRC_DB_NAME", "mysql"), "Database name")
	var serverName = flag.String("db.host", LookupEnvOrString("WRC_DB_HOST", "localhost:3306"), "Database host")
	var userName = flag.String("db.user", LookupEnvOrString("WRC_DB_USER", "root"), "Database user")
	var dbPassword = flag.String("db.password", LookupEnvOrString("WRC_DB_PASSWORD", "root"), "Database user")
	var apiKey = flag.String("apiKey", LookupEnvOrString("WRC_API_KEY", "qwerty123"), "Authentication API key")
	var authRole = flag.String("authRole", LookupEnvOrString("WRC_AUTH_ROLE", "analytic"), "Authorization role define api permission")
	flag.Parse()

	dbConfig := DBConfig{*serverName, *userName, *dbPassword, *dbName}
	authValidator := middleware.KeyAuthConfig{Validator: newKeyAuthValidator(*apiKey)}

	return &Config{
		ServerConfig{Port: *port},
		dbConfig,
		SecurityConfig{AuthRole: *authRole, KeyAuthConfig: authValidator},
	}
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}

func newKeyAuthValidator(apiKey string) func(s string, context echo.Context) (bool, error) {
	return func(s string, context echo.Context) (bool, error) {
		if len(s) < 8 {
			return false, InvalidKeyError
		}
		return s == apiKey, nil
	}
}
