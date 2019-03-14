package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/kwtucker/bitlyapi/api"
	"github.com/kwtucker/bitlyapi/models"
	"github.com/kwtucker/bitlyapi/repository/bitly"
)

// HandleGroupsCountries will get the average or the total number of clicks, per country, within the last number of days, for the Bitlinks in a user's default group.
func HandleGroupsCountries(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{}
	accessToken := r.Header.Get("Authorization")
	query := r.URL.Query()
	units := query.Get("units")
	unit := query.Get("unit")

	headers := map[string]string{
		"Authorization": accessToken,
	}

	user, byt, response, err := bitly.GetUser(models.BitlyAPIV4+"user", headers)
	if err != nil {
		if response != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.StatusCode)
			w.Write(byt)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	group := &models.Group{}

	url := models.BitlyAPIV4 + "groups/" + user.DefaultGroupGUID + "/bitlinks"
	parsedURL, err := api.AddURLQueries(url, map[string]string{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	byt, response, err = bitly.GetGroup(parsedURL, headers, group)
	if err != nil {
		if response != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(response.StatusCode)
			w.Write(byt)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if units != "" {
		params["units"] = units
	}

	if unit != "" {
		params["unit"] = unit
	}

	countryMetrics := []models.Metric{}
	average, _ := strconv.ParseBool(query.Get("average"))

	countryMetricsM := map[string][]int64{}
	countryMetricsT := map[string]int64{}

	for _, val := range group.Links {
		parsedCountryURL, err := api.AddURLQueries(models.BitlyAPIV4+"bitlinks/"+val.ID+"/countries", params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		metric, byt, response, err := bitly.GetBitlinkCountryClickMetrics(parsedCountryURL, headers)
		if err != nil {
			if response != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(response.StatusCode)
				w.Write(byt)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// For just total clicks
		for _, m := range metric.Metrics {
			countryMetricsT[m.Value] += m.Clicks
		}

		// For a average
		for _, m := range metric.Metrics {
			countryMetricsM[m.Value] = append(countryMetricsM[m.Value], m.Clicks)
		}

	}

	if average {
		countryMetrics = bitly.GetMetricsClickAveragePerCountry(countryMetricsM)
	} else {
		for country, clicks := range countryMetricsT {
			metric := models.Metric{}
			metric.Clicks = clicks
			metric.Value = country
			countryMetrics = append(countryMetrics, metric)
		}
	}

	unitReference := time.Now().UTC().Format("2006-02-01T15:04:05-0700")

	// If a error occurs it will set the value to zero as intented. Ignoring error.
	responseUnits, _ := strconv.Atoi(units)

	averageClickPerCountry := &models.ClickMetrics{
		Unit:          unit,
		Units:         int64(responseUnits),
		UnitReference: unitReference,
		Facet:         "countries",
		Metrics:       countryMetrics,
	}

	respByt, err := json.MarshalIndent(averageClickPerCountry, "", "   ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respByt)
}
