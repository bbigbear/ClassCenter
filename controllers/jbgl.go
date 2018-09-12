package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	"math"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type JbglController struct {
	BaseController
}

func (this *JbglController) AddAction() {
	fmt.Println("add case")
	//获取token
	var token models.Token
	json.Unmarshal(this.Ctx.Input.RequestBody, &token)

	if token.Token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}
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
	jbjh.Status = "已创建 待发起"
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
	//token
	token := this.Input().Get("token")

	if token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}
	o := orm.NewOrm()
	var maps []orm.Params
	jbjh := new(models.Jbjh)
	query := o.QueryTable(jbjh)

	//标题
	title := this.Input().Get("Title")
	if title != "" {
		query = query.Filter("Title", title)
	}
	//状态
	status := this.Input().Get("status")
	if status != "" {
		query = query.Filter("Status", status)
	}

	//我创建的
	founder := this.Input().Get("founder")
	if founder != "" {
		query = query.Filter("Founter", founder)
	}

	//我参加的
	part := this.Input().Get("part")
	if part != "" {
		query = query.Filter("Participant", part)
	}

	//index
	index, err := this.GetInt("index")
	if err != nil {
		fmt.Println("获取index下标错误")
	}

	//pagemax  一页多少
	pagemax, err := this.GetInt("pagemax")
	if err != nil {
		fmt.Println("获取每页数量为空")
	}

	//count
	count, err := query.Count()
	if err != nil {
		fmt.Println("获取数据总数为空")
		this.ajaxMsg("服务未知错误", MSG_ERR)
	}

	if pagemax != 0 {
		pagenum := int(math.Ceil(float64(count) / float64(pagemax)))

		if index > pagenum {
			//index = pagenum
			this.ajaxMsg("无法翻页了", MSG_ERR_Param)
		}
		fmt.Println("index&pagemax&pagenum", index, pagemax, pagenum)
	}
	query = query.Limit(pagemax, (index-1)*pagemax)

	num, err := query.OrderBy("-Id").Values(&maps)
	if err != nil {
		fmt.Println("get jbgl err", err.Error())
		this.ajaxMsg("get jbgl err", MSG_ERR_Resources)
	}
	fmt.Println("get jbgl reslut num:", num)
	this.ajaxList("get jbgl data success", MSG_OK, num, maps)

}

func (this *JbglController) EditAction() {
	fmt.Println("edit action")
	//获取token
	var token models.Token
	json.Unmarshal(this.Ctx.Input.RequestBody, &token)

	if token.Token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}

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
	//获取token
	var token models.Token
	json.Unmarshal(this.Ctx.Input.RequestBody, &token)

	if token.Token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}
	//id
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("get id err", err.Error())
		this.ajaxMsg("get id err", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	//status
	status := this.GetString("status")
	if status == "" {
		fmt.Println("status 不能为空")
	}
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
	//获取token
	var token models.Token
	json.Unmarshal(this.Ctx.Input.RequestBody, &token)

	if token.Token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}
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
