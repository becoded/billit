// Code generated by go-swagger; DO NOT EDIT.

package report

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

// NewReportReportListGetParams creates a new ReportReportListGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportReportListGetParams() *ReportReportListGetParams {
	return &ReportReportListGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportReportListGetParamsWithTimeout creates a new ReportReportListGetParams object
// with the ability to set a timeout on a request.
func NewReportReportListGetParamsWithTimeout(timeout time.Duration) *ReportReportListGetParams {
	return &ReportReportListGetParams{
		timeout: timeout,
	}
}

// NewReportReportListGetParamsWithContext creates a new ReportReportListGetParams object
// with the ability to set a context for a request.
func NewReportReportListGetParamsWithContext(ctx context.Context) *ReportReportListGetParams {
	return &ReportReportListGetParams{
		Context: ctx,
	}
}

// NewReportReportListGetParamsWithHTTPClient creates a new ReportReportListGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportReportListGetParamsWithHTTPClient(client *http.Client) *ReportReportListGetParams {
	return &ReportReportListGetParams{
		HTTPClient: client,
	}
}

/*
ReportReportListGetParams contains all the parameters to send to the API endpoint

	for the report report list get operation.

	Typically these are written to a http.Request.
*/
type ReportReportListGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report report list get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportReportListGetParams) WithDefaults() *ReportReportListGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report report list get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportReportListGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report report list get params
func (o *ReportReportListGetParams) WithTimeout(timeout time.Duration) *ReportReportListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report report list get params
func (o *ReportReportListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report report list get params
func (o *ReportReportListGetParams) WithContext(ctx context.Context) *ReportReportListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report report list get params
func (o *ReportReportListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report report list get params
func (o *ReportReportListGetParams) WithHTTPClient(client *http.Client) *ReportReportListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report report list get params
func (o *ReportReportListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReportReportListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
