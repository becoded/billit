// Code generated by go-swagger; DO NOT EDIT.

package accountant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new accountant API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new accountant API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new accountant API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for accountant API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// WithContentTypeApplicationxWwwFormUrlencoded sets the Content-Type header to "application/x-www-form-urlencoded".
func WithContentTypeApplicationxWwwFormUrlencoded(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/x-www-form-urlencoded"}
}

// WithContentTypeApplicationXML sets the Content-Type header to "application/xml".
func WithContentTypeApplicationXML(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/xml"}
}

// WithContentTypeTextJSON sets the Content-Type header to "text/json".
func WithContentTypeTextJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"text/json"}
}

// WithContentTypeTextXML sets the Content-Type header to "text/xml".
func WithContentTypeTextXML(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"text/xml"}
}

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptApplicationXML sets the Accept header to "application/xml".
func WithAcceptApplicationXML(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/xml"}
}

// WithAcceptTextJSON sets the Accept header to "text/json".
func WithAcceptTextJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/json"}
}

// WithAcceptTextXML sets the Accept header to "text/xml".
func WithAcceptTextXML(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/xml"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	AccountantDeleteFeeds(params *AccountantDeleteFeedsParams, opts ...ClientOption) (*AccountantDeleteFeedsOK, error)

	AccountantGetFeeds(params *AccountantGetFeedsParams, opts ...ClientOption) (*AccountantGetFeedsOK, error)

	AccountantGetFile(params *AccountantGetFileParams, opts ...ClientOption) (*AccountantGetFileOK, error)

	AccountantGetIndex(params *AccountantGetIndexParams, opts ...ClientOption) (*AccountantGetIndexOK, error)

	AccountantPostConfirm(params *AccountantPostConfirmParams, opts ...ClientOption) (*AccountantPostConfirmOK, error)

	AccountantPostFeeds(params *AccountantPostFeedsParams, opts ...ClientOption) (*AccountantPostFeedsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AccountantDeleteFeeds deletes the feed
*/
func (a *Client) AccountantDeleteFeeds(params *AccountantDeleteFeedsParams, opts ...ClientOption) (*AccountantDeleteFeedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountantDeleteFeedsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Accountant_DeleteFeeds",
		Method:             "DELETE",
		PathPattern:        "/v1/accountant/feeds/{feedName}",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountantDeleteFeedsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AccountantDeleteFeedsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Accountant_DeleteFeeds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountantGetFeeds accountant get feeds API
*/
func (a *Client) AccountantGetFeeds(params *AccountantGetFeedsParams, opts ...ClientOption) (*AccountantGetFeedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountantGetFeedsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Accountant_GetFeeds",
		Method:             "GET",
		PathPattern:        "/v1/accountant/feeds",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountantGetFeedsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AccountantGetFeedsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Accountant_GetFeeds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountantGetFile downloads a coda file from the feed
*/
func (a *Client) AccountantGetFile(params *AccountantGetFileParams, opts ...ClientOption) (*AccountantGetFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountantGetFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Accountant_GetFile",
		Method:             "GET",
		PathPattern:        "/v1/accountant/feeds/{feedName}/{feedItemID}",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountantGetFileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AccountantGetFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Accountant_GetFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountantGetIndex gets a list of all feeds to download only query this once per minute
*/
func (a *Client) AccountantGetIndex(params *AccountantGetIndexParams, opts ...ClientOption) (*AccountantGetIndexOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountantGetIndexParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Accountant_GetIndex",
		Method:             "GET",
		PathPattern:        "/v1/accountant/feeds/{feedName}",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountantGetIndexReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AccountantGetIndexOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Accountant_GetIndex: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountantPostConfirm confirms each succesfully downloaded feed item to remove it from the feedlist
*/
func (a *Client) AccountantPostConfirm(params *AccountantPostConfirmParams, opts ...ClientOption) (*AccountantPostConfirmOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountantPostConfirmParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Accountant_PostConfirm",
		Method:             "POST",
		PathPattern:        "/v1/accountant/feeds/{feedName}/{feedItemID}/confirm",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountantPostConfirmReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AccountantPostConfirmOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Accountant_PostConfirm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AccountantPostFeeds registers a new feed all newly exported orders or documents will be available in this new feed
*/
func (a *Client) AccountantPostFeeds(params *AccountantPostFeedsParams, opts ...ClientOption) (*AccountantPostFeedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccountantPostFeedsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Accountant_PostFeeds",
		Method:             "POST",
		PathPattern:        "/v1/accountant/feeds",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccountantPostFeedsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AccountantPostFeedsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Accountant_PostFeeds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
