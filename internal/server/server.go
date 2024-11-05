package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/meraiku/go_fundamentals/internal/client"
)

var (
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
)

type Api struct {
	client client.Client
}

func New(client client.Client) *Api {
	return &Api{
		client: client,
	}
}

func (api *Api) Run(ctx context.Context) error {

	if port == "" {
		port = "8080"
	}

	url := net.JoinHostPort(host, port)

	handler := api.SetupHandlers()

	srv := http.Server{
		Addr:         url,
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 120,
	}

	go func() {
		<-ctx.Done()
		srv.Close()
	}()

	fmt.Printf("Starting server on %s\n", url)

	return srv.ListenAndServe()
}
