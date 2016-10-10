package controllers

import (
	"encoding/json"
	"fmt"

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

func allowed(this *UserController) (models.User, bool) {
	sess := this.StartSession()
	//defer sess.SessionRelease()
	targetId64, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 32)
	targetId := int(targetId64)
	var target models.User

	user, loggedIn := sess.Get("user").(models.User)
	// not logged in
	if !loggedIn {
		defer this.DestroySession()
		this.Redirect("/", 302)
		return target, false
	}

	var err_target error
	target, err_target = models.GetUserById(targetId)
	// target not found
	if err_target == models.ErrNoRows {
		// por JSON deveria retornar "usuario_nao_existe", mas isso nunca acontecerá
		this.Redirect("/", 302)
		return target, false
	}

	this.Data["IsAuthorized"] = false
	if user.Id == targetId {
		this.Data["IsOwner"] = true
	} else if (user.User_Type != "moderator" || target.User_Type == "moderator") {
		// not allowed
		this.Redirect("/", 302)
		return target, false
	} else {
		this.Data["IsAuthorized"] = true
	}

	target.Password = ""
	this.Data["Target"] = target
	return target, true
}

func (this *UserController) GetEdit() {
	if _, allow := allowed(this); !allow {
		return
	}


	this.TplName = "user/edit.html"
	this.Data["HeadTitle"] = "Configurações da conta"
	this.Data["HeadStyles"] = []string{}
    this.Data["HeadScripts"] = []string{"user/edit.js"}
	this.Render()
}

func (this *UserController) ToggleAuthorization() {
	if target, allow := allowed(this); !allow {
		return
	}

	user := this.Data["User"]
	if user.User_Type != "moderator" || target.User_Type != "poster" {
		this.Redirect("/", 302)
		return
	}

	target.IsAuthorized = !target.IsAuthorized
	target.Update()

	status := struct{ Status string }{""}

	fmt.Println("editado com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()
}

func (this *UserController) PostAddress() {

	target, allow := allowed(this) 
	if !allow {
		return
	}

	if this.Data["isOwner"] == false || target.User_Type != "poster" {
		this.Redirect("/", 302)
		return
	}

	dado := struct {
	  	Street string
	  	Number string
	  	Complement string
	  	Neighborhood string
	  	City string
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dado)
	if err != nil {
		fmt.Println(err)
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("JSON invalido"))
		return
	}
	fmt.Println(dado)

	status := struct{ Status string }{""}

	target.Addr_Street = dado.Street
	target.Addr_Number = dado.Number
	target.Addr_Complement = dado.Complement
	target.Addr_Neighborhood = dado.Neighborhood
	target.Addr_City = dado.City

	target.Update()

	fmt.Println("editado com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()



}
func (this *UserController) PostInstitution() {
}
func (this *UserController) PostUser() {
}


