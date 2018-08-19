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

type JsbkController struct {
	BaseController
}

func (this *JsbkController) CaseAddAction() {
	fmt.Println("add case")
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
	fmt.Println("get case data")
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
	cases := new(models.Case)
	query := o.QueryTable(cases)
	filters := make([]interface{}, 0)
	//year
	year := this.Input().Get("year")
	if year != "" {
		filters = append(filters, "Year", year)
	}
	//courseName
	course := this.Input().Get("courseName")
	if course != "" {
		filters = append(filters, "CourseName", course)
	}
	//caseName
	caseName := this.Input().Get("caseName")
	if caseName != "" {
		filters = append(filters, "CaseName", caseName)
	}
	//status
	status := this.Input().Get("status")
	if status != "" {
		filters = append(filters, "Status", status)
	}

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
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
		fmt.Println("get case err", err.Error())
		this.ajaxMsg("get case err", MSG_ERR_Resources)
	}
	fmt.Println("get case reslut num:", num)
	this.ajaxList("get case data success", MSG_OK, num, maps)
}

func (this *JsbkController) EditAction() {
	fmt.Println("edit action")
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

	var case_info models.Case
	json.Unmarshal(this.Ctx.Input.RequestBody, &case_info)
	fmt.Println("case_info:", &case_info)
	id := case_info.Id
	status := case_info.Status
	if id == 0 || status == "" {
		this.ajaxMsg("更新失败,id或者status不能为空", MSG_ERR_Param)
	}

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

	var case_info models.Case
	json.Unmarshal(this.Ctx.Input.RequestBody, &case_info)
	fmt.Println("case_info:", &case_info)
	id := case_info.Id
	if id == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
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

func (this *JsbkController) TeachGroupAddAction() {
	fmt.Println("add group")
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
	fmt.Println("get teachGroup data")
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
	group := new(models.TeachGroup)
	query := o.QueryTable(group)
	filters := make([]interface{}, 0)

	//teachName
	groupname := this.Input().Get("name")
	if groupname != "" {
		filters = append(filters, "Name", groupname)
	}

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
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
		fmt.Println("get group err", err.Error())
		this.ajaxMsg("get group err", MSG_ERR_Resources)
	}
	fmt.Println("get group reslut num:", num)
	this.ajaxList("get group data success", MSG_OK, num, maps)
	return
}

func (this *JsbkController) TeachGroupEditAction() {
	fmt.Println("edit action")
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

	var group_info models.TeachGroup
	json.Unmarshal(this.Ctx.Input.RequestBody, &group_info)
	fmt.Println("group_info:", &group_info)
	id := group_info.Id
	if id == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
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
