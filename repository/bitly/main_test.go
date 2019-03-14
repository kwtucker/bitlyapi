package bitly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/kwtucker/bitlyapi/models"
)

func TestMain(m *testing.M) {
	fmt.Println("Setting Up Test Environment")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	// mock to list out the articles

	httpmock.RegisterResponder("GET", "https://api-ssl.bitly.com/v4/user",
		func(req *http.Request) (*http.Response, error) {

			u := &models.User{}
			json.Unmarshal([]byte(user), u)
			resp, err := httpmock.NewJsonResponse(200, u)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}

			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", "https://api-ssl.bitly.com/v4/group",
		func(req *http.Request) (*http.Response, error) {

			g := &models.Group{}
			err := json.Unmarshal([]byte(group), g)
			if err != nil {
				fmt.Println(fmt.Sprintf("%+v", err))
			}
			resp, err := httpmock.NewJsonResponse(200, g)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}

			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", "https://api-ssl.bitly.com/v4/metrics",
		func(req *http.Request) (*http.Response, error) {

			c := &models.ClickMetrics{}
			err := json.Unmarshal([]byte(metrics), c)
			if err != nil {
				fmt.Println(fmt.Sprintf("%+v", err))
			}
			resp, err := httpmock.NewJsonResponse(200, c)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}

			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", "https://api-ssl.bitly.com/v4/user/error",
		func(req *http.Request) (*http.Response, error) {

			e := &models.ErrorResponse{}
			json.Unmarshal([]byte(errUserResp), e)

			resp, err := httpmock.NewJsonResponse(403, e)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", "https://api-ssl.bitly.com/v4/error",
		func(req *http.Request) (*http.Response, error) {

			e := &models.ErrorResponse{}
			json.Unmarshal([]byte(errResp), e)
			resp, err := httpmock.NewJsonResponse(400, e)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}

			return resp, nil
		},
	)

	run := m.Run()
	fmt.Println("Tearing Down Test Environment")
	os.Exit(run)
}

const user string = `{
    "created": "2017-07-02T15:14:00+0000",
    "modified": "2019-03-12T02:33:03+0000",
    "login": "testUsername",
    "is_active": true,
    "is_2fa_enabled": false,
    "name": "testName",
    "emails": [
        {
            "email": "testEmail",
            "is_primary": true,
            "is_verified": true
        }
    ],
    "is_sso_user": false,
    "default_group_guid": "testguid"
}`

const errResp string = `{
    "message": "INVALID",
    "resource": "bitlinks",
    "description": "test",
    "errors": [
        {
            "field": "test",
            "error_code": "invalid"
        }
    ]
}`

const errUserResp string = `{"message": "FORBIDDEN"}`

const group string = `{
  "links": [
    {
      "created_at": "2018-06-29T19:55:07+0000",
      "id": "bit.ly/1234",
      "link": "http://bit.ly/1234",
      "custom_bitlinks": [
        "http://bit.ly/1234"
      ],
      "long_url": "https://www.test.com",
      "title": "test",
      "archived": false,
      "created_by": "test",
      "client_id": "1234",
      "tags": [],
			"deeplinks": [],
			"references":{}
    }
  ],
  "pagination": {
    "prev": "",
    "next": "",
    "size": 0,
    "page": 0,
    "total": 0
  }
}`

const metrics string = `{
    "unit_reference": "2019-03-12T23:40:36+0000",
    "metrics": [
        {
            "value": "US",
            "clicks": 0
        }
    ],
    "units": 0,
    "unit": "day",
    "facet": "countries"
}`
