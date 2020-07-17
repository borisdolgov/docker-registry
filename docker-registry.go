package dockerregistry

import (
	"fmt"
	"net/http"
	"net/url"
)

// A Client represents docker registry server
//     URL is a registry server url
type Client struct {
	URL *url.URL
}

// GetRepositoryList returns list of existings repositories
func (c *Client) GetRepositoryList() {
	method := APICallCatalog.method
	url := fmt.Sprintf("%v/v2/%v", c.URL, APICallCatalog.path)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
