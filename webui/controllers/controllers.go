package controllers

import (
	"strconv"
	"strings"

	"../bootstrap"
	"../models"
	"../services"
	"github.com/astaxie/beego"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	//beego.InsertFilter("", beego.BeforeRouter, func, "")
	//beego.FilterRouter{}
	beego.Router("/", &MainController{}, "*:Index")
	beego.Router("/login", &MainController{}, "*:Login")
	beego.Router("/logout", &MainController{}, "*:Logout")
	beego.Router("/gettime", &MainController{}, "*:GetTime")
	beego.AutoRouter(NewStorageController())

}

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.User
	userService    services.UserService
	userId         int64
	userName       string
	pageSize       int
}

func (this *BaseController) Prepare() {
	this.pageSize = 20
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.userService = services.NewUserService()

	this.auth()

	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
	this.Data["menuTag"] = this.controllerName
}

//登录状态验证
func (this *BaseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	if len(arr) == 2 {
		idstr, _ := arr[0], arr[1]
		userId, _ := strconv.ParseInt(idstr, 10, 64)

		if userId > 0 {
			user := this.userService.GetByID(userId)
			//if err == nil && password == libs.Md5([]byte(this.getClientIp()+"|"+user.Password+user.Salt)) {
			this.userId = user.ID
			this.userName = user.Username
			this.user = &user
			//}
		}
	}

	if this.userId == 0 && (this.controllerName != "main" ||
		(this.controllerName == "main" && this.actionName != "logout" && this.actionName != "login")) {
		this.redirect(beego.URLFor("MainController.Login"))
	}
}

//渲染模版
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}
	this.Layout = "public/layout.html"
	this.TplName = tplname
}

// 重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

// 显示错误信息
func (this *BaseController) showMsg(args ...string) {
	this.Data["message"] = args[0]
	redirect := this.Ctx.Request.Referer()
	if len(args) > 1 {
		redirect = args[1]
	}

	this.Data["redirect"] = redirect
	this.Data["pageTitle"] = "系统提示"
	this.display("error/message")
	this.Render()
	this.StopRun()
}

// 输出json
func (this *BaseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg

	this.jsonResult(out)
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
