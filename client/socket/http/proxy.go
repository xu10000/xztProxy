package http

import (
	"client/pkg/utils"
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
	stopCh := make(chan struct{})
	errCh := make(chan interface{})
	// timeout
	go func() {
		time.Sleep(5 * time.Second)
		stopCh <- struct{}{}
	}()
	return func(c *gin.Context) {
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
		b := make([]byte, 1024)
		sUrl := []byte("url:" + c.Request.URL.Host)
		sLen := int16(len(sUrl))
		// bigEndian
		b[0] = byte(sLen >> 8)
		b[1] = byte(sLen)
		copy(b[2:], sUrl)
		// fmt.Println("------ print len ", len(b))
		// defer destConn.Write(b)

		if err != nil {
			fmt.Println("srcConn err ", err)
			// panic(err)
			return
		}

		// _ = destConn
		// _ = srcConn
		// fmt.Println("------ begin proxy")
		// _, err = io.Copy(srcConn, destConn)
		// fmt.Println("------ print copy err ", err)

		go utils.IoCopy(errCh, srcConn, destConn)
		go utils.IoCopy(errCh, destConn, srcConn)

		select {
		case err := <-errCh:
			fmt.Println("------ errCh ", err.(error))
			c.Next()
			return
		case <-stopCh:
			fmt.Println("------ stopCh")
			c.Next()
			return
		}
		// time.Sleep(time.Second * 10)
		// fmt.Println("------ proxy success")
	}
}
