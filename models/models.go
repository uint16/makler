package main

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id               int64
	UserName         string `orm:"unique"`
	Email            string `orm:"unique"`
	RegistrationKey  string
	RegistrationDate time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}
