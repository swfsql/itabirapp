package controllers

import (
	"fmt"
	 "strconv"

	"encoding/json"
	"github.com/swfsql/itabirapp/models"
)


type PostController struct {
	BaseController
}

func (this *PostController) GetPost() {

	id64, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 32)
	id := int(id64)
	post, _ := models.GetPostById(id)
	this.Data["Post"]= post

	this.TplName = "post/post.html"
	this.Data["HeadTitle"] = "Visualizar Anúncio"
	this.Data["HeadStyles"] = []string{}
    this.Data["HeadScripts"] = []string{}
	this.Render()

}

func editPostAllowed(this *PostController) (models.Post, bool) {
	id64, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 32)
	id := int(id64)
	post, _ := models.GetPostById(id)



	sess := this.StartSession()
	//defer sess.SessionRelease()

	user, loggedIn := sess.Get("user").(models.User)
	// not logged in
	if !loggedIn {
		defer this.DestroySession()
		this.Redirect("/", 302)
		return post, false
	}

	if post.User.Id != user.Id {

		this.Redirect("/", 302)
		return post, false
	}

	return post, false

}

func (this *PostController) GetEdit() {
	var post models.Post
	var allow bool
	if post, allow = editPostAllowed(this); !allow {
		return
	}
 
	this.Data["Post"] = post

	this.TplName = "post/edit.html"
	this.Data["HeadTitle"] = "Configurações do post"
	this.Data["HeadStyles"] = []string{}
    this.Data["HeadScripts"] = []string{"post/edit.js"}
	this.Render()
}

func (this *PostController) PostEdit() {
	var post models.Post
	var allow bool
	if post, allow = editPostAllowed(this); !allow {
		return
	}

	dado := struct {
	  	Title string
	  	Subtitle string
	  	Text string
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

	post.Title = dado.Title
	post.Subtitle = dado.Subtitle
	post.Text = dado.Text


	post.Update()

	fmt.Println("editado com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()
}
