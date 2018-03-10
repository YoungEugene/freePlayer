package routers

import (
	"github.com/YoungEugene/freePlayer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.AdminController{})

}
