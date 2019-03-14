package bitly

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"

	"github.com/kwtucker/bitlyapi/api"
	"github.com/kwtucker/bitlyapi/models"
)

func GetBitlinkCountryClickMetrics(url *url.URL, headers map[string]string) (*models.ClickMetrics, []byte, *http.Response, error) {

	byt, response, err := api.SendGet(url.String(), headers)
	if err != nil {
		return nil, []byte{}, response, err
	}

	if response.StatusCode > 399 {
		return nil, byt, response, fmt.Errorf("Error geting Bitlink country click metrics the Bitlink")
	}

	clickMetrics := &models.ClickMetrics{}
	err = json.Unmarshal(byt, clickMetrics)
	if err != nil {
		return nil, nil, nil, err
	}

	return clickMetrics, byt, response, nil
}

func GetMetricsClickAveragePerCountry(countryMetrics map[string][]int64) []models.Metric {
	metricsAverage := []models.Metric{}

	for key, val := range countryMetrics {
		m := models.Metric{}
		m.Value = key

		var total float64
		for _, number := range val {
			total += float64(number)
		}

		m.Clicks = int64(math.Round(total / float64(len(val))))
		metricsAverage = append(metricsAverage, m)
	}
	return metricsAverage
}
