// Code generated by go-swagger; DO NOT EDIT.

package einvoice

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

// NewEinvoicePostRegistrationIntegrationParams creates a new EinvoicePostRegistrationIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEinvoicePostRegistrationIntegrationParams() *EinvoicePostRegistrationIntegrationParams {
	return &EinvoicePostRegistrationIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEinvoicePostRegistrationIntegrationParamsWithTimeout creates a new EinvoicePostRegistrationIntegrationParams object
// with the ability to set a timeout on a request.
func NewEinvoicePostRegistrationIntegrationParamsWithTimeout(timeout time.Duration) *EinvoicePostRegistrationIntegrationParams {
	return &EinvoicePostRegistrationIntegrationParams{
		timeout: timeout,
	}
}

// NewEinvoicePostRegistrationIntegrationParamsWithContext creates a new EinvoicePostRegistrationIntegrationParams object
// with the ability to set a context for a request.
func NewEinvoicePostRegistrationIntegrationParamsWithContext(ctx context.Context) *EinvoicePostRegistrationIntegrationParams {
	return &EinvoicePostRegistrationIntegrationParams{
		Context: ctx,
	}
}

// NewEinvoicePostRegistrationIntegrationParamsWithHTTPClient creates a new EinvoicePostRegistrationIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewEinvoicePostRegistrationIntegrationParamsWithHTTPClient(client *http.Client) *EinvoicePostRegistrationIntegrationParams {
	return &EinvoicePostRegistrationIntegrationParams{
		HTTPClient: client,
	}
}

/*
EinvoicePostRegistrationIntegrationParams contains all the parameters to send to the API endpoint

	for the einvoice post registration integration operation.

	Typically these are written to a http.Request.
*/
type EinvoicePostRegistrationIntegrationParams struct {

	// Post.
	Post *model.IntegrationPost

	// RegistrationID.
	//
	// Format: int32
	RegistrationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the einvoice post registration integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EinvoicePostRegistrationIntegrationParams) WithDefaults() *EinvoicePostRegistrationIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the einvoice post registration integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EinvoicePostRegistrationIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) WithTimeout(timeout time.Duration) *EinvoicePostRegistrationIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) WithContext(ctx context.Context) *EinvoicePostRegistrationIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) WithHTTPClient(client *http.Client) *EinvoicePostRegistrationIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPost adds the post to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) WithPost(post *model.IntegrationPost) *EinvoicePostRegistrationIntegrationParams {
	o.SetPost(post)
	return o
}

// SetPost adds the post to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) SetPost(post *model.IntegrationPost) {
	o.Post = post
}

// WithRegistrationID adds the registrationID to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) WithRegistrationID(registrationID int32) *EinvoicePostRegistrationIntegrationParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the einvoice post registration integration params
func (o *EinvoicePostRegistrationIntegrationParams) SetRegistrationID(registrationID int32) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *EinvoicePostRegistrationIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Post != nil {
		if err := r.SetBodyParam(o.Post); err != nil {
			return err
		}
	}

	// path param registrationID
	if err := r.SetPathParam("registrationID", swag.FormatInt32(o.RegistrationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
