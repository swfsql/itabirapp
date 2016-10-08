package models

import (
	"errors"
	"fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)


type Document struct {
	Id uint16 // 
	User *User `orm:"rel(fk)"`
	Name string //
	File []byte // (pdf, img, etc)
}


