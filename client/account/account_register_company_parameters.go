// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewAccountRegisterCompanyParams creates a new AccountRegisterCompanyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountRegisterCompanyParams() *AccountRegisterCompanyParams {
	return &AccountRegisterCompanyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountRegisterCompanyParamsWithTimeout creates a new AccountRegisterCompanyParams object
// with the ability to set a timeout on a request.
func NewAccountRegisterCompanyParamsWithTimeout(timeout time.Duration) *AccountRegisterCompanyParams {
	return &AccountRegisterCompanyParams{
		timeout: timeout,
	}
}

// NewAccountRegisterCompanyParamsWithContext creates a new AccountRegisterCompanyParams object
// with the ability to set a context for a request.
func NewAccountRegisterCompanyParamsWithContext(ctx context.Context) *AccountRegisterCompanyParams {
	return &AccountRegisterCompanyParams{
		Context: ctx,
	}
}

// NewAccountRegisterCompanyParamsWithHTTPClient creates a new AccountRegisterCompanyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountRegisterCompanyParamsWithHTTPClient(client *http.Client) *AccountRegisterCompanyParams {
	return &AccountRegisterCompanyParams{
		HTTPClient: client,
	}
}

/*
AccountRegisterCompanyParams contains all the parameters to send to the API endpoint

	for the account register company operation.

	Typically these are written to a http.Request.
*/
type AccountRegisterCompanyParams struct {

	// RegisterAccountRequest.
	RegisterAccountRequest *model.RegisterAccountRequestModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account register company params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountRegisterCompanyParams) WithDefaults() *AccountRegisterCompanyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account register company params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountRegisterCompanyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account register company params
func (o *AccountRegisterCompanyParams) WithTimeout(timeout time.Duration) *AccountRegisterCompanyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account register company params
func (o *AccountRegisterCompanyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account register company params
func (o *AccountRegisterCompanyParams) WithContext(ctx context.Context) *AccountRegisterCompanyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account register company params
func (o *AccountRegisterCompanyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account register company params
func (o *AccountRegisterCompanyParams) WithHTTPClient(client *http.Client) *AccountRegisterCompanyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account register company params
func (o *AccountRegisterCompanyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegisterAccountRequest adds the registerAccountRequest to the account register company params
func (o *AccountRegisterCompanyParams) WithRegisterAccountRequest(registerAccountRequest *model.RegisterAccountRequestModel) *AccountRegisterCompanyParams {
	o.SetRegisterAccountRequest(registerAccountRequest)
	return o
}

// SetRegisterAccountRequest adds the registerAccountRequest to the account register company params
func (o *AccountRegisterCompanyParams) SetRegisterAccountRequest(registerAccountRequest *model.RegisterAccountRequestModel) {
	o.RegisterAccountRequest = registerAccountRequest
}

// WriteToRequest writes these params to a swagger request
func (o *AccountRegisterCompanyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RegisterAccountRequest != nil {
		if err := r.SetBodyParam(o.RegisterAccountRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
