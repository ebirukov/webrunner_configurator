package main

import (
	"flag"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	api "webrunner_configurator/internal"
	"webrunner_configurator/internal/gen/server"
)

func main() {
	var port = flag.Int("port", 8080, "Port for HTTP server")
	var dbName = flag.String("db.name", "mysql", "Database name")
	var serverName = flag.String("db.host", "localhost:3306", "Database host")
	var userName = flag.String("db.user", "root", "Database user")
	var dbPassword = flag.String("db.password", "root", "Database user")
	flag.Parse()

	e := echo.New()
	e.Use(middleware.Logger())

	dbConfig := api.NewDBConfig(*serverName, *userName, *dbPassword, *dbName)
	db, err := api.Connect(dbConfig)
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	dbRepository := api.NewDBTaskConfig(db)
	handler := api.NewCRUDHandler(dbRepository)
	server.RegisterHandlers(e, handler)

	err = e.Start(fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		e.Logger.Fatal(err)
	}

}
