package domain

import "time"

type User struct {
	Id int `gorm:"column:id"`
	Username string `gorm:"column:username"`
	IsAdmin string `gorm:"column:is_admin"`
}

type Data struct {
	User User
	Sum float64 `gorm:"column:sum"`
	Type string `gorm:"column:expense_type"`
	Date time.Time `gorm:"column:expense_date"`
}