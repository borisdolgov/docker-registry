package dockerregistry

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

	repolist := &Repolist{}

	err = c.sendRequest(req, repolist)
	if err != nil {
		return nil, err
	}

	return repolist, nil
}
