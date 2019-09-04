package main

import (
	"github.com/astaxie/beego"
	_ "wel/models"
	_ "wel/routers"
)

func main() {
	beego.Run()
}

