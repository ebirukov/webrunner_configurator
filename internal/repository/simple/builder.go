package simple

import (
	"webrunner_configurator/internal/repository"
)

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b Builder) Build() (repository.TaskConfigRepository, error) {
	return NewTaskConfigStore(), nil
}
