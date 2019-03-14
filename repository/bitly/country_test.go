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

func TestGetMetricsClickAveragePerCountry(t *testing.T) {
	type args struct {
		countryMetrics map[string][]int64
	}
	tests := []struct {
		name string
		args args
		want []models.Metric
	}{
		{
			name: "test",
			args: args{
				countryMetrics: map[string][]int64{
					"US": []int64{4, 4},
				},
			},
			want: []models.Metric{
				models.Metric{
					Value:  "US",
					Clicks: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMetricsClickAveragePerCountry(tt.args.countryMetrics)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMetricsClickAveragePerCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBitlinkCountryClickMetrics(t *testing.T) {
	type args struct {
		url     *url.URL
		headers map[string]string
	}
	tests := []struct {
		name          string
		url           string
		args          args
		clicksMetrics *models.ClickMetrics
		wantbyts      bool
		response      *http.Response
		wantErr       bool
	}{
		{
			name: "Successful Metrics Retrieval",
			url:  models.BitlyAPIV4 + "metrics",
			args: args{
				headers: map[string]string{},
			},
			clicksMetrics: &models.ClickMetrics{
				Units:         0,
				Facet:         "countries",
				UnitReference: "2019-03-12T23:40:36+0000",
				Unit:          "day",
				Metrics:       []models.Metric{models.Metric{Value: "US", Clicks: 0}},
			},
			wantbyts: false,
			response: &http.Response{StatusCode: 200},
			wantErr:  false,
		},
		{
			name: "Error Metrics Retrieval",
			url:  models.BitlyAPIV4 + "error",
			args: args{
				headers: map[string]string{},
			},
			wantbyts: true,
			response: &http.Response{StatusCode: 400},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		parsedURL, _ := api.AddURLQuerys(tt.url, map[string]string{})
		tt.args.url = parsedURL
		t.Run(tt.name, func(t *testing.T) {
			metrics, byts, resp, err := GetBitlinkCountryClickMetrics(tt.args.url, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBitlinkCountryClickMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(metrics, tt.clicksMetrics) {
				t.Errorf("GetBitlinkCountryClickMetrics() got = %v, want %v", metrics, tt.clicksMetrics)
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
