package routers

import (
	"github.com/astaxie/beego/context"
	"shanghai/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//过滤器            过滤请求路径          过滤位置
	beego.InsertFilter("/article/*",beego.BeforeRouter,Filfter)

	beego.Router("/", &controllers.MainController{},"get:ShowGet;post:Post")

	//注册
	beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandlePost")

	//登陆
	beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HangLogin")

	//文章列表页访问
	beego.Router("/article/showArticleList",&controllers.ArticleController{},"get:ShowArticleList")

	//添加文章
	beego.Router("/article/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")

	//显示文章详情
	beego.Router("/article/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")

	//编辑文章
	beego.Router("/article/updateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")

	//删除文章
	beego.Router("/article/deleteArticle",&controllers.ArticleController{},"get:DeleteArticle")

	//添加分类
	beego.Router("/article/addType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")

	//退出登录
	beego.Router("/article/logout",&controllers.UserController{},"get:Logout")

	//删除类型
	beego.Router("/article/deleteType",&controllers.ArticleController{},"get:DeleteType")


	//给请求指定自定义方法 一个请求指定一个方法
	//beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:PostFunc")
	//给多个请求指定一个方法
	//beego.Router("/", &controllers.MainController{}, "get:post:HandleFunc")
	//给所有请求指定一个方法
	//beego.Router("/", &controllers.MainController{}, "*:HandleFunc")
	//当两种指定方法冲突的时候
	//beego.Router("/", &controllers.MainController{}, "*:HandleFunc;post:PostFunc")
}

//过滤器函数
 var Filfter= func(ctx * context.Context) {
 	userName:=ctx.Input.Session("userName")
 	if userName==nil{
 		ctx.Redirect(302,"/login")
		return
	}
 }








