package main

import (
	"os"

	_ "github.com/uint16/makler/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	web.Run(":" + port)
}
