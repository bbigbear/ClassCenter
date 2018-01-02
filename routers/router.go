package routers

import (
	"ClassCenter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/center", &controllers.CenterController{})
	beego.Router("/chooseClass", &controllers.ChooseClassController{})
	beego.Router("/queryClass", &controllers.QueryClassController{})
}
