package server

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerRun(t *testing.T) {
	api := New(nil)
	ctx, cancel := context.WithCancel(context.Background())

	var err error
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err = api.Run(ctx)
	}()

	cancel()

	require.NoError(t, err)
}
