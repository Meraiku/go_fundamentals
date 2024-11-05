package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPServer(t *testing.T) {

	c := http.Client{}

	resp, err := c.Get("http://localhost:8080/")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
}
