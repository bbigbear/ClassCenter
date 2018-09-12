package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	"math"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type JstkController struct {
	BaseController
}

func (this *JstkController) RyapAdd() {
	fmt.Println("人员安排-添加记录")
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
	var jstk models.Jstk
	json.Unmarshal(this.Ctx.Input.RequestBody, &jstk)

	fmt.Println("jstk_info:", &jstk)
	//insert
	_, err1 := o.Insert(&jstk)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["id"] = jstk.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *JsbkController) RyapAddPeople() {
	fmt.Println("add staff")
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

	var jstk_info models.Jstk
	json.Unmarshal(this.Ctx.Input.RequestBody, &jstk_info)
	fmt.Println("jstk_info:", &jstk_info)
	id := jstk_info.Id
	staff := jstk_info.Staffing
	if id == 0 || staff == "" {
		this.ajaxMsg("更新失败,id或者staff不能为空", MSG_ERR_Param)
	}

	o := orm.NewOrm()
	jstk := new(models.Jstk)
	//updata status db
	num, err := o.QueryTable(jstk).Filter("Id", id).Update(orm.Params{
		"Staffing": staff,
	})
	if err != nil {
		fmt.Println("add staff err", err.Error())
		this.ajaxMsg("add staff err", MSG_ERR_Resources)
	}
	fmt.Println("num", num)
	this.ajaxMsg("add staff success", MSG_OK)
	return
}

func (this *JstkController) RyapGetData() {
	fmt.Println("get jstk data")
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
	jstk := new(models.Jstk)
	query := o.QueryTable(jstk)
	filters := make([]interface{}, 0)
	//OrderName
	orderName := this.Input().Get("orderName")
	if orderName != "" {
		filters = append(filters, "OrderName", orderName)
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
		fmt.Println("get jstk err", err.Error())
		this.ajaxMsg("get jstk err", MSG_ERR_Resources)
	}
	fmt.Println("get jstk reslut num:", num)
	this.ajaxList("get jstk data success", MSG_OK, num, maps)
}

func (this *JstkController) RyapDel() {
	fmt.Println("del jstk")
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

	var jstk_info models.Jstk
	json.Unmarshal(this.Ctx.Input.RequestBody, &jstk_info)
	fmt.Println("jstk_info:", &jstk_info)
	id := jstk_info.Id
	if id == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
	o := orm.NewOrm()
	jstk := new(models.Jstk)
	num, err := o.QueryTable(jstk).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del jstk err", err.Error())
		this.ajaxMsg("del jstk err", MSG_ERR_Resources)
	}
	fmt.Println("del jstk reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del jstk success", MSG_OK)
	return
}

//调课申请
func (this *JstkController) TkAdd() {
	fmt.Println("调课-添加记录")
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
	var tksq models.Tksq
	json.Unmarshal(this.Ctx.Input.RequestBody, &tksq)
	fmt.Println("tksq_info:", &tksq)
	tksq.Status = "已提交 未审核"
	//insert
	_, err1 := o.Insert(&tksq)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["id"] = tksq.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}
