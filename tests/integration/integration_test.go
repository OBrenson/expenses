package integration

import (
	"expenses/internal/commands"
	"expenses/internal/dbaccess"
	"expenses/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	checkOperation(t, &res.RowsAffected, res.Error, "Create user")

	res = db.Where("username = ?", "test1").Find(&u)
	checkOperation(t, &res.RowsAffected, res.Error, "Where user")

	user := domain.User{Username: "test1"}
	res = db.First(&user)
	checkOperation(t, &res.RowsAffected, res.Error, "Find user")

	res = db.Model(&user).Update("username", "test2")
	checkOperation(t, &res.RowsAffected, res.Error, "Update user")

	d := domain.Data{User: &user, Sum: 100, Type: "shop", Date: time.Now()}
	res = db.Create(&d)
	checkOperation(t, &res.RowsAffected, res.Error, "Insert data")

	res = db.Where("user_id = ?", d.UserId).Delete(&d)
	checkOperation(t, &res.RowsAffected, res.Error, "Delete data")

	res = db.Delete(&user, user.Id)
	checkOperation(t, &res.RowsAffected, res.Error, "Delete user")
}

func TestDao(t *testing.T) {
	db := dbaccess.CreateDb(dsn, "expenses")
	dao := dbaccess.CreateDao(db)
	testInsert(t, dao)
	testGet(t, dao)
	testUpdate(t, dao)
	testDelete(t, dao)

	a := commands.Add{
		Action: commands.Action { 
			Next: commands.Delete{}}}
	a.GetType()
}

func testGet(t *testing.T, dao dbaccess.DaoApi) {
	err := dao.GetByField(user, "username", user.Username)
	checkOperation(t, nil, err, "Get " + user.ToString())
}

func testInsert(t *testing.T, dao dbaccess.DaoApi) {
	err := dao.Insert(user)
	checkOperation(t, nil, err, "Insert " + user.ToString())
}

func testUpdate(t *testing.T, dao dbaccess.DaoApi) {
	err := dao.UpdateField(user, "username", "testUpdated")
	checkOperation(t, nil, err, "Update" + user.ToString())
	assert.Equal(t, "testUpdated", user.Username)
}

func testDelete(t *testing.T, dao dbaccess.DaoApi) {
	err := dao.DeleteByField(user, "id", user.Id)
	checkOperation(t, nil, err, "Delete " + user.ToString())
}

func checkOperation(t *testing.T, af *int64, err error, op string) {
	if err != nil {
		t.Error("err in delete " + err.Error())
	}
	if af != nil && *af == 0 {
		t.Error("did not " + op)
	}
}
