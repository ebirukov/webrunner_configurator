package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"webrunner_configurator/internal"
	"webrunner_configurator/internal/repository"
	"webrunner_configurator/internal/repository/database"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Builder struct {
	config internal.DBConfig
}

func NewBuilder(config internal.DBConfig) *Builder {
	return &Builder{config: config}
}

func (b *Builder) Build() (repository.TaskConfigRepository, error) {
	db, err := gorm.Open("mysql", b.getConnectionString())
	if err != nil {
		return nil, err
	}
	log.Println("Connection to mysql was successful")
	return database.NewDBTaskConfig(db), nil
}

func (b *Builder) getConnectionString() string {
	config := b.config
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&amp;collation=utf8mb4_unicode_ci&amp;parseTime=true&amp;multiStatements=true",
		config.User, config.Password, config.ServerName, config.DB)
}
