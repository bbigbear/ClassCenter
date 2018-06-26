package controllers

//import (
//	"ClassCenter/models"
//	"encoding/json"
//	"fmt"
//	"time"

//	_ "github.com/Go-SQL-Driver/MySQL"
//	"github.com/astaxie/beego/orm"
//)

type KbglController struct {
	BaseController
}

func (this *KbglController) TeacherGet() {
	this.TplName = "kbgl_teacher.tpl"
}

func (this *KbglController) ClassRoomGet() {
	this.TplName = "kbgl_classroom.tpl"
}

func (this *KbglController) TeachingClassGet() {
	this.TplName = "kbgl_teachingclass.tpl"
}
