package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/damagination/makler/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// var o orm.Ormer

type MainController struct {
	beego.Controller
}

// func init() {
// 	o = orm.NewOrm()
// }

func (c *MainController) activeContent(view string) {
	c.Layout = "layout.tpl"
	c.TplName = view + ".html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["NavBar"] = "navbar.html"
	c.LayoutSections["MainContent"] = c.TplName
}

func (c *MainController) Get() {
	c.activeContent("index")
}

func (c *MainController) Profile() {
	c.activeContent("index")
	profileId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	o.Using("default")
	flash := beego.NewFlash()

	flash.Notice("Worraaaaa")
	var prof models.Swdata
	o.Raw("SELECT * FROM swdata WHERE id = ?", profileId).QueryRow(&prof)

	temp := strings.Replace(prof.Image, "%3A", ":", 1)
	prof.Image = temp

	//	prof = models.Swdata{Id: 29}
	//	err := o.Read(&prof)

	// if err == orm.ErrNoRows {
	// 	fmt.Println("No Data")
	// } else {
	fmt.Println(prof)
	// }

	c.Data["p"] = prof

}

func (c *MainController) Login() {
	c.activeContent("login")
}
