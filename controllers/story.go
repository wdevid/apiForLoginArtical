package controllers

import (
	"github.com/astaxie/beego"
	"apiForLoginArtical/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type StoryControllers struct {
	beego.Controller
}

func (this *StoryControllers) Post(){
	go story(this)
	this.ServeJSON()
}
func story(this *StoryControllers) {
	var artical models.Customer
	json.Unmarshal(this.Ctx.Input.RequestBody, &artical)
	o := orm.NewOrm()
	o.Using("myapp")
	_, err := o.Insert(&artical)
	if err != nil {
		this.Data["json"] = "{\"artical\":\"发布成功!!!\"}"
	} else {
		this.Data["json"] = "{\"artical\":\"发布失败!!!\"}"
	}
}
