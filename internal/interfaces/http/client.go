package http

import "io"

//go:generate mockgen -source=client.go -destination=mock/client.go
type Client interface {
	Post(path string, body io.Reader) ([]byte, error)
}
