package unit

import (
	"database/sql"
	"expenses/internal/dbaccess"
	"expenses/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	userApi dbaccess.Api[domain.User]
	user *domain.User
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	postgres.
	s.DB, err = gorm.Open()(postgres.New(postgres.Config{}))
	require.NoError(s.T(), err)

	s.api = CreateRepository(s.DB)
}

func TestInsert(t *testing.T) {
	dao := mockDao()
	user := &domain.User{Id: 0, Username: "test"}
	u,err := dao.Insert(user)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, u.Username, user.Username, "Usernames must be equal")
}
