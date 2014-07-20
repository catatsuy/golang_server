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
	path := "/tmp/server.sock"
	unix, err := net.Listen("unix", path)
	os.Chmod(path, 0777)

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
		log.Println(http.Serve(unix, nil))
	}()

	<-sigchan

	if err := os.Remove(path); err != nil {
		log.Fatal(err)
	}
}
