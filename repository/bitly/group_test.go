package bitly

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/kwtucker/bitlyapi/api"
	"github.com/kwtucker/bitlyapi/models"
)

func TestGetGroup(t *testing.T) {
	type args struct {
		url     *url.URL
		headers map[string]string
		group   *models.Group
	}

	tests := []struct {
		name     string
		url      string
		args     args
		wantbyts bool
		response *http.Response
		wantErr  bool
	}{
		{
			name: "Get Group Successful",
			url:  models.BitlyAPIV4 + "group",
			args: args{
				// url:     parsedURL,
				headers: map[string]string{},
				group: &models.Group{
					Pagination: models.Pagination{
						Total: 0,
						Size:  0,
						Prev:  "",
						Page:  0,
						Next:  "",
					},
					Links: []models.Link{models.Link{
						References:     map[string]string{},
						Archived:       false,
						Tags:           []string{},
						CreatedAt:      "2018-06-29T19:55:07+0000",
						Title:          "test",
						Deeplinks:      []models.Deeplink{},
						CreatedBy:      "test",
						LongURL:        "https://www.test.com",
						ClientID:       "1234",
						CustomBitlinks: []string{"http://bit.ly/1234"},
						Link:           "http://bit.ly/1234",
						ID:             "bit.ly/1234",
					},
					},
				},
			},
			wantbyts: false,
			response: &http.Response{StatusCode: 200},
			wantErr:  false,
		},
		{
			name: "Get Group Error",
			url:  models.BitlyAPIV4 + "error",
			args: args{
				// url:     parsedURL,
				headers: map[string]string{},
				group:   &models.Group{},
			},
			wantbyts: true,
			response: &http.Response{StatusCode: 400},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		parsedURL, _ := api.AddURLQueries(tt.url, map[string]string{})
		tt.args.url = parsedURL
		t.Run(tt.name, func(t *testing.T) {
			g := &models.Group{}
			byts, resp, err := GetGroup(tt.args.url, tt.args.headers, g)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(g, tt.args.group) {
				bytg, _ := json.MarshalIndent(g, "", "   ")
				bytt, _ := json.MarshalIndent(tt.args.group, "", "   ")

				t.Errorf("GetGroup() got = %+v, want %+v", string(bytg), string(bytt))
			}

			if resp.StatusCode != tt.response.StatusCode {
				t.Errorf("GetUser() response status code = %v, want %v", resp.StatusCode, tt.response.StatusCode)
			}

			if tt.wantbyts {
				errorResp := &models.ErrorResponse{}
				_ = json.Unmarshal(byts, errorResp)
				if errorResp.Message != "INVALID" {
					t.Errorf("GetUser() errorResp message = %+v, want INVALID", errorResp.Message)
				}
			}
		})
	}
}
