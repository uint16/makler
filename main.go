package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/damagination/makler/routers"
)

func main() {
	orm.Debug = true
	beego.Run()
}
