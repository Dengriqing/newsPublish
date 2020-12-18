package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"shanghai/models"
	"time"
)

type ArticleController struct {
	beego.Controller
}

//展示文章列表页 分页 index.html
func (this*ArticleController)ShowArticleList()  {
	//查询sessionon是否有值
	userName:=this.GetSession("userName")
	if userName==nil{
		//sessionon为空，跳转登录页面
		this.Redirect("/login",302)
		return
	}
	//获取数据
	//高级查询
	//指定表
	o:=orm.NewOrm()
	//查询所有
	qs:=o.QueryTable("Article")
	var articles []models.Article
	//_,err:=qs.All(&articles)
	//if err!=nil{
	//	beego.Info("chaxunshibai")
	//}

	//获取文章分类
	typeName:=this.GetString("select")

	//设置每页显示条数
	pageSize:=2

	//获取页码
	pageIndex,err:=this.GetInt("pageIndex")
	if err!=nil{
		pageIndex=1
	}

	//起始位置
	start:=(pageIndex-1)*pageSize

	//获取数据
	//主页文章分类
	//查询总记录数
	var count int64
	if typeName==""{
		//文章分类为空，则查询所有总记录数
		count,_=qs.Count()
	}else {
		//文章分类不为空，则按文章类型查询总记录数
		count,_=qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).Count()
	}
	//天花板函数：两个浮点数相除，得到一个浮点数，向上取整
	pageCount:=math.Ceil(float64(count)/float64(pageSize))


	//获取文章类型
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	this.Data["types"]=types

	//根据选中的类型查询相应的类型文章
	//typeName:=this.GetString("select")
	//条件查询 RelatedSel（查询外表名），Filter(过滤条件)，All（查询当前表）
	if typeName==""{
		//文章分类为空，则查询所有
		qs.Limit(pageSize,start).All(&articles)
		//作用获取数据库部分数据，第一个参数，获取几条数据，第二个数据，从哪条数据开始获取，返回值是querySeter
		//qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articles)
	}else{
		//文章分类不为空，则按文章类型查询所有
		qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articles)
	}


	//传递数据
	this.Data["typeName"]=typeName
	this.Data["userName"]=userName
	this.Data["pageIndex"]=pageIndex
	this.Data["pageCount"]=int(pageCount)
	this.Data["count"]=count
	this.Data["articles"]=articles
	//展示数据
	userlayoutName:=this.GetSession("userName")
	this.Data["userName"]=userlayoutName.(string)
	this.Layout="layout.html"
	this.TplName="index.html"
}

//展示添加文章页面 add.html 展示
func (this*ArticleController)ShowAddArticle()  {
	//查询所有类型并展示
	o:=orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	//传递数据
	this.Data["types"]=types
	userlayoutName:=this.GetSession("userName")
	this.Data["userName"]=userlayoutName.(string)
	this.Layout="layout.html"
	this.TplName="add.html"
}

//获取添加文章数据 add.html 处理数据 上传
func (this*ArticleController)HandleAddArticle()  {
	//获取数据
	//获取文章标题
	articleName:=this.GetString("articleName")
	//获取文章内容
	content:=this.GetString("content")

	//校验数据b
	if articleName==""||content==""{
		this.Data["errmsg"]="数据不完整"
		userlayoutName:=this.GetSession("userName")
		this.Data["userName"]=userlayoutName.(string)
		this.Layout="layout.html"
		this.TplName="add.html"
		return
	}
	//处理上传文件
	file,head,err:=this.GetFile("uploadname")
	defer file.Close()
	if err!=nil{
		this.Data["errmsg"]="文件上传失败"
		userlayoutName:=this.GetSession("userName")
		this.Data["userName"]=userlayoutName.(string)
		this.Redirect("/article/addArticle",302)
		return
	}


    //文件大小
    if head.Size>5000000{
		this.Data["errmsg"]="文件太大"
		userlayoutName:=this.GetSession("userName")
		this.Data["userName"]=userlayoutName.(string)
		this.Layout="layout.html"
		this.TplName="add.html"
		return
	}

    //文件格式
    //获取后缀名
	ext:=path.Ext(head.Filename)
	if ext!=".jpg"&&ext!=".png"{
		this.Data["errmsg"]="文件格式错误"
		userlayoutName:=this.GetSession("userName")
		this.Data["userName"]=userlayoutName.(string)
		this.Redirect("/article/addArticle",302)
		return
	}

    //防止重名
	fileName:=time.Now().Format("2006-01-02-15:04:05")+ext
	//保存文件:第一个参数是前端标签的name属性值，第二个参数是文件在服务器端存储的位置
	this.SaveToFile("uploadname","./static/img/"+fileName)

	//处理数据
	//插入操作
	o:=orm.NewOrm()
	var article models.Article
	article.ArtiName=articleName
	article.Acontent=content
	article.Aimg="/static/img/"+fileName
	//给文章添加类型
	//获取文章类型
	//根据名称查询类型
	typeName:=this.GetString("select")
	var articleType models.ArticleType
	articleType.TypeName=typeName
	o.Read(&articleType,"TypeName")

	article.ArticleType=&articleType
	//插入数据库
	o.Insert(&article)

	//返回页面
	this.Redirect("/article/showArticleList",302)
}

