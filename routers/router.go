package routers

import (
	"demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:Register")
		beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:Login")
		beego.Router("/info", &controllers.UserController{}, "get:Info")

}
