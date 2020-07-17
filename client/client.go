package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// A Registry represents docker registry server
//     URL is a registry server url
type Registry struct {
	APIVersion string
	URL        *url.URL
}

// GetRepositoryList returns list of existings repositories
func (r *Registry) GetRepositoryList() {
	method := APICallCatalog.method
	url := fmt.Sprintf("%v/%v/%v", r.URL, r.APIVersion, APICallCatalog.path)
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
