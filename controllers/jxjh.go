package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"
	"math"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type JxjhController struct {
	BaseController
}

func (this *JxjhController) AddAction() {
	fmt.Println("add plan")
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
	var jh models.Jxjh
	json.Unmarshal(this.Ctx.Input.RequestBody, &jh)

	fmt.Println("xjjh_info:", &jh)

	//time
	//	nowtime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	//	if err != nil {
	//		fmt.Println("time err", err.Error())
	//	}
	jh.Status = "未审核"

	//insert
	_, err1 := o.Insert(&jh)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["planId"] = jh.PlanId
	list["id"] = jh.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

func (this *JxjhController) GetData() {
	fmt.Println("get jxjh data")
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
	jh := new(models.Jxjh)
	query := o.QueryTable(jh)
	filters := make([]interface{}, 0)
	//major
	major := this.Input().Get("major")
	if major != "" {
		filters = append(filters, "Major", major)
	}
	//grade
	grade := this.Input().Get("grade")
	if grade != "" {
		filters = append(filters, "PlanGrade", grade)
	}
	//class
	class := this.Input().Get("class")
	if class != "" {
		filters = append(filters, "PlanClass", class)
	}
	fmt.Println("get class:", class)
	//planId
	planId := this.Input().Get("planId")
	if planId != "" {
		filters = append(filters, "PlanId", planId)
	}
	fmt.Println("get planId:", planId)

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
	num, err := query.OrderBy("-Id").Values(&maps)
	if err != nil {
		fmt.Println("get jxjh err", err.Error())
		this.ajaxMsg("get jxjh err", MSG_ERR_Resources)
	}
	fmt.Println("get jxjh reslut num:", num)
	this.ajaxList("get jxjh data success", MSG_OK, count, maps)
	return
}

func (this *JxjhController) EditAction() {
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
	var jh models.Jxjh
	json.Unmarshal(this.Ctx.Input.RequestBody, &jh)
	fmt.Println("jh_info:", &jh)
	//updata jxjh db
	_, err1 := o.Update(&jh)
	if err1 != nil {
		fmt.Println("updata jxjh err", err1.Error())
		this.ajaxMsg("updata jxjh err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update jxjh success", MSG_OK)
	return
}

func (this *JxjhController) ChangeStatus() {
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

	var jxjh_info models.Jxjh
	json.Unmarshal(this.Ctx.Input.RequestBody, &jxjh_info)
	fmt.Println("jxjh_info:", &jxjh_info)
	id := jxjh_info.Id
	status := jxjh_info.Status
	if id == 0 || status == "" {
		this.ajaxMsg("更新失败,id或者status不能为空", MSG_ERR_Param)
	}
	o := orm.NewOrm()
	jh := new(models.Jxjh)
	//updata status db
	num, err := o.QueryTable(jh).Filter("Id", id).Update(orm.Params{
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

func (this *JxjhController) Del() {
	fmt.Println("del jxjh")
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

	var jxjh_info models.Jxjh
	json.Unmarshal(this.Ctx.Input.RequestBody, &jxjh_info)
	fmt.Println("jxjh_info:", &jxjh_info)
	id := jxjh_info.Id
	if id == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
	o := orm.NewOrm()
	jh := new(models.Jxjh)
	num, err := o.QueryTable(jh).Filter("Id", id).Delete()
	if err != nil {
		fmt.Println("del jxjh err", err.Error())
		this.ajaxMsg("del jxjh err", MSG_ERR_Resources)
	}
	fmt.Println("del train reslut num:", num)
	//list["data"] = maps
	this.ajaxMsg("del jxjh success", MSG_OK)
	return
}

//分配
func (this *JxjhController) JhkcAddAction() {
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
	var jxrw models.Jxrw
	json.Unmarshal(this.Ctx.Input.RequestBody, &jxrw)
	//courseid
	if jxrw.CourseId == "" {
		fmt.Println("courseId is null")
		this.ajaxMsg("courseId is null", MSG_ERR_Param)
	}
	//planid
	if jxrw.PlanId == "" {
		fmt.Println("planId is null")
		this.ajaxMsg("planId is null", MSG_ERR_Param)
	}
	list := make(map[string]interface{})
	jxrw.Status = "生成成功"
	_, err1 := o.Insert(&jxrw)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["id"] = jxrw.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}

//编辑
func (this *JxjhController) JhkcEditAction() {
	fmt.Println("edit course")
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
	var jxrw models.Jxrw
	json.Unmarshal(this.Ctx.Input.RequestBody, &jxrw)
	//courseid
	if jxrw.CourseId == "" {
		fmt.Println("courseId is null")
		this.ajaxMsg("courseId is null", MSG_ERR_Param)
	}
	//planid
	if jxrw.PlanId == "" {
		fmt.Println("planId is null")
		this.ajaxMsg("planId is null", MSG_ERR_Param)
	}
	_, err1 := o.Update(&jxrw)
	if err1 != nil {
		fmt.Printf("update err", err1.Error())
		this.ajaxMsg("update err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update success", MSG_OK)
}

//查看
func (this *JxjhController) JhkcLook() {
	fmt.Println("add course")
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

	//planId为空
	planId := this.Input().Get("planId")
	if planId == "" {
		fmt.Println("planId 为空")
		this.ajaxMsg("planId 不能为空", MSG_ERR_Param)
	}
	o := orm.NewOrm()
	jxrw := new(models.Jxrw)
	var maps []orm.Params
	num, err := o.QueryTable(jxrw).Filter("PlanId", planId).Values(&maps)
	if err != nil {
		fmt.Println("get jhkc data err", err.Error())
	}
	this.ajaxList("get data success", MSG_OK, num, maps)
}

//获取
func (this *JxjhController) JhkcGetData() {
	fmt.Println("get jhkc data")
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
	courseName := this.Input().Get("courseName")
	if courseName != "" {
		filters = append(filters, "CourseName", courseName)
	}
	//term
	term := this.Input().Get("term")
	if term != "" {
		filters = append(filters, "Term", term)
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

	//get data dB
	num, err := query.OrderBy("-Id").Values(&maps)
	if err != nil {
		fmt.Println("get jhkc err", err.Error())
		this.ajaxMsg("get jhkc err", MSG_ERR_Resources)
	}
	fmt.Println("get jhkc reslut num:", num)
	this.ajaxList("get jhkc data success", MSG_OK, count, maps)
	return
}

//分配任务
func (this *JxjhController) JxrwTeacherAllotSave() {
	fmt.Println("add teacher")
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
	var jxrw models.Jxrw_teacher_allot
	json.Unmarshal(this.Ctx.Input.RequestBody, &jxrw)
	//TeacherId
	if jxrw.TeacherId == "" {
		fmt.Println("TeacherId is null")
		this.ajaxMsg("TeacherId is null", MSG_ERR_Param)
	}
	//planid
	if jxrw.PlanId == "" {
		fmt.Println("planId is null")
		this.ajaxMsg("planId is null", MSG_ERR_Param)
	}
	list := make(map[string]interface{})
	jxrw.Status = "已分配"
	_, err1 := o.Insert(&jxrw)
	if err1 != nil {
		fmt.Printf("insert err", err1.Error())
		this.ajaxMsg("insert err", MSG_ERR_Resources)
	}
	list["id"] = jxrw.Id
	this.ajaxList("add success", MSG_OK, 1, list)
	return
}
