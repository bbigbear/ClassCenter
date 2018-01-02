package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {

	uaccount := c.Input().Get("inputAccount")
	pwd := c.Input().Get("inputPassword")

	if beego.AppConfig.String("uaccount") == uaccount &&
		beego.AppConfig.String("pwd") == pwd {
		//		c.TplName = "studentCenter.tpl"
		c.Redirect("/center", 301)
	}

	c.Redirect("/login", 301)
	return

}
