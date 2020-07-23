package dockerregistry

import (
	"fmt"
	"net/http"
)

// APICallBase represents "base" api endpoint
var APICallBase = &RegAPIEndpoint{
	method: http.MethodGet,
	path:   "",
}

// Base API call is used for lightweight version checks and to validate registry authentication.
func (c *Client) Base() error {
	req, err := c.createRequest(APICallBase, nil)
	if err != nil {
		return err
	}

	resp, err := c.sendRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return fmt.Errorf("the registry \"%s\" does not implement the V2 API", c.BaseURL)
	}

	return nil
}
