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

type XzzglController struct {
	BaseController
}

func (this *XzzglController) AddAction() {
	fmt.Println("add case")
	var member string
	//获取token
	var token models.Token
	json.Unmarshal(this.Ctx.Input.RequestBody, &token)

	if token.Token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}

	name, err := this.Token_auth(token.Token, "ximi")
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)
	o := orm.NewOrm()
	list := make(map[string]interface{})
	//获取成员
	var group models.Group
	json.Unmarshal(this.Ctx.Input.RequestBody, &group)
	fmt.Println("group:", &group)
	l := len(group.GroupMember)
	for i := 0; i < l; i++ {
		member += group.GroupMember[i] + ","
	}
	var xzz models.Xzz
	json.Unmarshal(this.Ctx.Input.RequestBody, &xzz)
	fmt.Println("xzz_info:", &xzz)
	//time
	nowtime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("新建失败")
	}
	xzz.CreateTime = nowtime
	//member
	xzz.GroupMember = member
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
	//token
	token := this.Input().Get("token")

	if token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}

	name, err := this.Token_auth(token, "ximi")
	if err != nil {
		fmt.Println("token err", err.Error())
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)
	o := orm.NewOrm()
	var maps []orm.Params
	xzz := new(models.Xzz)
	query := o.QueryTable(xzz)
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
		fmt.Println("get xzz err", err.Error())
		this.ajaxMsg("get xzz err", MSG_ERR_Resources)
	}
	fmt.Println("get xzz reslut num:", num)
	this.ajaxList("get xzz data success", 0, num, maps)
	return
}

func (this *XzzglController) EditAction() {
	fmt.Println("edit action")
	var member string
	//获取token
	var token models.Token
	json.Unmarshal(this.Ctx.Input.RequestBody, &token)

	if token.Token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}

	name, err := this.Token_auth(token.Token, "ximi")
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)

	o := orm.NewOrm()
	var xzz models.Xzz
	//获取成员
	var group models.Group
	json.Unmarshal(this.Ctx.Input.RequestBody, &group)
	fmt.Println("group:", &group)
	l := len(group.GroupMember)
	for i := 0; i < l; i++ {
		member += group.GroupMember[i] + ","
	}
	xzz.GroupMember = member
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
	//获取token
	var token models.Token
	json.Unmarshal(this.Ctx.Input.RequestBody, &token)

	if token.Token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}

	name, err := this.Token_auth(token.Token, "ximi")
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)
	var xzz_info models.Xzz
	json.Unmarshal(this.Ctx.Input.RequestBody, &xzz_info)
	fmt.Println("xzz_info:", &xzz_info)
	id := xzz_info.Id
	if id == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
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

func (this *XzzglController) XzztlAdd() {
	fmt.Println("add xzztl")
	//获取token
	var token models.Token
	json.Unmarshal(this.Ctx.Input.RequestBody, &token)

	if token.Token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}

	name, err := this.Token_auth(token.Token, "ximi")
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)
	o := orm.NewOrm()
	list := make(map[string]interface{})
	var xzztl models.Xzztl
	json.Unmarshal(this.Ctx.Input.RequestBody, &xzztl)

	fmt.Println("xzztl_info:", &xzztl)
	//time
	nowtime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("新建失败")
	}
	xzztl.CreateTime = nowtime
	//insert
	_, err1 := o.Insert(&xzztl)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["id"] = xzztl.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}
