package controllers

import (
	"os"
	"strconv"

	"github.com/uint16/makler/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var profiles []models.Profile
var imagesURL string

//MainController abstraction of beego controller
type MainController struct {
	beego.Controller
}

func init() {

	imagesURL = os.Getenv("ASSETS_URL")
	orm := orm.NewOrm()
	orm.Using("default")

	orm.Raw("SELECT * FROM profile ORDER BY name ASC").QueryRows(&profiles)

}

func (controller *MainController) activeContent(view string) {
	controller.Layout = "layout.tpl"
	controller.TplName = view + ".html"
	controller.LayoutSections = make(map[string]string)
	controller.LayoutSections["NavBar"] = "navbar.html"
	controller.LayoutSections["MainContent"] = controller.TplName
}

//Get all users
func (controller *MainController) Get() {
	controller.activeContent("index")
	controller.Data["list"] = profiles
	controller.Data["assets"] = imagesURL
}

//Profile get user profile by id
func (controller *MainController) Profile() {
	controller.activeContent("profile")
	profileID, _ := strconv.Atoi(controller.Ctx.Input.Param(":id"))

	orm := orm.NewOrm()
	orm.Using("default")

	var profile models.Profile
	var educationHistory []models.EducationHistory
	var politicalExperienceHistory []models.PoliticalExperience
	var employmentHistory []models.EmploymentHistory

	orm.Raw("SELECT * FROM profile WHERE id = ?", profileID).QueryRow(&profile)
	orm.Raw("SELECT * FROM education_history WHERE mp_id = ?", profileID).QueryRows(&educationHistory)
	orm.Raw("SELECT * FROM political_experience WHERE mp_id = ?", profileID).QueryRows(&politicalExperienceHistory)
	orm.Raw("SELECT * FROM employment_history WHERE mp_id = ?", profileID).QueryRows(&employmentHistory)

	controller.Data["membersList"] = profiles
	controller.Data["memberProfile"] = profile
	controller.Data["memberEducation"] = educationHistory
	controller.Data["memberEmployment"] = employmentHistory
	controller.Data["memberExperience"] = politicalExperienceHistory
	controller.Data["assets"] = imagesURL

}

//Login login
func (controller *MainController) Login() {
	controller.activeContent("login")
}
