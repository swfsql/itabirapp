package models

import (
	_ "errors"
	_ "fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)


type Tag struct {
	Id int // 
	Post *Post `orm:"rel(fk)"`
	Name string  //
}


func tagForPost(post *Post, name string)(tag Tag) {
	tag = Tag{
		Post: post,
		Name: name,
	}
	return
}


func AppendTagsForPost(post *Post, strings []string) (num int64, err error) {
	var tags []Tag
	for _, s := range strings {
		tag := tagForPost(post, s)
		tags = append(tags, tag)
	}

	o := orm.NewOrm()
	num, err = o.InsertMulti(20, tags)

	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}
