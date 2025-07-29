package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port = 42069

func main() {
	// TODO: implement in internal/server/server.go
	server.ListenAndServe()
	server, err := server.Serve()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer server.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
