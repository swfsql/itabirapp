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
	fmt.Println("--------------")

	id64, _ := strconv.ParseUint(this.Ctx.Input.Param(":id"), 10, 32)
	id := int(id64)
	post, err_post := models.GetPostById(id)

	if err_post == models.ErrNoRows {
		this.Redirect("/", 302)
		return 
	}


	sess := this.StartSession()
	//defer sess.SessionRelease()

	fmt.Println("macacoide")
	user, loggedIn := sess.Get("user").(models.User)
	this.Data["IsOwner"] = false
	if loggedIn && post.User.Id == user.Id {
		this.Data["IsOwner"] = true
	}
	fmt.Println("macacoide")
	mrk_title := "#### " + post.Title + ": "
	mrk_subtitle := " *" + post.Subtitle + "*"
	mrk_text := "---\n\n" + post.Text
	//mrk_author := "" + post.User.Name
	//mrk_NameIdTag := "" + post.User.NameIdTag
	//mrk_Institution_Tag := "" + post.User.Institution_Tag

	fmt.Println("macacoide")
	mrk := mrk_title + mrk_subtitle + "\n\n" + mrk_text

	unsafe := blackfriday.MarkdownCommon([]byte(mrk))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	this.Data["PostHTML"] = template.HTML(html)
	this.Data["Post"] = post

	fmt.Println("macacoide")
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
	this.Data["HeadStyles"] = []string{"simplemde.min.css"}
    this.Data["HeadScripts"] = []string{"simplemde.min.js", "post/edit.js"}
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

func (this *PostController) GetDelete() {
	var post models.Post
	var allow bool
	if post, allow = editPostAllowed(this); !allow {
		return
	}

	post.Delete()

	status := struct{ Status string }{""}

	fmt.Println("removido com sucesso")

	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()
}

func (this *PostController) PostNew() {
	dado := struct {
		Title string
		Subtitle string
		Text string
		Tags []string
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dado)
	if err != nil {
		fmt.Println(err)
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("JSON invalido"))
		return
	}
	fmt.Println(dado)

	sess := this.StartSession()
	user, loggedIn := sess.Get("user").(models.User)
	// not logged in
	if !loggedIn {
		defer this.DestroySession()
		this.Redirect("/", 302)
		return
	}

	if user.User_Type != "poster" && user.IsAuthorized != true {
		fmt.Println("macaquice")
		this.Redirect("/", 302)
		return 
	}

	post := models.Post{
		User: &user,
		Title: dado.Title,
		Subtitle: dado.Subtitle,
		Text: dado.Text,
	}

	postId, err_post := post.New()
	if err_post != nil {
		fmt.Println("macaquice")
		this.Redirect("/", 302)
	}

	var tags []string
	tags = append(tags, user.NameIdTag)
	tags = append(tags, user.Institution_Tag)
	for _, t := range dado.Tags {
		if t != "" {
			tags = append(tags, t)
		}
	}
	models.AppendTagsForPost(&post, tags)

	fmt.Println("criado com sucesso")

	//this.Redirect("/anuncio/" + strconv.Itoa(int(postId)), 302)

	status := struct{ 
		Status string 
		PostId string 
	}{"", ""}

	status.Status = st_ok
	status.PostId = strconv.Itoa(int(postId))

	this.Data["json"] = status
	this.ServeJSON()
}

func (this *PostController) GetNew() {
	fmt.Println("hueee hue br")


	this.TplName = "post/new.html"
	this.Data["HeadTitle"] = "Criar novo anúncio"
	this.Data["HeadStyles"] = []string{"simplemde.min.css"}
    this.Data["HeadScripts"] = []string{"simplemde.min.js", "post/new.js"}
	this.Render()
}

func (this *PostController) GetSearch() {
	fmt.Println("hueee hue br")
	tags2 := this.Ctx.Input.Param(":search")

	_, posts, _ := models.GetPostsByTags(tags2)
	fmt.Println("~~~~~~~~~~~~~~~~~~")
	for _, p := range posts {
		fmt.Println(p.Title)
	}

	this.Data["Posts"] = posts

	this.TplName = "post/list.html"
	this.Data["HeadTitle"] = "Lista de anúncios"
	this.Data["HeadStyles"] = []string{"post/list.css"}
    this.Data["HeadScripts"] = []string{"post/list.js"}
	this.Render()

}
