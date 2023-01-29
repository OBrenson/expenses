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

func (m *MockedDb) Query(dbFunc func (args interface{}) *gorm.DB, args interface{}) *gorm.DB {
	m.Called(args)
	return &gorm.DB{
		RowsAffected: 1,
		Statement: &gorm.Statement{
			Dest: resUser,
		},
	}
}

func TestInsert(t *testing.T) {
	dao,testDb := createMockDao()
	testDb.On("Query", mock.Anything).Return(
		&gorm.DB{
			RowsAffected: 1,
		})
	
	u := &domain.User{Id: 0, Username: "test"}
	dao.Insert(u)
	testDb.AssertExpectations(t)
	assert.Equal(t, resUser.Id, u.Id, "User id should be changed after insert")
	
}

func createMockDao() (dbaccess.DaoApi[domain.User], *MockedDb) {
	testDb := &MockedDb{}
	return dbaccess.CreateDao[domain.User](testDb, &gorm.DB{}) ,testDb
}