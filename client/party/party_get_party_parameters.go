// Code generated by go-swagger; DO NOT EDIT.

package party

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
)

// NewPartyGetPartyParams creates a new PartyGetPartyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartyGetPartyParams() *PartyGetPartyParams {
	return &PartyGetPartyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartyGetPartyParamsWithTimeout creates a new PartyGetPartyParams object
// with the ability to set a timeout on a request.
func NewPartyGetPartyParamsWithTimeout(timeout time.Duration) *PartyGetPartyParams {
	return &PartyGetPartyParams{
		timeout: timeout,
	}
}

// NewPartyGetPartyParamsWithContext creates a new PartyGetPartyParams object
// with the ability to set a context for a request.
func NewPartyGetPartyParamsWithContext(ctx context.Context) *PartyGetPartyParams {
	return &PartyGetPartyParams{
		Context: ctx,
	}
}

// NewPartyGetPartyParamsWithHTTPClient creates a new PartyGetPartyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartyGetPartyParamsWithHTTPClient(client *http.Client) *PartyGetPartyParams {
	return &PartyGetPartyParams{
		HTTPClient: client,
	}
}

/*
PartyGetPartyParams contains all the parameters to send to the API endpoint

	for the party get party operation.

	Typically these are written to a http.Request.
*/
type PartyGetPartyParams struct {

	// PartyID.
	//
	// Format: int32
	PartyID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the party get party params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartyGetPartyParams) WithDefaults() *PartyGetPartyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the party get party params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartyGetPartyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the party get party params
func (o *PartyGetPartyParams) WithTimeout(timeout time.Duration) *PartyGetPartyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the party get party params
func (o *PartyGetPartyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the party get party params
func (o *PartyGetPartyParams) WithContext(ctx context.Context) *PartyGetPartyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the party get party params
func (o *PartyGetPartyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the party get party params
func (o *PartyGetPartyParams) WithHTTPClient(client *http.Client) *PartyGetPartyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the party get party params
func (o *PartyGetPartyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPartyID adds the partyID to the party get party params
func (o *PartyGetPartyParams) WithPartyID(partyID int32) *PartyGetPartyParams {
	o.SetPartyID(partyID)
	return o
}

// SetPartyID adds the partyId to the party get party params
func (o *PartyGetPartyParams) SetPartyID(partyID int32) {
	o.PartyID = partyID
}

// WriteToRequest writes these params to a swagger request
func (o *PartyGetPartyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param partyID
	if err := r.SetPathParam("partyID", swag.FormatInt32(o.PartyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
