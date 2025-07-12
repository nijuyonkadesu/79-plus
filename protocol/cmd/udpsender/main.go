package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// sending data to remote UDP server.
	// start the sever by running `nc -u -l 42069`
	upstream, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, upstream)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		data, err := reader.ReadString(';')
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte(data))
	}
}
