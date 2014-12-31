package main

import (
	_ "Hongic/models"
	_ "Hongic/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
