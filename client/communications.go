package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CloseCommunicationsPath computes a request path to the close action of communications.
func CloseCommunicationsPath(id string) string {
	return fmt.Sprintf("/communications/close/%v", id)
}

// recieve an email
func (c *Client) CloseCommunications(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCloseCommunicationsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCloseCommunicationsRequest create the request corresponding to the close action endpoint of the communications resource.
func (c *Client) NewCloseCommunicationsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListCommunicationsPath computes a request path to the list action of communications.
func ListCommunicationsPath(rid string) string {
	return fmt.Sprintf("/communications/councillor/%v", rid)
}

// read communications
func (c *Client) ListCommunications(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCommunicationsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCommunicationsRequest create the request corresponding to the list action endpoint of the communications resource.
func (c *Client) NewListCommunicationsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// RecieveEmailCommunicationsPath computes a request path to the recieveEmail action of communications.
func RecieveEmailCommunicationsPath() string {
	return fmt.Sprintf("/communications/email/recieve")
}

// recieve an email
func (c *Client) RecieveEmailCommunications(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRecieveEmailCommunicationsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRecieveEmailCommunicationsRequest create the request corresponding to the recieveEmail action endpoint of the communications resource.
func (c *Client) NewRecieveEmailCommunicationsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SendCommunicationsPath computes a request path to the send action of communications.
func SendCommunicationsPath() string {
	return fmt.Sprintf("/communications/send")
}

// send and email
func (c *Client) SendCommunications(ctx context.Context, path string, payload *Communication, contentType string) (*http.Response, error) {
	req, err := c.NewSendCommunicationsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSendCommunicationsRequest create the request corresponding to the send action endpoint of the communications resource.
func (c *Client) NewSendCommunicationsRequest(ctx context.Context, path string, payload *Communication, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
