package main

import (
	"ClassCenter/models"
	_ "classCenter/routers"
	"fmt"

	"github.com/astaxie/beego"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

func init() {
	DBConnection()
	RegisterModel()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

func DBConnection() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	host := beego.AppConfig.String("host")
	db := beego.AppConfig.String("database")
	user := beego.AppConfig.String("user")
	passwd := beego.AppConfig.String("passwd")
	maxOpenConns, err := beego.AppConfig.Int("MaxOpenConns")
	if err != nil {
		fmt.Println("MaxOpenConns is nil", err)
	}
	maxIdleConns, err := beego.AppConfig.Int("MaxIdleConns")
	if err != nil {
		fmt.Println("MaxIdleConns is nil", err)
	}

	sql := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, passwd, host, db)
	orm.RegisterDataBase("default", "mysql", sql, maxIdleConns, maxOpenConns)
}
func RegisterModel() {
	orm.RegisterModel(new(models.Jxkc), new(models.Jxjh),
		new(models.Jxrw), new(models.Jxrw_teacher_allot),
		new(models.XkSetting), new(models.Xkrw))
}
