package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"helloGo/controllers"
	"helloGo/models"
	_ "helloGo/routers"
	"reflect"
)

// 自定义标签
type CustomStruct struct {
	Name     string `customTag:"userName" tags:"user_name"`
	Password string `customTag:"pwd"`
	Age      int    `customTag:"userAge"`
}

func main() {

	cust := CustomStruct{"张三", "123456", 22}
	t := reflect.TypeOf(cust)
	val := reflect.ValueOf(cust)
	for i := 0; i < val.NumField(); i++ {
		println(">>>> field type=" + val.Field(i).Kind().String())
	}

	field, _ := t.FieldByName("Name")
	println(">>>> custom struct field Tag=" + field.Tag + " <<<<")

	println(">>>> struct field name's customTag value=" + field.Tag.Get("customTag") + " <<<<")
	v, _ := field.Tag.Lookup("tags")
	println(">>>> struct field name's tags value=" + v + " <<<<")
	field, _ = t.FieldByName("Password")
	println(">>>> struct field password's tag=" + field.Tag + " <<<<")
	field, _ = t.FieldByName("Age")
	println(">>>> struct field age's tag=" + field.Tag + " <<<<")

	//beego.InsertFilter过滤器是为某些router提供过滤功能，可以设置在某个路由访问时进行额外的操作。通过参数设置可以设置过滤器执行的时刻。
	//如下所示beego.InsertFilter有一下几个参数： InsertFilter(pattern string, pos int, filter FilterFunc, params …bool)
	//pattern：路由规则，过滤器作用的路由
    //pos : 过滤器执行的时刻，有以下5种
	//beego.BeforeStatic
	//beego.BeforeRouter 访问路由之前
	//beego.BeforeExec 访问路由之后执行controller之前
	//beego.AfterExec 执行controller之后调用
	//beego.FinishRouter 结束路由之后调用
	// 过滤器
	beego.InsertFilter("/user/*", beego.BeforeExec, controllers.BeforeExecFilter, false, false)
	beego.InsertFilter("/user/*", beego.BeforeRouter, controllers.BeforeRouterFilter, false, false)
	beego.InsertFilter("/user/*", beego.BeforeStatic, controllers.BeforeStaticFilter, false, false)
	beego.InsertFilter("/user/*", beego.AfterExec, controllers.AfterExecFilter, false, false)
	beego.InsertFilter("/user/*", beego.FinishRouter, controllers.FinishRouterFilter, false, false)

	// run方法前注册错误handler
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}

func init() {

	// 数据库初始化
	models.Init()
	// 初始化session配置
	models.InitSession()

}
