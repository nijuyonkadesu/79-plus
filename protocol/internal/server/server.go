package server

import "net"

type serverState string

var (
	stateInit serverState = "init"
	stateDone serverState = "done"
)

type Server struct {
	state serverState
}

func (*Server) Close() error {

	return nil
}

func (*Server) Listen() {
}

func (*Server) Serve(port int) (*Server, error) {
	// by returning the pointer, the memory is allocated in heap.
	// and avoid making a copy of struct before returning it.
	server := &Server{
		state: stateInit,
	}

	return server, nil
}

// dispatch this method using goroutine
func (*Server) handle(conn net.Conn) {
}
