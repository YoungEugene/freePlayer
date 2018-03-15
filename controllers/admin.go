package controllers

import (
	"strconv"

	"github.com/YoungEugene/freePlayer/models"
	"github.com/YoungEugene/freePlayer/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Manager() { //管理页
	if admin, ok := CheckLogin(this.Ctx); ok {
		this.Data["IsSuper"] = bool(admin.Type > 0)
		this.Data["Nickname"] = admin.Nickname
		this.TplName = "manager.html"
	} else {
		this.Data["msg"] = "未登录或登录超时，请重新登录"
		this.TplName = "admin_login.html"
	}
}

func (this *AdminController) ManagerHome() { //首页
	this.TplName = "manager_home.html"

}

func (this *AdminController) OpenChgPage() {
	this.TplName = "manager_changepwd.html"
}

func (this *AdminController) OpenUserListPage() {
	if _, ok := CheckLogin(this.Ctx); !ok {
		this.Redirect("/admin/login", 302)
		return
	}
	pageNoStr := this.Input().Get("pageNo")
	pageNo, err := strconv.Atoi(pageNoStr)
	if err != nil {
		pageNo = 1
	}
	userList, _ := models.GetUserList(pageNo)
	this.Data["UserList"] = userList
	this.TplName = "manager_users.html"
}

//============================
//============================
//============================
func (this *AdminController) Login() { //登录
	name := this.Input().Get("admin_name")
	pwd := this.Input().Get("admin_pwd")
	if name == "" && pwd == "" {
		this.TplName = "admin_login.html"
		return
	}
	//登录操作 orm表查
	admin, err := models.GetAdmin(name, pwd)
	if err != nil {
		this.Data["msg"] = "登录失败，请检查账号密码是否输入正确。"
		this.TplName = "admin_login.html"
		return
	}
	this.SetSession("admin", admin)
	this.Redirect("/admin/manager", 302)
}

func (this *AdminController) Logout() { //退出
	this.DelSession("admin")
	this.Data["msg"] = "已安全退出"
	this.TplName = "admin_login.html"
}

func (this *AdminController) ChgDefaultVedio() { //查看默认视频页和修改视频地址
	if _, ok := CheckLogin(this.Ctx); !ok {
		this.Data["msg"] = "未登录或登录超时，请重新登录"
		this.TplName = "admin_login.html"
		return
	}
	defaultUrl := this.Input().Get("DefaultUrl")
	if defaultUrl == "" {
		this.Data["DefaultUrl"], _ = models.GetConfigByName("DefaultVideoUrl")
	} else { //修改默认视频地址
		this.Data["msg"] = "修改成功！"
		this.Data["DefaultUrl"] = defaultUrl
		if _, err := models.SetConfig("DefaultVideoUrl", defaultUrl); err != nil {
			this.Data["msg"] = "修改失败！"
			this.Data["DefaultUrl"], _ = models.GetConfigByName("DefaultVideoUrl")
		}
	}
	this.TplName = "manager_defaultvideo.html"

}

func (this *AdminController) ChgPwd() {
	admin, ok := new(models.Admin), false
	if admin, ok = CheckLogin(this.Ctx); !ok {
		this.Data["msg"] = "未登录或登录超时，请重新登录"
		this.TplName = "admin_login.html"
	}
	nowPwd := this.Input().Get("nowPwd")
	newPwd := this.Input().Get("newPwd")
	reNewPwd := this.Input().Get("reNewPwd")

	if !utils.AllNotEmpty(nowPwd, newPwd, reNewPwd) || newPwd != reNewPwd {
		this.Data["msg"] = "修改失败!资料不全或两次密码不同"
	} else if nowPwd != admin.Password {
		this.Data["msg"] = "修改失败!当前的密码错误！"
	} else {
		this.Data["msg"] = "修改成功"
		admin.Password = newPwd
		num, err := models.UpdateAdmin(admin, "Password")
		if err != nil {
			beego.Error(err, num)
			this.Data["msg"] = "修改失败"
		}
	}
	this.TplName = "manager_changepwd.html"

}

func (this *AdminController) DelUser() {
	if _, ok := CheckLogin(this.Ctx); !ok {
		this.Redirect("/admin/login", 302)
		return
	}
	idStr := this.Input().Get("id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		err = models.DelUserById(int64(id))
		if err != nil {
			beego.Error(err)
		}
	}
	this.Redirect("/admin/openuserlistpage", 302)

}

//===========================
//===========================
//===========================
func CheckLogin(ctx *context.Context) (*models.Admin, bool) {
	i := ctx.Input.CruSession.Get("admin")
	if i == nil {
		return nil, false
	}
	return i.(*models.Admin), true
}
