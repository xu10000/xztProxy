package http

import (
	"client/socket"
	"net"
)

type ClientImpl struct {
	Conn net.Conn
}

// func (c *ClientImpl) SetConn(network string, addr string) (net.Conn, error) {
// 	// dial server
// 	var err error
// 	c.Conn, err = net.Dial(network, addr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return c.Conn, nil
// }

func (c *ClientImpl) Dial(addr string) (net.Conn, error) {
	var err error
	c.Conn, err = net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return c.Conn, nil
}

func NewClient() socket.Client {

	return &ClientImpl{nil}

}
