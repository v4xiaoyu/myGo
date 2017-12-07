package controllers

import (
	"github.com/astaxie/beego"
)

type HttpStatus struct {
	Id      int
	Running bool
}

func StartHttpListener(statuses chan HttpStatus) {
	//http+mysql
	beego.Router("/getNews", &NewsController{})
	beego.Router("/getUser", &UserController{})
	beego.Run(":8081")
}
