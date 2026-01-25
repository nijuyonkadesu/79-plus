package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"golang.org/x/time/rate"
)

type echoServer struct {
	id int16
	// Ooohhh... actually ingenius... route all logging from all servers instances into a single logger, nice.
	logf func(f string, v ...any)
}

func (s *echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"echo", "chill"},
	})
	if err != nil {
		s.logf("server-%02d: %v", s.id, err)
		return
	}
	defer c.CloseNow()

	l := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)

	err = echo(c, l)
	if err != nil {
		s.logf("server-%02d: echo: %w", s.id, err)
	}
}

func echo(c *websocket.Conn, l *rate.Limiter) error {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if !l.Allow() {
			reason := "rate limit exceeded for client"
			c.Close(websocket.StatusTryAgainLater, reason)
			return fmt.Errorf("%s", reason)
		}

		var msg incomingMessage
		err := wsjson.Read(ctx, c, &msg)
		if err != nil {
			return err
		}

		var response outgoingMessage
		switch msg.Type {
		case "echo":
			response = outgoingMessage{msg.Type, msg.Message}
		case "chill":
			response = outgoingMessage{msg.Type, "chill"}
		default:
			return fmt.Errorf("unsupported message type %v received", msg.Type)
		}

		err = wsjson.Write(ctx, c, &response)
		if err != nil {
			return err
		}
	}
}
