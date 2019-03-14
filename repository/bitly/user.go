package bitly

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kwtucker/bitlyapi/api"
	"github.com/kwtucker/bitlyapi/models"
)

func GetUser(url string, headers map[string]string) (*models.User, []byte, *http.Response, error) {

	byt, response, err := api.SendGet(url, headers)
	if err != nil {
		return nil, []byte{}, response, err
	}

	if response.StatusCode > 399 {
		return nil, byt, response, fmt.Errorf("Error getting user object")
	}

	user := &models.User{}
	err = json.Unmarshal(byt, user)
	if err != nil {
		return nil, nil, nil, err
	}

	return user, byt, response, nil
}
