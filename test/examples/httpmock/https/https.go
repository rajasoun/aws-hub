package https

import (
	"net/http"
)

//http client -> Do func (we need this to make request/calls)
//I created a Interface with Do func so that i can swap my objects in place of httpclient object for testing
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client HttpClient
