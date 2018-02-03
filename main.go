package main

import (
	"os"

	_ "github.com/uint16/makler/routers"

	"github.com/astaxie/beego"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	port = ":" + port
	beego.Run(port)
}
