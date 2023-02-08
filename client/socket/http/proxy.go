package http

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getClient(proxyUrl string) (net.Conn, error) {
	cli := NewClient()
	return cli.Dial(proxyUrl)
}
func randomPortAndPassword(port, portNumber int, passwordArr []string) (newPort string, password string) {
	_num := rand.Intn(portNumber)
	newPort = strconv.Itoa(port + _num)
	password = passwordArr[_num]
	fmt.Printf("---random port: %s\n", newPort)
	return
}

func NewProxy(portNumber int, ip string, port int, passwordArr []string) gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			fmt.Println("------ client proxy close", c.Request.URL.Host)
		}()

		// 开始代理
		fmt.Println("------ begin proxy", c.Request.URL.Host)
		var destConn net.Conn
		var srcConn net.Conn

		defer func() {
			destConn.Close()
			srcConn.Close()
		}()

		var newPort string
		var password string
		var err error
		for {
			newPort, password = randomPortAndPassword(port, portNumber, passwordArr)
			proxyUrl := ip + ":" + newPort
			destConn, err = getClient(proxyUrl)
			if err != nil {
				if destConn != nil {
					destConn.Close()
				}
				fmt.Println("destConn err ", err)
				continue
			}
			break
		}

		srcConn, _, err = c.Writer.(http.Hijacker).Hijack()
		if err != nil {
			if srcConn != nil {
				srcConn.Close()
			}
			fmt.Println("srcConn err ", err)
			return
		}

		// 写入代理数据
		Url := c.Request.URL.Host
		if c.Request.URL.Scheme == "http" && c.Request.URL.Port() == "" {
			Url = Url + ":80"
		}
		Url = Url + password
		var b [1024]byte
		n := copy(b[:], []byte(Url))
		// fmt.Println("------ print len ", len(b))
		destConn.Write(b[:n])

		// http
		if c.Request.Method != "CONNECT" {

			err := c.Request.WriteProxy(destConn)
			if err != nil {
				fmt.Println("WriteProxy err ", err)
				return
			}

		} else {
			// https
			srcConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
		}

		//srcConn -> destConn
		go func() {
			var b [1024 * 2]byte
			for {
				var n int
				var err error

				if n, err = srcConn.Read(b[:]); err != nil {
					fmt.Println("srcConn read over ", err)
					srcConn.Close()
					destConn.Close()
					return
				}
				if _, err = destConn.Write(b[:n]); err != nil {
					fmt.Println("destConn write err ", err)
					srcConn.Close()
					destConn.Close()
					return
				}
			}
		}()
		// destConn -> srcConn
		var b2 [1024 * 2]byte
		for {
			var n int
			var err error
			srcConn.SetDeadline(time.Now().Add(30 * time.Second))
			destConn.SetDeadline(time.Now().Add(3 * time.Second))

			if n, err = destConn.Read(b2[:]); err != nil {
				fmt.Println("destConn read over ", err)
				return
			}

			// fmt.Println("------ print n", n)

			if _, err = srcConn.Write(b2[:n]); err != nil {
				fmt.Println("srcConn write err ", err)
				return
			}
		}
	}
}
