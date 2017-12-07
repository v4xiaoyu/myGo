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
	flag_http := make(chan controllers.HttpStatus, 2)
	go controllers.StartHttpListener(flag_http)
	flag_socket := make(chan controllers.SocketStatus)
	controllers.StartSocket(flag_socket, "tcp", ":18188")
	//redis
	//key := "hello"
	//log.Println(quickDb.Get(key))
	//quickDb.Set(key, "welcome")
	//log.Println(quickDb.Get(key))
	//quickDb.Delete(key)
	//log.Println(quickDb.Get(key))

	//for true {
	//result :=<-flag_socket
	//if !result.Running {
	//	break
	//}
	//}
}

func stop() {
	db.CloseDb()
	quickDb.Quit()
}
