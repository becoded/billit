// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new order API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new order API client with basic auth credentials.
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

// New creates a new order API client with a bearer token for authentication.
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
Client for order API
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
	OrderDeleteOrder(params *OrderDeleteOrderParams, opts ...ClientOption) (*OrderDeleteOrderOK, error)

	OrderESignOrder(params *OrderESignOrderParams, opts ...ClientOption) (*OrderESignOrderOK, error)

	OrderGetDeleted(params *OrderGetDeletedParams, opts ...ClientOption) (*OrderGetDeletedOK, error)

	OrderGetOrder(params *OrderGetOrderParams, opts ...ClientOption) (*OrderGetOrderOK, error)

	OrderGetOrders(params *OrderGetOrdersParams, opts ...ClientOption) (*OrderGetOrdersOK, error)

	OrderPatchOrders(params *OrderPatchOrdersParams, opts ...ClientOption) (*OrderPatchOrdersOK, error)

	OrderPostOrderPayment(params *OrderPostOrderPaymentParams, opts ...ClientOption) (*OrderPostOrderPaymentOK, error)

	OrderPostOrders(params *OrderPostOrdersParams, opts ...ClientOption) (*OrderPostOrdersOK, error)

	OrderPostSend(params *OrderPostSendParams, opts ...ClientOption) (*OrderPostSendOK, error)

	OrderPutOrderBookings(params *OrderPutOrderBookingsParams, opts ...ClientOption) (*OrderPutOrderBookingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
OrderDeleteOrder order delete order API
*/
func (a *Client) OrderDeleteOrder(params *OrderDeleteOrderParams, opts ...ClientOption) (*OrderDeleteOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderDeleteOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_DeleteOrder",
		Method:             "DELETE",
		PathPattern:        "/v1/orders/{orderID}",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderDeleteOrderReader{formats: a.formats},
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
	success, ok := result.(*OrderDeleteOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_DeleteOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderESignOrder order e sign order API
*/
func (a *Client) OrderESignOrder(params *OrderESignOrderParams, opts ...ClientOption) (*OrderESignOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderESignOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_ESignOrder",
		Method:             "POST",
		PathPattern:        "/v1/orders/commands/eSign/{orderID}",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderESignOrderReader{formats: a.formats},
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
	success, ok := result.(*OrderESignOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_ESignOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderGetDeleted order get deleted API
*/
func (a *Client) OrderGetDeleted(params *OrderGetDeletedParams, opts ...ClientOption) (*OrderGetDeletedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderGetDeletedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_GetDeleted",
		Method:             "GET",
		PathPattern:        "/v1/orders/deleted",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderGetDeletedReader{formats: a.formats},
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
	success, ok := result.(*OrderGetDeletedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_GetDeleted: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderGetOrder order get order API
*/
func (a *Client) OrderGetOrder(params *OrderGetOrderParams, opts ...ClientOption) (*OrderGetOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderGetOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_GetOrder",
		Method:             "GET",
		PathPattern:        "/v1/orders/{orderID}",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderGetOrderReader{formats: a.formats},
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
	success, ok := result.(*OrderGetOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_GetOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderGetOrders order get orders API
*/
func (a *Client) OrderGetOrders(params *OrderGetOrdersParams, opts ...ClientOption) (*OrderGetOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderGetOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_GetOrders",
		Method:             "GET",
		PathPattern:        "/v1/orders",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderGetOrdersReader{formats: a.formats},
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
	success, ok := result.(*OrderGetOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_GetOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderPatchOrders order patch orders API
*/
func (a *Client) OrderPatchOrders(params *OrderPatchOrdersParams, opts ...ClientOption) (*OrderPatchOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderPatchOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_PatchOrders",
		Method:             "PATCH",
		PathPattern:        "/v1/orders/{orderID}",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderPatchOrdersReader{formats: a.formats},
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
	success, ok := result.(*OrderPatchOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_PatchOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderPostOrderPayment order post order payment API
*/
func (a *Client) OrderPostOrderPayment(params *OrderPostOrderPaymentParams, opts ...ClientOption) (*OrderPostOrderPaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderPostOrderPaymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_PostOrderPayment",
		Method:             "POST",
		PathPattern:        "/v1/orders/{orderID}/payments",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderPostOrderPaymentReader{formats: a.formats},
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
	success, ok := result.(*OrderPostOrderPaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_PostOrderPayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderPostOrders order post orders API
*/
func (a *Client) OrderPostOrders(params *OrderPostOrdersParams, opts ...ClientOption) (*OrderPostOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderPostOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_PostOrders",
		Method:             "POST",
		PathPattern:        "/v1/orders",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderPostOrdersReader{formats: a.formats},
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
	success, ok := result.(*OrderPostOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_PostOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderPostSend order post send API
*/
func (a *Client) OrderPostSend(params *OrderPostSendParams, opts ...ClientOption) (*OrderPostSendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderPostSendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_PostSend",
		Method:             "POST",
		PathPattern:        "/v1/orders/commands/send",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderPostSendReader{formats: a.formats},
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
	success, ok := result.(*OrderPostSendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_PostSend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrderPutOrderBookings order put order bookings API
*/
func (a *Client) OrderPutOrderBookings(params *OrderPutOrderBookingsParams, opts ...ClientOption) (*OrderPutOrderBookingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderPutOrderBookingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Order_PutOrderBookings",
		Method:             "PUT",
		PathPattern:        "/v1/orders/{orderID}/bookingEntries",
		ProducesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrderPutOrderBookingsReader{formats: a.formats},
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
	success, ok := result.(*OrderPutOrderBookingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Order_PutOrderBookings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
