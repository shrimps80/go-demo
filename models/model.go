package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Name     string
	PassWord string
}

func init() {
	//注册数据库,获取连接对象
	orm.RegisterDataBase("default","mysql","root@tcp(127.0.0.1:3306)/beego_demo?charset=utf8")
	//创建表格
	orm.RegisterModel(new(User))
	//生成表格
	orm.RunSyncdb("default",false,true)
}
