package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"webrunner_configurator/internal"
	"webrunner_configurator/internal/repository"
	"webrunner_configurator/internal/repository/database"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "gorm.io/driver/postgres"
)

type Builder struct {
	config   internal.DBConfig
	authRole string
}

func NewBuilder(config *internal.Config) *Builder {
	return &Builder{config: config.DBConfig, authRole: config.Security.AuthRole}
}

func (b *Builder) Build() (repository.TaskConfigRepository, error) {
	db, err := gorm.Open("postgres", b.getConnectionString())
	if err != nil {
		return nil, err
	}
	log.Println("Connection to postgresql was successful")
	return database.NewDBTaskConfig(db, b.authRole), nil
}

func (b *Builder) getConnectionString() string {
	config := b.config
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		config.User, config.Password, config.ServerName, config.DB)
}
