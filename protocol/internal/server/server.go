package server

import (
	"fmt"
	"log"
	"net"
	"sync/atomic"

	"me.httpfrom.tcp/internal/request"
)

type Server struct {
	listener net.Listener
	isAlive  atomic.Bool
}

func (s *Server) Close() error {
	s.isAlive.Store(false)
	return s.listener.Close()
}

func Serve(port uint16) (*Server, error) {
	// by returning the pointer, the memory is allocated in heap.
	// and avoid making a copy of struct before returning it.
	server := &Server{}
	server.isAlive.Store(true)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	server.listener = listener
	go server.listen(listener)
	return server, nil
}

func (s *Server) listen(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			if s.isAlive.Load() {
				log.Printf("Failed to accept connection: %v", err)
			}
			// return // never return, it stops the loop for everybody
		} else {
			// TODO: play with (context.Background() or smth) - for cancellation, timeouts, scopes, value propagation etc
			go s.handle(conn)
		}
	}
}

func (*Server) handle(conn net.Conn) {
	defer conn.Close()

	r, err := request.RequestFromReader(conn)
	fmt.Println(r)
	if err != nil {
		log.Printf("Failed to process connection %v", err)
	}
	response := []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello World!")
	conn.Write(response)
}
