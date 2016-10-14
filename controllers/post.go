package controllers

import (
	"fmt"
	 "strconv"

	"encoding/json"
	"github.com/swfsql/itabirapp/models"

	"github.com/microcosm-cc/bluemonday"
    "github.com/russross/blackfriday"
    "html/template"
)


type PostController struct {
	BaseController
}

func (this *PostController) GetPost() {

	id64, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 32)
	id := int(id64)
	post, _ := models.GetPostById(id)

	sess := this.StartSession()
	//defer sess.SessionRelease()

	user, loggedIn := sess.Get("user").(models.User)
	this.Data["IsOwner"] = false
	if loggedIn && post.User.Id == user.Id {
		this.Data["IsOwner"] = true
	}

	mrk_title := "# " + post.Title
	mrk_subtitle := "## " + post.Subtitle
	mrk_text := post.Text
	//mrk_author := "" + post.User.Name
	//mrk_NameIdTag := "" + post.User.NameIdTag
	//mrk_Institution_Tag := "" + post.User.Institution_Tag

	mrk := mrk_title + "\n" + mrk_subtitle + "\n" + mrk_text

	unsafe := blackfriday.MarkdownCommon([]byte(mrk))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	this.Data["PostHTML"] = template.HTML(html)

	this.TplName = "post/post.html"
	this.Data["HeadTitle"] = "Visualizar Anúncio"
	this.Data["HeadStyles"] = []string{}
    this.Data["HeadScripts"] = []string{}
	this.Render()

}

func editPostAllowed(this *PostController) (models.Post, bool) {
	fmt.Println("")
	fmt.Println("edit allowed?")
	fmt.Println("")
	id64, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 32)
	id := int(id64)
	post, _ := models.GetPostById(id)

	sess := this.StartSession()
	//defer sess.SessionRelease()

	user, loggedIn := sess.Get("user").(models.User)
	// not logged in
	if !loggedIn {
		fmt.Println("HUAHUAHAHA")
		defer this.DestroySession()
		this.Redirect("/", 302)
		return post, false
	}

	if post.User.Id != user.Id {
		fmt.Println("macaquice")
		this.Redirect("/", 302)
		return post, false
	}

	return post, true
}

func (this *PostController) GetEdit() {
	var post models.Post
	var allow bool
	if post, allow = editPostAllowed(this); !allow {
		return
	}
 
	this.Data["Post"] = post

	fmt.Println("")
	fmt.Println("vai seu macaco")
	fmt.Println("")
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

