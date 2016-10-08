package models

import (
	_ "errors"
	_ "fmt"
	//"reflect"
	_ "github.com/astaxie/beego/orm"
)


type Post struct {
	Id uint16 // 
	User *User `orm:"rel(fk)"`
	Title string  //
	Subtitle string  //
	Text string  //
	Thumbnail []byte `orm:"null"` // (imagem)
}


