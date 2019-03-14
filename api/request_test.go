package api

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestAddURLQuerys(t *testing.T) {
	type args struct {
		rawURL string
		params map[string]string
	}

	u, _ := url.Parse("http://www.test.com?test=testval")

	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name: "SuccessfulParse",
			args: args{
				rawURL: "http://www.test.com",
				params: map[string]string{
					"test": "testval",
				},
			},
			want:    u,
			wantErr: false,
		},
		{
			name: "Parse Error",
			args: args{
				rawURL: "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
				params: map[string]string{
					"test": "testval",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddURLQuerys(tt.args.rawURL, tt.args.params)

			if (err != nil) != tt.wantErr {
				t.Errorf("AddURLQuerys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddURLQuerys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendGet(t *testing.T) {
	type args struct {
		url     string
		headers map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   *http.Response
		wantErr bool
	}{
		{
			name: "Request",
			args: args{
				url:     "http://test.com/v1/sendget",
				headers: map[string]string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byts, resp, err := SendGet(tt.args.url, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(byts) == 0 {
				t.Errorf("SendGet() No bytes = %v, but want bytes", byts)
			}
			if resp.StatusCode != 400 {
				t.Errorf("SendGet() Response StatusCode = %v, but expected a 400", resp.StatusCode)
			}

		})
	}
}
