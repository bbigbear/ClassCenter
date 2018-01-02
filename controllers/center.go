package controllers

import (
	"github.com/astaxie/beego"
)

type CenterController struct {
	beego.Controller
}

func (c *CenterController) Get() {
	c.TplName = "studentCenter.tpl"
}
