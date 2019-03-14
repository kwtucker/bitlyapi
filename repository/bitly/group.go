package bitly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kwtucker/bitlyapi/api"
	"github.com/kwtucker/bitlyapi/models"
)

func GetGroup(url *url.URL, headers map[string]string, group *models.Group) ([]byte, *http.Response, error) {

	byt, response, err := api.SendGet(url.String(), headers)
	if err != nil {
		return []byte{}, response, err
	}

	if response.StatusCode > 399 {
		return byt, response, fmt.Errorf("Error getting group object")
	}

	newGroup := &models.Group{}
	err = json.Unmarshal(byt, newGroup)
	if err != nil {
		return nil, nil, err
	}

	// Add the all of the links from this page to the parent group.
	for _, link := range newGroup.Links {
		group.Links = append(group.Links, link)
	}

	if group.Pagination.Next != "" {
		params := map[string]string{
			"page": group.Pagination.Next,
		}
		parsedURL, err := api.AddURLQueries(url.String(), params)
		if err != nil {
			return nil, nil, err
		}

		return GetGroup(parsedURL, headers, group)
	}

	return byt, response, nil
}
