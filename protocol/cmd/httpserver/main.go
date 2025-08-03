package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
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
		// curl -X GET localhost:42069/httpbin/stream/100 -v --raw
		hasher := sha256.New()
		path := req.RequestLine.RequestTarget[len("/httpbin"):]
		length := 0
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
		h.Set("trailer", "X-Content-SHA256")
		h.Set("trailer", "X-Content-Length")
		w.WriteHeaders(h)

		// TODO: how to handle errors properly? (better check go's http library at this point)
		chunkedBody := make([]byte, 32)
		for {
			n, err := res.Body.Read(chunkedBody)
			if errors.Is(err, io.EOF) {
				break
			}
			chunk := chunkedBody[:n]
			length += len(chunk)
			// fun: try out io.MultiWriter
			hasher.Write(chunk)
			w.WriteChunkedBody(chunk)
		}
		hash := hasher.Sum(nil)
		h.Set("X-Content-SHA256", fmt.Sprintf("%x", hash))
		h.Set("X-Content-Length", fmt.Sprintf("%d", length))
		w.WriteChunkedBodyDone()
		w.WriteTrailers(h)

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
