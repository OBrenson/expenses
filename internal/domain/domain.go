package domain

import "time"

type User struct {
	Id int `json:"column:id" xorm:"serial pk"`
	Username string `json:"column:username"`
	IsAdmin bool `json:"column:is_admin"`
}

type Data struct {
	User User `xorm:"extends"`
	Sum float64 `json:"column:sum"`
	Type string `json:"column:expense_type"`
	Date time.Time `json:"column:expense_date"`
}