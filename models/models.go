package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id       int64
	Account  string `orm:"column(account);size(32)"`
	Password string `orm:size(64)`
	Nickname string `orm:"size(32)"`
	Phone    string `orm:size(16)`
	Type     int    `orm:"default(0)"`
}

type User struct {
	Id       int64
	Account  string    `orm:"size(32)"`
	Password string    `orm:"size(64)"`
	Nickname string    `orm:"size(32)"`
	Email    string    `orm:"size(64)"`
	Phone    string    `orm:"size(16)"`
	Type     int       `orm:"default(0)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

type Config struct {
	Name  string `orm:"pk"`
	Value string
}

func init() {
	orm.RegisterModel(new(Admin), new(User), new(Config))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/free_player?charset=utf8", 10)
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
	//全局ormer
	o = orm.NewOrm()
	o.Using("default")
}

var o orm.Ormer

/*
	通过账号密码获取Admin
*/
func GetAdmin(account, pwd string) (*Admin, error) {
	admin := new(Admin)
	err := o.QueryTable("admin").Filter("account", account).
		Filter("password", pwd).One(admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

/*
	根据ID查admin
*/
func GetAdminById(adminId int64) *Admin {
	admin := &Admin{Id: adminId}
	o.Read(admin)
	return admin
}

/*
	改admin信息
*/
func UpdateAdmin(admin *Admin, cols ...string) (int64, error) {
	return o.Update(admin, cols...)
}

/*
	按名称查config表
*/
func GetConfigByName(cname string) (string, error) {
	c := &Config{Name: cname}
	err := o.Read(c)
	return c.Value, err

}

/*
	按名称设置或增加config表行
*/
func SetConfig(cname, cvalue string) (int64, error) {
	c := &Config{Name: cname, Value: cvalue}
	return o.Update(c)
}

/*
	通过账号密码获取Admin
*/
func GetUser(account, pwd string) (*User, error) {
	user := new(User)
	err := o.QueryTable("user").Filter("account", account).
		Filter("password", pwd).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

/*
	注册User
*/
func AddUser(account, pwd, nickname, email, phone string) (int64, error) {
	user := &User{Account: account, Password: pwd, Nickname: nickname, Email: email, Phone: phone}
	if o.QueryTable("user").Filter("account", account).Exist() {
		return 0, nil
	}
	return o.Insert(user)
}

/*
	获取用户列表
*/
func GetUserList(pageNo int) ([]*User, error) {
	var userList []*User
	_, err := o.QueryTable("user").OrderBy("-id").Limit(10, (pageNo-1)*10).All(&userList)
	return userList, err
}

/*
	删除用户
*/
func DelUserById(userId int64) error {
	_, err := o.Delete(&User{Id: userId})
	return err
}

/*
	改User信息
*/
func UpdateUser(user *User, cols ...string) (int64, error) {
	return o.Update(user, cols...)
}
