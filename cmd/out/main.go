package main

import (
	"context"

	"github.com/meraiku/go_fundamentals/internal/client/google"
	"github.com/meraiku/go_fundamentals/internal/server"
)

func main() {
	client := google.New()
	api := server.New(client)

	ctx := context.TODO()

	api.Run(ctx)
}
