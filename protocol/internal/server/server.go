package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"sync/atomic"

	"me.httpfrom.tcp/internal/request"
	"me.httpfrom.tcp/internal/response"
)

type Server struct {
	listener net.Listener
	handler  response.Handler
	isAlive  atomic.Bool
}

func (s *Server) Close() error {
	s.isAlive.Store(false)
	return s.listener.Close()
}

func Serve(port uint16, handler response.Handler) (*Server, error) {
	// by returning the pointer, the memory is allocated in heap.
	// and avoid making a copy of struct before returning it.
	server := &Server{}
	server.isAlive.Store(true)
	server.handler = handler

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

/*
Creates buffer where response is written first. If no error happens, it is written to the connection.
This way, we can avoid a case we send 200 as response, and then server crashes right after.
*/
func (s *Server) handle(conn net.Conn) {
	defer conn.Close()
	buf := bytes.NewBuffer([]byte{})
	w := response.NewWriter(buf)
	h := response.GetDefaultHeaders(0)

	r, err := request.RequestFromReader(conn)
	if err != nil {
		w.WriteStatusLine(response.BadRequest)
		w.WriteHeaders(h)
		w.WriteBody([]byte(err.Error()))
		return
	}

	s.handler(w, r) // TODO: streaming response?
	buf.WriteTo(conn)
}

/*
Connection Managers : handle(), listen() - just manages tcp connections
Application Logic	: Handler() - decouples server logic and application logic
*/
