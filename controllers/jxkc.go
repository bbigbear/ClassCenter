package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type JxkcController struct {
	BaseController
}

func (this *JxkcController) ApplyGet() {
	this.TplName = "jxkc_apply.tpl"
}

func (this *JxkcController) CheckGet() {
	this.TplName = "jxkc_check.tpl"
}

func (this *JxkcController) MaintainGet() {
	this.TplName = "jxkc_maintain.tpl"
}

func (this *JxkcController) Add() {
	this.TplName = "jxkc_addclass.tpl"
}

func (this *JxkcController) AddAction() {
	fmt.Println("add course")
	o := orm.NewOrm()
	list := make(map[string]interface{})
	var kc models.Jxkc
	json.Unmarshal(this.Ctx.Input.RequestBody, &kc)

	fmt.Println("xjkc_info:", &kc)

	//time
	nowtime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("time err", err.Error())
	}
	kc.CreateTime = nowtime
	kc.Status = "未审核"

	//insert
	_, err1 := o.Insert(&kc)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["courseId"] = kc.CourseId
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *JxkcController) GetData() {
	fmt.Println("get jxkc data")
	o := orm.NewOrm()
	var maps []orm.Params
	kc := new(models.Jxkc)
	query := o.QueryTable(kc)
	filters := make([]interface{}, 0)
	//name
	name := this.Input().Get("name")
	if name != "" {
		filters = append(filters, "CourseName", name)
	}
	fmt.Println("get name:", name)
	//category
	category := this.Input().Get("category")
	if category != "" {
		filters = append(filters, "ElectiveCategory", category)
	}
	fmt.Println("get category:", category)

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	//get data dB
	num, err := query.Values(&maps)
	if err != nil {
		fmt.Println("get jxkc err", err.Error())
		this.ajaxMsg("get jxkc err", MSG_ERR_Resources)
	}
	fmt.Println("get jxkc reslut num:", num)
	this.ajaxList("get jxkc data success", 0, num, maps)
	return
}

func (this *JxkcController) Edit() {
	o := orm.NewOrm()
	var maps []orm.Params
	kc := new(models.Jxkc)

	id := this.Input().Get("id")
	fmt.Println("id:", id)

	num, err := o.QueryTable(kc).Filter("Id", id).Values(&maps)
	if err != nil {
		fmt.Println("edit jxkc err", err.Error())
		this.ajaxMsg("edit jxkc err", MSG_ERR_Resources)
	}
	fmt.Println("edit rq reslut num:", num)
	this.Data["m"] = maps
	fmt.Println("maps", maps)
	this.TplName = "jxkc_edit.tpl"
}

func (this *JxkcController) EditAction() {
	fmt.Println("edit action")
	o := orm.NewOrm()
	var kc models.Jxkc
	json.Unmarshal(this.Ctx.Input.RequestBody, &kc)
	fmt.Println("kc_info:", &kc)
	//updata jxkc db
	_, err := o.Update(&kc)
	if err != nil {
		fmt.Println("updata jxkc err", err.Error())
		this.ajaxMsg("updata jxkc err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update jxkc success", MSG_OK)
	return
}
