package main

import (
	clientHttp "client/socket/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	IP        = "127.0.0.1"
	PORT      = "8711"
	CONN_IP   = "127.0.0.1"
	CONN_PORT = "8712"
)

func main() {
	localUrl := IP + ":" + PORT
	proxyUrl := CONN_IP + ":" + CONN_PORT
	router := gin.Default()
	router.Use(clientHttp.NewProxy(proxyUrl))
	http.ListenAndServe(localUrl, router)

	// server := tcp.NewServer(host)
	// server.Start()
}
