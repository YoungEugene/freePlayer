package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsLogin"] = false
	c.Data["Username"] = "Eugene杨"
	c.TplName = "index.html"
}
