package main

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPing(t *testing.T) {

	//this will target http client and fools it
	httpmock.Activate()

	//this will undo fooling effect oresle it will always be fooled
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://google.com", func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") != "test" {
			t.Error("wanted exact Authorization")
		}
		resp, err := httpmock.NewJsonResponse(200, map[string]interface{}{
			"result": "success",
		})
		return resp, err
	})

	Ping()
}
