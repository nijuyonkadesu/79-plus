package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"me.httpfrom.tcp/internal/request"
	"me.httpfrom.tcp/internal/response"
	"me.httpfrom.tcp/internal/server"
)

const port = 42069

func simpleRouter(w *response.Writer, req *request.Request) {
	if req.RequestLine.RequestTarget == "/yourproblem" {
		message := "Your problem is not my problem\n"
		w.WriteStatusLine(response.BadRequest)
		w.WriteHeaders(response.GetDefaultHeaders(len(message)))
		w.WriteBody([]byte(message))

	} else if req.RequestLine.RequestTarget == "/myproblem" {
		message := "Woopsie, my bad\n"
		w.WriteStatusLine(response.Failure)
		w.WriteHeaders(response.GetDefaultHeaders(len(message)))
		w.WriteBody([]byte(message))

	} else if strings.HasPrefix(req.RequestLine.RequestTarget, "/httpbin") {
		path := req.RequestLine.RequestTarget[len("/httpbin"):]
		// using httpbin to get some huge data
		res, err := http.Get("https://httpbin.org" + path)
		if err != nil {
			w.WriteStatusLine(response.Failure)
			w.WriteHeaders(response.GetDefaultHeaders(len(err.Error())))
			w.WriteBody([]byte(err.Error()))
			return
		}
		w.WriteStatusLine(response.OK)
		h := response.GetDefaultHeaders(0)
		h.Delete("content-length")
		h.Set("transfer-encoding", "chunked")
		w.WriteHeaders(h)
		// TODO: how to handle errors properly? (better check go's http library at this point)

		chunkedBody := make([]byte, 32)
		for {
			n, err := res.Body.Read(chunkedBody)
			if errors.Is(err, io.EOF) {
				break
			}
			w.WriteChunkedBody(chunkedBody[:n])
		}
		w.WriteChunkedBodyDone()

	} else {
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
