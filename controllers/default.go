package controllers

import (
	"os"
	"strconv"

	"projects/makler/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var list []models.Profile
var imagesURL string

type MainController struct {
	beego.Controller
}

func init() {

	imagesURL = os.Getenv("ASSETS_URL")
	orm := orm.NewOrm()
	orm.Using("db2")

	orm.Raw("SELECT * FROM profile ORDER BY name ASC").QueryRows(&list)

}

func (controller *MainController) activeContent(view string) {
	controller.Layout = "layout.tpl"
	controller.TplName = view + ".html"
	controller.LayoutSections = make(map[string]string)
	controller.LayoutSections["NavBar"] = "navbar.html"
	controller.LayoutSections["MainContent"] = controller.TplName
}

func (controller *MainController) Get() {
	controller.activeContent("index")
	controller.Data["list"] = list
	controller.Data["assets"] = imagesURL
}

func (controller *MainController) Profile() {
	controller.activeContent("profile")
	profileId, _ := strconv.Atoi(controller.Ctx.Input.Param(":id"))

	orm := orm.NewOrm()
	orm.Using("db2")

	var prof models.Profile
	var edu []models.EducationHistory
	var pol []models.PoliticalExperience
	var emp []models.EmploymentHistory

	orm.Raw("SELECT * FROM profile WHERE id = ?", profileId).QueryRow(&prof)
	orm.Raw("SELECT * FROM education_history WHERE mp_id = ?", profileId).QueryRows(&edu)
	orm.Raw("SELECT * FROM political_experience WHERE mp_id = ?", profileId).QueryRows(&pol)
	orm.Raw("SELECT * FROM employment_history WHERE mp_id = ?", profileId).QueryRows(&emp)

	controller.Data["membersList"] = list
	controller.Data["memberProfile"] = prof
	controller.Data["memberEducation"] = edu
	controller.Data["memberEmployment"] = emp
	controller.Data["memberExperience"] = pol
	controller.Data["assets"] = imagesURL

}

func (c *MainController) Login() {
	c.activeContent("login")
}
