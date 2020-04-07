package controllers

import (
	"db/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	o :=orm.NewOrm()
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	id ,err := o.Insert(&user)
	pid,perr := o.Insert(&user.Profile)
	if err != nil {
		fmt.Println("insert user err :", err)
		return
	}
	if perr != nil {
		fmt.Println("insert profile err :", err)
		return
	}
	logs.Notice("insert profile id %v",pid)
	u.Data["json"] = id
	u.ServeJSON()
}
// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid, _ := u.GetInt(":uid")
	logs.Notice("Get By Uid :",uid)
	o :=orm.NewOrm()
	query := models.User{Id: uid}
	err :=o.Read(&query)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		u.Data["json"] = err.Error()
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		u.Data["json"] = err.Error()
	} else {
		fmt.Println(query.Id, query.Name)
		u.Data["json"] = query
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	o :=orm.NewOrm()
	query := models.User{Id: 1}
	err :=o.Read(&query)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		u.Data["json"] = err.Error()
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		u.Data["json"] = err.Error()
	} else {
		fmt.Println(query.Id, query.Name)
		u.Data["json"] = query
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		o :=orm.NewOrm()
		query := models.User{Id: 1}
		err :=o.Read(&query)
		if err == orm.ErrNoRows {
			fmt.Println("查询不到")
			u.Data["json"] = err.Error()
		} else if err == orm.ErrMissPK {
			fmt.Println("找不到主键")
			u.Data["json"] = err.Error()
		} else {
			fmt.Println(query.Id, query.Name)
			u.Data["json"] = query
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	logs.Notice("uid:"+uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	logs.Notice(username+":"+password)
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}