package models

import (
	_ "errors"
	_ "fmt"
	//"reflect"
	_ "github.com/astaxie/beego/orm"
)


type Document struct {
	Id int // 
	User *User `orm:"rel(fk)"`
	Name string //
	File []byte `orm:"-"`// (pdf, img, etc)
}


