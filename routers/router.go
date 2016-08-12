package routers

import (
	"github.com/astaxie/beego"
	"github.com/damagination/makler/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get,post:Login")
}
