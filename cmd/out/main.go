package main

import (
	"context"
	"log"

	"github.com/meraiku/go_fundamentals/internal/client/google"
	"github.com/meraiku/go_fundamentals/internal/server"
)

func main() {
	client := google.New()
	api := server.New(client)

	ctx := context.TODO()

	err := api.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
