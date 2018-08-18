package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	"math"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type XxkglController struct {
	BaseController
}

func (this *XxkglController) SaveSetting() {
	fmt.Println("add xk_setting")
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
	var setting models.XkSetting
	json.Unmarshal(this.Ctx.Input.RequestBody, &setting)
	fmt.Println("xk_setting_info:", &setting)
	setting.Id = 1
	//update
	_, err1 := o.Update(&setting)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["Id"] = setting.Id
	this.ajaxList("save success", MSG_OK, 1, list)
	return
}

func (this *XxkglController) GetSetting() {
	//获取token
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
	xk := new(models.XkSetting)
	var maps []orm.Params
	num, err := o.QueryTable(xk).Filter("Id", 1).Values(&maps)
	if err != nil {
		fmt.Println("get xk err", err.Error())
		this.ajaxMsg("get xk err", MSG_ERR_Resources)
	}
	this.ajaxList("get data success", MSG_OK, num, maps)
	return
}

func (this *XxkglController) AddAction() {
	fmt.Println("add task")
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
	var xkrw models.Xkrw
	json.Unmarshal(this.Ctx.Input.RequestBody, &xkrw)

	fmt.Println("xkrw_info:", &xkrw)

	//insert
	_, err1 := o.Insert(&xkrw)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["id"] = xkrw.Id
	list["taskId"] = xkrw.TaskId
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *XxkglController) GetData() {
	fmt.Println("get xkrw data")
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
	xkrw := new(models.Xkrw)
	query := o.QueryTable(xkrw)

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
	num, err := query.OrderBy("-Id").Values(&maps)
	if err != nil {
		fmt.Println("get xkrw err", err.Error())
		this.ajaxMsg("get xkrw err", MSG_ERR_Resources)
	}
	fmt.Println("get xkrw reslut num:", num)
	this.ajaxList("get xkrw data success", MSG_OK, count, maps)
	return
}

func (this *XxkglController) EditAction() {
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

func (this *XxkglController) Del() {
	fmt.Println("del xxkgl")
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

	var xk_info models.Xkrw
	json.Unmarshal(this.Ctx.Input.RequestBody, &xk_info)
	fmt.Println("xk_info:", &xk_info)
	id := xk_info.Id
	if id == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	o := orm.NewOrm()
	xk := new(models.Xkrw)
	num, err := o.QueryTable(xk).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del xkrw err", err.Error())
		this.ajaxMsg("del xkrw err", MSG_ERR_Resources)
	}
	fmt.Println("del xkrw reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del xkrw success", MSG_OK)
	return
}

func (this *XxkglController) OpenClassAddAction() {
	fmt.Println("add openclass apply")
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
	var kksq models.Kksq
	json.Unmarshal(this.Ctx.Input.RequestBody, &kksq)

	fmt.Println("kksq_info:", &kksq)

	//insert
	kksq.Status = "未审核"
	_, err1 := o.Insert(&kksq)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["Id"] = kksq.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}
func (this *XxkglController) OpenClassGetData() {
	fmt.Println("get kksq data")
	o := orm.NewOrm()
	var maps []orm.Params
	kksq := new(models.Kksq)
	query := o.QueryTable(kksq)
	//get data dB
	num, err := query.Values(&maps)
	if err != nil {
		fmt.Println("get kksq err", err.Error())
		this.ajaxMsg("get kksq err", MSG_ERR_Resources)
	}
	fmt.Println("get kksq reslut num:", num)
	this.ajaxList("get kksq data success", MSG_OK, num, maps)
	return
}

func (this *XxkglController) OpenClassChangeStatus() {
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

	var kksq_info models.Kksq
	json.Unmarshal(this.Ctx.Input.RequestBody, &kksq_info)
	fmt.Println("kksq_info:", &kksq_info)
	id := kksq_info.Id
	status := kksq_info.Status
	if id == 0 || status == "" {
		this.ajaxMsg("更新失败,id或者status不能为空", MSG_ERR_Param)
	}
	o := orm.NewOrm()
	kksq := new(models.Kksq)
	//updata status db
	num, err := o.QueryTable(kksq).Filter("Id", id).Update(orm.Params{
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

func (this *XxkglController) OpenClassDel() {
	fmt.Println("del xxkgl")
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

	var xk_info models.Kksq
	json.Unmarshal(this.Ctx.Input.RequestBody, &xk_info)
	fmt.Println("xk_info:", &xk_info)
	id := xk_info.Id
	if id == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	o := orm.NewOrm()
	xk := new(models.Kksq)
	num, err := o.QueryTable(xk).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del kksq err", err.Error())
		this.ajaxMsg("del kksq err", MSG_ERR_Resources)
	}
	fmt.Println("del kksq reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del kksq success", MSG_OK)
	return
}
