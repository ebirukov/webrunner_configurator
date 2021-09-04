package internal

import (
	"flag"
)

type ServerConfig struct {
	Port int
}

type Config struct {
	Server   ServerConfig
	DBConfig DBConfig
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
	flag.Parse()
	return &Config{
		ServerConfig{Port: *port},
		DBConfig{*serverName, *userName, *dbPassword, *dbName},
	}
}
