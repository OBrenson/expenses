package integration

import (
	"expenses/internal/dbaccess"
	"expenses/internal/domain"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	dsn = "host=localhost user=myusername password=mypassword dbname=postgres port=5432 sslmode=disable"
)
var user = &domain.User{Username: "test1", IsAdmin: false}

func TestGorm(t *testing.T) {
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "expenses.",
				SingularTable: true,
			}})
	if err != nil {
		panic(err)
	}

	u := &domain.User{Username: "test1", IsAdmin: false}
	res := db.Create(u)
	checkOperation(&res.RowsAffected, res.Error, "Create user")

	user := domain.User{Username: "test1"}
	res = db.First(&user)
	checkOperation(&res.RowsAffected, res.Error, "Find user")

	res = db.Model(&user).Update("username", "test2")
	checkOperation(&res.RowsAffected, res.Error, "Update user")

	d := domain.Data{User: &user, Sum: 100, Type: "shop", Date: time.Now()}
	res = db.Create(&d)
	checkOperation(&res.RowsAffected, res.Error, "Insert data")

	res = db.Where("user_id = ?", d.UserId).Delete(&d)
	checkOperation(&res.RowsAffected, res.Error, "Delete data")

	res = db.Delete(&user, user.Id)
	checkOperation(&res.RowsAffected, res.Error, "Delete user")
}

func TestInsert(t *testing.T) {
	db := dbaccess.CreateDb(dsn, "expenses")
	uDao := dbaccess.CreateDao[domain.User](&dbaccess.DbService{} ,db)
	err := uDao.Insert(user)
	if err != nil {
		panic(err)
	}
	if user.Id == nil {
		panic("Id field can not be nil after insert.")
	}
}

func TestDelete(t *testing.T) {

}

func checkOperation(af *int64, err error, op string) {
	if err != nil {
		panic("err in delete " + err.Error())
	}
	if af != nil && *af == 0 {
		panic("did not " + op)
	}
}
