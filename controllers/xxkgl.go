package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type XxkglController struct {
	BaseController
}

func (this *XxkglController) SettingGet() {
	this.TplName = "xxkgl_setting.tpl"
}

func (this *XxkglController) AddSetting() {
	fmt.Println("add xk_setting")
	o := orm.NewOrm()
	list := make(map[string]interface{})
	var setting models.XkSetting
	json.Unmarshal(this.Ctx.Input.RequestBody, &setting)

	fmt.Println("xk_setting_info:", &setting)

	//insert
	_, err := o.Insert(&setting)
	if err != nil {
		fmt.Printf("insert err", err.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["Id"] = setting.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *XxkglController) TaskGet() {
	this.TplName = "xxkgl_task.tpl"
}

func (this *XxkglController) TaskAddGet() {
	this.TplName = "xxkgl_task_add.tpl"
}

func (this *XxkglController) AddAction() {
	fmt.Println("add task")
	o := orm.NewOrm()
	list := make(map[string]interface{})
	var xkrw models.Xkrw
	json.Unmarshal(this.Ctx.Input.RequestBody, &xkrw)

	fmt.Println("xkrw_info:", &xkrw)

	//insert
	_, err1 := o.Insert(&xkrw)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["taskId"] = xkrw.TaskId
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *XxkglController) GetData() {
	fmt.Println("get xkrw data")
	o := orm.NewOrm()
	var maps []orm.Params
	xkrw := new(models.Xkrw)
	query := o.QueryTable(xkrw)
	//get data dB
	num, err := query.Values(&maps)
	if err != nil {
		fmt.Println("get xkrw err", err.Error())
		this.ajaxMsg("get xkrw err", MSG_ERR_Resources)
	}
	fmt.Println("get xkrw reslut num:", num)
	this.ajaxList("get xkrw data success", 0, num, maps)
	return
}

func (this *XxkglController) Edit() {
	o := orm.NewOrm()
	var maps []orm.Params
	xkrw := new(models.Xkrw)

	id := this.Input().Get("id")
	fmt.Println("id:", id)

	num, err := o.QueryTable(xkrw).Filter("Id", id).Values(&maps)
	if err != nil {
		fmt.Println("edit xkrw err", err.Error())
		this.ajaxMsg("edit xkrw err", MSG_ERR_Resources)
	}
	fmt.Println("edit xkrw reslut num:", num)
	this.Data["m"] = maps
	fmt.Println("maps", maps)
	this.TplName = "xxkgl_task_edit.tpl"
}

func (this *XxkglController) EditAction() {
	fmt.Println("edit action")
	o := orm.NewOrm()
	var xkrw models.Xkrw
	json.Unmarshal(this.Ctx.Input.RequestBody, &xkrw)
	fmt.Println("xkrw_info:", &xkrw)
	//updata xkrw db
	_, err1 := o.Update(&xkrw)
	if err1 != nil {
		fmt.Println("updata xkrw err", err1.Error())
		this.ajaxMsg("updata xkrw err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update xkrw success", MSG_OK)
	return
}

func (this *XxkglController) ApplyGet() {
	this.TplName = "xxkgl_openclass_apply.tpl"
}

func (this *XxkglController) OpenClassAdd() {
	this.TplName = "xxkgl_openclass_add.tpl"
}

func (this *XxkglController) ApproveGet() {
	this.TplName = "xxkgl_openclass_approve.tpl"
}
