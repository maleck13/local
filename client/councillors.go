package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ListForCountyAndAreaCouncillorsPath computes a request path to the listForCountyAndArea action of councillors.
func ListForCountyAndAreaCouncillorsPath() string {
	return fmt.Sprintf("/councillors")
}

// list councillors based on a users details
func (c *Client) ListForCountyAndAreaCouncillors(ctx context.Context, path string, area *string, county *string) (*http.Response, error) {
	req, err := c.NewListForCountyAndAreaCouncillorsRequest(ctx, path, area, county)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListForCountyAndAreaCouncillorsRequest create the request corresponding to the listForCountyAndArea action endpoint of the councillors resource.
func (c *Client) NewListForCountyAndAreaCouncillorsRequest(ctx context.Context, path string, area *string, county *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if area != nil {
		values.Set("area", *area)
	}
	if county != nil {
		values.Set("county", *county)
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

// ReadByIDCouncillorsPath computes a request path to the readById action of councillors.
func ReadByIDCouncillorsPath(id string) string {
	return fmt.Sprintf("/councillors/%v", id)
}

// read a councillor based on an id
func (c *Client) ReadByIDCouncillors(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewReadByIDCouncillorsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReadByIDCouncillorsRequest create the request corresponding to the readById action endpoint of the councillors resource.
func (c *Client) NewReadByIDCouncillorsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
