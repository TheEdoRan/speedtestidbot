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

const baseUrl = "https://www.speedtest.net/api/js/servers?engine=js&https_functional=true&limit=10&search="

// SearchByName uses Speedtest API to retrieve a list of servers relative to
// the query, for a maximum of 10 elements.
func SearchByName(query string) ([]ServerResponse, error) {
	reqUrl := baseUrl + url.QueryEscape(query)
	resp, err := http.Get(reqUrl)

	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", reqUrl, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("GET %s: reading body (status %s): %w", reqUrl, resp.Status, err)
	}

	var servers []ServerResponse
	if err := json.Unmarshal(body, &servers); err != nil {
		return nil, fmt.Errorf(
			"GET %s: decoding response failed (status %s, content-type %q): %w\nraw body: %s",
			reqUrl,
			resp.Status,
			resp.Header.Get("Content-Type"),
			err,
			truncate(string(body), 2000),
		)
	}

	return servers, nil
}

// truncate shortens s to at most max characters, appending an ellipsis with the
// original length when the body is cut, so logs stay readable.
func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return fmt.Sprintf("%s… (%d bytes total)", s[:max], len(s))
}
