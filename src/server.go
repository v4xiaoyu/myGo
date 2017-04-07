package main

import (
	"./controllers"
	"./controllers/db"
	"github.com/astaxie/beego"
)

func main() {
	db.InitDb()

	beego.Router("/getNews", &controllers.NewsController{})
	beego.Router("/getUser", &controllers.UserController{})
	beego.Run(":8081")
}
