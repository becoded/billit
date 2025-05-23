// Code generated by go-swagger; DO NOT EDIT.

package accountant

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
)

// NewAccountantGetIndexParams creates a new AccountantGetIndexParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountantGetIndexParams() *AccountantGetIndexParams {
	return &AccountantGetIndexParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountantGetIndexParamsWithTimeout creates a new AccountantGetIndexParams object
// with the ability to set a timeout on a request.
func NewAccountantGetIndexParamsWithTimeout(timeout time.Duration) *AccountantGetIndexParams {
	return &AccountantGetIndexParams{
		timeout: timeout,
	}
}

// NewAccountantGetIndexParamsWithContext creates a new AccountantGetIndexParams object
// with the ability to set a context for a request.
func NewAccountantGetIndexParamsWithContext(ctx context.Context) *AccountantGetIndexParams {
	return &AccountantGetIndexParams{
		Context: ctx,
	}
}

// NewAccountantGetIndexParamsWithHTTPClient creates a new AccountantGetIndexParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountantGetIndexParamsWithHTTPClient(client *http.Client) *AccountantGetIndexParams {
	return &AccountantGetIndexParams{
		HTTPClient: client,
	}
}

/*
AccountantGetIndexParams contains all the parameters to send to the API endpoint

	for the accountant get index operation.

	Typically these are written to a http.Request.
*/
type AccountantGetIndexParams struct {

	// FeedName.
	FeedName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accountant get index params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountantGetIndexParams) WithDefaults() *AccountantGetIndexParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accountant get index params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountantGetIndexParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accountant get index params
func (o *AccountantGetIndexParams) WithTimeout(timeout time.Duration) *AccountantGetIndexParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accountant get index params
func (o *AccountantGetIndexParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accountant get index params
func (o *AccountantGetIndexParams) WithContext(ctx context.Context) *AccountantGetIndexParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accountant get index params
func (o *AccountantGetIndexParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accountant get index params
func (o *AccountantGetIndexParams) WithHTTPClient(client *http.Client) *AccountantGetIndexParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accountant get index params
func (o *AccountantGetIndexParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeedName adds the feedName to the accountant get index params
func (o *AccountantGetIndexParams) WithFeedName(feedName string) *AccountantGetIndexParams {
	o.SetFeedName(feedName)
	return o
}

// SetFeedName adds the feedName to the accountant get index params
func (o *AccountantGetIndexParams) SetFeedName(feedName string) {
	o.FeedName = feedName
}

// WriteToRequest writes these params to a swagger request
func (o *AccountantGetIndexParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param feedName
	if err := r.SetPathParam("feedName", o.FeedName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
