package controllers

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net"
)

func StartSocket() {
	l, err := net.Listen("tcp", ":18888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go readFromConn(c)
	}
}

func handleConn(conn net.Conn, b []byte) {
	buffer := bytes.NewBuffer(b)
	s := buffer.String()

	j := simplejson.New()
	j.Set("status", 0)
	j.Set("message", "ok")
	j.Set("data", s)

	result, err := j.String()
	if err != nil {
		fmt.Println("json handle error:", err)
		return
	}
	go writeToConn(conn, result)
}

func readFromConn(conn net.Conn) {
	var b []byte
	_, err := conn.Read(b)
	if err != nil {
		fmt.Println("Read error:", err)
	}
	ch := make(chan interface{})
	ch <- b
	go handleConn(conn, b)
}

func writeToConn(conn net.Conn, str string) {
	buf := bytes.NewBufferString(str)
	conn.Write(buf.Bytes())
}
