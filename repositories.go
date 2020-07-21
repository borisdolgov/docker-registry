package dockerregistry

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
