package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"encoding/json"
	"github.com/swfsql/itabirapp/models"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"

	"io"
	"mime/multipart"
	"os"
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
		Title    string
		Subtitle string
		Text     string
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

	postId := post.Id

	// salva imagem de acordo com o ID
	sess := this.StartSession()
	file, hasFile := sess.Get("postImage").(multipart.File)
	if hasFile {
		defer file.Close()
		defer sess.Delete("postImage")
		out, _ := os.Create("static/images/post/" + strconv.Itoa(postId) + ".jpg")
		defer out.Close()
		io.Copy(out, file)
	}

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
		Title    string
		Subtitle string
		Text     string
		Tags     []string
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
		User:     &user,
		Title:    dado.Title,
		Subtitle: dado.Subtitle,
		Text:     dado.Text,
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

	postId_s := strconv.Itoa(int(postId))

	// salva imagem de acordo com o ID
	file, hasFile := sess.Get("postImage").(multipart.File)
	if hasFile {
		defer file.Close()
		defer sess.Delete("postImage")
		out, _ := os.Create("static/images/post/" + postId_s + ".jpg")
		defer out.Close()
		io.Copy(out, file)
	}

	status := struct {
		Status string
		PostId string
	}{"", ""}

	status.Status = st_ok
	status.PostId = postId_s

	this.Data["json"] = status
	this.ServeJSON()
}

func (this *PostController) GetNew() {
	fmt.Println("hueee hue br")

	sess := this.StartSession()
	//defer sess.SessionRelease()

	fmt.Println("macacoide")
	_, loggedIn := sess.Get("user").(models.User)
	if !loggedIn {
		defer this.DestroySession()
		this.Redirect("/usuario/criar", 302)
		return
	}

	this.TplName = "post/new.html"
	this.Data["HeadTitle"] = "Criar novo anúncio"
	this.Data["HeadStyles"] = []string{"simplemde.min.css"}
	this.Data["HeadScripts"] = []string{"simplemde.min.js", "post/new.js"}
	this.Render()
}

func (this *PostController) GetSearch() {
	fmt.Println("hueee hue br")
	tags2 := this.Ctx.Input.Param(":search")
	tags_str := strings.Split(tags2, ",")

	this.Data["TargetExists"] = false
	if len(tags_str) == 1 {
		target, err_target := models.GetUserByNameIdTag(tags_str[0])
		if err_target == nil {
			this.Data["TargetExists"] = true
			target.Password = ""
			this.Data["Target"] = target
			// addr
			addr := target.Addr_Street + ", "
			if target.Addr_Number != "" {
				addr += target.Addr_Number + ", "
			}
			if target.Addr_Complement != "" {
				addr += "ap. " + target.Addr_Complement + ", "
			}
			if target.Addr_Neighborhood != "" {
				addr += target.Addr_Neighborhood + ", "
			}
			if target.Addr_City != "" {
				addr += target.Addr_City + ", "
			}
			addr += "Minas Gerais, Brasil"
			this.Data["Address"] = addr
		}
	} else {
		// "a,b,c,d" => "a,b,c,d,*,*,*"
		count_names := 0
		count_operands := 0
		for _, s := range tags_str {
			if s != "*" && s != "+" {
				count_names++
			} else {
				count_operands++
			}
		}
		count_diff := count_names - count_operands
		for i := 0; i < count_diff-1; i++ {
			tags_str = append(tags_str, "*")
		}
	}

	_, posts, tags_all, _ := models.GetPostsByTags(tags_str)

	var tags []models.Tag
	count := 0
	for _, ta := range tags_all {
		for i, ts := range tags_str {
			if ta.Name == ts {
				break
			} else if i == len(tags_str)-1 {
				tags = append(tags, ta) // only consider brand new tags
				count++
			}
		}
		if count == 4 { // only consider the first 4
			break
		}
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~")
	for _, p := range posts {
		fmt.Println(p.Title)
	}

	type tag_url struct {
		Name string
		Url  string
	}
	var tags_url []tag_url
	for _, t := range tags {
		t_u := tag_url{Name: t.Name, Url: this.Ctx.Input.URL()}
		tags_url = append(tags_url, t_u)
	}

	this.Data["Posts"] = posts

	this.Data["Tags_URL"] = tags_url

	this.TplName = "post/list.html"
	this.Data["HeadTitle"] = "Lista de anúncios"
	this.Data["HeadStyles"] = []string{"post/list.css"}
	this.Data["HeadScripts"] = []string{"post/list.js"}
	this.Render()

}

func (this *PostController) PostPostImage() {

	file, _, err := this.GetFile("datafile")

	if file != nil {
		sess := this.StartSession()
		sess.Set("postImage", file)
		//sess.Set("userImageHeader", header)
	} else {
		fmt.Println(err)

	}

	status := struct{ Status string }{""}
	status.Status = st_ok
	this.Data["json"] = status
	this.ServeJSON()

}
