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
	c := make(chan string, 10)
	utils.GetDate(c)
	s := <-c
	fmt.Println(s)

	go db.InitDb()
	go quickDb.Init()
	go controllers.StartSocket()
	go controllers.StartHttpListener()
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
