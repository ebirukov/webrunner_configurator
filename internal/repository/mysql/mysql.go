package mysql

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"webrunner_configurator/internal/repository"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbName     *string
	serverName *string
	userName   *string
	dbPassword *string
)

func init() {
	fs := flag.NewFlagSet("DBConfig", flag.ExitOnError)
	dbName = fs.String("db.name", "mysql", "Database name")
	serverName = fs.String("db.host", "localhost:3306", "Database host")
	userName = fs.String("db.user", "root", "Database user")
	dbPassword = fs.String("db.password", "root", "Database user")
}

func (config *DBConfig) Build() (repository.TaskConfigRepository, error) {
	db, err := connect(config)
	if err != nil {
		return nil, err
	}
	return repository.NewDBTaskConfig(db), nil
}

type DBConfig struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

func NewDBConfig(serverName string, user string, password string, DB string) *DBConfig {
	return &DBConfig{serverName, user, password, DB}
	//return &DBConfig{ServerName: *serverName, User: *userName, Password: *dbPassword, DB: *dbName}
}

func (config *DBConfig) getConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&amp;collation=utf8mb4_unicode_ci&amp;parseTime=true&amp;multiStatements=true",
		config.User, config.Password, config.ServerName, config.DB)
}

//Connect creates MySQL connection
func connect(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.getConnectionString())
	if err != nil {
		return nil, err
	}
	log.Println("Connection to mysql was successful!!")
	return db, nil
}
