package tcp

import (
	"crypto/sha256"
	"fmt"
	"net"
	"server/pkg/utils"
	"strconv"
	"sync"
)

var (
	hashLen int = 64
)

func getClient(addr string) (net.Conn, error) {
	cli := NewClient()
	return cli.Dial(addr)
}

func NewProxy(srcConn net.Conn) {
	// tcp最大包是65536字节，所以server端一次是可以write到完整的url路径
	var b [1024]byte
	n, err := srcConn.Read(b[:])
	if err != nil {
		fmt.Println(" ReadAll err ", err)
		// panic(err)
		return
	}
	realUrl := string(b[:n-hashLen])
	password := string(b[n-64 : n])
	localPort := strconv.Itoa(srcConn.LocalAddr().(*net.TCPAddr).Port)
	LocalPassword := fmt.Sprintf("%x", sha256.Sum256([]byte("xztProxy"+localPort)))
	if password != LocalPassword {
		fmt.Println("------ password != localPassword", LocalPassword)
		return
	}
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
