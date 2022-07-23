package main

import (
	"log"
	"xgo/framework"
	"xgo/framework/middleware"
)

func main() {
	var x = framework.NewEngine()

	x.Use(middleware.Recovery())

	//framework.InitDb()
	//framework.InitRedis()

	x.GET("/", framework.IndexController)
	log.Fatalln(x.Run(":8000"))
}
