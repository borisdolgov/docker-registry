package dockerregistry

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// A Client represents docker registry server
//     BaseURL is a registry server url
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

// NewClient initialize parameters for client and return *Client
func NewClient(registryURL string) (*Client, error) {
	baseURL, err := url.Parse(registryURL)
	if err != nil {
		return nil, err
	}
	return &Client{BaseURL: baseURL, httpClient: http.DefaultClient}, nil
}

type repolist struct {
	Repositories []string
}

// GetRepositoryList returns list of existings repositories
func (c *Client) GetRepositoryList() {
	apiEndpoint := APICallCatalog
	url := fmt.Sprintf("%v/v2/%v", c.BaseURL, apiEndpoint.path)

	request, err := http.NewRequest(apiEndpoint.method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)

}
