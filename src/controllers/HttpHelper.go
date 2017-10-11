package controllers

import (
	"github.com/astaxie/beego"
)

func StartHttpListener() {
	//http+mysql
	beego.Router("/getNews", &NewsController{})
	beego.Router("/getUser", &UserController{})
	beego.Run(":8081")
}
