package controllers

import (
	"github.com/astaxie/beego"
)

type QueryClassController struct {
	beego.Controller
}

func (c *QueryClassController) Get() {
	c.TplName = "queryClass.tpl"
}
