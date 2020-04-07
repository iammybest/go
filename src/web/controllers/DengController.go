package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type DengController struct {
	beego.Controller
}

func (this *DengController)Prepare()  {
	log.Println("Init DengController")
}
func (this *DengController) Get() {
	v := this.GetSession("asta")
	if v == nil {
		this.SetSession("asta", int(1))
		this.Data["num"] = 0
	} else {
		this.SetSession("asta", v.(int)+1)
		this.Data["num"] = v.(int)
	}
	//日志展示
	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")

	this.Data["Website"] = "https://github.com/iammybest/go"
	this.Data["Email"] = "396186655@qq.com"
	this.Data["Admin"] = "Mr.Deng"
	this.TplName = "deng.tpl"
}
