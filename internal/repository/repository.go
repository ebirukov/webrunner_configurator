package repository

import (
	"github.com/jinzhu/gorm"
	"webrunner_configurator/internal/gen/model"
)

type DBTaskConfig struct {
	Connector *gorm.DB
}

func (db *DBTaskConfig) Create(config model.NewConfig) (int64, error) {
	panic("implement me")
}

func (db *DBTaskConfig) Get(id int64) (*model.TaskConfig, error) {
	taskConfig := new(model.TaskConfig)
	res := db.Connector.Where("id = ?", id).Find(taskConfig)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return taskConfig, nil
	}
}

func (db *DBTaskConfig) List() ([]model.TaskConfig, error) {
	var configs []model.TaskConfig
	res := db.Connector.Find(&configs)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return configs, nil
	}
}

func (db *DBTaskConfig) Update(config model.TaskConfig) {
	panic("implement me")
}

func (db *DBTaskConfig) Delete(id int64) error {
	panic("implement me")
}

func NewDBTaskConfig(connector *gorm.DB) *DBTaskConfig {
	return &DBTaskConfig{Connector: connector}
}
