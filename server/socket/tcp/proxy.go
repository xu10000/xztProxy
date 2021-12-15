package tcp

import (
	"fmt"
	"io"
	"net"
	"regexp"
	"server/pkg/utils"
	"sync"
)

func getClient(addr string) (net.Conn, error) {
	cli := NewClient()
	return cli.Dial(addr)
}

func getRealUrl(b []byte) string {
	// 自定义简单协议
	_bLen := b[0]<<8 + b[1]
	_b := string(b[2 : _bLen+2])
	// var realUrl string
	reg := regexp.MustCompile(`url:(\S+)`)
	regArr := reg.FindStringSubmatch(_b)
	fmt.Println("------ print", regArr)
	if len(regArr) == 0 {
		fmt.Println("------ regArr error", _b)
		return ""
	}
	fmt.Println("------getRealUrl function  proxy url", _b)
	return regArr[1]
}
func NewProxy(conn net.Conn) {
	// check proxy address
	b := make([]byte, 1024)
	_, err := io.ReadFull(conn, b)
	if err != nil {
		fmt.Println(" ReadAll err ", err)
		panic(err)
	}
	//
	realUrl := getRealUrl(b)

	// 开始代理
	// fmt.Println("------ begin server proxy ", "url:www.baidu.com:443")
	// // return
	destConn, err := getClient(realUrl)
	if err != nil {
		fmt.Println("destConn err ", err)
		panic(err)
	}
	defer destConn.Close()

	srcConn := conn

	if err != nil {
		fmt.Println("srcConn err ", err)
		panic(err)
	}
	defer srcConn.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go utils.IoCopy(&wg, srcConn, destConn)
	go utils.IoCopy(&wg, destConn, srcConn)
	wg.Wait()
	// fmt.Println("------ proxy success")
}
