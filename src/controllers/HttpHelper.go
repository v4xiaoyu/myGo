package controllers

import (
	"github.com/astaxie/beego"
	"main/controllers"
)

func StartHttpListener() {
	//http+mysql
	beego.Router("/getNews", &controllers.NewsController{})
	beego.Router("/getUser", &controllers.UserController{})
	beego.Run(":8081")
}
