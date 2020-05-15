package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"helloGo/models"
	"regexp"
	"strconv"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) URLMapping() {
	c.Mapping("/login", c.Login)
	c.Mapping("/register", c.Register)
	c.Mapping("/register/do", c.DoRegister)
	c.Mapping("/login/do", c.DoLogin)

}

// @router /login  [get]
func (c *LoginController) Login() {
	logs.Info(">>>> forward to login page start <<<<")
	c.TplName = "login.html"
}

// @router /register  [get]
func (c *LoginController) Register() {
	logs.Info(">>>> forward to Register page start <<<<")

	c.TplName = "register.html"
}

func (c *LoginController) DoRegisterNew() {
	var user models.User
	var result models.Result
	//Json Unmarshal：将json字符串解码到相应的数据结构
	//
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Info(err)
	logs.Info(c.Ctx.Input.RequestBody)
	checkError(err)

	userName := user.UserName
	password := user.Password
	email := user.Email
	logs.Info(">>>> user register information username=" + userName + ",email=" + email + ",password=" + password + " <<<<")

	pass, err := regexp.MatchString(`[a-zA-Z0-9]{8,16}`, password)
	checkError(err)
	if !pass {
		result.Code = 404
		result.Success = false
		result.Message = "密码格式错误"
	}

	em, err := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email)
	checkError(err)
	if !em {
		result.Code = 404
		result.Success = false
		result.Message = "邮箱格式错误"
	}
	// 对密码加密
	user.Password = passwordEncode(password)
	logs.Info("=========================================================================================")
	logs.Info(user.Password)
	id, e := orm.NewOrm().Insert(&user)
	if e != nil {
		result.Code = 404
		result.Success = false
		result.Message = "系统异常"
		logs.Error(e)
	} else {
		logs.Info(">>>> user register success,user id = " + strconv.FormatInt(id, 10) + " <<<<")
		result.Code = 200
		result.Success = true
		result.Message = "注册成功"
	}
	logs.Info("=========================================================================================")
	bytes, err := json.Marshal(&result)
	c.Ctx.ResponseWriter.Write(bytes)
}
// @router /register/do [post]
func (c *LoginController) DoRegister() {
	logs.Info("========DoRegisterDoRegisterDoRegisterDoRegisterDoRegisterDoRegisterDoRegisterDoRegisterDoRegister======")
	var user models.User
	var result models.Result
	//Json Unmarshal：将json字符串解码到相应的数据结构
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Info(err)
	logs.Info(c.Ctx.Input.RequestBody)
	checkError(err)

	userName := user.UserName
	password := user.Password
	email := user.Email
	logs.Info(">>>> user register information username=" + userName + ",email=" + email + ",password=" + password + " <<<<")

	pass, err := regexp.MatchString(`[a-zA-Z0-9]{8,16}`, password)
	checkError(err)
	if !pass {
		result.Code = 404
		result.Success = false
		result.Message = "密码格式错误"
	}

	em, err := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email)
	checkError(err)
	if !em {
		result.Code = 404
		result.Success = false
		result.Message = "邮箱格式错误"
	}
	// 对密码加密
	user.Password = passwordEncode(password)
	logs.Info("=========================================================================================")
	logs.Info(user.Password)
	id, e := orm.NewOrm().Insert(&user)
	if e != nil {
		result.Code = 404
		result.Success = false
		result.Message = "系统异常"
		logs.Error(e)
	} else {
		logs.Info(">>>> user register success,user id = " + strconv.FormatInt(id, 10) + " <<<<")
		result.Code = 200
		result.Success = true
		result.Message = "注册成功"
	}
	logs.Info("=========================================================================================")
	bytes, err := json.Marshal(&result)
	c.Ctx.ResponseWriter.Write(bytes)
}
func passwordEncode(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// @router /login/do  [post]
func (c *LoginController) DoLogin() {
	logs.Info(">>>> user do login start <<<<")
	var user models.User
	// 表单映射成struct
	// 实例化一个数据结构，用于生成json字符串
	// jsonStu是[]byte类型，转化成string类型便于查看
	// jsonStu是[]byte类型，转化成string类型便于查看
	// fmt.Println(string(jsonStu))
	// jsonStu是[]byte类型，转化为string类型便于查看
	// interface{}类型其实是一个空接口，即没有方法接口。go的每一种类型都实现了该接口。因此，任何其它类型的数据都可以复制给
	// interface{}类型
	// 无论是string、int、bool，还是指针类型，都可以赋值给interface{}类型，且正常编码，效果和前面的例子一样的。
	// 在实际项目中，编码成json串的数据结构，往往是切片类型。如下定义了一个[]StuRead类型的切片
	// 方式一
	// var stus1 []*StuRead
	// 只申明，不分配内存
	// 方式二
	// 分配初始值为0的内存
	// stus2 := make([]&StuRead，0)
	// stu1  := StuRead{成员赋值...}
	// stu2  := StuRead{成员赋值...}
	// 由方式1和2创建切片，都能成功追加数据
	// 方式2最好分配0长度，append是会自动增长。导致数据丢失
	// json1, _ := json.Marshal(stus1)
	// json2, _ := json.Marshal(stus2)
	// 解码时定义对应的切片接收即可。
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	user.Password = passwordEncode(user.Password)
	right := models.QueryByNamePwd(user.UserName, user.Password)
	var r models.Result
	if right {
		r.Code = 200
		r.Success = true
		r.Message = "success"
	} else {
		r.Code = 404
		r.Success = false
		r.Message = "用户或密码错误"
	}

	var bytes []byte
	bytes, err = json.Marshal(&r)
	if err != nil {
		logs.Error(err)
	}
	c.Ctx.ResponseWriter.Write(bytes)
}

// nameType:=reflect.TypeOf(stu.Name)
// nameType := reflect.TypeOf(stu.Name)
// ageType  := reflect.TypeOf(stu.Age)
// highType := reflect.TypeOf(stu.sex)
// sexType  := reflect.TypeOf(stu.)
// ageType:=reflect.TypeOf(stu.Age)
// highType:=reflect.TypeOf(stu.HIgh)
// sexType:=reflect.TypeOf(stu.sex)
// classType:=reflect.TypeOf(stu.Class)
// testType:=reflect.TypeOf(stu.Test)
// fmt.Println("nameType:",nameType)
// fmt.Println("ageType:",ageType)
// fmt.Println("highType:",highType)
// fmt.Println("sexType:",sexType)
// fmt.Println("classType:",classType)
// fmt.Println("testType:",testType)