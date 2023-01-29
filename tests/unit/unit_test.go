package unit

import (
	"expenses/internal/dbaccess"
	"expenses/internal/domain"
	"testing"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockedDb struct {
	mock.Mock
}

func (m *MockedDb) Query(dbFunc func (args interface{}) *gorm.DB, args interface{}) *gorm.DB {
	m.Called(args)
	return &gorm.DB{
		RowsAffected: 1,
	}
}

func TestInsert(t *testing.T) {
	dao,testDb := createMockDao()
	testDb.On("Query", mock.Anything).Return(
		&gorm.DB{
			RowsAffected: 1,
		})
	
	u := &domain.User{}
	dao.Insert(u)
	testDb.AssertExpectations(t)
}

func createMockDao() (dbaccess.DaoApi[domain.User], *MockedDb) {
	testDb := &MockedDb{}
	return dbaccess.CreateDao[domain.User](testDb, &gorm.DB{}) ,testDb
}