//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=conf/types.cfg.yaml ../../task-config-oapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=conf/server.cfg.yaml ../../task-config-oapi.yaml

package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type CRUDHandler struct {
	Lock sync.Mutex
}

func NewCRUDHandler() *CRUDHandler {
	return &CRUDHandler{}
}

func sendError(ctx echo.Context, code int, message string) error {
	opErr := OperationError{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, opErr)
	return err
}

func (h *CRUDHandler) DeleteTaskConfig(ctx echo.Context, id int64) (err error) {
	return sendError(ctx, http.StatusNotFound,
		fmt.Sprintf("Could not delete task config with ID %d", id))
}

func (h *CRUDHandler) FindTaskConfigById(ctx echo.Context, id int64) (err error) {
	return sendError(ctx, http.StatusNotFound,
		fmt.Sprintf("Could not find task config with ID %d", id))
}

func (h *CRUDHandler) FindTaskConfigs(ctx echo.Context, params FindTaskConfigsParams) error {

	return ctx.JSON(http.StatusOK, params.UrlPath)
}

func (h *CRUDHandler) AddTaskConfig(ctx echo.Context) error {
	//panic("implement me")
	err := ctx.String(http.StatusCreated, "created")
	return err
}
