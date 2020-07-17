package client

import (
	"net/url"
)

// A Registry represents docker registry server
//     URL is a registry server url
type Registry struct {
	URL *url.URL
}
