package routers

import (
	"github.com/uint16/makler/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/profile/:id([0-9]+)", &controllers.MainController{}, "get,post:Profile")
}
