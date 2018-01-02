package controllers

import (
	"github.com/astaxie/beego"
)

type ChooseClassController struct {
	beego.Controller
}

func (c *ChooseClassController) Get() {
	c.TplName = "chooseClassOnline.tpl"
}
