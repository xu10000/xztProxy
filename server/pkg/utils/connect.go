package utils

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

func IoCopy(wg *sync.WaitGroup, errCh chan interface{}, src, dest net.Conn, flag int) {

	//进行全双工的双向数据拷贝（中继）
	_, err := io.Copy(dest, src) //relay:dst->src
	t := time.Now()
	fmt.Println("------iocopy print", t, err, flag)
	if err != nil {
		errCh <- err
	}
	wg.Done()
}
