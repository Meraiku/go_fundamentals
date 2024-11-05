package client

type Client interface {
	Get(path string) ([]byte, error)
}
