package controllers

import (
	"encoding/json"
	"fmt"

	//"github.com/astaxie/beego/orm"
	"github.com/swfsql/itabirapp/models"
)

// ERRS
var (
	st_ok                   string = "ok"
	st_err_usuario_inexiste string = "err_usuario_inexiste"
	st_err_password_invalida   string = "err_password_invalida"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	this.TplName = "login.html"
	this.Data["HeadTitle"] = "Login Title"
	this.Data["HeadStyles"] = []string{}
    this.Data["HeadScripts"] = []string{"login.js"}
	this.Render()
}

func (this *LoginController) LoginPost() {
	dado := struct {
		Email string
		Password string
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dado)
	if err != nil {
		fmt.Println(err)
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("JSON invalido"))
		return
	}
	//fmt.Println(dado)

	//md5.Password := md5.New()
	//io.WriteString(md5Password, dado.Password)
	//buffer := bytes.NewBuffer(nil)
	//fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	//dado.Password = buffer.String()

	status := struct{ Status string }{""}
	user, err_user := models.GetUserByEmail(dado.Email)
	//o := orm.NewOrm()
	//o.QueryTable("conta").Filter("Usuario", dado.Email).RelatedSel().One(&conta)
	if err_user == models.ErrNoRows {

		status.Status = st_err_usuario_inexiste
		this.Data["json"] = status
		this.ServeJSON()
		return
	}

	if dado.Password != user.Password {
		fmt.Printf("%s nao bate com %s!\n", dado.Password, user.Password)
		status.Status = st_err_password_invalida
		this.Data["json"] = status
		this.ServeJSON()
		return
	}
	sess := this.StartSession()
	sess.Set("user", user)
	fmt.Println("logado com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()
}

func (this *LoginController) Logout() {
	//println("logout")
	//	sess := this.StartSession()
	this.DestroySession()
	//this.DelSession(sess)
	this.Redirect("/", 302)
}

