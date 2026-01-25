package main

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type testMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

var messages = map[int]testMessage{
	0: testMessage{Type: "echo", Message: "00-there"},
	1: testMessage{Type: "echo", Message: "01-there"},
	2: testMessage{Type: "echo", Message: "02-there"},
	3: testMessage{Type: "echo", Message: "03-there"},
	4: testMessage{Type: "chill", Message: "chill"},
}

func Test_echoServer(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(&echoServer{
		id:   1,
		logf: t.Logf,
	})
	defer s.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, s.URL, &websocket.DialOptions{
		Subprotocols: []string{"echo"},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close(websocket.StatusInternalError, "it's all over")

	for i := 0; i < 5; i++ {
		err = wsjson.Write(ctx, c, messages[i])
		if err != nil {
			t.Fatal(err)
		}

		v := make(map[string]string, 1)
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			t.Fatal(err)
		}

		if v["message"] != messages[i].Message {
			t.Fatalf("expected %v, got %v", v, messages[i])
		}
	}

	c.Close(websocket.StatusNormalClosure, "")
}
