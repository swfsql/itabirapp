package controllers

import (
	_ "fmt"
	 "strconv"

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