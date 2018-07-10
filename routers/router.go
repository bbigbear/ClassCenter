package routers

import (
	"classCenter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/put_file", &controllers.BaseController{}, "*:PutFile")
	//beego.Router("/v1/put_dbf_file", &controllers.BaseController{}, "*:PutDbf")
	//jxkc
	beego.Router("/v1/jxkc/apply", &controllers.JxkcController{}, "*:ApplyGet")
	beego.Router("/v1/jxkc/check", &controllers.JxkcController{}, "*:CheckGet")
	beego.Router("/v1/jxkc/maintain", &controllers.JxkcController{}, "*:MaintainGet")
	beego.Router("/v1/jxkc/add", &controllers.JxkcController{}, "*:Add")
	beego.Router("/v1/jxkc/add_action", &controllers.JxkcController{}, "post:AddAction")
	beego.Router("/v1/jxkc/getdata", &controllers.JxkcController{}, "Get:GetData")
	beego.Router("/v1/jxkc/look", &controllers.JxkcController{}, "*:Look")
	beego.Router("/v1/jxkc/edit", &controllers.JxkcController{}, "*:Edit")
	beego.Router("/v1/jxkc/del", &controllers.JxkcController{}, "post:Del")
	beego.Router("/v1/jxkc/edit_action", &controllers.JxkcController{}, "post:EditAction")
	beego.Router("/v1/jxkc/change", &controllers.JxkcController{}, "post:ChangeStatus")
	//jxjh
	beego.Router("/v1/jxjh/apply", &controllers.JxjhController{}, "*:ApplyGet")
	beego.Router("/v1/jxjh/check", &controllers.JxjhController{}, "*:CheckGet")
	beego.Router("/v1/jxjh/maintain", &controllers.JxjhController{}, "*:MaintainGet")
	beego.Router("/v1/jxjh/jxkc/allot", &controllers.JxjhController{}, "*:JhkcAllotGet")
	beego.Router("/v1/jxjh/jxkc/allot_save", &controllers.JxjhController{}, "post:JhkcAllotSave")
	beego.Router("/v1/jxjh/jxkc/add", &controllers.JxjhController{}, "*:JhkcAdd")
	beego.Router("/v1/jxjh/jxrw/teacher_allot", &controllers.JxjhController{}, "*:JxrwTeacherAllotGet")
	beego.Router("/v1/jxjh/jxrw/teacher_allot_save", &controllers.JxjhController{}, "post:JxrwTeacherAllotSave")
	beego.Router("/v1/jxjh/jxrw/teacher_add", &controllers.JxjhController{}, "*:JxrwTeacherAdd")
	beego.Router("/v1/jxjh/add", &controllers.JxjhController{}, "*:Add")
	beego.Router("/v1/jxjh/add_action", &controllers.JxjhController{}, "post:AddAction")
	beego.Router("/v1/jxjh/getdata", &controllers.JxjhController{}, "Get:GetData")
	beego.Router("/v1/jxjh/look", &controllers.JxjhController{}, "*:Look")
	beego.Router("/v1/jxjh/edit", &controllers.JxjhController{}, "*:Edit")
	beego.Router("/v1/jxjh/del", &controllers.JxjhController{}, "post:Del")
	beego.Router("/v1/jxjh/edit_action", &controllers.JxjhController{}, "post:EditAction")
	beego.Router("/v1/jxjh/change", &controllers.JxjhController{}, "post:ChangeStatus")
	//jxrw
	beego.Router("/v1/jxrw/allot", &controllers.JxjhController{}, "*:JxrwAllotGet")
	beego.Router("/v1/jxrw/generate", &controllers.JxjhController{}, "*:JxrwGenerateGet")
	beego.Router("/v1/jxrw/kc/look", &controllers.JxjhController{}, "*:Jhkclook")
	beego.Router("/v1/jxrw/check", &controllers.JxjhController{}, "*:JxrwCheckGet")
	//pkgl
	beego.Router("/v1/pkgl/setting", &controllers.PkglController{}, "*:SettingGet")
	beego.Router("/v1/pkgl/setting/save", &controllers.PkglController{}, "post:SettingSave")
	beego.Router("/v1/pkgl/classroom", &controllers.PkglController{}, "*:ClassRoomUnarrangingTimeGet")
	beego.Router("/v1/pkgl/teacher", &controllers.PkglController{}, "*:TeacherUnarrangingTimeGet")
	beego.Router("/v1/pkgl/eletive", &controllers.PkglController{}, "*:EletivePkGet")
	beego.Router("/v1/pkgl/compulsory", &controllers.PkglController{}, "*:CompulsoryPkGet")
	//sksj
	beego.Router("/v1/sksj/eletive", &controllers.SksjController{}, "*:EletiveGet")
	beego.Router("/v1/sksj/compulsory", &controllers.SksjController{}, "*:CompulsoryGet")
	beego.Router("/v1/sksj/getdata", &controllers.SksjController{}, "*:GetData")
	beego.Router("/v1/sksj/edit", &controllers.SksjController{}, "*:Edit")
	beego.Router("/v1/sksj/edit_action", &controllers.SksjController{}, "post:EditAction")
	beego.Router("/v1/sksj/save", &controllers.SksjController{}, "*:Save")
	//xxkgl
	beego.Router("/v1/xxkgl/setting", &controllers.XxkglController{}, "*:SettingGet")
	beego.Router("/v1/xxkgl/setting/add", &controllers.XxkglController{}, "post:AddSetting")
	beego.Router("/v1/xxkgl/task", &controllers.XxkglController{}, "*:TaskGet")
	beego.Router("/v1/xxkgl/task/add", &controllers.XxkglController{}, "*:TaskAddGet")
	beego.Router("/v1/xxkgl/task/add_action", &controllers.XxkglController{}, "post:AddAction")
	beego.Router("/v1/xxkgl/task/edit", &controllers.XxkglController{}, "*:Edit")
	beego.Router("/v1/xxkgl/task/edit_action", &controllers.XxkglController{}, "*:EditAction")
	beego.Router("/v1/xxkgl/task/getdata", &controllers.XxkglController{}, "*:GetData")
	beego.Router("/v1/xxkgl/apply", &controllers.XxkglController{}, "*:ApplyGet")
	beego.Router("/v1/xxkgl/openclass_add", &controllers.XxkglController{}, "*:OpenClassAdd")
	beego.Router("/v1/xxkgl/openclass_add_action", &controllers.XxkglController{}, "post:OpenClassAddAction")
	beego.Router("/v1/xxkgl/openclass/getdata", &controllers.XxkglController{}, "*:OpenClassGetData")
	beego.Router("/v1/xxkgl/openclass/change", &controllers.XxkglController{}, "*:OpenClassChangeStatus")
	beego.Router("/v1/xxkgl/approve", &controllers.XxkglController{}, "*:ApproveGet")
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

}
