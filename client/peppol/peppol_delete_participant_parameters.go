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

	"github.com/becoded/billit/model"
)

// NewPeppolDeleteParticipantParams creates a new PeppolDeleteParticipantParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPeppolDeleteParticipantParams() *PeppolDeleteParticipantParams {
	return &PeppolDeleteParticipantParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPeppolDeleteParticipantParamsWithTimeout creates a new PeppolDeleteParticipantParams object
// with the ability to set a timeout on a request.
func NewPeppolDeleteParticipantParamsWithTimeout(timeout time.Duration) *PeppolDeleteParticipantParams {
	return &PeppolDeleteParticipantParams{
		timeout: timeout,
	}
}

// NewPeppolDeleteParticipantParamsWithContext creates a new PeppolDeleteParticipantParams object
// with the ability to set a context for a request.
func NewPeppolDeleteParticipantParamsWithContext(ctx context.Context) *PeppolDeleteParticipantParams {
	return &PeppolDeleteParticipantParams{
		Context: ctx,
	}
}

// NewPeppolDeleteParticipantParamsWithHTTPClient creates a new PeppolDeleteParticipantParams object
// with the ability to set a custom HTTPClient for a request.
func NewPeppolDeleteParticipantParamsWithHTTPClient(client *http.Client) *PeppolDeleteParticipantParams {
	return &PeppolDeleteParticipantParams{
		HTTPClient: client,
	}
}

/*
PeppolDeleteParticipantParams contains all the parameters to send to the API endpoint

	for the peppol delete participant operation.

	Typically these are written to a http.Request.
*/
type PeppolDeleteParticipantParams struct {

	// RegisterRequest.
	RegisterRequest *model.DeleteParticipantRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the peppol delete participant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PeppolDeleteParticipantParams) WithDefaults() *PeppolDeleteParticipantParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the peppol delete participant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PeppolDeleteParticipantParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the peppol delete participant params
func (o *PeppolDeleteParticipantParams) WithTimeout(timeout time.Duration) *PeppolDeleteParticipantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peppol delete participant params
func (o *PeppolDeleteParticipantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peppol delete participant params
func (o *PeppolDeleteParticipantParams) WithContext(ctx context.Context) *PeppolDeleteParticipantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peppol delete participant params
func (o *PeppolDeleteParticipantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peppol delete participant params
func (o *PeppolDeleteParticipantParams) WithHTTPClient(client *http.Client) *PeppolDeleteParticipantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peppol delete participant params
func (o *PeppolDeleteParticipantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegisterRequest adds the registerRequest to the peppol delete participant params
func (o *PeppolDeleteParticipantParams) WithRegisterRequest(registerRequest *model.DeleteParticipantRequest) *PeppolDeleteParticipantParams {
	o.SetRegisterRequest(registerRequest)
	return o
}

// SetRegisterRequest adds the registerRequest to the peppol delete participant params
func (o *PeppolDeleteParticipantParams) SetRegisterRequest(registerRequest *model.DeleteParticipantRequest) {
	o.RegisterRequest = registerRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PeppolDeleteParticipantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RegisterRequest != nil {
		if err := r.SetBodyParam(o.RegisterRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
