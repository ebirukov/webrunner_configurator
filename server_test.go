package main

import (
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	api2 "webrunner_configurator/internal"
	"webrunner_configurator/internal/gen/server"
	"webrunner_configurator/internal/repository/simple"
)

func TestCRUDHandler(t *testing.T) {

	e := echo.New()

	/*	db, mock, err := sqlmock.New()
		connector, err := gorm.Open("mysql", db)
		connector.LogMode(true)*/
	repoBuilder := simple.NewBuilder()
	dbRepository, err := repoBuilder.Build()
	require.NoError(t, err)

	handler := api2.NewCRUDHandler(dbRepository)

	e.Use(middleware.Logger())
	server.RegisterHandlers(e, handler)

	result := testutil.NewRequest().Get("/configs").Go(t, e)
	assert.Equal(t, http.StatusInternalServerError, result.Code())

	//:TODO add cases for all methods
}
