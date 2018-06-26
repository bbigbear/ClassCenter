package controllers

//import (
//	"ClassCenter/models"
//	"encoding/json"
//	"fmt"
//	"time"

//	_ "github.com/Go-SQL-Driver/MySQL"
//	"github.com/astaxie/beego/orm"
//)

type XxkglController struct {
	BaseController
}

func (this *XxkglController) SettingGet() {
	this.TplName = "xxkgl_setting.tpl"
}

func (this *XxkglController) TaskGet() {
	this.TplName = "xxkgl_task.tpl"
}

func (this *XxkglController) ApplyGet() {
	this.TplName = "xxkgl_openclass_apply.tpl"
}

func (this *XxkglController) ApproveGet() {
	this.TplName = "xxkgl_openclass_approve.tpl"
}
