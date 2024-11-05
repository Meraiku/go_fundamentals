package tests

import (
	"io"
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

func TestAPISearch(t *testing.T) {

	c := http.Client{}

	resp, err := c.Get("http://localhost:8080/search?q=hello")
	require.NoError(t, err)
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Contains(t, string(data), "hello")
}
