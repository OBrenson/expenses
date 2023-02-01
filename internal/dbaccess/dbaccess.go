package dbaccess

import (
	"expenses/internal/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DaoApi interface {
	Insert(val domain.Entity) error
	GetByField(entity domain.Entity, field string, val interface{}) error
	DeleteByField(entity domain.Entity, field string, val interface{}) error
	UpdateField(entity domain.Entity, field string, val interface{}) error
}

type DaoService struct {
	db        *gorm.DB
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

func CreateDao(db *gorm.DB) DaoApi {
	return &DaoService{db: db}
}

func (ds *DaoService) Insert(val domain.Entity) (err error) {
	defer ds.recoverPanic(&err)
	res := ds.db.Create(val)
	return ds.handleError(res)
}

func (ds *DaoService) GetByField(entity domain.Entity, field string, val interface{}) (err error) {
	defer ds.recoverPanic(&err)
	res := ds.db.Where(fmt.Sprintf("%s = ?", field), val).Find(entity)
	return ds.handleError(res)
}

func (ds *DaoService) DeleteByField(entity domain.Entity, field string, val interface{}) (err error) {
	defer ds.recoverPanic(&err)
	res := ds.db.Where(fmt.Sprintf("%s = ?", field), val).Delete(entity)
	return ds.handleError(res)
}

func (ds *DaoService) UpdateField(entity domain.Entity, field string, val interface{}) (err error) {
	defer ds.recoverPanic(&err)
	res := ds.db.Model(entity).Update(field, val)
	return ds.handleError(res)
}

func (ds *DaoService) recoverPanic(err *error) {
	if r := recover(); r != nil {
		*err = &DbPanicErr{Method: "Insert", Message: fmt.Sprintf("%v",r)}
	}
}

func (ds *DaoService) handleError(g *gorm.DB) error {
	if g.RowsAffected == 0 {
		return &NoRowsAffected{Method: "Insert"}
	}
	return g.Error
}