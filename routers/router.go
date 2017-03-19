package routers

import (
	"projects/makler/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/search/:id", &controllers.MainController{}, "get,post:Search")
	//beego.Router("/login", &controllers.MainController{}, "get,post:Login")
	beego.Router("/profile/:id([0-9]+)", &controllers.MainController{}, "get,post:Profile")
}
