package main

import (
	"log"
	"xgo/framework"
)

func main() {
	var x = framework.NewEngine()
	framework.InitDb()
	framework.InitRedis()

	x.GET("/", framework.IndexController)
	log.Fatalln(x.Run(":8000"))
}
