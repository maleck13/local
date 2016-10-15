package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ListForCountyAndAreaCouncillorsPath computes a request path to the listForCountyAndArea action of councillors.
func ListForCountyAndAreaCouncillorsPath(county string) string {
	return fmt.Sprintf("/councillors/%v", county)
}

// list councillors based on a users details
func (c *Client) ListForCountyAndAreaCouncillors(ctx context.Context, path string, area *string) (*http.Response, error) {
	req, err := c.NewListForCountyAndAreaCouncillorsRequest(ctx, path, area)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListForCountyAndAreaCouncillorsRequest create the request corresponding to the listForCountyAndArea action endpoint of the councillors resource.
func (c *Client) NewListForCountyAndAreaCouncillorsRequest(ctx context.Context, path string, area *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if area != nil {
		values.Set("area", *area)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
