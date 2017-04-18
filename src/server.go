package main

import (
	"github.com/astaxie/beego"
	"./controllers"
	"./controllers/db"
)

func main() {
	db.ConnectDb()

	beego.Router("/getNews", &controllers.NewsController{})
	beego.Router("/getUser", &controllers.UserController{})
	beego.Run(":8080")
}
