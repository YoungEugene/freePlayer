package main

import (
	_ "github.com/YoungEugene/freePlayer/routers"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/common", "common")
}

func main() {
	beego.Debug(true)
	beego.Run()
}
