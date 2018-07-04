package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	//	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type PkglController struct {
	BaseController
}

func (this *PkglController) SettingGet() {
	this.TplName = "pkgl_arranging_setting.tpl"
}

func (this *PkglController) SettingSave() {
	fmt.Println("save PkSetting action")
	o := orm.NewOrm()
	var setting models.PkSetting
	json.Unmarshal(this.Ctx.Input.RequestBody, &setting)
	fmt.Println("pkset_info:", &setting)
	//updata pkgl db
	_, err1 := o.Update(&setting)
	if err1 != nil {
		fmt.Println("updata PkSetting err", err1.Error())
		this.ajaxMsg("updata PkSetting err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update PkSetting success", MSG_OK)
	return
}

func (this *PkglController) ClassRoomUnarrangingTimeGet() {
	this.TplName = "pkgl_classroom_unarranging_time.tpl"
}

func (this *PkglController) TeacherUnarrangingTimeGet() {
	this.TplName = "pkgl_teacher_unarranging_time.tpl"
}

func (this *PkglController) EletivePkGet() {
	this.TplName = "pkgl_eletive_pk.tpl"
}

func (this *PkglController) CompulsoryPkGet() {
	this.TplName = "pkgl_compulsory_pk.tpl"
}
