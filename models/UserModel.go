package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"log"
)

type User struct {
	Id       int    `form:"-" json:"id"`
	UserName string `form:"username"  orm:"column(username)" json:"userName"`
	Age      int    `form:"age" json:"age"`
	Sex      string `form:"sex" json:"sex"`
	Mobile   string `form:"mobile" json:"mobile"`
	Password string `json:"password"`
	Email    string `form:"email" json:"email"`
}

type Result struct {
	Id       int    `form:"-" json:"id"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}
type Cao struct {
	Id       int    `form:"-" json:"id"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}

type Article struct {
	Id int
	Title string
	Content string
	User *User `orm:"rel(fk)"`
	//这样写，默认在Article中应该有一个user_id的字段，如果没有，
	//通过beego自动建表的话，它会生成；如果是手动建表，记得把user_id 这个字段加上去
}
type Baby struct {
	Id int64
	Name string `json:"name" orm:"size(50)"`
	User *User `json:"user" orm:"rel(fk);index"`
}

func QueryUserById(id int) *User {
	var user User
	orm := orm.NewOrm()
	orm.QueryTable("user").Filter("id", id).One(&user, "username", "age", "sex", "mobile")
	logs.Info(">>>> query user by user id from database <<<<")
	return &user
}

func QueryUserList() []*User {
	var list []*User
	orm := orm.NewOrm()
	orm.QueryTable("user").All(&list, "username", "age", "sex", "mobile")
	return list
}

func InsertUser(u *User) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func InsertUserNew(u *User) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func QueryByNamePwd(username, password string) bool {
	logs.Info(">>>> query user by user name and password from database <<<<")
	var user User
	orm := orm.NewOrm()
	err := orm.QueryTable("user").Filter("username", username).Filter("password", password).One(&user, "username", "age", "sex", "mobile", "email")
	var result bool
	if err != nil {
		logs.Error(err)
		result = false
	} else {
		result = true
	}
	return result
}
