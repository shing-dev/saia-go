package saia

import "net/http"

type ClientOptions struct {
	APIHost    string
	APIKey     string
	HttpClient *http.Client
	Debug      bool
}

func newDefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		APIHost:    "https://saia.3dlook.me/api/v2",
		HttpClient: http.DefaultClient,
		Debug:      false,
	}
}

// ClientOption is a option to change client configuration.
type ClientOption interface {
	apply(*ClientOptions)
}

type clientOptionFunc struct {
	f func(config *ClientOptions)
}

func (c *clientOptionFunc) apply(pc *ClientOptions) {
	c.f(pc)
}

func newClientOptionFunc(f func(pc *ClientOptions)) *clientOptionFunc {
	return &clientOptionFunc{
		f: f,
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return newClientOptionFunc(func(c *ClientOptions) {
		c.HttpClient = httpClient
	})
}

func WithAPIHost(apiHost string) ClientOption {
	return newClientOptionFunc(func(c *ClientOptions) {
		c.APIHost = apiHost
	})
}

// WithDebugEnabled enable debug logs
// this option is should be used in development only.
func WithDebugEnabled() ClientOption {
	return newClientOptionFunc(func(c *ClientOptions) {
		c.Debug = true
	})
}

func withAPIKey(authToken string) ClientOption {
	return newClientOptionFunc(func(c *ClientOptions) {
		c.APIKey = authToken
	})
}
