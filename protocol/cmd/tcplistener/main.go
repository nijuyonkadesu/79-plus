package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"

	"me.httpfrom.tcp/internal/request"
)

/*
Flow: 

1. Read 8 bytes from a file and print it print as "Read: %s"
2. Aggregate the bytes till \n is found 
3. Use tcp connection (instead of reading files) 
	a. use net.Listen (tcp) connection to get the input from netcat (echo) command
	b. channel to emit the lines to a loop running in cmd/tcplistener/main.go (without "Read") prefix
4. infinite for loop in main, pass the io.ReadCloser to new method, emit the value, print the value
5. Define `internal/request` package, which can handle the HTTP protocol. 
	a. Request Line Parsing (method, path, version)  GET /api/data HTTP/1.1
	b. Field Line Parsing (headers)
	c. body
6. Create a HTTP server to return response to the callers.

*/
func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	var lineBuffer string
	data := make([]byte, 8)

	go func() {
		defer f.Close()
		defer close(out)
		for {
			n, err := f.Read(data)
			if err != nil {
				break
			}
			newData := data[:n]

			// TODO: cannot capture multiple new lines
			newlinePos := bytes.IndexByte(newData, '\n')

			if newlinePos == -1 {
				lineBuffer += string(newData)
			} else {
				out <- lineBuffer + string(newData[:newlinePos])
				lineBuffer = string(newData[newlinePos+1:])
			}
		}
		out <- lineBuffer
	}()

	return out
}

func main() {

	// send data to port using
	// echo "Do you have what it takes to be an engineer at TheStartup™?" | nc -w 1 127.0.0.1 42069
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: this is not concurrent??!!
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// TODO: move this passing connection and handling request object to a new routine by `go request.RequestFromReader(conn)`
		// TODO: verify whether this is a proper way to launch a goroutine. 
		r, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Request line:\n")
		fmt.Printf("- Method: %s\n", r.RequestLine.Method)
		fmt.Printf("- Target: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", r.RequestLine.HttpVersion)

		fmt.Printf("Headers\n")
		r.Headers.ForEach(func(n, v string) {
			fmt.Printf(" - %s: %s\n", n, v)
		})
		fmt.Printf("Body:\n")
		fmt.Print(string(r.Body))
	}
}


/* simplified view of go's net/http/server.go
func (srv *Server) Serve(l net.Listener) error {
    for {
        rw, err := l.Accept()
        if err != nil {
            // ... 
            continue
        }
        // create a new connection object
        c := srv.newConn(rw)
        // serve the connection in a new goroutine - the same pattern!
        go c.serve(context.Background())
    }
}
*/
