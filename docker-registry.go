package dockerregistry

import (
	"encoding/json"
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
	if baseURL.Scheme != "https" {
		// TODO: create package error variable and return it here
		return nil, fmt.Errorf("docker registry url scheme should be https")
	}
	return &Client{BaseURL: baseURL, httpClient: http.DefaultClient}, nil
}

// A Repolist represents docker registry repository list
type Repolist struct {
	Repositories []string `json:"repositories"`
}

// GetRepositoryList returns list of existings repositories
func (c *Client) GetRepositoryList() (*Repolist, error) {
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

	repolist := &Repolist{}
	err = json.NewDecoder(response.Body).Decode(repolist)
	if err != nil {
		return nil, err
	}

	return repolist, nil

}
