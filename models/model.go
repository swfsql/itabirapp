package models

import (
	"errors"
	_ "fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)

// ERRS
var (
	ErrNoRows = errors.New("<QuerySeter> no row found")
)

func init() {
	orm.Debug = false
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "itabirapp_user:itabirapp_senha@tcp(localhost:3306)/itabirapp_db?charset=utf8", 30)
	orm.RegisterModel(new(User), new(Tag), new(Post), new(Document), new(Institution_type))
	orm.RunSyncdb("default", true, false)
	createData()
}
