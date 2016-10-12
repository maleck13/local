package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// CreateRequest returns a mocked request
func CreateRequest(t *testing.T, responseStatus int, path, response string) func() (*http.Response, error) {
	return func() (*http.Response, error) {
		bodyRC := ioutil.NopCloser(bytes.NewReader([]byte(response)))
		resp := &http.Response{StatusCode: responseStatus, Body: bodyRC, Header: http.Header{}}
		return resp, nil
	}
}
