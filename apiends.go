package dockerregistry

import "net/http"

// A RegAPIEndpoint represents docker registry API endpoint
type RegAPIEndpoint struct {
	method string
	path   string
}

// APICallCatalog provides list of repositories stored in docker registry
var APICallCatalog = RegAPIEndpoint{method: http.MethodGet, path: "_catalog"}
