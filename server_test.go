package main

import (
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	api2 "webrunner_configurator/internal"
	"webrunner_configurator/internal/gen/model"
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

	newConfig := model.NewConfig{
		Category:   new(string),
		Cluster:    new(string),
		Container:  "Container",
		Desc:       new(string),
		ScriptPath: "ScriptPath",
		TaskDef:    "TaskDef",
		UrlPath:    "UrlPath",
	}
	result := testutil.NewRequest().Post("/configs").WithJsonBody(newConfig).Go(t, e)
	// We expect 201 code on successful pet insertion
	assert.Equal(t, http.StatusCreated, result.Code())

	result = testutil.NewRequest().Get("/configs").Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
	var configs []model.TaskConfig
	err = result.UnmarshalBodyToObject(&configs)
	assert.NoError(t, err, "error unmarshaling response")
	assert.Equal(t, 1, len(configs))

	id := configs[0].Id
	result = testutil.NewRequest().Delete(fmt.Sprintf("/configs/%d", id)).Go(t, e)
	assert.Equal(t, http.StatusNoContent, result.Code())

	result = testutil.NewRequest().Delete(fmt.Sprintf("/configs/%d", id)).Go(t, e)
	assert.Equal(t, http.StatusNotFound, result.Code())

	//:TODO add cases for all methods
}
