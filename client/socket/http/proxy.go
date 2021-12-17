package http

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getClient(proxyUrl string) (net.Conn, error) {
	cli := NewClient()
	return cli.Dial(proxyUrl)
}

func NewProxy(proxyUrl string) gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			fmt.Println("------ client proxy close", c.Request.URL.Host)
		}()

		if c.Request.Method != "CONNECT" {
			// fmt.Println(r.Method, r.RequestURI)
			c.JSON(404, gin.H{
				"msg": "NOT FOUND, PLS USE HTTPS",
			})
			return
		}

		// c.Request.Method
		c.JSON(200, gin.H{
			"msg": "success",
		})
		// 开始代理
		fmt.Println("------ begin proxy", c.Request.URL.Host)
		destConn, err := getClient(proxyUrl)
		if err != nil {
			fmt.Println("destConn err ", err)
			panic(err)
		}
		defer destConn.Close()

		srcConn, _, err := c.Writer.(http.Hijacker).Hijack()
		if err != nil {
			fmt.Println("srcConn err ", err)
			// panic(err)
			return
		}
		defer srcConn.Close()

		// 写入代理数据
		sUrl := []byte("url:" + c.Request.URL.Host)
		sLen := int16(len(sUrl))
		// bigEndian tcp最大包是65536字节，所以server端一次是可以write到完整的url路径
		var b [1024]byte
		b[0] = byte(sLen >> 8)
		b[1] = byte(sLen)
		copy(b[2:], sUrl)
		// fmt.Println("------ print len ", len(b))
		destConn.Write(b[:sLen+2])

		//srcConn -> destConn
		go func() {
			var b [1024 * 2]byte
			for {
				var n int
				var err error

				srcConn.SetDeadline(time.Now().Add(10 * time.Second))
				destConn.SetDeadline(time.Now().Add(10 * time.Second))

				if n, err = srcConn.Read(b[:]); err != nil {
					fmt.Println("srcConn read over ", err)
					return
				}
				if _, err = destConn.Write(b[:n]); err != nil {
					fmt.Println("destConn write err ", err)
					return
				}
			}
		}()
		// destConn -> srcConn
		var b2 [1024 * 2]byte
		for {
			var n int
			var err error
			srcConn.SetDeadline(time.Now().Add(10 * time.Second))
			destConn.SetDeadline(time.Now().Add(10 * time.Second))

			if n, err = destConn.Read(b2[:]); err != nil {
				fmt.Println("destConn read over ", err)
				return
			}

			if _, err = srcConn.Write(b2[:n]); err != nil {
				fmt.Println("srcConn write err ", err)
				return
			}
		}
	}
}
