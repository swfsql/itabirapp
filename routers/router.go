package routers

import (
	"github.com/astaxie/beego"
	"github.com/swfsql/itabirapp/controllers"
)

func init() {
    beego.Router("/", &controllers.IndexController{})

    beego.Router("/anuncio/:id", &controllers.PostController{}, "get:GetPost")
    //beego.Router("/anuncio/criar", &controllers.PostController{}, "get:blabla;post:blabla")
    beego.Router("/anuncio/:id/editar", &controllers.PostController{}, "get:GetEdit;post:PostEdit")
    //beego.Router("/anuncio/:id/remover", &controllers.PostController{})

    //beego.Router("/busca/*", &controllers.SearchController{})
    //beego.Router("/buscajs/*", &controllers.SearchController{})

	beego.Router("/usuario/criar", &controllers.UserController{}, "get:GetNew;post:PostNew")
	beego.Router("/usuario/listar", &controllers.UserController{}, "get:GetList")
	beego.Router("/usuario/:id/editar", &controllers.UserController{}, "get:GetEdit")
	beego.Router("/usuario/:id/editar/usuario", &controllers.UserController{}, "post:PostEditUser")
	beego.Router("/usuario/:id/editar/instituicao", &controllers.UserController{}, "post:PostEditInstitution")
	beego.Router("/usuario/:id/editar/instituicao/autorizacao", &controllers.UserController{}, "get:ToggleAuthorization")
	beego.Router("/usuario/:id/editar/endereco", &controllers.UserController{}, "post:PostEditAddress")
	beego.Router("/usuario/:id/remover", &controllers.UserController{}, "get:GetDelete")

	//beego.Router("/documento/:id", &controllers.DocumentController{}, "get:Login;post:Post")
	//beego.Router("/documento/adicionar", &controllers.DocumentController{}, "get:Login;post:Post")
	//beego.Router("/documento/:id/remover", &controllers.DocumentController{}, "get:Login;post:Post")

	beego.Router("/login", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}
