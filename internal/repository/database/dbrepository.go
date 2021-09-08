package database

import (
	"github.com/jinzhu/gorm"
	"webrunner_configurator/internal/gen/model"
)

const tableName = "restconf"

type DBTaskConfig struct {
	Connector *gorm.DB
}

func (db *DBTaskConfig) Update(config model.NewConfig, id int64) (*model.TaskConfig, error) {
	taskConfig := model.TaskConfig{
		NewConfig: config,
		Id:        id,
	}
	res := db.Connector.Table(tableName).Save(taskConfig)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return &taskConfig, nil
	}
}

func (db *DBTaskConfig) Create(config model.NewConfig) (int64, error) {
	taskConfig := model.TaskConfig{
		NewConfig: config,
	}
	res := db.Connector.Table(tableName).Create(&taskConfig)
	if res.Error != nil {
		return 0, res.Error
	} else {
		return taskConfig.Id, nil
	}
}

func (db *DBTaskConfig) Get(id int64) (*model.TaskConfig, error) {
	taskConfig := new(model.TaskConfig)
	res := db.Connector.Table(tableName).Where("id = ?", id).Find(taskConfig)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return taskConfig, nil
	}
}

func (db *DBTaskConfig) List() ([]model.TaskConfig, error) {
	var configs []model.TaskConfig
	res := db.Connector.Table(tableName).Find(&configs)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return configs, nil
	}
}

func (db *DBTaskConfig) Delete(id int64) error {
	var taskConfig model.TaskConfig
	res := db.Connector.Table(tableName).Where("id = ?", id).Delete(taskConfig)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

func NewDBTaskConfig(connector *gorm.DB) *DBTaskConfig {
	return &DBTaskConfig{Connector: connector}
}
