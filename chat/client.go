package main

import (
	"io"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	log.SetFlags(log.Lshortfile)
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
			var message incomingMessage
			err := websocket.JSON.Receive(conn, &message)
			if err != nil {
				if err == io.EOF {
					log.Println("Connection closed by the server.")
				}
				log.Printf("Error reading message: %v\n", err)
				return
			}
			log.Printf("Received message: %s\n", message)
		}
	}()

	go func() {
		messageToSend := outgoingMessage{Type: "echo", Message: "Hello from the client!"}
		err = websocket.JSON.Send(conn, messageToSend)
		if err != nil {
			log.Printf("Error sending message: %v\n", err)
			return
		}
		log.Printf("Sent message: %s\n", messageToSend)
		time.Sleep(3 * time.Second)

		messageToSend = outgoingMessage{Type: "echo", Message: "hello again?"}
		err = websocket.JSON.Send(conn, messageToSend)
		if err != nil {
			log.Printf("Error sending message: %v\n", err)
			return
		}
		log.Printf("Sent message: %v\n", messageToSend)
		time.Sleep(3 * time.Second)
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// Hmmm... this blocks execution..., and when one of the case happens, the program resumes...
	// maybe try context.NotifyContext instead of manually sending values using channels
	select {
	case <-done:
		log.Println("Client finished.")
	case <-sig:
		log.Println("Interrupt received, closing connection.")
		err := conn.Close()
		if err != nil {
			log.Printf("Error sending close message: %v\n", err)
		}
	}
}
