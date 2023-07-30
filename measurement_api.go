package saia

import (
	"context"
	"fmt"
	"net/http"
)

type MeasurementAPI interface {
	GetMeasurement(ctx context.Context, measurementID int) (*Measurement, error)
}

type measurementAPI struct {
	*apiClient
}

func newMeasurementAPI(apiClient *apiClient) *measurementAPI {
	return &measurementAPI{apiClient}
}

func (m *measurementAPI) GetMeasurement(ctx context.Context, measurementID int) (*Measurement, error) {
	url := m.buildURL(fmt.Sprintf("/measurements/mtm-widgets/%d/", measurementID))
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var measurement Measurement
	if err := m.do(req, &measurement); err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return &measurement, nil
}
