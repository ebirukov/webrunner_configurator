package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	controller "webrunner_configurator/internal"
	"webrunner_configurator/internal/gen/server"
	//"webrunner_configurator/internal/repository/simple"
	db "webrunner_configurator/internal/repository/database/mysql"
)

func main() {
	config := controller.NewCommandlineConfig()
	repoBuilder := db.NewBuilder(config)
	//repoBuilder := simple.NewBuilder()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.KeyAuthWithConfig(config.Security.KeyAuthConfig))

	dbRepository, err := repoBuilder.Build()
	if err != nil {
		e.Logger.Fatal(err)
		os.Exit(1)
	}

	handler := controller.NewCRUDHandler(dbRepository)
	server.RegisterHandlers(e, handler)

	err = e.Start(fmt.Sprintf("0.0.0.0:%d", config.Server.Port))
	if err != nil {
		e.Logger.Fatal(err)
	}

}
