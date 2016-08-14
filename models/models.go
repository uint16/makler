package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
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

type EducationHistory struct {
	MpId       int
	SchoolName string `orm:"column(schoolName)"`
	Level      string
	Award      string
	From       int
	To         int
	Id         int
}

type EmploymentHistory struct {
	MpId        int
	Institution string
	Position    string
	From        int
	To          int
	Id          int
}

type PoliticalExperience struct {
	MpId        int
	Institution string
	Position    string
	From        int
	To          int
	Id          int
}

type User struct {
	Id       int
	UserName string `orm:"unique"`
	Email    string
	Password string
	RegKey   string
	RegDate  time.Time `orm:"auto_now_add;type(datetime)"`
	//	Comments []*Comment `orm:"reverse(many)"`
}

type Comment struct {
	Id          int
	CommenterId int // `orm:"rel(fk);null;on_delete(set_null)"`
	ProfileId   int
	Content     string
	CommentTime time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "sqlite3", "database/scraperwiki.sqlite")
	orm.RegisterDataBase("db2", "postgres", "user=damas dbname=data sslmode=disable")
	orm.RegisterModel(new(Swdata), new(PoliticalExperience), new(EducationHistory), new(EmploymentHistory), new(User), new(Comment))
}
