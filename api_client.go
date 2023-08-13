package saia

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type apiClient struct {
	httpClient *http.Client
	apiHost    string
	apiKey     string
	debug      bool
}

func newAPIClient(httpClient *http.Client, apiHost string, apiKey string, debug bool) *apiClient {
	return &apiClient{
		apiKey:     apiKey,
		httpClient: httpClient,
		apiHost:    apiHost,
		debug:      debug,
	}
}

func (a *apiClient) request(req *http.Request, v any) error {
	resp, err := a.do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err == nil {
			return fmt.Errorf("failed to send request status: %s, body: %s", resp.Status, string(bodyBytes))
		}
		return fmt.Errorf("failed to send request: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}
	return nil
}

func (a *apiClient) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "APIKey "+a.apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.httpClient.Do(req)
	return resp, err
}

func (a *apiClient) buildURL(path string) (*url.URL, error) {
	u, err := url.Parse(a.apiHost + path)
	return u, err
}
