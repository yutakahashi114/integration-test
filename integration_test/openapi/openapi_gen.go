// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// ID defines model for ID.
type ID struct {
	Id int64 `json:"id"`
}

// User defines model for User.
type User struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody User

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody User

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetUsers request
	GetUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateUser request with any body
	CreateUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateUser(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteUser request
	DeleteUser(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindUserByID request
	FindUserByID(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateUser request with any body
	UpdateUserWithBody(ctx context.Context, id int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateUser(ctx context.Context, id int64, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateUserRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateUser(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateUserRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteUser(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteUserRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FindUserByID(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFindUserByIDRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUserWithBody(ctx context.Context, id int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUser(ctx context.Context, id int64, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetUsersRequest generates requests for GetUsers
func NewGetUsersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateUserRequest calls the generic CreateUser builder with application/json body
func NewCreateUserRequest(server string, body CreateUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateUserRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateUserRequestWithBody generates requests for CreateUser with any type of body
func NewCreateUserRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteUserRequest generates requests for DeleteUser
func NewDeleteUserRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFindUserByIDRequest generates requests for FindUserByID
func NewFindUserByIDRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateUserRequest calls the generic UpdateUser builder with application/json body
func NewUpdateUserRequest(server string, id int64, body UpdateUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateUserRequestWithBody(server, id, "application/json", bodyReader)
}

// NewUpdateUserRequestWithBody generates requests for UpdateUser with any type of body
func NewUpdateUserRequestWithBody(server string, id int64, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetUsers request
	GetUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsersResponse, error)

	// CreateUser request with any body
	CreateUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateUserResponse, error)

	CreateUserWithResponse(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateUserResponse, error)

	// DeleteUser request
	DeleteUserWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*DeleteUserResponse, error)

	// FindUserByID request
	FindUserByIDWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*FindUserByIDResponse, error)

	// UpdateUser request with any body
	UpdateUserWithBodyWithResponse(ctx context.Context, id int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error)

	UpdateUserWithResponse(ctx context.Context, id int64, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error)
}

type GetUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]User
}

// Status returns HTTPResponse.Status
func (r GetUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ID
}

// Status returns HTTPResponse.Status
func (r CreateUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type FindUserByIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
}

// Status returns HTTPResponse.Status
func (r FindUserByIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FindUserByIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r UpdateUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetUsersWithResponse request returning *GetUsersResponse
func (c *ClientWithResponses) GetUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsersResponse, error) {
	rsp, err := c.GetUsers(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsersResponse(rsp)
}

// CreateUserWithBodyWithResponse request with arbitrary body returning *CreateUserResponse
func (c *ClientWithResponses) CreateUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateUserResponse, error) {
	rsp, err := c.CreateUserWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateUserResponse(rsp)
}

func (c *ClientWithResponses) CreateUserWithResponse(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateUserResponse, error) {
	rsp, err := c.CreateUser(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateUserResponse(rsp)
}

// DeleteUserWithResponse request returning *DeleteUserResponse
func (c *ClientWithResponses) DeleteUserWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*DeleteUserResponse, error) {
	rsp, err := c.DeleteUser(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteUserResponse(rsp)
}

// FindUserByIDWithResponse request returning *FindUserByIDResponse
func (c *ClientWithResponses) FindUserByIDWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*FindUserByIDResponse, error) {
	rsp, err := c.FindUserByID(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFindUserByIDResponse(rsp)
}

// UpdateUserWithBodyWithResponse request with arbitrary body returning *UpdateUserResponse
func (c *ClientWithResponses) UpdateUserWithBodyWithResponse(ctx context.Context, id int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error) {
	rsp, err := c.UpdateUserWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserResponse(rsp)
}

func (c *ClientWithResponses) UpdateUserWithResponse(ctx context.Context, id int64, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error) {
	rsp, err := c.UpdateUser(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserResponse(rsp)
}

// ParseGetUsersResponse parses an HTTP response from a GetUsersWithResponse call
func ParseGetUsersResponse(rsp *http.Response) (*GetUsersResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateUserResponse parses an HTTP response from a CreateUserWithResponse call
func ParseCreateUserResponse(rsp *http.Response) (*CreateUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ID
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteUserResponse parses an HTTP response from a DeleteUserWithResponse call
func ParseDeleteUserResponse(rsp *http.Response) (*DeleteUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &DeleteUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseFindUserByIDResponse parses an HTTP response from a FindUserByIDWithResponse call
func ParseFindUserByIDResponse(rsp *http.Response) (*FindUserByIDResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &FindUserByIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateUserResponse parses an HTTP response from a UpdateUserWithResponse call
func ParseUpdateUserResponse(rsp *http.Response) (*UpdateUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &UpdateUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
