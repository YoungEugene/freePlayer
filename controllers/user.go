package controllers

import (
	"github.com/YoungEugene/freePlayer/models"
	"github.com/YoungEugene/freePlayer/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type UserController struct {
	beego.Controller
}

//============================
//============================
//============================

func (this *UserController) OpenLoginPage() {
	this.TplName = "user_login.html"
}

func (this *UserController) OpenRegisterPage() {
	this.TplName = "user_register.html"
}

func (this *UserController) UserLogin() { //登录
	name := this.Input().Get("user_name")
	pwd := this.Input().Get("user_pwd")
	if name == "" || pwd == "" {
		this.TplName = "user_login.html"
		this.Data["msg"] = "账号或密码不能为空！"
		return
	}
	//登录操作 orm表查
	user, err := models.GetUser(name, pwd)
	if err != nil {
		this.Data["msg"] = "登录失败，请检查账号密码是否输入正确。"
		this.TplName = "user_login.html"
		return
	}
	this.SetSession("user", user)
	this.Redirect("/", 302)
}

func (this *UserController) UserLogout() { //退出
	this.DelSession("user")
	this.Redirect("/", 302)
}

func (this *UserController) UserRegister() { //登录
	account := this.Input().Get("account")
	nickname := this.Input().Get("nickname")
	email := this.Input().Get("email")
	phone := this.Input().Get("phone")
	password := this.Input().Get("password")
	rePassword := this.Input().Get("rePassword")

	if !utils.AllNotEmpty(account, nickname, email, phone, password, rePassword) || password != rePassword {
		this.TplName = "user_register.html"
		this.Data["account"] = account
		this.Data["nickname"] = nickname
		this.Data["email"] = email
		this.Data["phone"] = phone
		this.Data["msg"] = "资料不全或两次输入密码不同！"
		return
	}
	_, err := models.AddUser(account, password, nickname, email, phone)
	if err != nil {
		beego.Error(err)
		this.Data["account"] = account
		this.Data["nickname"] = nickname
		this.Data["email"] = email
		this.Data["phone"] = phone
		this.Data["msg"] = "注册失败，账号已经被注册，请更换别的试试。"
		this.TplName = "user_register.html"
		return
	}
	this.Data["msg"] = "注册成功，请登录。"
	this.TplName = "user_login.html"
}

//===========================
//===========================
//===========================
func CheckUserLogin(ctx *context.Context) (*models.User, bool) {
	i := ctx.Input.CruSession.Get("user")
	if i == nil {
		return nil, false
	}
	return i.(*models.User), true
}
