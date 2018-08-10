package controllers

import (
	"ClassCenter/models"
	"encoding/json"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type SksjController struct {
	BaseController
}

func (this *SksjController) GetData() {
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

	style, err := this.GetInt("style")
	if err != nil {
		fmt.Println("get sksj style err", err.Error())
	}
	o := orm.NewOrm()
	var maps []orm.Params
	kbsz := new(models.Kbsz)
	num, err := o.QueryTable(kbsz).Filter("Style", style).Values(&maps)
	if err != nil {
		fmt.Println("get kbsz err", err.Error())
		this.ajaxMsg("get kbsz err", MSG_ERR_Resources)
	}
	fmt.Println("get kbsz reslut num:", num)
	this.ajaxList("get kbsz data success", MSG_OK, num, maps)
	return
}

func (this *SksjController) EditAction() {

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

	fmt.Println("edit action")
	o := orm.NewOrm()
	var kbsz models.Kbsz
	json.Unmarshal(this.Ctx.Input.RequestBody, &kbsz)
	fmt.Println("kbsz_info:", &kbsz)
	//updata kbsz db
	_, err1 := o.Update(&kbsz)
	if err1 != nil {
		fmt.Println("updata kbsz err", err1.Error())
		this.ajaxMsg("updata kbsz err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update kbsz success", MSG_OK)
	return
}

func (this *SksjController) Save() {
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

	fmt.Println("edit action")
	o := orm.NewOrm()
	var kbsz models.KbszSetting
	json.Unmarshal(this.Ctx.Input.RequestBody, &kbsz)
	fmt.Println("kbsz_info:", &kbsz)
	//updata kbsz db
	_, err1 := o.Update(&kbsz)
	if err1 != nil {
		fmt.Println("updata kbsz err", err1.Error())
		this.ajaxMsg("updata kbsz err", MSG_ERR_Resources)
	}
	this.ajaxMsg("update kbsz success", MSG_OK)
	return
}
