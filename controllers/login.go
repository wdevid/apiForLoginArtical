package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"samples/apiForLoginArtical/models"
	"samples/apiForLoginArtical/utils"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post()  {
	var ob models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	passWd ,err:= utils.Base64Decode([]byte(ob.PassWord))
	if err==nil {
		o := orm.NewOrm()
		o.Using("myapp")
		u := new(models.User)
		qs := o.QueryTable(&u)
		if qs != nil {
			exist := qs.Exclude("UserName",ob.UserName).Filter("PassWord",passWd).Exist()
			if exist {
				this.Data["json"] = "{\"login\":\"登陆成功!!!\"}"
			}else {
				this.Data["json"] = "{\"login\":\"用户名密码错误!!!\"}"
			}
			existName := qs.Exclude("UserName",ob.UserName).Exist()
			if !existName {
				this.Data["json"] = "{\"login\":\"用户名不存在!!!\"}"
			}
		}
	}
	this.ServeJSON()
}
