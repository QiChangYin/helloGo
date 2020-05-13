package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 数据库
func Init() {
	logs.Info(">>>> database init start <<<<")

	//  数据库驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//// 注册数据库
	//username := beego.AppConfig.String("username")
	//password := beego.AppConfig.String("password")
	//url := beego.AppConfig.String("url")
	//dbname := beego.AppConfig.String("dbname")
	//datasource := "mysql://" + username + ":" + password + "@" + url + "/" + dbname
	//orm.RegisterDataBase("mysql","mysql",datasource)
	//// 最大连接和空闲连接
	//orm.SetMaxOpenConns("mysql",30)
	//orm.SetMaxIdleConns("mysql",10)
	//// 注册model
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//
	//orm.RegisterModel(new(User))
	//
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(123.123.123.123:3306)/test?charset=utf8")
	//————————————————
	//版权声明：本文为CSDN博主「loongshawn」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
	//原文链接：https://blog.csdn.net/loongshawn/article/details/54973203
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/helloGo?charset=utf8")
	//o := orm.NewOrm()
	orm.RegisterModel(new(User),new(Result))
	//configInfo := beego.AppConfig.String("MYSQL::helloGo")
	//o.Using(configInfo)
	//orm.RunSyncdb("default", true, true
	//dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	//orm.RegisterDataBase("default", "mysql", dsn)
	//orm.RegisterModel(new(User), new(Result))

}

