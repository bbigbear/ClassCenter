package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type SksjController struct {
	BaseController
}

func (this *SksjController) CompulsoryGet() {
	this.TplName = "sksj_compulsory_setting.tpl"
}

func (this *SksjController) EletiveGet() {
	this.TplName = "sksj_eletive_setting.tpl"
}

func (this *SksjController) GetData() {
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
func (this *SksjController) Edit() {
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

func (this *SksjController) EditAction() {
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

func (this *SksjController) Save() {
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
