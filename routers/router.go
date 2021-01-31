package routers

import (
	"todolist/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/item/?:id", &controllers.ItemController{})
	beego.Router("/history", &controllers.HistoryController{})
	beego.SetStaticPath("/node_modules", "node_modules")
}
