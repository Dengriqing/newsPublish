package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghai/models"
)

type UserController struct {
	beego.Controller
}

//展示注册页面
func (this*UserController)ShowRegister() {
	this.TplName="register.html"
}

//处理注册数据
func (this*UserController)HandlePost()  {
	//获取数据
	userName:=this.GetString("userName")
	pwd:=this.GetString("password")
	//校验数据
	if userName==""||pwd==""{
		this.Data["errmsg"]="数据不完整"
		beego.Info("数据不完整")
		this.TplName="register.html"
		return
	}

	//操作数据
	//获取orm对象
	o:=orm.NewOrm()
	//获取插入对象
	var user models.User
	//给插入对象赋值
	user.Name=userName
	user.Password=pwd
	//插入
	o.Insert(&user)
	//返回结果


	//返回页面
	//this.Ctx.WriteString("注册成功")
	//http状态码
	//1XX 100 请求成功，但需要继续发送
	//2XX 200 请求成功
	//3XX 300 302资源被转移，重定向
	//4XX 404 请求错误
	//5XX 500 服务端崩溃
	//this.TplName="login.html" 可以传递数据 转发
	this.Redirect("/login",302)//重定向
}

//展示登陆页面
func (this*UserController)ShowLogin() {
	userName:=this.Ctx.GetCookie("userName")
	if userName==""{
		this.Data["userName"]=""
		this.Data["checked"]=""
	}else {
		this.Data["userName"]=userName
		this.Data["checked"]="checked"
	}
	this.TplName="login.html"
}

//登陆
func(this*UserController)HangLogin(){
	//获取数据
	userName:=this.GetString("userName")
	pwd:=this.GetString("password")
	//校验数据
	if userName==""||pwd==""{
		this.Data["errmsg"]="登陆数据不完整"
		this.TplName="login.html"
		return
	}
	//操作数据
	//获取orm对象
	o:=orm.NewOrm()
	var user models.User
	user.Name=userName
	user.Password=pwd
	err:=o.Read(&user,"Name")
	if user.Password!=pwd{
		this.Data["errmsg"]="密码错误"
		this.TplName="login.html"
		return
	}
	if err!=nil{
		this.Data["errmsg"]="用户不存在"
		this.TplName="login.html"
		return
	}


	//返回页面
	//this.Ctx.WriteString("登陆成功")
	data:=this.GetString("remember")

	if data=="on"{
		this.Ctx.SetCookie("userName",userName,100)
	}else {
		this.Ctx.SetCookie("userName",userName,-1)
	}
	this.SetSession("userName",userName)
	this.Redirect("/article/showArticleList",302)
}

//退出登录
func (this*UserController)Logout()  {
	//删除session
	this.DelSession("userName")
	//跳转登录页面
	this.Redirect("/login",302)

}