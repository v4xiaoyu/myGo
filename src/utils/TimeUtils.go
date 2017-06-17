package utils

import (
	"time"
	"fmt"
)

var now time.Time = time.Now()

func GetDate(c chan string) {
	var s string;
	s = fmt.Sprintf("%d/%d/%d", now.Year(), now.Month(), now.Day())
	c <- s
}

func GetTime(c chan string) {
	var s string;
	s = fmt.Sprintf("%2d:%2d:%2d", now.Hour(), now.Minute(), now.Second())
	c <- s
}

func CompareNow(t time.Time, c chan int64) {
	go CompareTime(t, now, c)
}

func CompareTime(t1, t2 time.Time, c chan int64) {
	if t1.UnixNano() < t2.UnixNano() {
		c <- -1
	} else if t1.UnixNano() > t2.UnixNano() {
		c <- 1
	} else {
		c <- 0
	}
}