package models

import (
	_ "errors"
	_ "fmt"
	//"reflect"
	_ "github.com/astaxie/beego/orm"
)


type Tag struct {
	Id uint16 // 
	Post *Post `orm:"rel(fk)"`
	Name string  //
}


