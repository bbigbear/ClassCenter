package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	"time"

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

func (this *JsbkController) CaseAddAction() {
	fmt.Println("add case")
	o := orm.NewOrm()
	list := make(map[string]interface{})
	var cases models.Case
	json.Unmarshal(this.Ctx.Input.RequestBody, &cases)

	fmt.Println("case_info:", &cases)
	cases.Status = "未审核"
	//time
	nowtime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("新建失败")
	}
	cases.CreateTime = nowtime
	//insert
	_, err1 := o.Insert(&cases)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["caseId"] = cases.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *JsbkController) CaseGetData() {
	o := orm.NewOrm()
	var maps []orm.Params
	cases := new(models.Case)
	num, err := o.QueryTable(cases).Values(&maps)
	if err != nil {
		fmt.Println("get case err", err.Error())
		this.ajaxMsg("get case err", MSG_ERR_Resources)
	}
	fmt.Println("get case reslut num:", num)
	this.ajaxList("get case data success", 0, num, maps)
	return
}

func (this *JsbkController) Edit() {
	o := orm.NewOrm()
	var maps []orm.Params
	cases := new(models.Case)

	id := this.Input().Get("id")
	fmt.Println("id:", id)

	num, err := o.QueryTable(cases).Filter("Id", id).Values(&maps)
	if err != nil {
		fmt.Println("edit cases err", err.Error())
		this.ajaxMsg("edit cases err", MSG_ERR_Resources)
	}
	fmt.Println("edit cases reslut num:", num)
	this.Data["m"] = maps
	fmt.Println("maps", maps)
	this.TplName = "jsbk_teachcase_edit.tpl"
}

func (this *JsbkController) EditAction() {
	fmt.Println("edit action")
	o := orm.NewOrm()
	var cases models.Case
	json.Unmarshal(this.Ctx.Input.RequestBody, &cases)
	fmt.Println("cases_info:", &cases)
	//updata jxjh db
	_, err1 := o.Update(&cases)
	if err1 != nil {
		fmt.Println("updata cases err", err1.Error())
		this.ajaxMsg("updata cases err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update cases success", MSG_OK)
	return
}
func (this *JsbkController) ChangeStatus() {
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
	cases := new(models.Case)
	//updata status db
	num, err := o.QueryTable(cases).Filter("Id", id).Update(orm.Params{
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

func (this *JsbkController) Del() {
	fmt.Println("del cases")
	//id
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("del cases err", err.Error())
		this.ajaxMsg("del cases err", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	o := orm.NewOrm()
	cases := new(models.Case)
	num, err := o.QueryTable(cases).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del cases err", err.Error())
		this.ajaxMsg("del cases err", MSG_ERR_Resources)
	}
	fmt.Println("del cases reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del cases success", MSG_OK)
	return
}

func (this *JsbkController) TeachGroup() {
	this.TplName = "jsbk_teachgroup.tpl"
}

func (this *JsbkController) TeachGroupAdd() {
	this.TplName = "jsbk_teachgroup_add.tpl"
}

func (this *JsbkController) TeachGroupAddAction() {
	fmt.Println("add group")
	o := orm.NewOrm()
	list := make(map[string]interface{})
	var group models.TeachGroup
	json.Unmarshal(this.Ctx.Input.RequestBody, &group)

	fmt.Println("group_info:", &group)
	//insert
	_, err1 := o.Insert(&group)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["groudId"] = group.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *JsbkController) TeachGroupGetData() {
	o := orm.NewOrm()
	var maps []orm.Params
	group := new(models.TeachGroup)
	num, err := o.QueryTable(group).Values(&maps)
	if err != nil {
		fmt.Println("get group err", err.Error())
		this.ajaxMsg("get group err", MSG_ERR_Resources)
	}
	fmt.Println("get group reslut num:", num)
	this.ajaxList("get group data success", 0, num, maps)
	return
}

func (this *JsbkController) TeachGroupEdit() {
	o := orm.NewOrm()
	var maps []orm.Params
	group := new(models.TeachGroup)

	id := this.Input().Get("id")
	fmt.Println("id:", id)

	num, err := o.QueryTable(group).Filter("Id", id).Values(&maps)
	if err != nil {
		fmt.Println("edit group err", err.Error())
		this.ajaxMsg("edit group err", MSG_ERR_Resources)
	}
	fmt.Println("edit group reslut num:", num)
	this.Data["m"] = maps
	fmt.Println("maps", maps)
	this.TplName = "jsbk_teachgroup_edit.tpl"
}

func (this *JsbkController) TeachGroupEditAction() {
	fmt.Println("edit action")
	o := orm.NewOrm()
	var group models.TeachGroup
	json.Unmarshal(this.Ctx.Input.RequestBody, &group)
	fmt.Println("group_info:", &group)
	//updata jxjh db
	_, err1 := o.Update(&group)
	if err1 != nil {
		fmt.Println("updata group err", err1.Error())
		this.ajaxMsg("updata group err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update group success", MSG_OK)
	return
}

func (this *JsbkController) TeachGroupDel() {
	fmt.Println("del group")
	//id
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("del group err", err.Error())
		this.ajaxMsg("del group err", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	o := orm.NewOrm()
	group := new(models.TeachGroup)
	num, err := o.QueryTable(group).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del group err", err.Error())
		this.ajaxMsg("del group err", MSG_ERR_Resources)
	}
	fmt.Println("del group reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del group success", MSG_OK)
	return
}

func (this *JsbkController) TeachGroupAddPepole() {
	this.TplName = "jsbk_teachgroup_addpeople.tpl"
}

func (this *JsbkController) TeachGroupAddLeader() {
	this.TplName = "jsbk_teachgroup_addleader.tpl"
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
