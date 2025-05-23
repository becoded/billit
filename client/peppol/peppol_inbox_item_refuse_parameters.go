// Code generated by go-swagger; DO NOT EDIT.

package peppol

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

// NewPeppolInboxItemRefuseParams creates a new PeppolInboxItemRefuseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPeppolInboxItemRefuseParams() *PeppolInboxItemRefuseParams {
	return &PeppolInboxItemRefuseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPeppolInboxItemRefuseParamsWithTimeout creates a new PeppolInboxItemRefuseParams object
// with the ability to set a timeout on a request.
func NewPeppolInboxItemRefuseParamsWithTimeout(timeout time.Duration) *PeppolInboxItemRefuseParams {
	return &PeppolInboxItemRefuseParams{
		timeout: timeout,
	}
}

// NewPeppolInboxItemRefuseParamsWithContext creates a new PeppolInboxItemRefuseParams object
// with the ability to set a context for a request.
func NewPeppolInboxItemRefuseParamsWithContext(ctx context.Context) *PeppolInboxItemRefuseParams {
	return &PeppolInboxItemRefuseParams{
		Context: ctx,
	}
}

// NewPeppolInboxItemRefuseParamsWithHTTPClient creates a new PeppolInboxItemRefuseParams object
// with the ability to set a custom HTTPClient for a request.
func NewPeppolInboxItemRefuseParamsWithHTTPClient(client *http.Client) *PeppolInboxItemRefuseParams {
	return &PeppolInboxItemRefuseParams{
		HTTPClient: client,
	}
}

/*
PeppolInboxItemRefuseParams contains all the parameters to send to the API endpoint

	for the peppol inbox item refuse operation.

	Typically these are written to a http.Request.
*/
type PeppolInboxItemRefuseParams struct {

	// InboxItemID.
	//
	// Format: int32
	InboxItemID int32

	// Request.
	Request *model.RefusePeppolItemRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the peppol inbox item refuse params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PeppolInboxItemRefuseParams) WithDefaults() *PeppolInboxItemRefuseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the peppol inbox item refuse params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PeppolInboxItemRefuseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) WithTimeout(timeout time.Duration) *PeppolInboxItemRefuseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) WithContext(ctx context.Context) *PeppolInboxItemRefuseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) WithHTTPClient(client *http.Client) *PeppolInboxItemRefuseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInboxItemID adds the inboxItemID to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) WithInboxItemID(inboxItemID int32) *PeppolInboxItemRefuseParams {
	o.SetInboxItemID(inboxItemID)
	return o
}

// SetInboxItemID adds the inboxItemId to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) SetInboxItemID(inboxItemID int32) {
	o.InboxItemID = inboxItemID
}

// WithRequest adds the request to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) WithRequest(request *model.RefusePeppolItemRequest) *PeppolInboxItemRefuseParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the peppol inbox item refuse params
func (o *PeppolInboxItemRefuseParams) SetRequest(request *model.RefusePeppolItemRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PeppolInboxItemRefuseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inboxItemId
	if err := r.SetPathParam("inboxItemId", swag.FormatInt32(o.InboxItemID)); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
