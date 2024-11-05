package google

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type GoogleFetcher struct {
	url string
}

func New() *GoogleFetcher {
	return &GoogleFetcher{
		url: "https://google.com",
	}
}

func (g *GoogleFetcher) Get(path string) ([]byte, error) {
	c := http.Client{
		Timeout: time.Second * 10,
	}

	url := fmt.Sprintf("%s/%s", g.url, path)

	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode >= 400 && resp.StatusCode < 500:
		return nil, fmt.Errorf("client error: %s", resp.Status)
	case resp.StatusCode >= 500:
		return nil, fmt.Errorf("server error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
