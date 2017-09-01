package main

import (
	"../controllers"
	"../data/db"
	"../data/quickDb"
	"../utils"
	"fmt"
	//"github.com/astaxie/beego"
	//"log"
)

func main() {
	db.InitDb()

	quickDb.Init()

	c := make(chan string, 10)
	utils.GetDate(c)
	s := <-c
	fmt.Println(s)

	controllers.StartSocket()
	//http+mysql
	//beego.Router("/getNews", &controllers.NewsController{})
	//beego.Router("/getUser", &controllers.UserController{})
	//beego.Run(":8081")

	//redis
	//key := "hello"
	//log.Println(quickDb.Get(key))
	//quickDb.Set(key, "welcome")
	//log.Println(quickDb.Get(key))
	//quickDb.Delete(key)
	//log.Println(quickDb.Get(key))
}

func stop() {
	db.CloseDb()
	quickDb.Quit()
}
