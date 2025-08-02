package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"me.httpfrom.tcp/internal/request"
	"me.httpfrom.tcp/internal/response"
	"me.httpfrom.tcp/internal/server"
)

const port = 42069

func simpleRouter(w io.Writer, req *request.Request) *response.HandlerError {
	switch req.RequestLine.RequestTarget {
	case "/yourproblem":
		message := "Your problem is not my problem\n"
		return &response.HandlerError{
			Code: response.BadRequest,
			Message: message,
		}

	case "/myproblem":
		message := "Woopsie, my bad\n"
		return &response.HandlerError{
			Code: response.Failure,
			Message: message,
		}

	default:
		body := "All good, frfr\n"
		w.Write([]byte(body))
	}

	return nil
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
