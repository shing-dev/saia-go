package saia

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_personAPI_GetPerson(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx      context.Context
		personID int
	}
	tests := []struct {
		name           string
		args           args
		resp           string
		respStatusCode int
		want           *Person
		wantErr        bool
	}{
		{
			name: "Successful response",
			args: args{ctx: context.Background(), personID: 1021366},
			// TODO: Add more fields to the response
			resp: `{
  "id": 1021366
}`,
			want: &Person{
				ID: 1021366,
			},
		},
		{
			name:           "Error response",
			args:           args{ctx: context.Background(), personID: 0000},
			resp:           `{"error": "invalid api key"}`,
			respStatusCode: 409,
			wantErr:        true,
		},
		{
			name:    "Invalid json response",
			args:    args{ctx: context.Background(), personID: 0000},
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
			m := mockPersonAPI(t, tt.resp, statusCode)

			got, err := m.GetPerson(tt.args.ctx, tt.args.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetPerson() (-got, +want)\n%s", diff)
			}
		})
	}
}

func mockPersonAPI(t *testing.T, response string, status int) *personAPI {
	t.Helper()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, response)
	})
	s := httptest.NewServer(h)
	return &personAPI{
		&apiClient{
			httpClient: http.DefaultClient,
			apiHost:    s.URL,
		},
	}
}
