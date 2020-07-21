package dockerregistry

import "net/http"

// A Repolist represents docker registry repository list
type Repolist struct {
	Repositories []string `json:"repositories"`
}

// Repositories returns list of existings repositories
func (c *Client) Repositories() (*Repolist, error) {
	req, err := c.createRequest(&APICallCatalog, nil)
	if err != nil {
		return nil, err
	}

	data := &APIData{
		Content: &Repolist{},
		Header:  http.Header{},
		Status:  "",
	}

	err = c.sendRequest(req, data)
	if err != nil {
		return nil, err
	}

	return data.Content.(*Repolist), nil
}
