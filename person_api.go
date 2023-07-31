package saia

import (
	"context"
	"fmt"
	"net/http"
)

type PersonAPI interface {
	GetPerson(ctx context.Context, personID int) (*Person, error)
}

type personAPI struct {
	*apiClient
}

func newPersonAPI(apiClient *apiClient) *personAPI {
	return &personAPI{apiClient}
}

func (m *personAPI) GetPerson(ctx context.Context, personID int) (*Person, error) {
	url := m.buildURL(fmt.Sprintf("/persons/%d/", personID))
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var person Person
	if err := m.do(req, &person); err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return &person, nil
}