//展示文章详情页面 content.html 多对多
func (this*ArticleController)ShowArticleDetail()  {
	//获取数据
	//获取展示文章id
	id,err:=this.GetInt("articleId")
	//数据校验
	if err!=nil{
		beego.Info("传递连接错误")
	}
	//操作查询
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	//o.Read(&article)
	//根据id，文字类型查询文字详情
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("Id",id).One(&article)

	//修改阅读量
	article.Acount+=1
	//更新
	o.Update(&article)

	//多对多插入浏览记录
	//获取多对多操作对象
	m2m:=o.QueryM2M(&article,"Users")
	//获取当前浏览用户名
	userName:=this.GetSession("userName")
	if userName==nil{
		this.Redirect("/login",302)
		return
	}
	var user models.User
	user.Name=userName.(string)
	//数据库查询当前用户
	o.Read(&user,"Name")

	//插入操作
	m2m.Add(user)

	//查询最近浏览
	//o.LoadRelated(&article,"Users")
	var users []models.User
	//.Distinct 去重
	o.QueryTable("User").Filter("Article__Article__Id",id).Distinct().All(&users)

	//返回数据
	this.Data["users"]=users
	this.Data["article"]=article


	//返回视图
	userlayoutName:=this.GetSession("userName")
	this.Data["userName"]=userlayoutName.(string)
	this.Layout="layout.html"
	this.TplName="content.html"
}

//显示编辑页面 展示数据 update.html
func (this*ArticleController)ShowUpdateArticle()  {
	//获取数据
	//获取编辑用户的id
	id,err:=this.GetInt("articleId")
	//检验数据
	if err!=nil{
		beego.Info("请求文章错误")
		return
	}
	//处理数据
	//根据id查询相应的文章
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	o.Read(&article)

	//返回视图
	this.Data["article"]=article
	userlayoutName:=this.GetSession("userName")
	this.Data["userName"]=userlayoutName.(string)
	this.Layout="layout.html"
	this.TplName="update.html"
}

//封装上传文件函数
func UploadFile(this*beego.Controller,filePath string)string  {
	//处理上传文件
	file,head,err:=this.GetFile(filePath)
	//不更新图片
	if head.Filename==""{
		return "Noimg"
	}

	if err!=nil{
		this.Data["errmsg"]="文件上传失败"
		userlayoutName:=this.GetSession("userName")
		this.Data["userName"]=userlayoutName.(string)
		this.Layout="layout.html"
		this.TplName="add.html"
		return ""
	}
	defer file.Close()

	//文件大小
	if head.Size>5000000{
		this.Data["errmsg"]="文件太大"
		userlayoutName:=this.GetSession("userName")
		this.Data["userName"]=userlayoutName.(string)
		this.Layout="layout.html"
		this.TplName="add.html"
		return ""
	}

	//文件格式
	//获取后缀名
	ext:=path.Ext(head.Filename)
	if ext!=".jpg"&&ext!=".png"{
		this.Data["errmsg"]="文件格式错误"
		userlayoutName:=this.GetSession("userName")
		this.Data["userName"]=userlayoutName.(string)
		this.Redirect("/article/addArticle",302)
		return ""
	}


	//防止重名
	fileName:=time.Now().Format("2006-01-02-15:04:05")+ext
	//存储
	this.SaveToFile(filePath,"./static/img/"+fileName)
	return "/static/img/"+fileName
}

//处理编辑界面数据 更新数据处理 update.html
func (this*ArticleController)HandleUpdateArticle(){
	//获取数据
	id,err:=this.GetInt("articleId")
	articleName:=this.GetString("articleName")
	content:=this.GetString("content")
	filePath:=UploadFile(&this.Controller,"uploadname")


	//检验数据
	if err!=nil || articleName=="" || content=="" || filePath==""{
		beego.Info("请求路径错误")
	}
	//处理数据
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	err=o.Read(&article)
	if err!=nil{
		beego.Info("更新文章不存在")
		return
	}
	article.ArtiName=articleName
	article.Acontent=content
	//如果filePath等于Noing则不需要更新图片，不等于Noing则需要更新
	if filePath !="Noimg"{
		article.Aimg=filePath
	}
	o.Update(&article)

	//返回视图
	this.Redirect("/article/showArticleList",302)
}

//删除处理文章 index.html
func (this*ArticleController)DeleteArticle()  {
	//获取数据
	id,err:=this.GetInt("articleId")
	//校验数据
	if err!=nil{
		beego.Info("删除文章请求路径错误")
		return
	}
	//处理数据
	//删除操作
	o:=orm.NewOrm()
	var article models.Article
	article.Id=id
	o.Delete(&article)
	//返回视图
	this.Redirect("/article/showArticleList",302)
}

//展示添加类型页面 展示 addType.html
func (this*ArticleController) ShowAddType() {
	//查询所有类型
	o:=orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	//传递数据
	this.Data["types"]=types
	userlayoutName:=this.GetSession("userName")
	this.Data["userName"]=userlayoutName.(string)
	this.Layout="layout.html"
	this.TplName="addType.html"
}

//处理添加类型数据 数据处理 addType
func (this*ArticleController)HandleAddType()  {
	//获取数据
	typeName:=this.GetString("typeName")
	//数据校验
	if typeName==""{
		beego.Info("信息数据不完整")
		return
	}
	//处理数据
	o:=orm.NewOrm()
	var atcitleType models.ArticleType
	atcitleType.TypeName=typeName
	o.Insert(&atcitleType)
	//返回视图
	this.Redirect("/article/addType",302)
}

//删除类型
func (this*ArticleController)DeleteType()  {
	//获取数据
	id,err:=this.GetInt("id")
	//校验数据
	if err!=nil{
		beego.Info("删除类型错误")
		return
	}
	//处理数据
	o:=orm.NewOrm()
	var articleType models.ArticleType
	articleType.Id=id
	o.Delete(&articleType)
	//返回视图
	this.Redirect("/article/addType",302)
}


















