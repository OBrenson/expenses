package unit

import (
	"expenses/internal/dbaccess"
	"expenses/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var resId int = 1
var resUser = domain.User{
	Id: &resId,
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

func (m *MockedDb) ExpandedQuery(dbFunc func (val interface{}, conds ...interface{}) *gorm.DB, 
									value interface{}, conds ...interface{}) *gorm.DB {
	m.Called(value)
	if value == nil {
		panic("arg is nil")
	}
	u := value.(*domain.User)
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
	var id int = 0;
	u := &domain.User{Id: &id, Username: "test"}

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

func TestDelete(t *testing.T) {
	var id int = 0;
	u := &domain.User{Id: &id, Username: "test"}

	dao,testDb := createMockDao()
	testDb.On("ExpandedQuery", u).Return(
		&gorm.DB{
			RowsAffected: 1,
		})
	err := dao.DeleteByField(u, "id", "1")
	assert.Nil(t, err)
	testDb.AssertExpectations(t)
	
	err = dao.DeleteByField(u, "id", "1")
	assert.IsType(t, &dbaccess.NoRowsAffected{}, err)

	err = dao.DeleteByField(u, "id", "1")
	assert.IsType(t, &dbaccess.DbPanicErr{}, err)
}

func createMockDao() (dbaccess.DaoApi[domain.User], *MockedDb) {
	testDb := &MockedDb{}
	return dbaccess.CreateDao[domain.User](testDb, &gorm.DB{}) ,testDb
}