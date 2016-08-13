package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/damagination/makler/routers"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.HttpPort = port
	}
	beego.Run()
}
