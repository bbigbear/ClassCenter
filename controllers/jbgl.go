package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type JbglController struct {
	BaseController
}

func (this *JbglController) Get() {
	this.TplName = "jbgl_plan_maintain.tpl"
}

func (this *JbglController) Add() {
	this.TplName = "jbgl_plan_add.tpl"
}

func (this *JbglController) AddAction() {
	fmt.Println("add case")
	o := orm.NewOrm()
	list := make(map[string]interface{})
	var jbjh models.Jbjh
	json.Unmarshal(this.Ctx.Input.RequestBody, &jbjh)

	fmt.Println("jbjh_info:", &jbjh)
	//time
	nowtime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("新建失败")
	}
	jbjh.CreateTime = nowtime
	jbjh.Status = "已创建"
	//insert
	_, err1 := o.Insert(&jbjh)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["jbjhId"] = jbjh.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *JbglController) GetData() {
	o := orm.NewOrm()
	var maps []orm.Params
	jbjh := new(models.Jbjh)
	num, err := o.QueryTable(jbjh).Values(&maps)
	if err != nil {
		fmt.Println("get jbjh err", err.Error())
		this.ajaxMsg("get jbjh err", MSG_ERR_Resources)
	}
	fmt.Println("get jbjh reslut num:", num)
	this.ajaxList("get jbjh data success", 0, num, maps)
	return
}

func (this *JbglController) Edit() {
	o := orm.NewOrm()
	var maps []orm.Params
	jbjh := new(models.Jbjh)

	id := this.Input().Get("id")
	fmt.Println("id:", id)

	num, err := o.QueryTable(jbjh).Filter("Id", id).Values(&maps)
	if err != nil {
		fmt.Println("edit jbjh err", err.Error())
		this.ajaxMsg("edit jbjh err", MSG_ERR_Resources)
	}
	fmt.Println("edit jbjh reslut num:", num)
	this.Data["m"] = maps
	fmt.Println("maps", maps)
	this.TplName = "jbgl_plan_edit.tpl"
}

func (this *JbglController) EditAction() {
	fmt.Println("edit action")
	o := orm.NewOrm()
	var jbjh models.Jbjh
	json.Unmarshal(this.Ctx.Input.RequestBody, &jbjh)
	fmt.Println("jbjh_info:", &jbjh)
	//updata jxjh db
	_, err1 := o.Update(&jbjh)
	if err1 != nil {
		fmt.Println("updata jbjh err", err1.Error())
		this.ajaxMsg("updata jbjh err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update jbjh success", MSG_OK)
	return
}
func (this *JbglController) ChangeStatus() {
	fmt.Println("change status")
	//id
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("get id err", err.Error())
		this.ajaxMsg("get id err", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	//status
	status := this.GetString("status")
	fmt.Println("status is", status)

	o := orm.NewOrm()
	jbjh := new(models.Jbjh)
	//updata status db
	num, err := o.QueryTable(jbjh).Filter("Id", id).Update(orm.Params{
		"Status": status,
	})
	if err != nil {
		fmt.Println("change status err", err.Error())
		this.ajaxMsg("change status err", MSG_ERR_Resources)
	}
	fmt.Println("num", num)
	this.ajaxMsg("update status success", MSG_OK)
	return
}

func (this *JbglController) Del() {
	fmt.Println("del jbjh")
	//id
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("del jbjh err", err.Error())
		this.ajaxMsg("del jbjh err", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	o := orm.NewOrm()
	jbjh := new(models.Jbjh)
	num, err := o.QueryTable(jbjh).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del jbjh err", err.Error())
		this.ajaxMsg("del jbjh err", MSG_ERR_Resources)
	}
	fmt.Println("del jbjh reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del jbjh success", MSG_OK)
	return
}

func (this *JbglController) WaitLaunch() {
	this.TplName = "jbgl_plan_waitlaunch.tpl"
}

func (this *JbglController) Launch() {
	this.TplName = "jbgl_plan_launch.tpl"
}

func (this *JbglController) Join() {
	this.TplName = "jbgl_plan_join.tpl"
}

func (this *JbglController) View() {
	this.TplName = "jbgl_plan_view.tpl"
}

func (this *JbglController) Check() {
	this.TplName = "jbjc_plan_check.tpl"
}

func (this *JbglController) CheckJoin() {
	this.TplName = "jbjc_plan_join.tpl"
}

func (this *JbglController) CheckView() {
	this.TplName = "jbjc_plan_view.tpl"
}
