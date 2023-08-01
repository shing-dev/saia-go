package saia

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/shing-dev/saia-go/pkg/convutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_measurementAPI_GetMeasurementList(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name           string
		args           args
		resp           string
		respStatusCode int
		want           *GetMeasurementListResponse
		wantErr        bool
	}{
		{
			name: "Successful response",
			args: args{ctx: context.Background()},
			// TODO: Add more fields to the response
			resp: `{
  "count": 22,
  "next": "https://saia.3dlook.me/api/v2/measurements/mtm-widgets/?is_archived=false&page=1&page_size=20",
  "results": []
}`,
			want: &GetMeasurementListResponse{
				Count:   22,
				Next:    convutil.ToPointer("https://saia.3dlook.me/api/v2/measurements/mtm-widgets/?is_archived=false&page=1&page_size=20"),
				Results: []*Measurement{},
			},
		},
		{
			name:           "Error response",
			args:           args{ctx: context.Background()},
			resp:           `{"error": "invalid api key"}`,
			respStatusCode: 409,
			wantErr:        true,
		},
		{
			name:    "Invalid json response",
			args:    args{ctx: context.Background()},
			resp:    "{",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			statusCode := 200
			if tt.respStatusCode > 0 {
				statusCode = tt.respStatusCode
			}
			m := mockMeasurementAPI(t, tt.resp, statusCode)

			got, err := m.GetMeasurementList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMeasurementList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetMeasurementList() (-got, +want)\n%s", diff)
			}
		})
	}
}

func Test_measurementAPI_GetMeasurement(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx           context.Context
		measurementID int
	}
	tests := []struct {
		name           string
		args           args
		resp           string
		respStatusCode int
		want           *Measurement
		wantErr        bool
	}{
		{
			name: "Successful response",
			args: args{ctx: context.Background(), measurementID: 1021366},
			// TODO: Add more fields to the response
			resp: `{
  "id": 1021366
}`,
			want: &Measurement{
				ID: 1021366,
			},
		},
		{
			name:           "Error response",
			args:           args{ctx: context.Background(), measurementID: 0000},
			resp:           `{"error": "invalid api key"}`,
			respStatusCode: 409,
			wantErr:        true,
		},
		{
			name:    "Invalid json response",
			args:    args{ctx: context.Background(), measurementID: 0000},
			resp:    "{",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			statusCode := 200
			if tt.respStatusCode > 0 {
				statusCode = tt.respStatusCode
			}
			m := mockMeasurementAPI(t, tt.resp, statusCode)

			got, err := m.GetMeasurement(tt.args.ctx, tt.args.measurementID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMeasurement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetMeasurement() (-got, +want)\n%s", diff)
			}
		})
	}
}

func mockMeasurementAPI(t *testing.T, response string, status int) *measurementAPI {
	t.Helper()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, response)
	})
	s := httptest.NewServer(h)
	return &measurementAPI{
		&apiClient{
			httpClient: http.DefaultClient,
			apiHost:    s.URL,
		},
	}
}
