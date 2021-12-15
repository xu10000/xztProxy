package tcp

import (
	"net"
	"server/socket"
)

type ClientImpl struct {
	Conn net.Conn
}

func (c *ClientImpl) Dial(addr string) (net.Conn, error) {
	var err error

	c.Conn, err = net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return c.Conn, nil
}

func NewClient() socket.Client {

	return &ClientImpl{}

}
