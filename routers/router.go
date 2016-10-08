package routers

import (
	"github.com/swfsql/itabirapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})

    beego.Router("/anuncio/:id", &controllers.AnuncioController{})
    beego.Router("/anuncio/criar", &controllers.AnuncioController{}, "get:blabla;post:blabla")
    beego.Router("/anuncio/:id/editar", &controllers.AnuncioController{})
    beego.Router("/anuncio/:id/remover", &controllers.AnuncioController{})

    beego.Router("/busca/*", &controllers.BuscaController{})
    beego.Router("/buscajs/*", &controllers.BuscaController{})

	beego.Router("/user/criar", &controllers.UserController{}, "get:Login;post:Post")
	beego.Router("/user/listar", &controllers.UserController{}, "get:Login;post:Post")
	beego.Router("/user/?:id", &controllers.UserController{}, "get:Login;post:Post")
	beego.Router("/user/:id/editar", &controllers.UserController{}, "get:Login;post:Post")
	beego.Router("/user/:id/remover", &controllers.UserController{}, "get:Login;post:Post")

	beego.Router("/documento/:id", &controllers.DocumentController{}, "get:Login;post:Post")
	beego.Router("/documento/adicionar", &controllers.DocumentController{}, "get:Login;post:Post")
	beego.Router("/documento/:id/remover", &controllers.DocumentController{}, "get:Login;post:Post")

	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:Post")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}
