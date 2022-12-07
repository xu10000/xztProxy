package main

import (
	"client/config"
	clientHttp "client/socket/http"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	IP          = "0.0.0.0"
	PORT        = "8711"
	CONN_IP     string
	CONN_PORT   int
	Config      config.Config
	PORT_NUMBER int
	PROXY_URL   = "http://127.0.0.1:8711"
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
	CONN_PORT = Config.BeginPort
	PORT_NUMBER = len(Config.PasswordArr)
}

func main() {

	go requestTask()

	localUrl := IP + ":" + PORT
	router := gin.Default()
	router.Use(clientHttp.NewProxy(PORT_NUMBER, CONN_IP, CONN_PORT, Config.PasswordArr))
	http.ListenAndServe(localUrl, router)

}

func xztProxy(_ *http.Request) (*url.URL, error) {
	return url.Parse(PROXY_URL)
}

func requestTask() {
	rand.Seed(time.Now().Unix())
	transport := &http.Transport{Proxy: xztProxy}
	for {
		time.Sleep(30 * time.Minute)
		client := &http.Client{Transport: transport}
		resp, err := client.Get("http://www.google.com")
		if err == nil {
			resp.Body.Close()
		}
		fmt.Printf("resp %+v err %+v", resp, err)
	}
}
