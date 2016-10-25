package main

import (
	"github.com/astaxie/beego"
	_ "github.com/swfsql/itabirapp/routers"
)

func main() {
	beego.Run()
}
