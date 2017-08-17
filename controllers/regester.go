package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"apiForLoginArtical/models"
	"apiForLoginArtical/utils"
)

type RegesterController struct {
	beego.Controller
}


func (this *RegesterController) Post() {
	var ob models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	passWd := utils.Base64Encode([]byte(ob.PassWord))
	ob.PassWord = string(passWd)
	o := orm.NewOrm()
	o.Using("myapp")
	_,err := o.Insert(&ob)
	if err !=nil {
		this.Data["json"] = "{\"state\":\"0\"}"
	}else {
		this.Data["json"] = "{\"state\":\"1\"}"
	}
	this.ServeJSON()
}
