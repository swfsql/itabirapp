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
	st_err_password_diferente string = "err_password_diferente"
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
	var target models.User
	var allow bool
	if target, allow = allowed(this); !allow {
		return
	}

	sess := this.StartSession()
	//defer sess.SessionRelease()
	user, _ := sess.Get("user").(models.User)

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
	target, allow := allowed(this) 
	if !allow {
		return
	}

	if this.Data["isOwner"] == false || target.User_Type != "poster" {
		this.Redirect("/", 302)
		return
	}

	dado := struct {
        Description string
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


	target.Institution_Description = dado.Description;

	target.Update()

	fmt.Println("editado com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()
}
func (this *UserController) PostUser() {
	target, allow := allowed(this) 
	if !allow {
		return
	}

	if this.Data["isOwner"] == false {
		this.Redirect("/", 302)
		return
	}

	dado := struct {
		Name string
		Email string
		Password string
		Password2 string
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

	if (dado.Password != dado.Password2) {
		status.Status = st_err_password_diferente
		this.Data["json"] = status
		this.ServeJSON()
		return
	}

	target.Name = dado.Name
	target.Email = dado.Email
	target.Password = dado.Password

	target.Update()

	fmt.Println("editado com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()
}


