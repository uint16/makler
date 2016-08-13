package main

import (
	"os"

	"github.com/astaxie/beego"
	_ "github.com/damagination/makler/routers"
)

func main() {
	port := os.Getenv("PORT")
	port = ":" + port
	beego.Run(port)
}
