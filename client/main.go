package main

import (
	"client/config"
	clientHttp "client/socket/http"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	IP        = "127.0.0.1"
	PORT      = "8711"
	CONN_IP   string
	CONN_PORT string
	Config    config.Config
)

func init() {
	viper.AddConfigPath("./")
	// viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("------ReadInConfig error", err)
		panic(err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println("------Unmarshal error", err)
		panic(err)
	}

	CONN_IP = Config.Host
	CONN_PORT = strconv.Itoa(Config.Port)
}

func main() {
	localUrl := IP + ":" + PORT
	proxyUrl := CONN_IP + ":" + CONN_PORT
	router := gin.Default()
	router.Use(clientHttp.NewProxy(proxyUrl, Config.Password))
	http.ListenAndServe(localUrl, router)

}
