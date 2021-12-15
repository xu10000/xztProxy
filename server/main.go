package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"server/socket/tcp"
)

var (
	CONN_IP   = "0.0.0.0"
	CONN_PORT = "8712"
)

func main() {
	cert, err := tls.LoadX509KeyPair("./config/server.pem", "./config/server.key")
	if err != nil {
		log.Println("load cert error ", err)
		panic(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	localUrl := CONN_IP + ":" + CONN_PORT
	// 拦截所有请求再做转发
	listener, err := tls.Listen("tcp", localUrl, config)
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
