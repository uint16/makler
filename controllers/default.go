package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) activeContent(view string) {
	c.Layout = "layout.tpl"
	c.TplName = view + ".html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["NavBar"] = "navbar.html"
	c.LayoutSections["MainContent"] = c.TplName
}

func (c *MainController) Get() {
	c.activeContent("index")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

}

func (c *MainController) Login() {
	c.activeContent("login")
}
