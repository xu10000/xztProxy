package tcp

import (
	"fmt"
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
	_b := string(b[2:])
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
	// tcp最大包是65536字节，所以server端一次是可以write到完整的url路径
	var b [1024]byte
	n, err := conn.Read(b[:])
	if err != nil {
		fmt.Println(" ReadAll err ", err)
		// panic(err)
		return
	}
	//
	realUrl := getRealUrl(b[:n])

	defer func() {
		fmt.Println("------ server proxy close", realUrl)
	}()
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
	go utils.IoCopy(&wg, srcConn, destConn)
	go utils.IoCopy(&wg, destConn, srcConn)

	wg.Wait()

}
