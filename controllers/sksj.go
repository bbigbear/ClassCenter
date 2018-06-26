package controllers

//import (
//	"ClassCenter/models"
//	"encoding/json"
//	"fmt"
//	"time"

//	_ "github.com/Go-SQL-Driver/MySQL"
//	"github.com/astaxie/beego/orm"
//)

type SKsjController struct {
	BaseController
}

func (this *SKsjController) CompulsoryGet() {
	this.TplName = "sksj_compulsory_setting.tpl"
}

func (this *SKsjController) EletiveGet() {
	this.TplName = "sksj_eletive_setting.tpl"
}
