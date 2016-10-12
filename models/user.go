package models


import (
	"github.com/astaxie/beego/orm"
	_ "errors"
	//"reflect"
    _ "fmt"
    "strings"
    "unicode"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
    "regexp"
	"strconv"
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
	IsAuthorized bool `orm:"null"` // pode postar ou nao
	InstitutionType *InstitutionType `orm:"rel(fk);null"` // [1=republica, 2=professor, ...]
	Institution_Tag string `orm:"null"` // [republica, professor, ...] (informação duplicada)
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

func GetUsers() (users []*User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	_, err = qs.All(&users)
	if err == orm.ErrNoRows {
		err = ErrNoRows
	}
	return
}

func isMn(r rune) bool {
    return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
func (this User) Update() (num int64, err error) {
	// for nameTag and nameIdTag
    t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	reg, _ := regexp.Compile("[^a-z]+")
    nameTag, _, _ := transform.String(t, this.Name) // É2 -> E2
    nameTag2 := strings.ToLower(nameTag) // E2 -> e2
    this.NameTag = reg.ReplaceAllString(nameTag2, "") // e2 -> e
    j, _ := CountNameTag(this.NameTag);
	js := strconv.FormatInt(j,10)
    this.NameIdTag = this.NameTag + "_" + js

	o := orm.NewOrm()
	num, err = o.Update(&this)
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