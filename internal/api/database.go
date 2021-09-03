package api

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type DBConfig struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

func NewDBConfig(serverName string, user string, password string, DB string) *DBConfig {
	return &DBConfig{ServerName: serverName, User: user, Password: password, DB: DB}
}

func (config *DBConfig) getConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&amp;collation=utf8mb4_unicode_ci&amp;parseTime=true&amp;multiStatements=true",
		config.User, config.Password, config.ServerName, config.DB)
}

//Connect creates MySQL connection
func Connect(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.getConnectionString())
	if err != nil {
		return nil, err
	}
	log.Println("Connection to mysql was successful!!")
	return db, nil
}
