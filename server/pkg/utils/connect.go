package utils

import (
	"fmt"
	"io"
	"net"
	"sync"
)

func IoCopy(wg *sync.WaitGroup, src, dest net.Conn) {
	defer wg.Done()
	//进行全双工的双向数据拷贝（中继）
	_, err := io.Copy(src, dest) //relay:dst->src
	if err != nil {
		fmt.Println("iocopy err ", err)
		// panic(err)
	}

}
