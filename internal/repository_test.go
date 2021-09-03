package internal

import (
	"database/sql"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"regexp"
	"testing"
	"webrunner_configurator/internal/gen/model"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository TaskConfigRepository
	config     *model.NewConfig
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = NewDBTaskConfig(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Get() {
	var (
		id   = int64(1)
		name = model.NewConfig{}
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `task_configs` WHERE (id = ?)")).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "newConfig"}).
			AddRow(id, name))

	res, err := s.repository.Get(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&model.TaskConfig{
		NewConfig: model.NewConfig{},
		Id:        1,
	}, res))
}
