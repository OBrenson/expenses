package unit

import (
	"expenses/internal/dbaccess"
	"expenses/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var resUser = domain.User{
	Id: 1,
	Username: "test",
}

type MockedDb struct {
	mock.Mock
}

func (m *MockedDb) Query(dbFunc func (args interface{}) *gorm.DB, arg interface{}) *gorm.DB {
	m.Called(arg)
	if arg == nil {
		panic("arg is nil")
	}
	u := arg.(*domain.User)
	if u.Id != resUser.Id {
		u.Id = resUser.Id
		return &gorm.DB{
			RowsAffected: 1,
			Statement: &gorm.Statement{
				Dest: resUser,
			},
		}
	}
	return &gorm.DB{
		RowsAffected: 0,
	}
}

func TestInsert(t *testing.T) {
	u := &domain.User{Id: 0, Username: "test"}

	dao,testDb := createMockDao()
	testDb.On("Query", u).Return(
		&gorm.DB{
			RowsAffected: 1,
		})
	err := dao.Insert(u)
	assert.Nil(t, err)
	testDb.AssertExpectations(t)
	assert.Equal(t, resUser.Id, u.Id, "User id should be changed after insert")
	
	err = dao.Insert(u)
	assert.IsType(t, &dbaccess.NoRowsAffected{}, err)

	err = dao.Insert(nil)
	assert.IsType(t, &dbaccess.DbPanicErr{}, err)
}

func createMockDao() (dbaccess.DaoApi[domain.User], *MockedDb) {
	testDb := &MockedDb{}
	return dbaccess.CreateDao[domain.User](testDb, &gorm.DB{}) ,testDb
}