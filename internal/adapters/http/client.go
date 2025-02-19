package http

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sony/gobreaker/v2"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	cb         *gobreaker.CircuitBreaker[[]byte]
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	settings := gobreaker.Settings{
		Name:        "HTTP Client",
		MaxRequests: 5,
		Interval:    60 * time.Second,
		Timeout:     30 * time.Second,
	}
	cb := gobreaker.NewCircuitBreaker[[]byte](settings)

	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		cb: cb,
	}
}

func (c *Client) Post(path string, body io.Reader) ([]byte, error) {
	fullURL := c.BaseURL + path

	responseBody, err := c.cb.Execute(func() ([]byte, error) {
		resp, err := c.HTTPClient.Post(fullURL, "application/json", body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return responseBody, nil
	})

	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
