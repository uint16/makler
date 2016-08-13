package routers

import (
	"github.com/astaxie/beego"
	"github.com/damagination/makler/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get,post:Login")
	beego.Router("/profile/:id([0-9]+)", &controllers.MainController{}, "get,post:Profile")
}
