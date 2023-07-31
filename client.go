package saia

// Client is a SAIA API client.
type Client struct {
	apiClient *apiClient

	Person         PersonAPI
	MeasurementAPI MeasurementAPI
}

// NewClient creates a new SAIA client.
func NewClient(apiKey string, opt ...ClientOption) *Client {
	opts := newDefaultClientOptions()
	for _, o := range append(opt, withAPIKey(apiKey)) {
		o.apply(opts)
	}
	apiClient := newAPIClient(opts.HttpClient, opts.APIHost, opts.APIKey, opts.Debug)
	return &Client{
		apiClient:      apiClient,
		Person:         newPersonAPI(apiClient),
		MeasurementAPI: newMeasurementAPI(apiClient),
	}
}
