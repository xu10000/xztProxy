package tcp

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net"
	"strconv"
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
	if n < hashLen {
		fmt.Println("------ print n < hashLen", hashLen)
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

	// 开始代理
	fmt.Println("------ begin server proxy ", realUrl)
	// // return
	destConn, err := getClient(realUrl)
	if err != nil {
		fmt.Println("destConn err ", err)
		if srcConn != nil {
			srcConn.Close()
		}
		if destConn != nil {
			destConn.Close()
		}
		return
	}

	defer func() {
		if srcConn != nil {
			srcConn.Close()
		}
		if destConn != nil {
			destConn.Close()
		}
		fmt.Println("------ server proxy close", realUrl)
	}()

	if err != nil {
		fmt.Println("srcConn err ", err)
		// panic(err)
		return
	}

	stopChan := make(chan int)
	go func() {
		if _, err := io.Copy(srcConn, destConn); err != nil {
			fmt.Println("------iocopy print1", err)
			stopChan <- 1
		}
	}()
	go func() {
		//进行全双工的双向数据拷贝（中继）
		if _, err := io.Copy(destConn, srcConn); err != nil {
			fmt.Println("------iocopy print2", err)
			stopChan <- 1
		} //relay:dst->src
		fmt.Println("------ end server proxy ", realUrl)
		stopChan <- 1
	}()

	<-stopChan
	fmt.Println("------ end server proxy ", realUrl)

}
