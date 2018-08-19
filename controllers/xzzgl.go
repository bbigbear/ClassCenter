package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type XzzglController struct {
	BaseController
}

func (this *XzzglController) AddAction() {
	fmt.Println("add case")
	o := orm.NewOrm()
	list := make(map[string]interface{})
	var xzz models.Xzz
	json.Unmarshal(this.Ctx.Input.RequestBody, &xzz)

	fmt.Println("xzz_info:", &xzz)
	//time
	nowtime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("新建失败")
	}
	xzz.CreateTime = nowtime
	//insert
	_, err1 := o.Insert(&xzz)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["xzzId"] = xzz.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *XzzglController) GetData() {
	o := orm.NewOrm()
	var maps []orm.Params
	xzz := new(models.Xzz)
	num, err := o.QueryTable(xzz).Values(&maps)
	if err != nil {
		fmt.Println("get xzz err", err.Error())
		this.ajaxMsg("get xzz err", MSG_ERR_Resources)
	}
	fmt.Println("get xzz reslut num:", num)
	this.ajaxList("get xzz data success", 0, num, maps)
	return
}

func (this *XzzglController) Edit() {
	o := orm.NewOrm()
	var maps []orm.Params
	xzz := new(models.Xzz)

	id := this.Input().Get("id")
	fmt.Println("id:", id)

	num, err := o.QueryTable(xzz).Filter("Id", id).Values(&maps)
	if err != nil {
		fmt.Println("edit xzz err", err.Error())
		this.ajaxMsg("edit xzz err", MSG_ERR_Resources)
	}
	fmt.Println("edit xzz reslut num:", num)
	this.Data["m"] = maps
	fmt.Println("maps", maps)
	this.TplName = "xzzgl_edit.tpl"
}

func (this *XzzglController) EditAction() {
	fmt.Println("edit action")
	o := orm.NewOrm()
	var xzz models.Xzz
	json.Unmarshal(this.Ctx.Input.RequestBody, &xzz)
	fmt.Println("xzz_info:", &xzz)
	//updata xzz db
	_, err1 := o.Update(&xzz)
	if err1 != nil {
		fmt.Println("updata xzz err", err1.Error())
		this.ajaxMsg("updata xzz err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update xzz success", MSG_OK)
	return
}

func (this *XzzglController) Del() {
	fmt.Println("del xzz")
	//id
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("del xzz err", err.Error())
		this.ajaxMsg("del xzz err", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	o := orm.NewOrm()
	xzz := new(models.Xzz)
	num, err := o.QueryTable(xzz).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del xzz err", err.Error())
		this.ajaxMsg("del xzz err", MSG_ERR_Resources)
	}
	fmt.Println("del xzz reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del xzz success", MSG_OK)
	return
}
