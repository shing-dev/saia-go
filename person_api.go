package saia

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp"
	"io"
	"net/http"
	"regexp"
)

var uuidRegexp = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")

type PersonAPI interface {
	GetPerson(ctx context.Context, personID int) (*Person, error)
	CreatePerson(ctx context.Context, params *CreatePersonParams) (*CreatePersonResponse, error)
	CreatePersonWithImages(ctx context.Context, params *CreatePersonWithImagesParams) (*CreatePersonWithImagesResponse, error)
	StartCalculation(ctx context.Context, personID int) (*StartCalculationResponse, error)
	GetTaskSet(ctx context.Context, taskSetID string) (*GetTaskSetResponse, error)
	// TODO: Add PartialUpdatePerson method to add images after creating a person
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
		return nil, fmt.Errorf("build url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var person Person
	if err := m.request(req, &person); err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	return &person, nil
}

type PhotoFlowType string

const (
	PhotoFlowTypeFriend PhotoFlowType = "friend"
	PhotoFlowTypeHand   PhotoFlowType = "hand"
)

type DeviceCoordinates struct {
	FrontPhoto *DeviceCoordinate `json:"frontPhoto"`
	SidePhoto  *DeviceCoordinate `json:"sidePhoto"`
}

type DeviceCoordinate struct {
	BetaX  float64 `json:"betaX"`
	GammaY float64 `json:"gammaY"`
	AlphaZ float64 `json:"alphaZ"`
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
	req, err := http.NewRequestWithContext(ctx, "POST", url.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var resp CreatePersonResponse
	if err := m.request(req, &resp); err != nil {
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
	SideImage         io.Reader
	DeviceCoordinates *DeviceCoordinates
	PhotoFlowType     PhotoFlowType
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
		"gender":         c.Gender,
		"height":         c.Height,
		"weight":         c.Weight,
		"front_image":    base64.StdEncoding.EncodeToString(frontImageBytes),
		"side_image":     base64.StdEncoding.EncodeToString(sideImageBytes),
		"phone_position": c.DeviceCoordinates,
		"photo_flow":     c.PhotoFlowType,
	})
}

type CreatePersonWithImagesResponse struct {
	TaskSetURL string `json:"task_set_url"`
	TaskSetID  string `json:"-"`
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
	req, err := http.NewRequestWithContext(ctx, "POST", url.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var resp CreatePersonWithImagesResponse
	if err := m.request(req, &resp); err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	taskSetID := uuidRegexp.FindStringSubmatch(resp.TaskSetURL)[0]
	resp.TaskSetID = taskSetID

	return &resp, nil
}

type StartCalculationResponse struct {
	TaskSetURL string `json:"task_set_url"`
	TaskSetID  string `json:"-"`
}

func (m *personAPI) StartCalculation(ctx context.Context, personID int) (*StartCalculationResponse, error) {
	url, err := m.buildURL(fmt.Sprintf("/persons/%d/calculate/?measurements_type=all", personID))
	if err != nil {
		return nil, fmt.Errorf("build url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	var resp StartCalculationResponse
	if err := m.request(req, &resp); err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	taskSetID := uuidRegexp.FindStringSubmatch(resp.TaskSetURL)[0]
	resp.TaskSetID = taskSetID

	return &resp, nil
}

// GetTaskSetResponse is the response of task set endpoint
// It returns taskSet when task is not finished or failed
// It returns measured person when it's successful
type GetTaskSetResponse struct {
	TaskSet *TaskSet
	Person  *Person
}

func (m *personAPI) GetTaskSet(ctx context.Context, taskSetID string) (*GetTaskSetResponse, error) {
	url, err := m.buildURL(fmt.Sprintf("/queue/%s/", taskSetID))
	if err != nil {
		return nil, fmt.Errorf("build url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := m.do(req)
	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}
	if resp.StatusCode >= 500 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err == nil {
			return nil, fmt.Errorf("failed to send request status: %s, body: %s", resp.Status, string(bodyBytes))
		}
		return nil, fmt.Errorf("failed to send request: %s", resp.Status)
	}

	respBody := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}
	respJSON, err := json.Marshal(respBody)
	if err != nil {
		return nil, err
	}

	if respBody["is_ready"] != nil {
		var taskSet TaskSet
		if err := json.Unmarshal(respJSON, &taskSet); err != nil {
			return nil, err
		}
		return &GetTaskSetResponse{TaskSet: &taskSet}, nil
	} else {
		pp.Println(respBody)
		var person Person
		if err := json.Unmarshal(respJSON, &person); err != nil {
			return nil, err
		}
		return &GetTaskSetResponse{Person: &person}, nil
	}
}
