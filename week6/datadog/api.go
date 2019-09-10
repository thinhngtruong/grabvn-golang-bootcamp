package datadog

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

const (
	baseURL       = "https://app.datadoghq.com"
	eventEndpoint = "/api/v1/events?"
	gaugeEndpoint = "/api/v1/series?"
)

// Datadog -> Struct for the datadog API
type Datadog struct {
	apiKey string
	appKey string
}

type event struct {
	Title     string
	Text      string
	Tags      []string
	AlertType string
}

type series struct {
	Series []gauge
}

type gauge struct {
	Metric string
	Points float64
	Type   string
	Tags   []string
}

// Init -> This is the constructor function.
func Init(apiKey, appKey string) (dd Datadog) {
	d := Datadog{
		apiKey: apiKey,
		appKey: appKey,
	}
	return d
}

func queryBuilder(baseURL, apiKey, appKey string) string {
	params := url.Values{}
	params.Add("api_key", apiKey)
	params.Add("application_key", appKey)
	finalURL := baseURL + params.Encode()
	return finalURL
}

func buildEvent(title, text string, tags []string, alertType string) event {
	if alertType != "info" && alertType != "warning" && alertType != "error" {
		log.Println("Invalid alert type.")
		panic("Invalid alert type.")
	}
	evt := event{
		Title:     title,
		Text:      text,
		Tags:      tags,
		AlertType: alertType,
	}
	return evt
}

func buildMetric(metric string, points float64, ty string, tags []string) series {
	g := gauge{
		Metric: metric,
		Points: points,
		Type:   ty,
		Tags:   tags,
	}
	arr := []gauge{g}
	s := series{
		Series: arr,
	}
	return s
}

// PostEvent -> Post event to datadog using API endpoint.
func (dd *Datadog) PostEvent(title, text string, tags []string, alertType string) bool {
	evt := buildEvent(title, text, tags, alertType)
	jsonStr, err := json.Marshal(evt)
	if err != nil {
		log.Println(err)
		return false
	}
	url := queryBuilder(baseURL+eventEndpoint, dd.apiKey, dd.appKey)
	log.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	if (resp.StatusCode)/10 == 20 {
		return true
	}
	return false
}

// SendMetric -> Send out a metric to datadog
func (dd *Datadog) SendMetric(metric string, points float64, ty string, tags []string) bool {
	series := buildMetric(metric, points, ty, tags)
	jsonStr, err := json.Marshal(series)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(string(jsonStr))
	url := queryBuilder(baseURL+gaugeEndpoint, dd.apiKey, dd.appKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	if (resp.StatusCode)/10 == 20 {
		return true
	}
	return false
}
