package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetupHandlers(t *testing.T) {
	api := New(nil)

	r := api.SetupHandlers()

	require.NotNil(t, r)
}
