package tcp

import (
	"fmt"
	"net"
	"os"
	"server/socket"
)

type serverImpl struct {
	Listener socket.Listener
}

func (s *serverImpl) Accept() (net.Conn, error) {
	return s.Listener.Accept()
}

func (s *serverImpl) Start() {
	for {
		_, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		// go s.handle(conn)
		fmt.Println("------ start proxy-----")

	}
}

func NewServer(HOST string) socket.Server {
	l, err := net.Listen("tcp", HOST)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	// defer l.Close()
	fmt.Println("Listening on " + HOST)

	return &serverImpl{l}
}
