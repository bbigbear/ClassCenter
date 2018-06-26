package controllers

//import (
//	"ClassCenter/models"
//	"encoding/json"
//	"fmt"
//	"time"

//	_ "github.com/Go-SQL-Driver/MySQL"
//	"github.com/astaxie/beego/orm"
//)

type XkglController struct {
	BaseController
}

func (this *XkglController) SettingGet() {
	this.TplName = "xkgl_arranging_setting.tpl"
}

func (this *XkglController) ClassRoomUnarrangingTimeGet() {
	this.TplName = "xkgl_classroom_unarranging_time.tpl"
}

func (this *XkglController) TeacherUnarrangingTimeGet() {
	this.TplName = "xkgl_teacher_unarranging_time.tpl"
}

func (this *XkglController) EletivePkGet() {
	this.TplName = "xkgl_eletive_pk.tpl"
}

func (this *XkglController) CompulsoryPkGet() {
	this.TplName = "xkgl_compulsory_pk.tpl"
}
