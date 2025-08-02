package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"me.httpfrom.tcp/internal/request"
	"me.httpfrom.tcp/internal/response"
	"me.httpfrom.tcp/internal/server"
)

const port = 42069

func simpleRouter(w *response.Writer, req *request.Request) {
	switch req.RequestLine.RequestTarget {
	case "/yourproblem":
		message := "Your problem is not my problem\n"
		w.WriteStatusLine(response.BadRequest)
		w.WriteHeaders(response.GetDefaultHeaders(len(message)))
		w.WriteBody([]byte(message))

	case "/myproblem":
		message := "Woopsie, my bad\n"
		w.WriteStatusLine(response.Failure)
		w.WriteHeaders(response.GetDefaultHeaders(len(message)))
		w.WriteBody([]byte(message))

	default:
		message := "All good, frfr\n"
		w.WriteStatusLine(response.OK)
		w.WriteHeaders(response.GetDefaultHeaders(len(message)))
		w.WriteBody([]byte(message))
	}
}

func main() {
	s, err := server.Serve(port, simpleRouter)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer s.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
