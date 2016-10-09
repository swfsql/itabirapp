package controllers

import (
	_ "encoding/json"
	_ "fmt"

	//"github.com/astaxie/beego/orm"
	"github.com/swfsql/itabirapp/models"
	"strconv"
)

// ERRS
var (
)

type UserController struct {
	BaseController
}

func (this *UserController) GetEdit() {
	sess := this.StartSession()
	//defer sess.SessionRelease()
	targetId64, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 32)
	targetId := int(targetId64)

	user, loggedIn := sess.Get("user").(models.User)
	if !loggedIn {
		defer this.DestroySession()
		this.Redirect("/", 302)
	}

	target, err_target := models.GetUserById(targetId)
	if err_target == models.ErrNoRows {
		this.Redirect("/", 302)
	}

	this.Data["IsAuthorized"] = false
	if user.Id == targetId {
		this.Data["IsOwner"] = true
	} else if (user.User_Type != "moderator" || target.User_Type == "moderator") {
		this.Redirect("/", 302)
	} else {
		this.Data["IsAuthorized"] = true
	}

	target.Password = ""
	this.Data["Target"] = target

	this.TplName = "user/edit.html"
	this.Data["HeadTitle"] = "Configurações da conta"
	this.Data["HeadStyles"] = []string{}
    this.Data["HeadScripts"] = []string{"user/edit.js"}
	this.Render()
}


