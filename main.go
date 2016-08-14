package main

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/damagination/makler/routers"
)

func main() {

	err := orm.RunSyncdb("db2", false, false)
	if err != nil {

	}
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	port = ":" + port
	beego.Run(port)
}
