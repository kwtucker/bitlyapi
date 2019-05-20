package bitly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kwtucker/bitlyapi/api"
	"github.com/kwtucker/bitlyapi/models"
)

// GetGroup will request a single group and fill the group object passed in. It will call for all the pages and fill the links slice.
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

	if newGroup.Pagination.Next != "" {
		params := map[string]string{
			"page": newGroup.Pagination.Next,
		}
		parsedURL, err := api.AddURLQueries(url.String(), params)
		if err != nil {
			return nil, nil, err
		}

		return GetGroup(parsedURL, headers, newGroup)
	}

	return byt, response, nil
}
