// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/becoded/billit/model"
)

// NewOrderPostOrderPaymentParams creates a new OrderPostOrderPaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrderPostOrderPaymentParams() *OrderPostOrderPaymentParams {
	return &OrderPostOrderPaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrderPostOrderPaymentParamsWithTimeout creates a new OrderPostOrderPaymentParams object
// with the ability to set a timeout on a request.
func NewOrderPostOrderPaymentParamsWithTimeout(timeout time.Duration) *OrderPostOrderPaymentParams {
	return &OrderPostOrderPaymentParams{
		timeout: timeout,
	}
}

// NewOrderPostOrderPaymentParamsWithContext creates a new OrderPostOrderPaymentParams object
// with the ability to set a context for a request.
func NewOrderPostOrderPaymentParamsWithContext(ctx context.Context) *OrderPostOrderPaymentParams {
	return &OrderPostOrderPaymentParams{
		Context: ctx,
	}
}

// NewOrderPostOrderPaymentParamsWithHTTPClient creates a new OrderPostOrderPaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrderPostOrderPaymentParamsWithHTTPClient(client *http.Client) *OrderPostOrderPaymentParams {
	return &OrderPostOrderPaymentParams{
		HTTPClient: client,
	}
}

/*
OrderPostOrderPaymentParams contains all the parameters to send to the API endpoint

	for the order post order payment operation.

	Typically these are written to a http.Request.
*/
type OrderPostOrderPaymentParams struct {

	// Data.
	Data *model.PostPaymentRequest

	// OrderID.
	//
	// Format: int32
	OrderID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the order post order payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderPostOrderPaymentParams) WithDefaults() *OrderPostOrderPaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the order post order payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrderPostOrderPaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the order post order payment params
func (o *OrderPostOrderPaymentParams) WithTimeout(timeout time.Duration) *OrderPostOrderPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order post order payment params
func (o *OrderPostOrderPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order post order payment params
func (o *OrderPostOrderPaymentParams) WithContext(ctx context.Context) *OrderPostOrderPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order post order payment params
func (o *OrderPostOrderPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the order post order payment params
func (o *OrderPostOrderPaymentParams) WithHTTPClient(client *http.Client) *OrderPostOrderPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the order post order payment params
func (o *OrderPostOrderPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the order post order payment params
func (o *OrderPostOrderPaymentParams) WithData(data *model.PostPaymentRequest) *OrderPostOrderPaymentParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the order post order payment params
func (o *OrderPostOrderPaymentParams) SetData(data *model.PostPaymentRequest) {
	o.Data = data
}

// WithOrderID adds the orderID to the order post order payment params
func (o *OrderPostOrderPaymentParams) WithOrderID(orderID int32) *OrderPostOrderPaymentParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the order post order payment params
func (o *OrderPostOrderPaymentParams) SetOrderID(orderID int32) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *OrderPostOrderPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param orderID
	if err := r.SetPathParam("orderID", swag.FormatInt32(o.OrderID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
