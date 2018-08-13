package statuspage

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// DefaultAPIURL should be the current URL prefix as described in the docs:
// https://doers.statuspage.io/api/basics/
const DefaultAPIURL = "https://api.statuspage.io/v1/pages/"

// Client contains the information needed to find and consume the StatusPage API
type Client struct {
	apiKey     string
	pageID     string
	httpClient *http.Client
	url        *url.URL
}

// NewClient creates a new Client from a StatusPage API key and page ID
// See https://doers.statuspage.io/api/authentication/ on how to obtain
// an API key.
func NewClient(apiKey, pageID string) (*Client, error) {
	u, err := url.Parse(DefaultAPIURL + pageID + "/")
	if err != nil {
		return nil, fmt.Errorf("url error parsing (%s): %s", pageID, err)
	}
	c := Client{
		apiKey:     apiKey,
		pageID:     pageID,
		httpClient: &http.Client{Timeout: 5 * time.Second},
		url:        u,
	}
	return &c, nil
}
