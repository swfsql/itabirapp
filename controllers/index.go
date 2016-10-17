package controllers

import (
	_ "fmt"
	_ "strconv"
)



type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.TplName = "index.html"
	this.Data["HeadTitle"] = "itabirApp"
	this.Data["HeadStyles"] = []string{"index.css"}
    this.Data["HeadScripts"] = []string{}
	this.Render()
}