package repository

import "webrunner_configurator/internal/gen/model"

type Builder interface {
	Build() (TaskConfigRepository, error)
}

type TaskConfigRepository interface {
	Create(config model.NewConfig) (int64, error)
	Get(id int64) (*model.TaskConfig, error)
	List() ([]model.TaskConfig, error)
	Update(config model.NewConfig, id int64) (*model.TaskConfig, error)
	Delete(id int64) error
}
