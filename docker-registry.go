package dockerregistry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
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

	httpclient := &http.Client{
		Timeout: 15 * time.Second,
	}

	return &Client{BaseURL: baseURL, httpClient: httpclient}, nil
}

func (c *Client) createRequest(api *RegAPIEndpoint, reqBody io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s/v2/%s", c.BaseURL, api.path)
	req, err := http.NewRequest(api.method, url, reqBody)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// A Repolist represents docker registry repository list
type Repolist struct {
	Repositories []string `json:"repositories"`
}

// GetRepositoryList returns list of existings repositories
func (c *Client) GetRepositoryList() (*Repolist, error) {
	req, err := c.createRequest(&APICallCatalog, nil)
	if err != nil {
		return nil, err
	}

	repolist := &Repolist{}

	err = c.sendRequest(req, repolist)
	if err != nil {
		return nil, err
	}

	return repolist, nil
}
