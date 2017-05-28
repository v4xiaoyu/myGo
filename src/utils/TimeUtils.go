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

func CompareTime(t time.Time, c chan int64) {
	if t.UnixNano() < now.UnixNano() {
		c <- -1
	} else if t.UnixNano() > now.UnixNano() {
		c <- 1
	} else {
		c <- 0
	}
}