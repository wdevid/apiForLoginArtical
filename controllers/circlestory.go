package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"apiForLoginArtical/models"
	"encoding/json"
	"apiForLoginArtical/mylog"
)

type CircleControllers struct {
	beego.Controller
}
func (this *CircleControllers) Post(){
	a,err := this.GetInt("page",-1)
	mylog.LogersError(err.Error())
	go circles(err, a, this)
	this.ServeJSON()
}
func circles(err error, a int, this *CircleControllers) {
	if err == nil {
		o := orm.NewOrm()
		o.Using("myapp")
		customers := make([]*models.Customer, 20)
		ids := make([]int64, a, a+20)
		for _, num := range ids {
			cus := new(models.Customer)
			cus.Id = num
			err := o.Read(cus)

			if err != nil {
				this.Data["json"] = "{\"circle\":\"没有更多数据!!!\"}"
			} else {
				customers[num] = cus
				mylog.Logers(string(num))
			}
		}
		this.Data["json"], err = json.Marshal(&customers)

	} else {
		this.Data["json"] = "{\"circle\":\"没有更多数据!!!\"}"
	}
}
