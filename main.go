package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"samples/apiForLoginArtical/models"
	"samples/apiForLoginArtical/controllers"
)

func main() {
	models.RegisterDB()
	//// 开启 ORM 调试模式
	orm.Debug = true
	//// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/example/shorten", &controllers.ShortController{})
	beego.Router("/example/expand", &controllers.ExpandController{})
	beego.Router("/example/regester", &controllers.RegesterController{})
	beego.Router("/example/login", &controllers.LoginController{})
	beego.Run()
}
