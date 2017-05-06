package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//Profile member of parliament profile structure
type Profile struct {
	Name       string
	Group      string
	Area       string
	Phone      string
	Email      string
	Image      string
	MemberType string
	Address    string
	BirthDate  string
	Id         int
	Term       int
	Source     string
}

//EducationHistory member of parliament education history structure
type EducationHistory struct {
	MpId        int
	Institution string `db:"schoolName"`
	Level       string
	Award       string
	From        int
	To          int
	Id          int
}

//EmploymentHistory member of parliament employment history structure
type EmploymentHistory struct {
	MpId        int
	Institution string
	Position    string
	From        int
	To          int
	Id          int
}

//PoliticalExperience member of parliament political experience structure
type PoliticalExperience struct {
	MpId        int
	Institution string
	Position    string
	From        int
	To          int
	Id          int
}

func init() {
	db := os.Getenv("DB_NAME")
	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PG_PASSWORD")

	connection := fmt.Sprintf("user=%s  password=%s dbname=%s sslmode=disable", dbUser, dbPassword, db)
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "sqlite3", "database/scraperwiki0.sqlite")
	orm.RegisterDataBase("db2", "postgres", connection)
	orm.RegisterModel(new(Profile), new(PoliticalExperience), new(EducationHistory), new(EmploymentHistory))
}
