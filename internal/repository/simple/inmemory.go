package simple

import (
	"errors"
	"webrunner_configurator/internal/gen/model"
)

type TaskConfigStore struct {
}

func NewTaskConfigStore() *TaskConfigStore {
	return &TaskConfigStore{}
}

func (t TaskConfigStore) Create(config model.NewConfig) (int64, error) {
	return 0, errors.New("implement me")
}

func (t TaskConfigStore) Get(id int64) (*model.TaskConfig, error) {
	return nil, errors.New("implement me")
}

func (t TaskConfigStore) List() ([]model.TaskConfig, error) {
	return nil, errors.New("implement me")
}

func (t TaskConfigStore) Update(config model.TaskConfig) {
	panic("implement me")
}

func (t TaskConfigStore) Delete(id int64) error {
	return errors.New("implement me")
}
