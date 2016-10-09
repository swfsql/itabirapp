package models

import (
	_ "errors"
	_ "fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)


type User struct {
	Id int // 
	User_Type string // [poster, moderator]
	User_Privilege uint8 `orm:"-"` // [low, high]
	Name string  //
	NameTag string // (normalized name)
	NameIdTag string `orm:"unique"` // (normalized name_id)
	Email string `orm:"unique"`// 
	Password string // 
	//	
	IsAuthorized bool // pode postar ou nao
	Institution_Tag string `orm:"null"` // [republica, professor, ...]
	Institution_Description string `orm:"null"` // (somos a UP e tals)
	Institution_Thumbnail []byte `orm:"-"` // (imagem)
	Addr_Street string `orm:"null"` // Girassol
	Addr_Number string `orm:"null"` // 123
	Addr_Complement string `orm:"null"` // 103
	Addr_Neighborhood string `orm:"null"` // santo antonio
	Addr_City string `orm:"null"` // itabira
}


func GetUserByEmail(email string) (user User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err = qs.Filter("Email", email).RelatedSel().One(&user)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func GetUserById(id int) (user User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err = qs.Filter("Id", id).RelatedSel().One(&user)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func CountNameTag(nametag string) (quantity int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	quantity, err = qs.Filter("NameTag", nametag).Count()
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}