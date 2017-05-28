package main

import (
	"fmt"
	"./utils"
	"time"
)

func client() {
	c := make(chan string, 10)
	comp := make(chan int64)
	go utils.GetDate(c)
	go utils.GetTime(c)
	t := time.Now()
	go utils.CompareTime(t, comp)

	b := <-comp

	if b > 0 {
		fmt.Println("after")
	} else if b < 0 {

		fmt.Println("before")
	} else {

		fmt.Println("now")
	}

	fmt.Println(<-c, <-c)
}

func main() {
	//db.ConnectDb()
	client()
}
