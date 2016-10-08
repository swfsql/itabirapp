package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Layout = "layout.html"
}
