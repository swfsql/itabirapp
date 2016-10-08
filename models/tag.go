package models

import (
	"errors"
	"fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)


type Tag struct {
	Id uint16 // 
	Post *Post `orm:"rel(fk)"`
	Name string  //
}


