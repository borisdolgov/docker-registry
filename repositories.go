package dockerregistry

import (
	"encoding/json"
)

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

	resp, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	repolist := &Repolist{}
	err = json.NewDecoder(resp.Body).Decode(repolist)
	if err != nil {
		return nil, err
	}

	return repolist, nil
}
