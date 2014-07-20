package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 3000})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT)
	signal.Notify(sigchan, syscall.SIGTERM)

	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/ng", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	go func() {
		log.Println(http.Serve(l, nil))
	}()

	<-sigchan
}
