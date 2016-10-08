package models

import (
	"errors"
	"fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)


type User struct {
	Id uint16 // 
	User_Type string // [poster, moderator]
	User_Privilege uint8 `orm:"-"` // [low, high]
	Name string  //
	NameTag string // (normalized name)
	NameIdTag string `orm:"unique"` // (normalized name_id)
	Email string // 
	Password string // 
	//
	Institution_Tag string `orm:"null"` // [republica, professor, ...]
	Institution_Description string `orm:"null"` // (somos a UP e tals)
	Institution_Thumbnail []byte `orm:"null"` // (imagem)
	Addr_Street string `orm:"null"` // Girassol
	Addr_Number string `orm:"null"` // 123
	Addr_Complement string `orm:"null"` // 103
	Addr_Neighborhood string `orm:"null"` // santo antonio
	Addr_City string `orm:"null"` // itabira
}


