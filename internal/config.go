package internal

import (
	"errors"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	var port = flag.Int("port", 8080, "Port for HTTP server")
	var dbName = flag.String("db.name", "mysql", "Database name")
	var serverName = flag.String("db.host", "localhost:3306", "Database host")
	var userName = flag.String("db.user", "root", "Database user")
	var dbPassword = flag.String("db.password", "root", "Database user")
	var apiKey = flag.String("apiKey", "qwerty123", "Authentication API key")
	var authRole = flag.String("authRole", "analytic", "Authorization role define api permission")
	flag.Parse()

	dbConfig := DBConfig{*serverName, *userName, *dbPassword, *dbName}
	authValidator := middleware.KeyAuthConfig{Validator: newKeyAuthValidator(*apiKey)}

	return &Config{
		ServerConfig{Port: *port},
		dbConfig,
		SecurityConfig{AuthRole: *authRole, KeyAuthConfig: authValidator},
	}
}

func newKeyAuthValidator(apiKey string) func(s string, context echo.Context) (bool, error) {
	return func(s string, context echo.Context) (bool, error) {
		if len(s) < 8 {
			return false, InvalidKeyError
		}
		return s == apiKey, nil
	}
}
