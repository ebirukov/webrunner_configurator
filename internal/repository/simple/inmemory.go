package simple

import (
	"errors"
	"webrunner_configurator/internal/gen/model"
	"webrunner_configurator/internal/repository"
)

var (
	TaskConfigNotFound = errors.New("task config not found")
)

type taskConfigStore map[int64]model.TaskConfig

func (s taskConfigStore) values() []model.TaskConfig {
	values := make([]model.TaskConfig, len(s))
	i := 0
	for _, value := range s {
		values[i] = value
		i++
	}
	return values
}

func NewTaskConfigStore() repository.TaskConfigRepository {
	var store = make(taskConfigStore, 5)
	return repository.TaskConfigRepository(store)
}

func (s taskConfigStore) Create(config model.NewConfig) (int64, error) {
	id := int64(len(s))
	v := model.TaskConfig{
		NewConfig: config,
		Id:        id,
	}
	s[id] = v
	return id, nil
}

func (s taskConfigStore) Get(id int64) (*model.TaskConfig, error) {
	if taskConfig, ok := s[id]; ok {
		return &taskConfig, nil
	} else {
		return nil, TaskConfigNotFound
	}

}

func (s taskConfigStore) List() ([]model.TaskConfig, error) {
	return s.values(), nil
}

func (s taskConfigStore) Update(config model.TaskConfig) {
	panic("implement me")
}

func (s taskConfigStore) Delete(id int64) error {
	if _, ok := s[id]; ok {
		delete(s, id)
		return nil
	} else {
		return TaskConfigNotFound
	}
}
