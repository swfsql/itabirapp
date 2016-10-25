package models

import (
	_ "fmt"
	"strconv"
	_ "strings"

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
		j, _ := CountNameTag(N + s)
		js := strconv.FormatInt(j, 10)

		o.Insert(&User{User_Type: "moderator", Name: N + s, NameTag: N + s, NameIdTag: N + s + "_" + js,
			Email: N + s + "@s", Password: N + s, Institution_Tag: "", Institution_Description: "",
			Addr_Street: "", Addr_Number: "", Addr_Complement: "", Addr_Neighborhood: "", Addr_City: ""})
	}

	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		j, _ := CountNameTag(N + s)
		js := strconv.FormatInt(j, 10)
		var authorized bool
		if i%3 == 0 {
			authorized = true
		} else {
			authorized = false
		}

		o.Insert(&User{User_Type: "poster", Name: N + s, NameTag: N + s, NameIdTag: N + s + "_" + js,
			Email: N + s + "@a", Password: N + s, IsAuthorized: authorized, Institution_Tag: "instTag" + s,
			Institution_Description: "descricao" + s, Addr_Street: "addrSt" + s, Addr_Number: "n." + s,
			Addr_Complement: "compl." + s, Addr_Neighborhood: "neigh" + s, Addr_City: "city" + s})
	}

	var posters []*User
	qs := o.QueryTable("user")
	_, _ = qs.Filter("User_Type", "poster").All(&posters)

	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		user := posters[(i*2)%10]
		post := &Post{User: user, Title: "título_" + s,
			Subtitle: "subtítulo_" + s, Text: "text" + s}
		if i == 5 {
			post.Text = "![errou](https://i.makeagif.com/media/5-23-2014/uCZNYo.gif)"
		}
		o.Insert(post)

		var tags []string
		tags = append(tags, user.NameIdTag)
		tags = append(tags, user.Institution_Tag)

		if i == 0 {
			tags = append(tags, "rep")
			tags = append(tags, "masc")
			tags = append(tags, "feminino")
			tags = append(tags, "professor")

		}
		if i == 1 {
			tags = append(tags, "rep")
			tags = append(tags, "masc")
			tags = append(tags, "cetrulo")
			tags = append(tags, "thiago")

		}
		if i == 2 {
			tags = append(tags, "rep")
			tags = append(tags, "cetrulo")
			tags = append(tags, "nova")
			tags = append(tags, "velha")

		}
		AppendTagsForPost(post, tags)
	}

}
