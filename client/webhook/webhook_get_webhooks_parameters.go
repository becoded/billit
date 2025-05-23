// Code generated by go-swagger; DO NOT EDIT.

package webhook

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

// NewWebhookGetWebhooksParams creates a new WebhookGetWebhooksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWebhookGetWebhooksParams() *WebhookGetWebhooksParams {
	return &WebhookGetWebhooksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWebhookGetWebhooksParamsWithTimeout creates a new WebhookGetWebhooksParams object
// with the ability to set a timeout on a request.
func NewWebhookGetWebhooksParamsWithTimeout(timeout time.Duration) *WebhookGetWebhooksParams {
	return &WebhookGetWebhooksParams{
		timeout: timeout,
	}
}

// NewWebhookGetWebhooksParamsWithContext creates a new WebhookGetWebhooksParams object
// with the ability to set a context for a request.
func NewWebhookGetWebhooksParamsWithContext(ctx context.Context) *WebhookGetWebhooksParams {
	return &WebhookGetWebhooksParams{
		Context: ctx,
	}
}

// NewWebhookGetWebhooksParamsWithHTTPClient creates a new WebhookGetWebhooksParams object
// with the ability to set a custom HTTPClient for a request.
func NewWebhookGetWebhooksParamsWithHTTPClient(client *http.Client) *WebhookGetWebhooksParams {
	return &WebhookGetWebhooksParams{
		HTTPClient: client,
	}
}

/*
WebhookGetWebhooksParams contains all the parameters to send to the API endpoint

	for the webhook get webhooks operation.

	Typically these are written to a http.Request.
*/
type WebhookGetWebhooksParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the webhook get webhooks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebhookGetWebhooksParams) WithDefaults() *WebhookGetWebhooksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the webhook get webhooks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WebhookGetWebhooksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the webhook get webhooks params
func (o *WebhookGetWebhooksParams) WithTimeout(timeout time.Duration) *WebhookGetWebhooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the webhook get webhooks params
func (o *WebhookGetWebhooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the webhook get webhooks params
func (o *WebhookGetWebhooksParams) WithContext(ctx context.Context) *WebhookGetWebhooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the webhook get webhooks params
func (o *WebhookGetWebhooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the webhook get webhooks params
func (o *WebhookGetWebhooksParams) WithHTTPClient(client *http.Client) *WebhookGetWebhooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the webhook get webhooks params
func (o *WebhookGetWebhooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WebhookGetWebhooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
