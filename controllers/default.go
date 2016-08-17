package controllers

import (
	"os"
	"strconv"

	"github.com/damagination/makler/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var list []models.Profile
var imagesURL string

type MainController struct {
	beego.Controller
}

func init() {

	imagesURL = os.Getenv("S3_BUCKET_URL")
	o := orm.NewOrm()
	o.Using("default")

	o.Raw("SELECT name, id FROM profile ORDER BY name ASC").QueryRows(&list)

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
	c.Data["l"] = list
	c.Data["z"] = imagesURL
}

func (c *MainController) Profile() {
	c.activeContent("profile")
	profileId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	o.Using("default")
	flash := beego.NewFlash()

	flash.Notice("Worraaaaa")
	var prof models.Profile
	var edu []models.EducationHistory
	var pol []models.PoliticalExperience
	var emp []models.EmploymentHistory

	o.Raw("SELECT * FROM profile WHERE id = ?", profileId).QueryRow(&prof)
	o.Raw("SELECT * FROM education_history WHERE mp_id = ?", profileId).QueryRows(&edu)
	o.Raw("SELECT * FROM political_experience WHERE mp_id = ?", profileId).QueryRows(&pol)
	o.Raw("SELECT * FROM employment_history WHERE mp_id = ?", profileId).QueryRows(&emp)

	c.Data["l"] = list
	c.Data["p"] = prof
	c.Data["e"] = edu
	c.Data["i"] = emp
	c.Data["j"] = pol
	c.Data["z"] = imagesURL

}

func (c *MainController) Login() {
	c.activeContent("login")
}
