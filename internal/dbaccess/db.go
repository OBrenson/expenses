package dbaccess

import "gorm.io/gorm"

type DbApi interface {
	Query(dbFunc func(arg interface{}) *gorm.DB, arg interface{}) *gorm.DB
	ExpandedQuery(dbFunc func (val interface{}, conds ...interface{}) *gorm.DB, value interface{}, conds ...interface{}) *gorm.DB
}

type DbService struct {
}

func (db *DbService) Query(dbFunc func(arg interface{}) *gorm.DB, arg interface{}) *gorm.DB {
	return dbFunc(arg)
}

func (db *DbService) ExpandedQuery(dbFunc func (val interface{}, conds ...interface{}) *gorm.DB, 
									value interface{}, conds ...interface{}) *gorm.DB {
	return dbFunc(value, conds...)
}