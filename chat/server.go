package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"golang.org/x/time/rate"
)

type echoServer struct {
	id int16
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
	
	l := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)

	switch c.Subprotocol() {
	default:
		// not sure if this is needed...
		s.logf("server-%02d: unknown subprotocol %q", s.id, c.Subprotocol())
		c.Close(websocket.StatusPolicyViolation, "unknown message type")

	case "echo":
		// wait.. this is rate-limiter for http clients... what? no, it's for both server and client.
		err := echo(c, l)
		if err != nil {
			s.logf("server-%02d: echo: %w", s.id, err)
		}
		c.CloseNow()
	}
}

func echo(c *websocket.Conn, l *rate.Limiter) error {
	// hmm... so this context should arrive from something (encapsulation of goroutine?), using backgroud for a simple server for now.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// TODO: read about context usage in go docs
	defer cancel()

	// rate limiter for client, Use Allow for servers.
	// coz, .Wait() blocks, and Allow() returns immediately.
	// and dam.. Reserve() is cool
	if !l.Allow() {
		reason := "rate limit exceeded for client"
		c.Close(websocket.StatusTryAgainLater, reason)
		return fmt.Errorf("%s", reason)
	}
	// allows all acceptable message types? (mainly the custom designed ones?)
	msgType, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}

	w, err := c.Writer(ctx, msgType)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to echo: %w", err)
	}

	err = w.Close()
	return err
}
