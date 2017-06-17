package main

import (
	"fmt"
	"../utils"
	"time"
)

func client() {
	c := make(chan string, 10)
	comp := make(chan int64)
	go utils.GetDate(c)
	go utils.GetTime(c)

	fmt.Println(<-c, <-c)

	t := time.Now()
	go utils.CompareNow(t, comp)

	b := <-comp

	if b > 0 {
		fmt.Println("after")
	} else if b < 0 {
		fmt.Println("before")
	} else {
		fmt.Println("now")
	}

	time.Sleep(3000)
	t2 := time.Now()
	go utils.CompareTime(t, t2, comp)

	a := <-comp

	if a > 0 {
		fmt.Println("t after t2")
	} else if a < 0 {
		fmt.Println("t before t2")
	} else {
		fmt.Println("t is t2")
	}

}

func main() {
	//db.ConnectDb()
	client()
}
