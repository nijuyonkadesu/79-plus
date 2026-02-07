package main

import (
	"context"
	"encoding/json"
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

	for {
		err = echo(c, l)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			s.logf("server-%02d: echo: %w", s.id, err)
			return
		}

	}
}

func echo(c *websocket.Conn, l *rate.Limiter) error {
	ctx, cancel := context.WithCancel(context.Background())
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
	return nil
}

type ForwardMessageServer struct {
	id   int16
	logf func(f string, v ...any)
}

func (s *ForwardMessageServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"subscribe"},
	})
	if err != nil {
		s.logf("server-%02d: %v", s.id, err)
		return
	}
	ctx := r.Context()
	// TODO: perform all the checks, liveTransit should simply forward the messages to connection
	err = liveTransit(ctx, c)
	if err != nil {
		s.logf("server-%02d: %v", s.id, err)
		return
	}
}

func liveTransit(ctx context.Context, c *websocket.Conn) error {
	// TODO: ideally do the transit till the connection is closed
	// one goroutine to receive from queue, and another to listen for connection close?
	for {
		// TODO: these messages should ideally be received from a queue
		dummy := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

		for _, content := range dummy {
			msg := map[string]string{
				"message": content,
			}
			data, _ := json.Marshal(msg)
			response := WSMessage{
				Type:    "msessage",
				Payload: json.RawMessage(data),
			}
			err := wsjson.Write(ctx, c, &response)
			if err != nil {
				return err
			}
		}
	}
}
