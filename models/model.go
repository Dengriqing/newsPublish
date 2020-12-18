package models

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

//定义一个结构体对象
type User struct {
	Id int
	Name string
	Password string
	//多对多
	Article []*Article `orm:"reverse(many)"`

}

type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string `orm:"size(500)"`
	Aimg string `orm:"size(100)"`
 	//一对多 多表
	ArticleType *ArticleType `orm:"rel(fk)"`
	//多对多
	Users []*User`orm:"rel(m2m)"`
}

//类型表
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	//一对多 1表
	Articles []*Article `orm:"reverse(many)"`
}


func init()  {
	//操作数据库代码
	//第一个参数数据库驱动
	//连接数据库字符串
	//“root:123456 @tcp(127.0.0.1:3306)"
	//用户名：密码 @tcp（127.0.0.1：3306）/数据库名称？charset=utf8
	/*
	conn,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err!=nil{
		beego.Info("连接错误",err)
		return
	}
	//创建表
	//_,err=conn.Exec("create table itcast(name VARCHAR(40) , password VARCHAR(40) );")
	//if err!=nil{
	//	beego.Error("创建失败",err)
	//}
	//关闭数据库
	defer conn.Close()

	//插入数据
	//conn.Exec("insert into itcast (name,password) values (?,?)","chuanzhi","heima")

	//查询
	res,err:=conn.Query("select name from itcast")
	var name string
	for res.Next(){
		res.Scan(&name)
		beego.Info(name)
	}
   */
   //ORM操作数据库
   //获取连接对象
   orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

   //创建表
   orm.RegisterModel(new(User),new(Article),new(ArticleType))

   //生成表
   //第一个参数数据库别名
   //第二个参数是否强制更新 true 删除重新创建
   orm.RunSyncdb("default",false,true)

   //操作表


}