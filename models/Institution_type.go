package models

import (
	_ "errors"
	_ "fmt"
	//"reflect"
	"github.com/astaxie/beego/orm"
)


type Institution_type struct {
	Id int // 
	Name string  //
}

func GetInstitutions() (insts []*Institution_type, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("institution_type")
	_, err = qs.All(&insts)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func GetInstitutionByName(name string) (inst Institution_type, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("institution_type")
	err = qs.Filter("Name", name).RelatedSel().One(&inst)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}
