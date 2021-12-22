package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"server/socket/tcp"
	"strconv"
)

var (
	CONN_IP       = "0.0.0.0"
	CONN_PORT     = 8712
	CONN_MAX_PORT = 8912
)

func main() {
	stopCh := make(chan struct{})
	cert, err := tls.LoadX509KeyPair("./config/server.pem", "./config/server.key")
	if err != nil {
		log.Println("load cert error ", err)
		panic(err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// KeepAlive: 1000 * time.Second
	}
	for i := CONN_PORT; i < CONN_MAX_PORT; i++ {
		port := i
		localUrl := CONN_IP + ":" + strconv.Itoa(port)
		// 拦截所有请求再做转发
		listener, err := tls.Listen("tcp", localUrl, config)
		if err != nil {
			fmt.Println("Listen err ", err)
			panic(err)
		}

		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Accept err ", err)
					panic(err)

				}
				go tcp.NewProxy(conn)
			}
		}()
	}

	// stuck
	<-stopCh
	fmt.Println("------ print", "already stop")

}
