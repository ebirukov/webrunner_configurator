package database

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
	repository2 "webrunner_configurator/internal/repository"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository repository2.TaskConfigRepository
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

	s.repository = NewDBTaskConfig(s.DB, "analytic")
}

func (s *Suite) AfterTest(_, _ string) {
	//	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Update() {
	var (
		config = model.NewConfig{}
		id     = int64(1)
		role   = "analytic"
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `restconf` WHERE (id = ? and access like ?)")).
		WithArgs(id, role).
		WillReturnRows(sqlmock.NewRows([]string{"id", "role", "newConfig"}).
			AddRow(id, role, config))
	s.mock.ExpectBegin()
	s.mock.ExpectExec(
		"UPDATE `restconf` SET").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	_, err := s.repository.Update(config, id)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_Create() {
	var (
		config = model.NewConfig{}
	)
	s.mock.ExpectBegin()
	s.mock.ExpectExec(
		"INSERT INTO `restconf`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	_, err := s.repository.Create(config)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_Get() {
	var (
		id   = int64(1)
		role = "analytic"
		name = model.NewConfig{}
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `restconf` WHERE (id = ? and access like ?)")).
		WithArgs(id, role).
		WillReturnRows(sqlmock.NewRows([]string{"id", "role", "newConfig"}).
			AddRow(id, role, name))

	res, err := s.repository.Get(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&model.TaskConfig{
		NewConfig: model.NewConfig{},
		Id:        1,
	}, res))
}

func (s *Suite) Test_repository_Delete() {
	var (
		id = int64(1)
		//name = model.NewConfig{}
	)
	s.mock.ExpectBegin()
	s.mock.ExpectExec(
		"DELETE FROM `restconf` WHERE").WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.Delete(id)

	require.NoError(s.T(), err)
}
