//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=conf/types.cfg.yaml ../task-config-oapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=conf/server.cfg.yaml ../task-config-oapi.yaml

package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"webrunner_configurator/internal/gen/model"
	"webrunner_configurator/internal/repository"
	"webrunner_configurator/internal/repository/simple"
)

type CRUDHandler struct {
	repository repository.TaskConfigRepository
}

func NewCRUDHandler(repository repository.TaskConfigRepository) *CRUDHandler {
	return &CRUDHandler{repository: repository}
}

func sendError(ctx echo.Context, code int, message string) error {
	opErr := model.OperationError{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, opErr)
	return err
}

func (h *CRUDHandler) UpdateTaskConfig(ctx echo.Context, id int64) (err error) {
	var newConfig model.NewConfig
	err = ctx.Bind(&newConfig)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format for NewConfig")
	}
	_, err = h.repository.Update(newConfig, id)
	if err == simple.TaskConfigNotFound {
		return sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find task config with ID %d", id))
	}
	if err != nil {
		return sendError(ctx, http.StatusUnprocessableEntity, "Can't update task config")
	}
	return ctx.JSON(http.StatusOK, model.TaskConfig{NewConfig: newConfig, Id: id})
}

func (h *CRUDHandler) DeleteTaskConfig(ctx echo.Context, id int64) (err error) {
	err = h.repository.Delete(id)
	if err == simple.TaskConfigNotFound {
		return sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find task config with ID %d", id))
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h *CRUDHandler) FindTaskConfigById(ctx echo.Context, id int64) error {
	taskConfig, err := h.repository.Get(id)
	if err == simple.TaskConfigNotFound {
		return sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("Could not find task config with ID %d", id))
	}
	return ctx.JSON(http.StatusOK, taskConfig)
}

func (h *CRUDHandler) FindTaskConfigs(ctx echo.Context) error {
	tasks, err := h.repository.List()
	if err != nil {
		return sendError(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, tasks)
}

func (h *CRUDHandler) AddTaskConfig(ctx echo.Context) (err error) {
	var (
		newConfig model.NewConfig
		id        int64
	)
	err = ctx.Bind(&newConfig)
	if err != nil {
		return sendError(ctx, http.StatusBadRequest, "Invalid format for NewConfig")
	}
	id, err = h.repository.Create(newConfig)
	if err != nil {
		return sendError(ctx, http.StatusInternalServerError, err.Error())
	}
	err = ctx.JSON(http.StatusCreated, model.TaskConfig{NewConfig: newConfig, Id: id})
	return nil
}
