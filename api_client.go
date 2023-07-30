package saia

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (a *apiClient) do(req *http.Request, v any) error {
	req.Header.Set("Authorization", "APIKey "+a.apiKey)
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to send request: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("failed to decode response body: %w", err)
	}
	return nil
}

func (a *apiClient) buildURL(path string) string {
	return a.apiHost + path
}
