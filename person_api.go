package saia

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PersonAPI interface {
	GetPerson(ctx context.Context, personID int) (*Person, error)
	CreatePerson(ctx context.Context, params *CreatePersonParams) (*CreatePersonResponse, error)
	CreatePersonWithImages(ctx context.Context, params *CreatePersonWithImagesParams) (*CreatePersonWithImagesResponse, error)
}

type personAPI struct {
	*apiClient
}

func newPersonAPI(apiClient *apiClient) *personAPI {
	return &personAPI{apiClient}
}

func (m *personAPI) GetPerson(ctx context.Context, personID int) (*Person, error) {
	url, err := m.buildURL(fmt.Sprintf("/persons/%d/", personID))
	if err != nil {
		return nil, fmt.Errorf("failed to build url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var person Person
	if err := m.do(req, &person); err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return &person, nil
}

type CreatePersonParams struct {
	// Gender of person, male or female
	Gender Gender `json:"gender"`
	// Height of person, in cm
	Height int `json:"height"`
	// Weight of person, in kg
	Weight float64 `json:"weight"`
}

type CreatePersonResponse struct {
	ID     int     `json:"id"`
	URL    string  `json:"url"`
	Gender Gender  `json:"gender"`
	Height int     `json:"height"`
	Weight float64 `json:"weight"`
}

func (m *personAPI) CreatePerson(ctx context.Context, params *CreatePersonParams) (*CreatePersonResponse, error) {
	url, err := m.buildURL("/persons/?measurements_type=all")
	if err != nil {
		return nil, fmt.Errorf("build url: %w", err)
	}
	reqBody, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("marshal params to json: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var resp CreatePersonResponse
	if err := m.do(req, &resp); err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	return &resp, nil
}

type CreatePersonWithImagesParams struct {
	// Gender of person, male or female
	Gender Gender
	// Height of person, in cm
	Height int
	// Weight of person, in kg
	Weight float64
	// FrontImage is front image file
	FrontImage io.Reader
	// SideImage is side image file
	SideImage io.Reader
}

func (c *CreatePersonWithImagesParams) toJSON() ([]byte, error) {
	frontImageBytes, err := io.ReadAll(c.FrontImage)
	if err != nil {
		return nil, fmt.Errorf("read front image: %w", err)
	}
	sideImageBytes, err := io.ReadAll(c.SideImage)
	if err != nil {
		return nil, fmt.Errorf("read side image: %w", err)
	}

	return json.Marshal(map[string]any{
		"gender":     c.Gender,
		"height":     c.Height,
		"weight":     c.Weight,
		"frontImage": base64.StdEncoding.EncodeToString(frontImageBytes),
		"sideImage":  base64.StdEncoding.EncodeToString(sideImageBytes),
	})
}

type CreatePersonWithImagesResponse struct {
	TaskSetURL string `json:"task_set_url"`
}

func (m *personAPI) CreatePersonWithImages(ctx context.Context, params *CreatePersonWithImagesParams) (*CreatePersonWithImagesResponse, error) {
	url, err := m.buildURL("/persons/?measurements_type=all")
	if err != nil {
		return nil, fmt.Errorf("build url: %w", err)
	}
	reqBody, err := params.toJSON()
	if err != nil {
		return nil, fmt.Errorf("convert params to json: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var resp CreatePersonWithImagesResponse
	if err := m.do(req, &resp); err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	return &resp, nil
}
