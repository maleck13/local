package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateCouncillorAdminPath computes a request path to the createCouncillor action of admin.
func CreateCouncillorAdminPath() string {
	return fmt.Sprintf("/admin/councillor")
}

// admin api to add a councillor
func (c *Client) CreateCouncillorAdmin(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCreateCouncillorAdminRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCouncillorAdminRequest create the request corresponding to the createCouncillor action endpoint of the admin resource.
func (c *Client) NewCreateCouncillorAdminRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
