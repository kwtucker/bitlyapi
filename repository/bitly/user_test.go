package bitly

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/kwtucker/bitlyapi/models"
)

func TestGetUser(t *testing.T) {
	type args struct {
		headers map[string]string
		url     string
	}
	tests := []struct {
		name     string
		args     args
		user     *models.User
		wantbyts bool
		response *http.Response
		wantErr  bool
	}{
		{
			name: "Successful User Response",
			args: args{
				headers: map[string]string{},
				url:     models.BitlyAPIV4 + "user",
			},
			user: &models.User{
				Created:      "2017-07-02T15:14:00+0000",
				Modified:     "2019-03-12T02:33:03+0000",
				Login:        "testUsername",
				IsActive:     true,
				Is2FaEnabled: false,
				Name:         "testName",
				Emails: []models.Email{
					models.Email{
						Email:      "testEmail",
						IsPrimary:  true,
						IsVerified: true,
					},
				},
				IsSsoUser:        false,
				DefaultGroupGUID: "testguid",
			},
			response: &http.Response{StatusCode: 200},
			wantErr:  false,
			wantbyts: false,
		},
		{
			name: "Error User Authorization",
			args: args{
				headers: map[string]string{},
				url:     models.BitlyAPIV4 + "user/error",
			},
			user:     nil,
			response: &http.Response{StatusCode: 403},
			wantErr:  true,
			wantbyts: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			user, byts, resp, err := GetUser(tt.args.url, tt.args.headers)

			if !reflect.DeepEqual(user, tt.user) {
				t.Errorf("GetUser() got = %+v, want %+v", user, tt.user)
			}

			if len(byts) == 0 && tt.wantbyts {
				t.Errorf("GetUser() []byte slice is empty, but wanted bytes %v", tt.wantbyts)
			}

			if resp.StatusCode != tt.response.StatusCode {
				t.Errorf("GetUser() response status code = %v, want %v", resp.StatusCode, tt.response.StatusCode)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantbyts {
				errorResp := &models.ErrorResponse{}
				_ = json.Unmarshal(byts, errorResp)

				if errorResp.Message != "FORBIDDEN" {
					t.Errorf("GetUser() errorResp message = %+v, want FORBIDDEN", errorResp.Message)
				}

			}
		})
	}
}
