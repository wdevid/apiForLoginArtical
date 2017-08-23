package main

import (
	"apiForLoginArtical/controllers"
	"apiForLoginArtical/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.SetStaticPath("/", "FileForPortrait")
	beego.BConfig.WebConfig.Session.SessionOn = true
	models.RegisterDB()
	//// 开启 ORM 调试模式
	orm.Debug = true
	//// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/example/shorten", &controllers.ShortController{})
	beego.Router("/example/regester", &controllers.RegesterController{})
	beego.Router("/example/login", &controllers.LoginController{})
	beego.Router("/example/artical", &controllers.StoryControllers{})
	beego.Router("/example/circle", &controllers.CircleControllers{})
	beego.Run()
}
