package dbaccess

import (
	"expenses/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DaoApi[T domain.Data | domain.User] interface {
	Insert(val *T) (*T, error)
	GetByField(field string, val interface{}) (*T, error)
	DeleteByField(field string, val interface{}) (*T, error)
	UpdateField(id int, field string, val interface{}) (*T, error)
}

type DbApi interface {
	Query(func (args... interface{}) *gorm.DB) *gorm.DB
}

type DbService struct {
	gorm.DB
}

type DaoService[T domain.Data | domain.User] struct {
	db *DbService
}

func CreateDb(dsn string, schemaName string) (db DbApi) {
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
	return &DbService{*gdb}
}

func CreateDao[T domain.Data | domain.User](db DbApi) DaoApi[T] {
	db.(DbService)
	return &DaoService[T]{db: db.(DbService)}
}

func (db *DbService) Query(func(args... interface{}) *gorm.DB) *gorm.DB {
	
}

func (ds *DaoService[T]) Insert(val *T) (*T, error) {
	return nil, nil
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
