package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghai/models"
)

type MainController struct {
	beego.Controller
}

func (c*MainController) Get(){
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["data"]="china"
	c.TplName = "test.html"
}

func (c*MainController) Post(){
	c.Data["data"]="上海一期最棒"
	c.TplName = "test.html"
}

func (c*MainController) ShowGet(){
	//获取ORM对象
	O:=orm.NewOrm()
	//执行某个函数 增删改查
	//插入操作
	/*
	//插入操作
	var user models.User
	user.Name="heima"
	user.Password="chuanzhi"

	//返回结果

	count,err:=O.Insert(&user)
	if err!=nil{
		beego.Error("失败")
	}
	beego.Info(count)
	*/

	//查询操作
	/*
	var user models.User
	user.Id=1
	err:=O.Read(&user,"Id")
	if err!=nil{
		beego.Error("查询失败")
	}
	beego.Info(user)
	*/

	//更新操作
	/*
	var user models.User
	user.Id=1
	err:=O.Read(&user)
	if err!=nil{
		beego.Info("更新数据不存在")
	}
	user.Name="sh"
	count,err:=O.Update(&user)
	if err!=nil{
		beego.Info("更新失败")
	}
	beego.Info(count)
	*/

	//删除操作
	var user models.User
	user.Id=1
	count,err:=O.Delete(&user)
	if err!=nil{
		beego.Info("失败")
	}
	beego.Info(count)

	c.Data["data"]="上海"
	c.TplName = "test.html"
}
