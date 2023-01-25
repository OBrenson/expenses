package integration

import (
	"expenses/internal/domain"
	"fmt"
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

func TestGorm(t *testing.T) {
  	ds := "postgres://myusername:mypassword@localhost:5432/postgres?sslmode=disable"
	db, err := xorm.NewEngine("postgres", ds)
	if err != nil {
		panic(err)
	}

  	if err != nil {
    	panic("failed to connect database")
  	}
	db.SetSchema("expenses")
  
	u := domain.User{Username: "test1", IsAdmin: false}
	_,err = db.Insert(u)
	if err != nil {
    	panic("failed to insert" + err.Error())
  	}

	res := domain.User{Username: "test1"}
	ok,err := db.Get(&res)

	if err != nil {
    	panic("err in get" + err.Error())
  	}

	if !ok {
		panic("did not find user")
	}

	fmt.Println(res.Id)

	res.Username = "test2"
	af,err := db.ID(res.Id).Update(res)
	if err != nil {
    	panic("err in update" + err.Error())
  	}
	if af == 0 {
		panic("did not update")
	}

	af,err = db.Delete(res)

	if err != nil {
    	panic("err in delete" + err.Error())
  	}
	if af == 0 {
		panic("did not delete")
	}

}