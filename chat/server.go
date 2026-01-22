package main

import (
	"context"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"golang.org/x/time/rate"
)

type echoServer struct {
	id   int16
	// Ooohhh... actually ingenius... route all logging from all servers instances into a single logger, nice.
	logf func(f string, v ...any)
}

func (s *echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"echo"},
	})
	if err != nil {
		s.logf("server-%02d: %v", s.id, err)
		return
	}
	defer c.CloseNow()

	switch c.Subprotocol() {
	case "echo": 
		l := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
		for {

		}
	}
}

func echo(c *websocket.Conn, l *rate.Limiter) error {
	// hmm... so this context should arrive from something (encapsulation of goroutine?), using backgroud for a simple server for now. 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
}
