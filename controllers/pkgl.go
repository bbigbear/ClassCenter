package controllers

//import (
//	"ClassCenter/models"
//	"encoding/json"
//	"fmt"
//	"time"

//	_ "github.com/Go-SQL-Driver/MySQL"
//	"github.com/astaxie/beego/orm"
//)

type PkglController struct {
	BaseController
}

func (this *PkglController) SettingGet() {
	this.TplName = "pkgl_arranging_setting.tpl"
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
