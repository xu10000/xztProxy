package main

import (
	"fmt"
	"net"
	"server/socket/tcp"
)

var (
	CONN_IP   = "127.0.0.1"
	CONN_PORT = "8712"
)

func main() {
	localUrl := CONN_IP + ":" + CONN_PORT
	// 拦截所有请求再做转发
	listener, err := net.Listen("tcp", localUrl)
	if err != nil {
		fmt.Println("Listen err ", err)
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err ", err)
			panic(err)

		}
		go tcp.NewProxy(conn)
	}

	// server := tcp.NewServer(host)
	// server.Start()
}
