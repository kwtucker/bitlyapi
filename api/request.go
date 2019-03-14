package api

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SendGet will send a get requesst to the passed in url including the headers. It will response with the raw bytes body.
func SendGet(url string, headers map[string]string) ([]byte, *http.Response, error) {
	var httpClient http.Client
	var request *http.Request
	var response *http.Response
	var contentBuffer io.Reader
	var err error

	if request, err = http.NewRequest(http.MethodGet, url, contentBuffer); err != nil {
		errorMsg := fmt.Sprintf("Unable to create request: %v\n", err)
		return nil, nil, errors.New(errorMsg)
	}

	// Add custom headers
	for key, val := range headers {
		request.Header.Add(key, val)
	}

	if response, err = httpClient.Do(request); err != nil {
		errorMsg := fmt.Sprintf("Unable to send request: %v\n", err)
		return nil, nil, errors.New(errorMsg)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	return body, response, err
}

// AddURLQueries will parse the raw string url and encode the query parameters.
func AddURLQueries(rawURL string, params map[string]string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return u, err
	}

	// Setting url queries.
	query := u.Query()
	for key, value := range params {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()
	return u, err
}
