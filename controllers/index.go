package controllers

import (
	"github.com/YoungEugene/freePlayer/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	user := new(models.User)
	user, this.Data["IsLogin"] = CheckUserLogin(this.Ctx)
	if user != nil {
		this.Data["Nickname"] = user.Nickname
	}
	this.Data["IsHome"] = true
	this.Data["DefaultVideoUrl"], _ = models.GetConfigByName("DefaultVideoUrl")
	this.TplName = "index.html"
}
