package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	"github.com/swfsql/itabirapp/models"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Layout = "layout.html"

	sess := this.StartSession()
	//defer sess.SessionRelease()

	if user, loggedIn := sess.Get("user").(models.User); loggedIn == true {
		this.Data["IsLoggedIn"] = loggedIn
		user.Password = ""
		this.Data["User"] = user
	}
}
