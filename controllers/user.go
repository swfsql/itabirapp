package controllers

import (
	"encoding/json"
	"fmt"

	//"github.com/astaxie/beego/orm"
	"github.com/swfsql/itabirapp/models"
	"strconv"
	"bytes"
)

// ERRS
var (
	st_err_password_diferente string = "err_password_diferente"
)

type UserController struct {
	BaseController
}

func editAllowed(this *UserController) (models.User, bool) {
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
		// not editAllowed
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
	if _, allow := editAllowed(this); !allow {
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
	if target, allow = editAllowed(this); !allow {
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

func (this *UserController) PostEditAddress() {
	target, allow := editAllowed(this) 
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

func (this *UserController) PostEditInstitution() {
	target, allow := editAllowed(this) 
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

func (this *UserController) PostEditUser() {
	target, allow := editAllowed(this) 
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


func (this *UserController) GetList() {
	sess := this.StartSession()
	//defer sess.SessionRelease()

	user, loggedIn := sess.Get("user").(models.User)
	// not logged in
	if !loggedIn {
		defer this.DestroySession()
		this.Redirect("/", 302)
		return 
	} else if user.User_Type != "moderator" {
		this.Redirect("/", 302)
		return 
	}

	users, err_users := models.GetUsers()
	if err_users == models.ErrNoRows {
		users = []*models.User{}
	}

	var moderators []*models.User 
	var authorized []*models.User 
	var unauthorized []*models.User 
	fmt.Println("os usuários pêgos:")
	for _, u := range users {
		fmt.Println("~")
		// limit the description to 15 chars
		runes := bytes.Runes([]byte(u.Institution_Description))
		if len(runes) > 15 {
			u.Institution_Description = string(runes[:15])
		}
		// filter the users
		if u.User_Type == "moderator" {
			moderators = append(moderators, u)
		} else if u.IsAuthorized == true {
			authorized = append(authorized, u)
		} else {
			unauthorized = append(unauthorized, u)
		}
	} 

	this.Data["Moderators"] = moderators;
	this.Data["Authorized"] = authorized;
	this.Data["Unauthorized"] = unauthorized;

	this.TplName = "user/list.html"
	this.Data["HeadTitle"] = "Listagem das contas"
	this.Data["HeadStyles"] = []string{"datatables.min.css"}
    this.Data["HeadScripts"] = []string{"user/list.js", "datatables.min.js"}
	this.Render()
}


func (this *UserController) GetDelete() {
	var target models.User
	var allow bool
	if target, allow = editAllowed(this); !allow {
		return
	}

	target.Delete()
	if this.Data["IsOwner"] == true {
		fmt.Println("SE AUTO-REMOVEU!")
		this.DestroySession()
		//this.Redirect("/", 302)
		//return
	} 

	status := struct{ Status string }{""}

	fmt.Println("removido com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()
}


func (this *UserController) GetNew() {
	fmt.Println("hueee hue br")

	insts, err_insts := models.GetInstitutions()
	if err_insts == models.ErrNoRows {
		insts = []*models.Institution_type{}
	}

	this.Data["Institutions"] = insts

	this.TplName = "user/new.html"
	this.Data["HeadTitle"] = "Criar nova conta"
	this.Data["HeadStyles"] = []string{}
    this.Data["HeadScripts"] = []string{"user/new.js"}
	this.Render()
}

func (this *UserController) PostNew() {
	dado := struct {
		Name string 
		Email string
		Password string
		Password2 string // -
		//	
		Institution_Description string
		Institution_Tag string
		//
		Addr_Street string
		Addr_Number  string
		Addr_Complement  string
		Addr_Neighborhood  string
		Addr_City  string
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dado)
	if err != nil {
		fmt.Println(err)
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("JSON invalido"))
		return
	}
	fmt.Println(dado)

	if dado.Password != dado.Password2 {
		return
	}

	fmt.Println("criado com sucesso")
	inst, _ := models.GetInstitutionByName(dado.Institution_Tag)
	
	fmt.Println("criado com sucesso")
	var user models.User
	user.User_Type = "poster"
	user.IsAuthorized = false
	user.Institution_type = &inst
	//
	user.Name = dado.Name 
	user.Email = dado.Email 
	user.Password = dado.Password 
	//	
	user.Institution_Description = dado.Institution_Description 
	user.Institution_Tag = dado.Institution_Tag 
	//
	user.Addr_Street = dado.Addr_Street 
	user.Addr_Number = dado.Addr_Number  
	user.Addr_Complement = dado.Addr_Complement  
	user.Addr_Neighborhood = dado.Addr_Neighborhood  
	user.Addr_City = dado.Addr_City  

	fmt.Println("criado com sucesso")
	user.New()

	fmt.Println("criado com sucesso")
	status := struct{ Status string }{""}

	fmt.Println("criado com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()
}
