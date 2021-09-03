package main

import (
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"regexp"
	"testing"
	"webrunner_configurator/webservice/api" //inline
)

func TestCRUDHandler(t *testing.T) {

	e := echo.New()

	db, mock, err := sqlmock.New()
	connector, err := gorm.Open("mysql", db)
	connector.LogMode(true)
	require.NoError(t, err)

	dbRepository := api.NewDBTaskConfig(connector)
	handler := api.NewCRUDHandler(dbRepository)

	require.NoError(t, err)
	e.Use(middleware.Logger())
	api.RegisterHandlers(e, handler)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `task_configs`")).WillReturnRows(sqlmock.NewRows([]string{}))
	result := testutil.NewRequest().Get("/configs").Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())

	//:TODO add cases for all methods
}
