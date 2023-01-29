package dbaccess

import (
	"errors"
	"expenses/internal/domain"
	"reflect"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DaoApi[T domain.Data | domain.User] interface {
	Insert(val *T) error
	GetByField(field string, val interface{}) (*T, error)
	DeleteByField(field string, val interface{}) (*T, error)
	UpdateField(id int, field string, val interface{}) (*T, error)
}

type DbApi interface {
	Query(dbFunc func (args interface{}) *gorm.DB, args interface{}) *gorm.DB
}

type DbService struct {

}

type DaoService[T domain.Data | domain.User] struct {
	db *gorm.DB
	dbService DbApi
	tableName string
}

func CreateDb(dsn string, schemaName string) (*gorm.DB) {
	gdb, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   schemaName + ".",
				SingularTable: true,
			}})
	if err != nil {
		panic(err)
	}
	return gdb
}

func CreateDao[T domain.Data | domain.User](dbApi DbApi, db *gorm.DB) DaoApi[T] {
	
	return &DaoService[T]{dbService: dbApi, db: db}
}

func (db *DbService) Query(dbFunc func (args interface{}) *gorm.DB, args interface{}) *gorm.DB {
	return nil
}

func (ds *DaoService[T]) Insert(val *T) error {
	f := ds.db.Create
	res := ds.dbService.Query(f, val)
	if res.RowsAffected != 0 {
		return &NoRowsAffected{EntityName: val.(domain.User).Id}
	}
	return res.Error
}

func (ds *DaoService[T]) GetByField(field string, val interface{}) (*T, error) {
	return nil, nil
}

func (ds *DaoService[T]) DeleteByField(field string, val interface{}) (*T, error) {
	return nil, nil
}

func (ds *DaoService[T]) UpdateField(id int, field string, val interface{}) (*T, error) {
	return nil, nil
}
