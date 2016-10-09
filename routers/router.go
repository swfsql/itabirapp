package routers

import (
	"github.com/astaxie/beego"
	"github.com/swfsql/itabirapp/controllers"
)

func init() {
    beego.Router("/", &controllers.IndexController{})

    //beego.Router("/anuncio/:id", &controllers.PostController{})
    //beego.Router("/anuncio/criar", &controllers.PostController{}, "get:blabla;post:blabla")
    //beego.Router("/anuncio/:id/editar", &controllers.PostController{})
    //beego.Router("/anuncio/:id/remover", &controllers.PostController{})

    //beego.Router("/busca/*", &controllers.SearchController{})
    //beego.Router("/buscajs/*", &controllers.SearchController{})

	//beego.Router("/user/criar", &controllers.UserController{}, "get:Login;post:Post")
	//beego.Router("/user/listar", &controllers.UserController{}, "get:Login;post:Post")
	//beego.Router("/user/?:id", &controllers.UserController{}, "get:Login;post:Post")
	beego.Router("/user/:id/editar", &controllers.UserController{}, "get:GetEdit")
	//beego.Router("/user/:id/remover", &controllers.UserController{}, "get:Login;post:Post")

	//beego.Router("/documento/:id", &controllers.DocumentController{}, "get:Login;post:Post")
	//beego.Router("/documento/adicionar", &controllers.DocumentController{}, "get:Login;post:Post")
	//beego.Router("/documento/:id/remover", &controllers.DocumentController{}, "get:Login;post:Post")

	beego.Router("/login", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}
