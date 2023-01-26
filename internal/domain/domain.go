package domain

import "time"

type User struct {
	Id int `gorm:"primaryKey column:id"`
	Username string `gorm:"column:username;unique"`
	IsAdmin bool `gorm:"column:is_admin"`
}

type Data struct {
	UserId int `gorm:"column:user_id"`
	User *User `gorm:"foreignKey:UserId;references:Id"`
	Sum float64 `gorm:"column:sum"`
	Type string `gorm:"column:expense_type"`
	Date time.Time `gorm:"column:expense_date"`
}