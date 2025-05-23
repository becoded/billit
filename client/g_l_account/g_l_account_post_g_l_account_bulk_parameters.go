// Code generated by go-swagger; DO NOT EDIT.

package g_l_account

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

// NewGLAccountPostGLAccountBulkParams creates a new GLAccountPostGLAccountBulkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGLAccountPostGLAccountBulkParams() *GLAccountPostGLAccountBulkParams {
	return &GLAccountPostGLAccountBulkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGLAccountPostGLAccountBulkParamsWithTimeout creates a new GLAccountPostGLAccountBulkParams object
// with the ability to set a timeout on a request.
func NewGLAccountPostGLAccountBulkParamsWithTimeout(timeout time.Duration) *GLAccountPostGLAccountBulkParams {
	return &GLAccountPostGLAccountBulkParams{
		timeout: timeout,
	}
}

// NewGLAccountPostGLAccountBulkParamsWithContext creates a new GLAccountPostGLAccountBulkParams object
// with the ability to set a context for a request.
func NewGLAccountPostGLAccountBulkParamsWithContext(ctx context.Context) *GLAccountPostGLAccountBulkParams {
	return &GLAccountPostGLAccountBulkParams{
		Context: ctx,
	}
}

// NewGLAccountPostGLAccountBulkParamsWithHTTPClient creates a new GLAccountPostGLAccountBulkParams object
// with the ability to set a custom HTTPClient for a request.
func NewGLAccountPostGLAccountBulkParamsWithHTTPClient(client *http.Client) *GLAccountPostGLAccountBulkParams {
	return &GLAccountPostGLAccountBulkParams{
		HTTPClient: client,
	}
}

/*
GLAccountPostGLAccountBulkParams contains all the parameters to send to the API endpoint

	for the g l account post g l account bulk operation.

	Typically these are written to a http.Request.
*/
type GLAccountPostGLAccountBulkParams struct {

	// Accounts.
	Accounts []*model.GLAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the g l account post g l account bulk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GLAccountPostGLAccountBulkParams) WithDefaults() *GLAccountPostGLAccountBulkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the g l account post g l account bulk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GLAccountPostGLAccountBulkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the g l account post g l account bulk params
func (o *GLAccountPostGLAccountBulkParams) WithTimeout(timeout time.Duration) *GLAccountPostGLAccountBulkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g l account post g l account bulk params
func (o *GLAccountPostGLAccountBulkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g l account post g l account bulk params
func (o *GLAccountPostGLAccountBulkParams) WithContext(ctx context.Context) *GLAccountPostGLAccountBulkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g l account post g l account bulk params
func (o *GLAccountPostGLAccountBulkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g l account post g l account bulk params
func (o *GLAccountPostGLAccountBulkParams) WithHTTPClient(client *http.Client) *GLAccountPostGLAccountBulkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g l account post g l account bulk params
func (o *GLAccountPostGLAccountBulkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccounts adds the accounts to the g l account post g l account bulk params
func (o *GLAccountPostGLAccountBulkParams) WithAccounts(accounts []*model.GLAccount) *GLAccountPostGLAccountBulkParams {
	o.SetAccounts(accounts)
	return o
}

// SetAccounts adds the accounts to the g l account post g l account bulk params
func (o *GLAccountPostGLAccountBulkParams) SetAccounts(accounts []*model.GLAccount) {
	o.Accounts = accounts
}

// WriteToRequest writes these params to a swagger request
func (o *GLAccountPostGLAccountBulkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Accounts != nil {
		if err := r.SetBodyParam(o.Accounts); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
