package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		for line := range getLinesChannel(conn) {
			fmt.Println("read:", line)
		}
	}
}
