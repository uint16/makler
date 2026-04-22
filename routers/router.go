package routers

import (
	"github.com/uint16/makler/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/profile/:id([0-9]+)", &controllers.MainController{}, "get,post:Profile")
}
