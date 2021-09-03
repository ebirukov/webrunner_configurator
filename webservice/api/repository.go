package api

import (
	"github.com/jinzhu/gorm"
)

type TaskConfigRepository interface {
	Create(config NewConfig) (int64, error)
	Get(id int64) (*TaskConfig, error)
	List() ([]TaskConfig, error)
	Update(config TaskConfig)
	Delete(id int64) error
}

type DBTaskConfig struct {
	Connector *gorm.DB
}

func (db *DBTaskConfig) Create(config NewConfig) (int64, error) {
	panic("implement me")
}

func (db *DBTaskConfig) Get(id int64) (*TaskConfig, error) {
	taskConfig := new(TaskConfig)
	res := db.Connector.Where("id = ?", id).Find(taskConfig)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return taskConfig, nil
	}
}

func (db *DBTaskConfig) List() ([]TaskConfig, error) {
	var configs []TaskConfig
	res := db.Connector.Find(&configs)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return configs, nil
	}
}

func (db *DBTaskConfig) Update(config TaskConfig) {
	panic("implement me")
}

func (db *DBTaskConfig) Delete(id int64) error {
	panic("implement me")
}

func NewDBTaskConfig(connector *gorm.DB) *DBTaskConfig {
	return &DBTaskConfig{Connector: connector}
}
