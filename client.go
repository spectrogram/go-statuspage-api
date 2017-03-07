package statuspage

import (
	"net/http"
	"net/url"
	"time"
)

const DefaultAPIURL = "https://api.statuspage.io/v2/"

type Client struct {
	apiKey     string
	pageID     string
	httpClient *http.Client
	url        *url.URL
}

func NewClient(apiKey, pageID string) *Client {
	u, err := url.Parse(DefaultAPIURL)
	if err != nil {
		panic(err)
	}
	return &Client{
		apiKey:     apiKey,
		pageID:     pageID,
		httpClient: &http.Client{Timeout: 5 * time.Second},
		url:        u,
	}
}
