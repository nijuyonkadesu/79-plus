package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	u, err := url.Parse("ws://127.0.0.1:7777")
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	config, err := websocket.NewConfig(u.String(), "http://127.0.0.1:7777")
	if err != nil {
		log.Fatalf("Failed to create websocket config: %v", err)
	}
	config.Protocol = []string{"echo"}

	conn, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to server.")

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			var message string
			err := websocket.Message.Receive(conn, &message)
			if err != nil {
				if err.Error() == "websocket: close sent" || err.Error() == "websocket: read limit exceeded" {
					log.Println("Server closed the connection.")
				} else {
					log.Printf("Error reading message: %v\n", err)
				}
				return
			}
			log.Printf("Received message: %s\n", message)
		}
	}()

	messageToSend := "Hello from the client!"
	err = websocket.Message.Send(conn, messageToSend)
	if err != nil {
		log.Printf("Error sending message: %v\n", err)
		return
	}
	log.Printf("Sent message: %s\n", messageToSend)
	time.Sleep(3 * time.Second)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-done:
		log.Println("Client finished.")
	case <-sig:
		log.Println("Interrupt received, closing connection.")
		err := websocket.Message.Send(conn, websocket.CloseFrame)
		if err != nil {
			log.Printf("Error sending close message: %v\n", err)
		}
	}
}
