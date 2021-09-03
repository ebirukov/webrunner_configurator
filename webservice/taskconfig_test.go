package main

import (
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"webrunner_configurator/webservice/api"
)

func TestCRUDHandler(t *testing.T) {
	var err error
	e := echo.New()
	handler := api.NewCRUDHandler()

	require.NoError(t, err)
	api.RegisterHandlers(e, handler)

	result := testutil.NewRequest().Get("/configs").Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())

	//:TODO add cases for all methods
}
