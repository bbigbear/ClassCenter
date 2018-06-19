package routers

import (
	"classCenter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//jxkc
	beego.Router("/v1/jxkc/apply", &controllers.JxkcController{}, "*:ApplyGet")
	beego.Router("/v1/jxkc/check", &controllers.JxkcController{}, "*:CheckGet")
	beego.Router("/v1/jxkc/maintain", &controllers.JxkcController{}, "*:MaintainGet")
	beego.Router("/v1/jxkc/add", &controllers.JxkcController{}, "*:Add")
	beego.Router("/v1/jxkc/add_action", &controllers.JxkcController{}, "post:AddAction")
	beego.Router("/v1/jxkc/getdata", &controllers.JxkcController{}, "Get:GetData")
	beego.Router("/v1/jxkc/edit", &controllers.JxkcController{}, "*:Edit")
	beego.Router("/v1/jxkc/edit_action", &controllers.JxkcController{}, "post:EditAction")
}
