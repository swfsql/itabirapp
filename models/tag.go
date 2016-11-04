package models

import (
	_ "errors"
	_ "fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id    int     //
	Posts []*Post `orm:"reverse(many)"`
	Name  string  //
}

func GetTagByName(name string) (tag Tag, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("tag")
	err = qs.Filter("Name", name).One(&tag)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func AppendTagsForPost(post *Post, strings []string) (num int64, err error) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(post, "Tags")
	for _, s := range strings {

		tag, err := GetTagByName(s)
		if err != nil {
			tag = Tag{Name: s}
			o.Insert(&tag)
		}

		_, err2 := m2m.Add(&tag)
		if err2 != nil {
			err2 = ErrNoRows
		}
	}

	return
}

func RemoveUserTagsForPost(post *Post) (num int64, err error) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(post, "Tags")
	var tags []*Tag = post.Tags[2:]
	num, err = m2m.Remove(tags)

	return
}
