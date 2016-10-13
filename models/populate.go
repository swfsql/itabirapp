package models

import (
	 "fmt"
	"strconv"

	 "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


func createData() {
	o := orm.NewOrm()
	N := "n"

	for i := 0; i < 5; i++ {
		s := strconv.Itoa(i)
		o.Insert(&Institution_type{Name: "instituição_" + s})
	}


	for i := 0; i < 4; i++ {
		s := strconv.Itoa(i)
		j, _ := CountNameTag(N + s);
		js := strconv.FormatInt(j,10)

		o.Insert(&User{User_Type: "moderator", Name: N + s, NameTag: N + s, NameIdTag: N + s + "_" + js,
		Email: N + s + "@s", Password: N + s, Institution_Tag: "", Institution_Description: "", 
		Addr_Street: "", Addr_Number: "", Addr_Complement: "", Addr_Neighborhood: "", Addr_City: "",})
	}

	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		j, _ := CountNameTag(N + s);
		js := strconv.FormatInt(j,10)
		var authorized bool
		if i % 3 == 0 { 
			authorized = true 
		} else { 
			authorized = false
		};

		o.Insert(&User{User_Type: "poster", Name: N + s, NameTag: N + s, NameIdTag: N + s + "_" + js,
		Email: N + s + "@a", Password: N + s, IsAuthorized: authorized,  Institution_Tag: "instTag" + s, 
		Institution_Description: "descricao" + s, Addr_Street: "addrSt" + s, Addr_Number: "n." + s, 
		Addr_Complement: "compl." + s, Addr_Neighborhood: "neigh" + s, Addr_City: "city" + s,})
	}

	var users []*User
	qs := o.QueryTable("user")
	qs.All(&users)

	for _, u := range users {
	fmt.Println(u)

	}


}
