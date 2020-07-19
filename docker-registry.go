package dockerregistry

import (
	"fmt"
	"net/http"
	"net/url"
)

// A Client represents docker registry server
//     BaseURL is a registry server url
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

// GetRepositoryList returns list of existings repositories
func (c *Client) GetRepositoryList() {
	apiEndpoint := APICallCatalog
	url := fmt.Sprintf("%v/v2/%v", c.BaseURL, apiEndpoint.path)

	request, err := http.NewRequest(apiEndpoint.method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

}
