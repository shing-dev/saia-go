package saia

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type MeasurementAPI interface {
	GetMeasurementList(ctx context.Context, options ...GetMeasurementListOption) (*GetMeasurementListResponse, error)
	GetMeasurement(ctx context.Context, measurementID int) (*Measurement, error)
}

type measurementAPI struct {
	*apiClient
}

func newMeasurementAPI(apiClient *apiClient) *measurementAPI {
	return &measurementAPI{apiClient}
}

type GetMeasurementListParams struct {
	Page         int
	PageSize     int
	PersonGender *Gender
	Status       *string
	IsArchived   bool
}

func newGetMeasurementListParams() *GetMeasurementListParams {
	return &GetMeasurementListParams{
		Page:     1,
		PageSize: 20,
	}
}

func (g *GetMeasurementListParams) toQueryParams() url.Values {
	queryParams := url.Values{
		"page":        {strconv.Itoa(g.Page)},
		"page_size":   {strconv.Itoa(g.PageSize)},
		"is_archived": {strconv.FormatBool(g.IsArchived)},
	}
	if g.PersonGender != nil {
		queryParams["person_gender"] = []string{string(*g.PersonGender)}
	}
	if g.Status != nil {
		queryParams["status"] = []string{string(*g.Status)}
	}
	return queryParams
}

type GetMeasurementListOption func(*GetMeasurementListParams)

func GetMeasurementListOptionLimit(page int) GetMeasurementListOption {
	return func(p *GetMeasurementListParams) {
		p.Page = page
	}
}

func GetMeasurementListOptionPresence(pageSize int) GetMeasurementListOption {
	return func(p *GetMeasurementListParams) {
		p.PageSize = pageSize
	}
}

func GetMeasurementListOptionPersonGender(gender Gender) GetMeasurementListOption {
	return func(p *GetMeasurementListParams) {
		p.PersonGender = &gender
	}
}

func GetMeasurementListOptionStatus(status string) GetMeasurementListOption {
	return func(p *GetMeasurementListParams) {
		p.Status = &status
	}
}

func GetMeasurementListOptionIsArchived(isArchived bool) GetMeasurementListOption {
	return func(p *GetMeasurementListParams) {
		p.IsArchived = isArchived
	}
}

type GetMeasurementListResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []*Measurement `json:"results"`
}

func (m *measurementAPI) GetMeasurementList(ctx context.Context, options ...GetMeasurementListOption) (*GetMeasurementListResponse, error) {
	params := newGetMeasurementListParams()
	for _, opt := range options {
		opt(params)
	}

	url, err := m.buildURL("/measurements/mtm-widgets/")
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %w", err)
	}
	url.RawQuery = params.toQueryParams().Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var resp GetMeasurementListResponse
	if err := m.do(req, &resp); err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return &resp, nil
}

func (m *measurementAPI) GetMeasurement(ctx context.Context, measurementID int) (*Measurement, error) {
	url, err := m.buildURL(fmt.Sprintf("/measurements/mtm-widgets/%d/", measurementID))
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var measurement Measurement
	if err := m.do(req, &measurement); err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return &measurement, nil
}
