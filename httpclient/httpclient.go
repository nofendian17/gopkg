package httpclient

import (
	"io"
	"net/http"
	"time"
)

// Config holds the configuration for the HTTP client.
type Config struct {
	Timeout int
}

// HTTPClient defines the interface for an HTTP client.
type HTTPClient interface {
	Get(url string, headers http.Header) (resp *http.Response, err error)
	Post(url string, headers http.Header, body io.Reader) (resp *http.Response, err error)
	Put(url string, headers http.Header, body io.Reader) (resp *http.Response, err error)
	Delete(url string, headers http.Header, body io.Reader) (resp *http.Response, err error)
}

// httpClient implements the HTTPClient interface using net/http package.
type httpClient struct {
	client *http.Client
}

// New creates a new instance of httpClient with the provided configuration.
func New(cfg Config) HTTPClient {
	return &httpClient{
		client: &http.Client{
			Transport: &http.Transport{},
			Timeout:   time.Duration(cfg.Timeout) * time.Second,
		},
	}
}

// Get sends an HTTP GET request with headers.
func (h *httpClient) Get(url string, headers http.Header) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	return h.client.Do(req)
}

// Post sends an HTTP POST request with headers.
func (h *httpClient) Post(url string, headers http.Header, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	return h.client.Do(req)
}

// Put sends an HTTP PUT request with headers.
func (h *httpClient) Put(url string, headers http.Header, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	return h.client.Do(req)
}

// Delete sends an HTTP DELETE request with headers.
func (h *httpClient) Delete(url string, headers http.Header, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	return h.client.Do(req)
}
