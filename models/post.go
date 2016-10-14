package models

import (
	_ "errors"
	_ "fmt"
	//"reflect"
	 "github.com/astaxie/beego/orm"
)


type Post struct {
	Id int // 
	User *User `orm:"rel(fk)"`
	Title string  //
	Subtitle string  //
	Text string  //
	Thumbnail []byte `orm:"-"` // (imagem)
}

func GetPostById(id int) (post Post, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("post")
	err = qs.Filter("Id", id).RelatedSel().One(&post)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}
