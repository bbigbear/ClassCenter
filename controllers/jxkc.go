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

type JxkcController struct {
	BaseController
}

func (this *JxkcController) AddAction() {
	fmt.Println("add course")
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
	kc := new(models.Jxkc)
	query := o.QueryTable(kc)
	filters := make([]interface{}, 0)
	//name
	courseName := this.Input().Get("name")
	if courseName != "" {
		filters = append(filters, "CourseName", courseName)
	}
	fmt.Println("get name:", courseName)
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

	//get data dB
	_, err1 := query.OrderBy("-Id").Values(&maps)
	if err1 != nil {
		fmt.Println("get jxkc err", err.Error())
		this.ajaxMsg("get jxkc err", MSG_ERR_Resources)
	}
	this.ajaxList("get jxkc data success", MSG_OK, count, maps)
	return
}

func (this *JxkcController) EditAction() {
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
	var kc models.Jxkc
	json.Unmarshal(this.Ctx.Input.RequestBody, &kc)
	fmt.Println("kc_info:", &kc)
	//updata jxkc db
	//time
	nowtime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("time err", err.Error())
	}
	kc.CreateTime = nowtime
	_, err1 := o.Update(&kc)
	if err1 != nil {
		fmt.Println("updata jxkc err", err1.Error())
		this.ajaxMsg("updata jxkc err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update jxkc success", MSG_OK)
	return
}

func (this *JxkcController) ChangeStatus() {
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

	var kc_info models.Jxkc
	json.Unmarshal(this.Ctx.Input.RequestBody, &kc_info)
	fmt.Println("kc_info:", &kc_info)
	id := kc_info.Id
	status := kc_info.Status
	if id == 0 || status == "" {
		this.ajaxMsg("更新失败,id或者status不能为空", MSG_ERR_Param)
	}
	o := orm.NewOrm()
	kc := new(models.Jxkc)
	//updata status db
	num, err := o.QueryTable(kc).Filter("Id", id).Update(orm.Params{
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

func (this *JxkcController) Del() {
	fmt.Println("del jxkc")
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

	var kc_info models.Jxkc
	json.Unmarshal(this.Ctx.Input.RequestBody, &kc_info)
	fmt.Println("kc_info:", &kc_info)
	id := kc_info.Id
	if id == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	o := orm.NewOrm()
	kc := new(models.Jxkc)
	num, err := o.QueryTable(kc).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del jxkc err", err.Error())
		this.ajaxMsg("del jxkc err", MSG_ERR_Resources)
	}
	fmt.Println("del train reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del jxkc success", MSG_OK)
	return
}
