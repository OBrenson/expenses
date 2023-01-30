package dbaccess

import (
	"expenses/internal/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DaoApi[T domain.Data | domain.User] interface {
	Insert(val *T) error
	GetByField(field string, val interface{}) (*T, error)
	DeleteByField(entity *T, field string, val interface{}) error
	UpdateField(id int, field string, val interface{}) (*T, error)
}

type DaoService[T domain.Data | domain.User] struct {
	db        *gorm.DB
	dbService DbApi
	tableName string
}

func CreateDb(dsn string, schemaName string) *gorm.DB {
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

func (ds *DaoService[T]) Insert(val *T) (err error) {
	defer ds.recoverPanic(&err)
	res := ds.dbService.Query(ds.db.Create, val)
	return ds.handleError(res)
}

func (ds *DaoService[T]) GetByField(field string, val interface{}) (*T, error) {
	return nil, nil
}

func (ds *DaoService[T]) DeleteByField(entity *T, field string, val interface{}) (err error) {
	defer ds.recoverPanic(&err)
	cond := ds.db.Where(fmt.Sprintf("%s = ?", field), val)
	res := ds.dbService.ExpandedQuery(cond.Delete, entity, nil)
	return ds.handleError(res)
}

func (ds *DaoService[T]) UpdateField(id int, field string, val interface{}) (*T, error) {
	return nil, nil
}

func (ds *DaoService[T]) recoverPanic(err *error) {
	if r := recover(); r != nil {
		*err = &DbPanicErr{Method: "Insert", Message: fmt.Sprintf("%v",r)}
	}
}

func (ds *DaoService[T]) handleError(g *gorm.DB) error {
	if g.RowsAffected == 0 {
		return &NoRowsAffected{Method: "Insert"}
	}
	return g.Error
}