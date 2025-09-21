package server

import (
	"fmt"
	"io"
	"net"

	"github.com/AdonaIsium/tcp_to_http_review/internal/response"
)

type Server struct {
	closed bool
}

func runConnection(_s *Server, conn io.ReadWriteCloser) {
	defer conn.Close()
	headers := response.GetDefaultHeaders(0)
	response.WriteStatusLine(conn, response.StatusOK)
	response.WriteHeaders(conn, headers)
}

func runServer(s *Server, listener net.Listener) error {

	go func() {
		for {
			conn, err := listener.Accept()
			if s.closed {
				return
			}
			if err != nil {
				return
			}
			go runConnection(s, conn)
		}

	}()

	return nil
}

func Serve(port uint16) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	server := &Server{closed: false}
	go runServer(server, listener)
	return server, nil
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}
