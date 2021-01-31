package main

import (
	_ "todolist/models"
	_ "todolist/routers"
	_ "todolist/utils"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
