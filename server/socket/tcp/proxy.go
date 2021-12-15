package tcp

import (
	"fmt"
	"io"
	"net"
	"regexp"
	"server/pkg/utils"
	"sync"
	"time"
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
	return regArr[1]
}
func NewProxy(conn net.Conn) {
	stopCh := make(chan struct{})
	errCh := make(chan interface{})
	// timeout
	go func() {
		time.Sleep(5 * time.Second)
		stopCh <- struct{}{}
	}()
	// check proxy address
	b := make([]byte, 1024)
	_, err := io.ReadFull(conn, b)
	if err != nil {
		fmt.Println(" ReadAll err ", err)
		// panic(err)
		return
	}
	//
	realUrl := getRealUrl(b)

	// 开始代理
	fmt.Println("------ begin server proxy ", realUrl)
	// // return
	destConn, err := getClient(realUrl)
	if err != nil {
		fmt.Println("destConn err ", err)
		// panic(err)
		return
	}
	defer destConn.Close()

	srcConn := conn

	if err != nil {
		fmt.Println("srcConn err ", err)
		// panic(err)
		return
	}
	defer srcConn.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go utils.IoCopy(errCh, srcConn, destConn)
	go utils.IoCopy(errCh, destConn, srcConn)
	select {
	case err := <-errCh:
		fmt.Println("------ errCh ", err.(error))
		return
	case <-stopCh:
		fmt.Println("------ stopCh")
		return
	}
	// time.Sleep(time.Second * 10)
	fmt.Println("------ proxy success")
}
