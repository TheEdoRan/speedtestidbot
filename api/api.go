package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ServerResponse is the response object from Speedtest API.
type ServerResponse struct {
	URL             string `json:"url"`
	Lat             string `json:"lat"`
	Lon             string `json:"lon"`
	Distance        int    `json:"distance"`
	Name            string `json:"name"`
	Country         string `json:"country"`
	Cc              string `json:"cc"`
	Sponsor         string `json:"sponsor"`
	ID              string `json:"id"`
	Preferred       int    `json:"preferred"`
	HTTPSFunctional int    `json:"https_functional"`
	Host            string `json:"host"`
}

const baseUrl = "https://www.speedtest.net/api/js/servers?engine=js&https_functional=true&limit=10&search=%s"

// SearchByName uses Speedtest API to retrieve a list of servers relative to
// the query, for a maximum of 10 elements.
func SearchByName(query string) ([]ServerResponse, error) {
	url := fmt.Sprintf(baseUrl, url.QueryEscape(query))
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var servers []ServerResponse
	if err := json.Unmarshal(body, &servers); err != nil {
		return nil, err
	}

	return servers, nil
}
