package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
)

func Run() {

	if port == "" {
		port = "8080"
	}

	url := net.JoinHostPort(host, port)

	handler := SetupHandlers()

	srv := http.Server{
		Addr:         url,
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 120,
	}

	fmt.Printf("Starting server on %s\n", url)

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
