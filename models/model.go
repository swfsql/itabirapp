package models

import (
	"errors"
	"fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)

// ERRS
var (
	ErrNoRows = errors.New("<QuerySeter> no row found")
)

func init() {
	fmt.Println("HUEHUEHUE")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "itabirapp_user:itabirapp_senha@itabirapp_host/itabirapp_db?charset=utf8", 30)
	orm.RegisterModel(new(User), new(Tag), new(Post), new(Document))
	orm.RunSyncdb("default", true, true)
	orm.Debug = true
	createData()
}
