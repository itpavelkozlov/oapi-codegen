// Package client provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package client

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// HTTP client with any customized settings, such as certificate chains.
	Client http.Client

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor func(req *http.Request, ctx context.Context) error
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostBoth request with JSON body
	PostBoth(ctx context.Context, body SchemaObject) (*http.Response, error)

	// PostBothWithBody request with arbitrary body
	PostBothWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	// GetBoth request
	GetBoth(ctx context.Context) (*http.Response, error)

	// PostJson request with JSON body
	PostJson(ctx context.Context, body SchemaObject) (*http.Response, error)

	// GetJson request
	GetJson(ctx context.Context) (*http.Response, error)

	// PostOther request with arbitrary body
	PostOther(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	// GetOther request
	GetOther(ctx context.Context) (*http.Response, error)
}

// PostBoth request with JSON body
func (c *Client) PostBoth(ctx context.Context, body SchemaObject) (*http.Response, error) {
	req, err := NewPostBothRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// PostBothWithBody request with arbitrary body
func (c *Client) PostBothWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPostBothRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// GetBoth request
func (c *Client) GetBoth(ctx context.Context) (*http.Response, error) {
	req, err := NewGetBothRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// PostJson request with JSON body
func (c *Client) PostJson(ctx context.Context, body SchemaObject) (*http.Response, error) {
	req, err := NewPostJsonRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// GetJson request
func (c *Client) GetJson(ctx context.Context) (*http.Response, error) {
	req, err := NewGetJsonRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// PostOther request with arbitrary body
func (c *Client) PostOther(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPostOtherRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// GetOther request
func (c *Client) GetOther(ctx context.Context) (*http.Response, error) {
	req, err := NewGetOtherRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewPostBothRequest generates requests for PostBoth with JSON body
func NewPostBothRequest(server string, body SchemaObject) (*http.Request, error) {
	var bodyReader io.Reader

	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)

	return NewPostBothRequestWithBody(server, "application/json", bodyReader)
}

// NewPostBothRequestWithBody generates requests for PostBoth with non-JSON body
func NewPostBothRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/with_both_bodies", server)

	req, err := http.NewRequest("POST", queryUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetBothRequest generates requests for GetBoth
func NewGetBothRequest(server string) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/with_both_responses", server)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostJsonRequest generates requests for PostJson with JSON body
func NewPostJsonRequest(server string, body SchemaObject) (*http.Request, error) {
	var bodyReader io.Reader

	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)

	return NewPostJsonRequestWithBody(server, "application/json", bodyReader)
}

// NewPostJsonRequestWithBody generates requests for PostJson with non-JSON body
func NewPostJsonRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/with_json_body", server)

	req, err := http.NewRequest("POST", queryUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetJsonRequest generates requests for GetJson
func NewGetJsonRequest(server string) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/with_json_response", server)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostOtherRequestWithBody generates requests for PostOther with non-JSON body
func NewPostOtherRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/with_other_body", server)

	req, err := http.NewRequest("POST", queryUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetOtherRequest generates requests for GetOther
func NewGetOtherRequest(server string) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/with_other_response", server)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses returns a ClientWithResponses with a default Client:
func NewClientWithResponses(server string) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client: http.Client{},
			Server: server,
		},
	}
}

type postBothResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r postBothResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r postBothResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type getBothResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r getBothResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r getBothResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type postJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r postJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r postJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type getJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r getJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r getJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type postOtherResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r postOtherResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r postOtherResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type getOtherResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r getOtherResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r getOtherResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostBoth request with JSON body returning *PostBothResponse
func (c *ClientWithResponses) PostBothWithResponse(ctx context.Context, body SchemaObject) (*postBothResponse, error) {
	rsp, err := c.PostBoth(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParsepostBothResponse(rsp)
}

// PostBothWithBody request with arbitrary body returning *PostBothResponse
func (c *ClientWithResponses) PostBothWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*postBothResponse, error) {
	rsp, err := c.PostBothWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsepostBothResponse(rsp)
}

// GetBoth request returning *GetBothResponse
func (c *ClientWithResponses) GetBothWithResponse(ctx context.Context) (*getBothResponse, error) {
	rsp, err := c.GetBoth(ctx)
	if err != nil {
		return nil, err
	}
	return ParsegetBothResponse(rsp)
}

// PostJson request with JSON body returning *PostJsonResponse
func (c *ClientWithResponses) PostJsonWithResponse(ctx context.Context, body SchemaObject) (*postJsonResponse, error) {
	rsp, err := c.PostJson(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParsepostJsonResponse(rsp)
}

// GetJson request returning *GetJsonResponse
func (c *ClientWithResponses) GetJsonWithResponse(ctx context.Context) (*getJsonResponse, error) {
	rsp, err := c.GetJson(ctx)
	if err != nil {
		return nil, err
	}
	return ParsegetJsonResponse(rsp)
}

// PostOther request with arbitrary body returning *PostOtherResponse
func (c *ClientWithResponses) PostOtherWithResponse(ctx context.Context, contentType string, body io.Reader) (*postOtherResponse, error) {
	rsp, err := c.PostOther(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsepostOtherResponse(rsp)
}

// GetOther request returning *GetOtherResponse
func (c *ClientWithResponses) GetOtherWithResponse(ctx context.Context) (*getOtherResponse, error) {
	rsp, err := c.GetOther(ctx)
	if err != nil {
		return nil, err
	}
	return ParsegetOtherResponse(rsp)
}

// ParsepostBothResponse parses an HTTP response from a PostBothWithResponse call
func ParsepostBothResponse(rsp *http.Response) (*postBothResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &postBothResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParsegetBothResponse parses an HTTP response from a GetBothWithResponse call
func ParsegetBothResponse(rsp *http.Response) (*getBothResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &getBothResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.StatusCode == 200:
		break // No content-type
	}

	return response, nil
}

// ParsepostJsonResponse parses an HTTP response from a PostJsonWithResponse call
func ParsepostJsonResponse(rsp *http.Response) (*postJsonResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &postJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParsegetJsonResponse parses an HTTP response from a GetJsonWithResponse call
func ParsegetJsonResponse(rsp *http.Response) (*getJsonResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &getJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.StatusCode == 200:
		break // No content-type
	}

	return response, nil
}

// ParsepostOtherResponse parses an HTTP response from a PostOtherWithResponse call
func ParsepostOtherResponse(rsp *http.Response) (*postOtherResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &postOtherResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParsegetOtherResponse parses an HTTP response from a GetOtherWithResponse call
func ParsegetOtherResponse(rsp *http.Response) (*getOtherResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &getOtherResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.StatusCode == 200:
		break // No content-type
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// (POST /with_both_bodies)
	PostBoth(ctx echo.Context) error
	// (GET /with_both_responses)
	GetBoth(ctx echo.Context) error
	// (POST /with_json_body)
	PostJson(ctx echo.Context) error
	// (GET /with_json_response)
	GetJson(ctx echo.Context) error
	// (POST /with_other_body)
	PostOther(ctx echo.Context) error
	// (GET /with_other_response)
	GetOther(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostBoth converts echo context to params.
func (w *ServerInterfaceWrapper) PostBoth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostBoth(ctx)
	return err
}

// GetBoth converts echo context to params.
func (w *ServerInterfaceWrapper) GetBoth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBoth(ctx)
	return err
}

// PostJson converts echo context to params.
func (w *ServerInterfaceWrapper) PostJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostJson(ctx)
	return err
}

// GetJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetJson(ctx)
	return err
}

// PostOther converts echo context to params.
func (w *ServerInterfaceWrapper) PostOther(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostOther(ctx)
	return err
}

// GetOther converts echo context to params.
func (w *ServerInterfaceWrapper) GetOther(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOther(ctx)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST("/with_both_bodies", wrapper.PostBoth)
	router.GET("/with_both_responses", wrapper.GetBoth)
	router.POST("/with_json_body", wrapper.PostJson)
	router.GET("/with_json_response", wrapper.GetJson)
	router.POST("/with_other_body", wrapper.PostOther)
	router.GET("/with_other_response", wrapper.GetOther)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8yTQW8TMRCF/8pq4LgkW7jtEQ6oSFBEI3EAhBzvS+xq1zbjaatVlf+OxknIRlRRJFTU",
	"SzSOZ0bvvc/7QDYOKQYEydQ+ULYOgynldSmvljewoufEMYHFo9yuPGf5ZAboQcYEaikL+7CmTU0c+8cu",
	"9Aa/bj2jo/bbtquerPqx0RYfVlGHO2TLPomPgVpaOJ8rQZZc3TuIA1fiUL3rPYJUJnS78qsX9wU5xZCR",
	"K8Oo1ghgI+gqG5lhpR+/B6qp9xYhF52hGKGPlwtVL15UPi2QpboG34Gppjtw3kq5mDWzRhtjQjDJU0tv",
	"Zs3sgmpKRlzJZ37vxf1cxvLT7UJLMZcoNUijvi47aulzzPI2iqNtOtBTN2qfjUEQyohJqfe2DM1vssrY",
	"w9LqJWNFLb2YH2jOdyjnRxw13+mqaAXyKgvDDMcrV5EHI9TS0gfDI9V/wTyiKXyL8tfEOO8x6L41HrH+",
	"Hgfnk97XTfNcPU88qiSFO55G+0GV/xe0p4AUsfuQT/H4I/cJeUxTjPopnxHjlfadneNTveut2nNyPOg9",
	"HeS/vsbN5ncAAAD//29vOELEBQAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
