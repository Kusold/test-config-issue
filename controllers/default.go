package controllers

import (
	"github.com/astaxie/beego"
)

var testvariable = beego.AppConfig.String("testvariable")

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["TestVariable"] = testvariable
	c.TplNames = "index.tpl"
}
