package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"apiForLoginArtical/models"
	"apiForLoginArtical/utils"
	"apiForLoginArtical/mylog"
	"strings"
)

type RegesterController struct {
	beego.Controller
}


func (this *RegesterController) Post() {
	go regester(this)
}
func regester(this *RegesterController) {
	var ob models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	passWd := utils.Base64Encode([]byte(ob.PassWord))
	ob.PassWord = string(passWd)
	if ob.IsHead == 0 {
		f,h,_:=this.GetFile("filename")
		path := h.Filename
		arr := strings.Split(path,":")
		if len(arr)>1 {
			index := len(arr) - 1
			path = arr[index]
		}
		f.Close()
		this.SaveToFile("FileForPortrait",path)
		ob.HeadPath = "/FileForPortrait"+path
	}
	o := orm.NewOrm()
	o.Using("myapp")
	_, err := o.Insert(&ob)
	mylog.LogersError(err.Error())
	if err != nil {
		this.Data["json"] = "{\"state\":\"0\"}"
	} else {
		this.Data["json"] = "{\"state\":\"1\"}"
	}
	this.ServeJSON()
}
