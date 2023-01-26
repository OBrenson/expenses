package dbaccess

import (
	"expenses/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Api[T domain.Data | domain.User] interface {

	Insert(val *T) (*T,error)
	GetByField (field string, val interface{}) (*T,error)
	DeleteByField (field string, val interface{}) (*T,error)
	UpdateField(id int, field string, val interface{}) (*T, error)

}

type DaoService[T domain.Data | domain.User] struct {
	db *gorm.DB
}

func CreateDb(dsn string, schemaName string) (db *gorm.DB) {
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   schemaName + ".",
				SingularTable: true,
			}})
	if err != nil {
		panic(err)
	}
	return db
}

func CreateDao[T domain.Data | domain.User](db *gorm.DB) Api[T] {
	return &DaoService[T]{db: db}
}

func (ds *DaoService[T]) Insert(val *T) (*T,error) {
	return nil, nil;
}

func (ds *DaoService[T]) GetByField (field string, val interface{}) (*T,error) {
	return nil, nil;
}

func (ds *DaoService[T]) DeleteByField (field string, val interface{}) (*T,error) {
	return nil, nil;
}

func (ds *DaoService[T]) UpdateField(id int, field string, val interface{}) (*T, error){
	return nil, nil;
}