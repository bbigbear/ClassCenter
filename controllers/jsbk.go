package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type JsbkController struct {
	BaseController
}

func (this *JsbkController) CaseGet() {
	this.TplName = "jsbk_teachcase.tpl"
}

func (this *JsbkController) CaseAdd() {
	this.TplName = "jsbk_teachcase_add.tpl"
}

func (this *JsbkController) TeachGroup() {
	this.TplName = "jsbk_teachgroup.tpl"
}

func (this *JsbkController) TeachCaseApprove() {
	this.TplName = "jsbk_teachcase_approve.tpl"
}

func (this *JsbkController) TeachCaseQuery() {
	this.TplName = "jsbk_teachcase_query.tpl"
}

func (this *JsbkController) TimeManage() {
	this.TplName = "jsbk_timemanage.tpl"
}

func (this *JsbkController) TeachSchedule() {
	this.TplName = "jsbk_teachschedule_statistics.tpl"
}

func (this *JsbkController) GetData() {
	//
	style, err := this.GetInt("style")
	if err != nil {
		fmt.Println("get sksj style err", err.Error())
	}
	o := orm.NewOrm()
	var maps []orm.Params
	kbsz := new(models.Kbsz)
	num, err := o.QueryTable(kbsz).Filter("Style", style).Values(&maps)
	if err != nil {
		fmt.Println("get kbsz err", err.Error())
		this.ajaxMsg("get kbsz err", MSG_ERR_Resources)
	}
	fmt.Println("get kbsz reslut num:", num)
	this.ajaxList("get kbsz data success", 0, num, maps)
	return
}
func (this *JsbkController) Edit() {
	o := orm.NewOrm()
	var maps []orm.Params
	kbsz := new(models.Kbsz)

	id := this.Input().Get("id")
	fmt.Println("id:", id)

	num, err := o.QueryTable(kbsz).Filter("Id", id).Values(&maps)
	if err != nil {
		fmt.Println("edit kbsz err", err.Error())
		this.ajaxMsg("edit kbsz err", MSG_ERR_Resources)
	}
	fmt.Println("edit kbsz reslut num:", num)
	this.Data["m"] = maps
	fmt.Println("maps", maps)
	this.TplName = "sksj_kbsz_edit.tpl"
}

func (this *JsbkController) EditAction() {
	fmt.Println("edit action")
	o := orm.NewOrm()
	var kbsz models.Kbsz
	json.Unmarshal(this.Ctx.Input.RequestBody, &kbsz)
	fmt.Println("kbsz_info:", &kbsz)
	//updata kbsz db
	_, err1 := o.Update(&kbsz)
	if err1 != nil {
		fmt.Println("updata kbsz err", err1.Error())
		this.ajaxMsg("updata kbsz err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update kbsz success", MSG_OK)
	return
}

func (this *JsbkController) Save() {
	fmt.Println("edit action")
	o := orm.NewOrm()
	var kbsz models.KbszSetting
	json.Unmarshal(this.Ctx.Input.RequestBody, &kbsz)
	fmt.Println("kbsz_info:", &kbsz)
	//updata kbsz db
	_, err1 := o.Update(&kbsz)
	if err1 != nil {
		fmt.Println("updata kbsz err", err1.Error())
		this.ajaxMsg("updata kbsz err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update kbsz success", MSG_OK)
	return
}
