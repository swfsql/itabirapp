package controllers

import (
	_ "fmt"
	_ "github.com/swfsql/itabirapp/models"
	_ "strconv"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {

}
