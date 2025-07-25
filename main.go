package main

import (
	"os"

	_ "github.com/uint16/makler/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
		web.BConfig.RunMode = "dev"
	} else {
		web.BConfig.RunMode = "prod"
	}
	port = ":" + port
	web.Run(port)
}
