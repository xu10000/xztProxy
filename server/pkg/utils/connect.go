package utils

import (
	"fmt"
	"io"
	"net"
	"sync"
)

func IoCopy(wg *sync.WaitGroup, src, dest net.Conn) {

	//进行全双工的双向数据拷贝（中继）
	if _, err := io.Copy(dest, src); err != nil {
		fmt.Println("------iocopy print", err)
	} //relay:dst->src
	wg.Done()
}
