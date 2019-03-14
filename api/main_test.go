package api

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

	httpmock.RegisterResponder("GET", "http://test.com/v1/sendget",
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

const errResp string = `{"message": "BAD_REQUEST"}`
