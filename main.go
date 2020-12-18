package main

import (
	_ "shanghai/routers"
	"github.com/astaxie/beego"
	_"shanghai/models"
)

func main() {
	beego.AddFuncMap("prepage",ShowPrePage)
	beego.AddFuncMap("nextpage",ShowNext)
	beego.Run()
}

//后台定义一个函数
func ShowPrePage(pageIndex int)int  {
	if pageIndex==1{
		return  pageIndex
	}
	return pageIndex-1
}

func ShowNext(pageIndex int ,pageCount int)int  {
	if pageIndex==pageCount {
		return pageCount
	}
	return pageIndex+1
}

/*
	作用：处理视图中简单业务逻辑
	1：创建后台函数
	2：在视图中定义函数名
	3.在beego.run之前关联起来
*/