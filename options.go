package terminal

import (
	"net/http"
	"strings"
	"time"
)

// ClientOption is a function that modifies a Client configuration.
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL for the Terminal Africa API client.
func WithBaseURL(urlStr string) ClientOption {
	return func(c *Client) {
		if urlStr != "" {
			c.baseURL = strings.TrimRight(urlStr, "/")
		}
	}
}

// WithHTTPClient sets a custom http.Client for network requests.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		if client != nil {
			c.httpClient = client
		}
	}
}

// WithTimeout sets a timeout duration for the default HTTP client.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		if c.httpClient != nil {
			c.httpClient.Timeout = timeout
		}
	}
}

// WithSecretKey sets or overrides the API secret key.
func WithSecretKey(secretKey string) ClientOption {
	return func(c *Client) {
		c.secretKey = secretKey
	}
}
