package main

import (
	"../controllers"
	"../data/db"
	"github.com/astaxie/beego"
	"../utils"
	"fmt"
)

func main() {
	db.InitDb()

	c := make(chan string, 10)
	utils.GetDate(c)
	s := <-c
	fmt.Println(s)

	beego.Router("/getNews", &controllers.NewsController{})
	beego.Router("/getUser", &controllers.UserController{})
	beego.Run(":8081")
}
