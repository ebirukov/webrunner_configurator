package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"webrunner_configurator/webservice/api"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	e := echo.New()
	e.Use(middleware.Logger())

	handler := api.NewCRUDHandler()
	api.RegisterHandlers(e, handler)

	err := e.Start(fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		e.Logger.Fatal(err)
	}

}
