package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.SetFlags(0)
	log.SetFlags(log.Lshortfile)
	err := run()
	if err != nil {
		log.Fatal(err)
	}

}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("usage: chat <port>")
	}

	l, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		return err
	}
	defer l.Close()

	log.Printf("starting server at ws://%v", l.Addr())

	s := &http.Server{
		Handler: &echoServer{
			id:   1,
			logf: log.Printf,
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	errc := make(chan error, 1)
	go func() {
		errc <- s.Serve(l)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		log.Printf("server exited: %v", err)
	case sig := <-sigs:
		log.Printf("caught ctrl+c: %v, terminating server", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return s.Shutdown(ctx)
}

//  go run main.go server.go localhost:7777

