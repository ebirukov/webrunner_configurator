package database

import (
	"github.com/jinzhu/gorm"
	"webrunner_configurator/internal/gen/model"
)

const tableName = "restconf"

type DBTaskConfig struct {
	Connector *gorm.DB
	authRole  string
}

func (db *DBTaskConfig) Update(config model.NewConfig, id int64) (*model.TaskConfig, error) {
	tc, err := db.Get(id)
	if tc == nil {
		return nil, err
	}
	taskConfig := model.TaskConfig{
		NewConfig: config,
		Id:        id,
	}
	res := db.Connector.Table(tableName).Where("access like ?", db.authRole).Save(taskConfig)
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
	res := db.Connector.Table(tableName).Where("id = ? and access like ?", id, db.authRole).Find(taskConfig)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return taskConfig, nil
	}
}

func (db *DBTaskConfig) List() ([]model.TaskConfig, error) {
	var configs []model.TaskConfig
	res := db.Connector.Table(tableName).Where("access like ?", db.authRole).Find(&configs)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return configs, nil
	}
}

func (db *DBTaskConfig) Delete(id int64) error {
	var taskConfig model.TaskConfig
	res := db.Connector.Table(tableName).Where("id = ? and access like ?", id, db.authRole).Delete(taskConfig)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

func NewDBTaskConfig(connector *gorm.DB, role string) *DBTaskConfig {
	return &DBTaskConfig{Connector: connector, authRole: role}
}
