package socket

import "net"

type Listener interface {
	Accept() (net.Conn, error)
}

// type Dialer interface {
// 	Listen(string) (Listener, error)
// }

type Server interface {
	Listener
	Start()
}

type Dialer interface {
	Dial(string) (net.Conn, error)
}

type Client interface {
	Dialer
}
