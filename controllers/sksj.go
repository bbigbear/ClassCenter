package controllers

//import (
//	"ClassCenter/models"
//	"encoding/json"
//	"fmt"
//	"time"

//	_ "github.com/Go-SQL-Driver/MySQL"
//	"github.com/astaxie/beego/orm"
//)

type SksjController struct {
	BaseController
}

func (this *SksjController) CompulsoryGet() {
	this.TplName = "sksj_compulsory_setting.tpl"
}

func (this *SksjController) EletiveGet() {
	this.TplName = "sksj_eletive_setting.tpl"
}
