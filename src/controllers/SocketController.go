package controllers

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net"
)

type SocketStatus struct {
	Id      int
	Running bool
}

func StartSocket(running chan SocketStatus, protocol string, port string) {
	l, err := net.Listen(protocol, port)
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
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

	result, err := j.Encode()
	if err != nil {
		panic(err)
	}
	go writeToConn(conn, bytes.NewBuffer(result).String())
}

func readFromConn(conn net.Conn) {
	b := make([]byte, 32)
	_, err := conn.Read(b)
	if err != nil {
		panic(err)
	}

	go handleConn(conn, b)
}

func writeToConn(conn net.Conn, str string) {
	buf := bytes.NewBufferString(str)
	conn.Write(buf.Bytes())
}
