package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// DeleteUserPath computes a request path to the delete action of user.
func DeleteUserPath(id string) string {
	return fmt.Sprintf("/user/%v", id)
}

// delete a user
func (c *Client) DeleteUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteUserRequest create the request corresponding to the delete action endpoint of the user resource.
func (c *Client) NewDeleteUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListUserPath computes a request path to the list action of user.
func ListUserPath() string {
	return fmt.Sprintf("/user")
}

// get a list user
func (c *Client) ListUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUserRequest create the request corresponding to the list action endpoint of the user resource.
func (c *Client) NewListUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// LoginUserPath computes a request path to the login action of user.
func LoginUserPath() string {
	return fmt.Sprintf("/user/login")
}

// login user
func (c *Client) LoginUser(ctx context.Context, path string, payload *Login, contentType string) (*http.Response, error) {
	req, err := c.NewLoginUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginUserRequest create the request corresponding to the login action endpoint of the user resource.
func (c *Client) NewLoginUserRequest(ctx context.Context, path string, payload *Login, contentType string) (*http.Request, error) {
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
	return req, nil
}

// ReadUserPath computes a request path to the read action of user.
func ReadUserPath(id string) string {
	return fmt.Sprintf("/user/%v", id)
}

// get a user
func (c *Client) ReadUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewReadUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReadUserRequest create the request corresponding to the read action endpoint of the user resource.
func (c *Client) NewReadUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ResetpasswordUserPayload is the user resetpassword action payload.
type ResetpasswordUserPayload struct {
	Newpassword string `form:"newpassword" json:"newpassword" xml:"newpassword"`
}

// ResetpasswordUserPath computes a request path to the resetpassword action of user.
func ResetpasswordUserPath() string {
	return fmt.Sprintf("/user/resetpassword")
}

// resets the users password
func (c *Client) ResetpasswordUser(ctx context.Context, path string, payload *ResetpasswordUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewResetpasswordUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewResetpasswordUserRequest create the request corresponding to the resetpassword action endpoint of the user resource.
func (c *Client) NewResetpasswordUserRequest(ctx context.Context, path string, payload *ResetpasswordUserPayload, contentType string) (*http.Request, error) {
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

// SignUpCouncillorUserPayload is the user signUpCouncillor action payload.
type SignUpCouncillorUserPayload struct {
	// The email of the user
	Email string `form:"email" json:"email" xml:"email"`
}

// SignUpCouncillorUserPath computes a request path to the signUpCouncillor action of user.
func SignUpCouncillorUserPath() string {
	return fmt.Sprintf("/user/councillor/signup")
}

// handles a councillor signup. By verify the email address is a councillors email and sending out a verification email
func (c *Client) SignUpCouncillorUser(ctx context.Context, path string, payload *SignUpCouncillorUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewSignUpCouncillorUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSignUpCouncillorUserRequest create the request corresponding to the signUpCouncillor action endpoint of the user resource.
func (c *Client) NewSignUpCouncillorUserRequest(ctx context.Context, path string, payload *SignUpCouncillorUserPayload, contentType string) (*http.Request, error) {
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
	return req, nil
}

// SignupUserPath computes a request path to the signup action of user.
func SignupUserPath() string {
	return fmt.Sprintf("/user/signup")
}

// Signup a user
func (c *Client) SignupUser(ctx context.Context, path string, payload *User, contentType string) (*http.Response, error) {
	req, err := c.NewSignupUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSignupUserRequest create the request corresponding to the signup action endpoint of the user resource.
func (c *Client) NewSignupUserRequest(ctx context.Context, path string, payload *User, contentType string) (*http.Request, error) {
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
	return req, nil
}

// UpdateUserPath computes a request path to the update action of user.
func UpdateUserPath(id string) string {
	return fmt.Sprintf("/user/%v", id)
}

// update a user
func (c *Client) UpdateUser(ctx context.Context, path string, payload *UpdateUser, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUserRequest create the request corresponding to the update action endpoint of the user resource.
func (c *Client) NewUpdateUserRequest(ctx context.Context, path string, payload *UpdateUser, contentType string) (*http.Request, error) {
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

// VerifySignupUserPath computes a request path to the verifySignup action of user.
func VerifySignupUserPath() string {
	return fmt.Sprintf("/user/signup/verify")
}

// verifies a signup using a token in the  url
func (c *Client) VerifySignupUser(ctx context.Context, path string, key *string, uid *string) (*http.Response, error) {
	req, err := c.NewVerifySignupUserRequest(ctx, path, key, uid)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewVerifySignupUserRequest create the request corresponding to the verifySignup action endpoint of the user resource.
func (c *Client) NewVerifySignupUserRequest(ctx context.Context, path string, key *string, uid *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if key != nil {
		values.Set("key", *key)
	}
	if uid != nil {
		values.Set("uid", *uid)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
