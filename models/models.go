package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

// type User struct {
// 	Id               int64
// 	UserName         string `orm:"unique"`
// 	Email            string `orm:"unique"`
// 	RegistrationKey  string
// 	RegistrationDate time.Time `orm:"auto_now_add;type(datetime)"`
// }

type Swdata struct {
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

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/scraperwiki.sqlite")
	orm.RegisterModel(new(Swdata))
}
