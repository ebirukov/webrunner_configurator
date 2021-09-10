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
	config   internal.DBConfig
	authRole string
}

func NewBuilder(config *internal.Config) *Builder {
	return &Builder{config: config.DBConfig, authRole: config.Security.AuthRole}
}

func (b *Builder) Build() (repository.TaskConfigRepository, error) {
	db, err := gorm.Open("mysql", b.getConnectionString())
	if err != nil {
		return nil, err
	}
	log.Println("Connection to mysql was successful")
	return database.NewDBTaskConfig(db, b.authRole), nil
}

func (b *Builder) getConnectionString() string {
	config := b.config
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.User, config.Password, config.ServerName, config.DB)
}
