package utils

import (
	"io"
	"net"
)

func IoCopy(errCh chan interface{}, src, dest net.Conn) {

	//进行全双工的双向数据拷贝（中继）
	_, err := io.Copy(src, dest) //relay:dst->src
	if err != nil {
		errCh <- err
	}
}
