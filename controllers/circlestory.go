package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"apiForLoginArtical/models"
	"encoding/json"
)

type CircleControllers struct {
	beego.Controller
}

func (this *CircleControllers) Post(){
	a,err := this.GetInt("page",-1)
	if err!=nil {
		o := orm.NewOrm()
		o.Using("myapp")
		c := new(models.Customer)
		qs := o.QueryTable(&c).Filter("id__in",a,a+20)
		if qs != nil {
			this.Data["json"],err = json.Marshal(&c)
		}

	}else {
		this.Data["json"] = "{\"circle\":\"没有更多数据!!!\"}"
	}
	this.ServeJSON()
}
