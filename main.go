package main

import (
	"github.com/astaxie/beego"
	_ "github.com/damagination/makler/routers"
)

func main() {
	beego.Run()
}
