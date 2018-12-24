package controllers

import (
  "fmt"
  "strconv"
  "demo/models"
  "demo/util"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) ShowRegister() {
	this.TplName = "register.tpl"
}

func (this *UserController) Register() {
	//接受页面数据
	userName := this.GetString("userName")
	password := this.GetString("password")
	beego.Info(userName)
	beego.Info(password)
	//进行判断
	if userName == "" || password == "" {
		beego.Error("注册不合法,请重新输入")
		this.TplName = "register.tpl"
		return
	}
	//如果注册格式合法,将数据填充到数据库中
	//1.获取orm对象
	o := orm.NewOrm()
	//2.定义用户
	var user models.User
  o.QueryTable("User").Filter("Name", userName).One(&user)
  if user.Id > 0 {
    beego.Error("用户名已经存在")
		this.TplName = "register.tpl"
		return
  }
	//赋值
	user.Name = userName
	//加密用户密码
  //user.PassWord = password
  user.PassWord = util.SHA256([]byte(password))
	//插入数据库
	i, e := o.Insert(&user)
	if e != nil {
		beego.Error("用户数据插入失败:",e)
	}
	beego.Info("用户数据插入成功",i)
	//如果登录成功返回数据
	//this.Ctx.WriteString("恭喜你,注册成功 ~ ")
	//页面重定向 重定向:内部重新再发送一次.
	this.Redirect("/login",302)

}

func (this *UserController) ShowLogin() {
	this.TplName = "login.tpl"
}

func (this *UserController) Login() {
	//接受数据
	userName := this.GetString("userName")
	password := this.GetString("password")
	beego.Info(userName)
	beego.Info(password)
	//判断条件
	if userName == "" || password == "" {
		beego.Error("数据不合法,请重新输入")
		this.Data["errmsg"] = "数据不合法,请重新输入"
		this.TplName="login.tpl"
		return
	}
	//操作数据
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Error("用户不存在")
		this.Data["errmsg"]="用户不存在,请重新输入"
		this.TplName="login.tpl"
		return
	}
	//将用户密码加密
	user.PassWord = util.SHA256([]byte(password))
	err = o.Read(&user, "PassWord")
	if err!=nil{
		beego.Error("密码错误")
		this.Data["errmsg"]="密码错误,请重新输入"
		this.TplName="login.tpl"
		return
	}

  // cookie
	checked := this.GetString("remember")
	if checked == "on" {
		this.Ctx.SetCookie("userName", userName, 60*60*24*7)
	}else {
		this.Ctx.SetCookie("userName", userName, -1)
	}

  // 保存下给info使用
  fmt.Printf("userId: %d\n", user.Id)
  this.Ctx.SetCookie("userId", strconv.Itoa(user.Id), 60*60*24*7)
	//this.SetSession("userName",userName)
	// this.Ctx.WriteString("恭喜您,登录成功")
  this.Redirect("/info",302)
}

func (this *UserController) Info()  {
  userId := this.Ctx.GetCookie("userId")
  fmt.Printf("userId: %d\n", userId)

  o := orm.NewOrm()
	var user models.User
  o.QueryTable("User").Filter("Id", userId).One(&user)

  fmt.Printf("userId: %d\n", user.Id)
  if user.Id == 0 {
    beego.Error("找不到用户id ----- %d", userId)
    this.Data["errmsg"]="找不到用户id，重新登录"
		this.TplName="login.tpl"
		return
  }

  this.Data["json"] = user
  this.Ctx.Output.JSON(this.Data["json"], true, false)
}
