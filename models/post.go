package models

import (
	_ "errors"
	 "fmt"
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
	Tags []*Tag `orm:"rel(m2m)"`

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

func (this Post) Update() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Update(&this)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func (this *Post) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func (this *Post) New() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Insert(this)

	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func GetPostsByTag(tag string) (num int64, posts []Post, err error) {
	o := orm.NewOrm()
//	o.QueryTable("post").Filter("Tag__Name", tag).Distinct().RelatedSel().All(&posts)


	//"n0_1"
	//"instTag0"


	cond := orm.NewCondition()
	//cond1 := cond.And("Tags__Tag__Name", "instTag0")
	cond2 := cond.And("Tags__Tag__Name", "n0_1")
	//cond3 := cond1.AndCond(cond2)
	//o.QueryTable("post").Filter("Tags__Tag__Name", "n0_1").All(&posts)
	o.QueryTable("post").SetCond(cond2).All(&posts)
	

	for i,p := range posts {
		fmt.Println(i, ":", p.Id)
		fmt.Println(i, ":", p.Title)
		//fmt.Println(i, ":", p.User.Name)
	}




	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

