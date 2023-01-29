package unit

import (
	"expenses/internal/dbaccess"
	"expenses/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockedDb struct {
	mock.Mock
	gorm.DB
}

func (m *MockedDb) Query(dbFunc func (args interface{}) *gorm.DB, args interface{}) *gorm.DB {
	m.Called(1)
	return &gorm.DB{
		RowsAffected: 1,
	}
}

func TestInsert(t *testing.T) {
	dao,testDb := createMockDao()
	// testDb.On("Query", mock.Anything).Return(
	// 	&gorm.DB{
	// 		RowsAffected: 1,
	// 	})
	
	u := &domain.User{}
	dao.Insert(u)
	c := testDb.Calls
	fmt.Println(c)
	//testDb.AssertExpectations(t)
}

func createMockDao() (dbaccess.DaoApi[domain.User], *MockedDb) {
	testDb := new(MockedDb)
	return dbaccess.CreateDao[domain.User](testDb),testDb
}