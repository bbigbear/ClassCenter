package routers

import (
	"classCenter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/v1/put_file", &controllers.BaseController{}, "*:PutFile")
	//beego.Router("/v1/put_dbf_file", &controllers.BaseController{}, "*:PutDbf")
	beego.Router("/v1/login", &controllers.LoginController{}, "post:LoginAction")
	//jxkc
	beego.Router("/v1/jxkc/add_action", &controllers.JxkcController{}, "post:AddAction")
	beego.Router("/v1/jxkc/getdata", &controllers.JxkcController{}, "Get:GetData")
	beego.Router("/v1/jxkc/del", &controllers.JxkcController{}, "post:Del")
	beego.Router("/v1/jxkc/edit_action", &controllers.JxkcController{}, "post:EditAction")
	beego.Router("/v1/jxkc/change", &controllers.JxkcController{}, "post:ChangeStatus")
	//jxjh
	beego.Router("/v1/jxjh/jhkc/allot", &controllers.JxjhController{}, "post:JhkcAddAction")
	beego.Router("/v1/jxjh/jhkc/edit_action", &controllers.JxjhController{}, "post:JhkcEditAction")
	beego.Router("/v1/jxjh/jhkc/getdata", &controllers.JxjhController{}, "*:JhkcGetData")
	beego.Router("/v1/jxjh/jxrw/allot", &controllers.JxjhController{}, "post:JxrwTeacherAllotSave")
	beego.Router("/v1/jxjh/add_action", &controllers.JxjhController{}, "post:AddAction")
	beego.Router("/v1/jxjh/getdata", &controllers.JxjhController{}, "Get:GetData")
	beego.Router("/v1/jxjh/del", &controllers.JxjhController{}, "post:Del")
	beego.Router("/v1/jxjh/edit_action", &controllers.JxjhController{}, "post:EditAction")
	beego.Router("/v1/jxjh/change", &controllers.JxjhController{}, "post:ChangeStatus")
	beego.Router("/v1/jxjh/jhkc/getcoursedata", &controllers.JxjhController{}, "*:JhkcLook")

	//pkgl
	beego.Router("/v1/pkgl/setting", &controllers.PkglController{}, "*:SettingGet")
	beego.Router("/v1/pkgl/setting/save", &controllers.PkglController{}, "post:SettingSave")
	beego.Router("/v1/pkgl/classroom", &controllers.PkglController{}, "*:ClassRoomUnarrangingTimeGet")
	beego.Router("/v1/pkgl/teacher", &controllers.PkglController{}, "*:TeacherUnarrangingTimeGet")
	beego.Router("/v1/pkgl/eletive", &controllers.PkglController{}, "*:EletivePkGet")
	beego.Router("/v1/pkgl/compulsory", &controllers.PkglController{}, "*:CompulsoryPkGet")
	//sksj
	beego.Router("/v1/sksj/getdata", &controllers.SksjController{}, "*:GetData")
	beego.Router("/v1/sksj/edit_action", &controllers.SksjController{}, "post:EditAction")
	beego.Router("/v1/sksj/save", &controllers.SksjController{}, "*:Save")
	beego.Router("/v1/sksj/getSetting", &controllers.SksjController{}, "*:GetSetting")

	//xxkgl
	beego.Router("/v1/xxkgl/setting/getdata", &controllers.XxkglController{}, "*:GetSetting")
	beego.Router("/v1/xxkgl/setting/save", &controllers.XxkglController{}, "post:SaveSetting")
	beego.Router("/v1/xxkgl/task/add_action", &controllers.XxkglController{}, "post:AddAction")
	beego.Router("/v1/xxkgl/task/edit_action", &controllers.XxkglController{}, "*:EditAction")
	beego.Router("/v1/xxkgl/task/del", &controllers.XxkglController{}, "post:Del")
	beego.Router("/v1/xxkgl/task/getdata", &controllers.XxkglController{}, "*:GetData")
	beego.Router("/v1/xxkgl/openclass_add_action", &controllers.XxkglController{}, "post:OpenClassAddAction")
	beego.Router("/v1/xxkgl/openclass/getdata", &controllers.XxkglController{}, "*:OpenClassGetData")
	beego.Router("/v1/xxkgl/openclass/change", &controllers.XxkglController{}, "post:OpenClassChangeStatus")
	beego.Router("/v1/xxkgl/openclass/del", &controllers.XxkglController{}, "post:OpenClassDel")
	//kbgl
	beego.Router("/v1/kbgl/classroom", &controllers.KbglController{}, "*:ClassRoomGet")
	beego.Router("/v1/kbgl/teacher", &controllers.KbglController{}, "*:TeacherGet")
	beego.Router("/v1/kbgl/teachingclass", &controllers.KbglController{}, "*:TeachingClassGet")
	//jsbk
	beego.Router("/v1/jsbk/case", &controllers.JsbkController{}, "*:CaseGet")
	beego.Router("/v1/jsbk/case/add", &controllers.JsbkController{}, "*:CaseAdd")
	beego.Router("/v1/jsbk/case/add_action", &controllers.JsbkController{}, "post:CaseAddAction")
	beego.Router("/v1/jsbk/case/getdata", &controllers.JsbkController{}, "*:CaseGetData")
	beego.Router("/v1/jsbk/case/edit", &controllers.JsbkController{}, "*:Edit")
	beego.Router("/v1/jsbk/case/edit_action", &controllers.JsbkController{}, "post:EditAction")
	beego.Router("/v1/jsbk/case/del", &controllers.JsbkController{}, "post:Del")
	beego.Router("/v1/jsbk/case/change", &controllers.JsbkController{}, "*:ChangeStatus")
	beego.Router("/v1/jsbk/group", &controllers.JsbkController{}, "*:TeachGroup")
	beego.Router("/v1/jsbk/group/add", &controllers.JsbkController{}, "*:TeachGroupAdd")
	beego.Router("/v1/jsbk/group/add_action", &controllers.JsbkController{}, "post:TeachGroupAddAction")
	beego.Router("/v1/jsbk/group/edit", &controllers.JsbkController{}, "*:TeachGroupEdit")
	beego.Router("/v1/jsbk/group/edit_action", &controllers.JsbkController{}, "post:TeachGroupEditAction")
	beego.Router("/v1/jsbk/group/del", &controllers.JsbkController{}, "post:TeachGroupDel")
	beego.Router("/v1/jsbk/group/getdata", &controllers.JsbkController{}, "*:TeachGroupGetData")
	beego.Router("/v1/jsbk/group/addpeople", &controllers.JsbkController{}, "*:TeachGroupAddPepole")
	beego.Router("/v1/jsbk/group/addleader", &controllers.JsbkController{}, "*:TeachGroupAddLeader")
	beego.Router("/v1/jsbk/case/approve", &controllers.JsbkController{}, "*:TeachCaseApprove")
	beego.Router("/v1/jsbk/case/query", &controllers.JsbkController{}, "*:TeachCaseQuery")
	beego.Router("/v1/jsbk/time_manage", &controllers.JsbkController{}, "*:TimeManage")
	beego.Router("/v1/jsbk/schedule", &controllers.JsbkController{}, "*:TeachSchedule")
	//jbgl
	beego.Router("/v1/jbgl/plan", &controllers.JbglController{}, "*:Get")
	beego.Router("/v1/jbgl/plan/add", &controllers.JbglController{}, "*:Add")
	beego.Router("/v1/jbgl/plan/add_action", &controllers.JbglController{}, "post:AddAction")
	beego.Router("/v1/jbgl/plan/getdata", &controllers.JbglController{}, "*:GetData")
	beego.Router("/v1/jbgl/plan/edit", &controllers.JbglController{}, "*:Edit")
	beego.Router("/v1/jbgl/plan/edit_action", &controllers.JbglController{}, "post:EditAction")
	beego.Router("/v1/jbgl/plan/del", &controllers.JbglController{}, "post:Del")
	beego.Router("/v1/jbgl/plan/wait", &controllers.JbglController{}, "*:WaitLaunch")
	beego.Router("/v1/jbgl/plan/launch", &controllers.JbglController{}, "*:Launch")
	beego.Router("/v1/jbgl/plan/join", &controllers.JbglController{}, "*:Join")
	beego.Router("/v1/jbgl/plan/view", &controllers.JbglController{}, "*:View")
	beego.Router("/v1/jbjc/plan/launch", &controllers.JbglController{}, "*:Check")
	beego.Router("/v1/jbjc/plan/join", &controllers.JbglController{}, "*:CheckJoin")
	beego.Router("/v1/jbjc/plan/view", &controllers.JbglController{}, "*:CheckView")
	//xzz
	beego.Router("/v1/xzz", &controllers.XzzglController{}, "*:Get")
	beego.Router("/v1/xzz/add", &controllers.XzzglController{}, "*:Add")
	beego.Router("/v1/xzz/add_action", &controllers.XzzglController{}, "post:AddAction")
	beego.Router("/v1/xzz/getdata", &controllers.XzzglController{}, "*:GetData")
	beego.Router("/v1/xzz/edit", &controllers.XzzglController{}, "*:Edit")
	beego.Router("/v1/xzz/edit_action", &controllers.XzzglController{}, "post:EditAction")
	beego.Router("/v1/xzz/del", &controllers.XzzglController{}, "post:Del")
	beego.Router("/v1/xzz_discuss", &controllers.XzzglController{}, "*:GetDiscuss")

}
