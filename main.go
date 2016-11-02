package main

import (
	"github.com/astaxie/beego"
	_ "github.com/swfsql/itabirapp/routers"
)

func add(x, y int) int {
	return x + y
}

func main() {
	beego.AddFuncMap("add", add)
	beego.Run()
}
