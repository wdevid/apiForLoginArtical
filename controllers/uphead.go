package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"strings"
	"apiForLoginArtical/models"
	"github.com/astaxie/beego/orm"
	mylog "apiForLoginArtical/mylog"
)

type UpHeadController struct {
	beego.Controller
}
/**
更新或者上传用户头像
 */
func Post(this *UpHeadController)  {
	var ob models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
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
	_, err := o.Update(&ob)
	mylog.LogersError(err.Error())
	if err != nil {
		this.Data["json"] = "{\"state\":\"0\"}"
	} else {
		this.Data["json"] = "{\"state\":\"1\"}"
	}
	this.ServeJSON()
}
