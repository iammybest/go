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
	this.Data["Website"] = "https://github.com/iammybest/go"
	this.Data["Email"] = "396186655@qq.com"
	this.Data["Admin"] = "Mr.Deng"
	this.TplName = "deng.tpl"
}
