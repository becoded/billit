// Code generated by go-swagger; DO NOT EDIT.

package daily_receipt

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

// NewDailyReceiptGetDailyReceiptEntryParams creates a new DailyReceiptGetDailyReceiptEntryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDailyReceiptGetDailyReceiptEntryParams() *DailyReceiptGetDailyReceiptEntryParams {
	return &DailyReceiptGetDailyReceiptEntryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDailyReceiptGetDailyReceiptEntryParamsWithTimeout creates a new DailyReceiptGetDailyReceiptEntryParams object
// with the ability to set a timeout on a request.
func NewDailyReceiptGetDailyReceiptEntryParamsWithTimeout(timeout time.Duration) *DailyReceiptGetDailyReceiptEntryParams {
	return &DailyReceiptGetDailyReceiptEntryParams{
		timeout: timeout,
	}
}

// NewDailyReceiptGetDailyReceiptEntryParamsWithContext creates a new DailyReceiptGetDailyReceiptEntryParams object
// with the ability to set a context for a request.
func NewDailyReceiptGetDailyReceiptEntryParamsWithContext(ctx context.Context) *DailyReceiptGetDailyReceiptEntryParams {
	return &DailyReceiptGetDailyReceiptEntryParams{
		Context: ctx,
	}
}

// NewDailyReceiptGetDailyReceiptEntryParamsWithHTTPClient creates a new DailyReceiptGetDailyReceiptEntryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDailyReceiptGetDailyReceiptEntryParamsWithHTTPClient(client *http.Client) *DailyReceiptGetDailyReceiptEntryParams {
	return &DailyReceiptGetDailyReceiptEntryParams{
		HTTPClient: client,
	}
}

/*
DailyReceiptGetDailyReceiptEntryParams contains all the parameters to send to the API endpoint

	for the daily receipt get daily receipt entry operation.

	Typically these are written to a http.Request.
*/
type DailyReceiptGetDailyReceiptEntryParams struct {

	// DailyReceiptEntryID.
	//
	// Format: int32
	DailyReceiptEntryID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the daily receipt get daily receipt entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DailyReceiptGetDailyReceiptEntryParams) WithDefaults() *DailyReceiptGetDailyReceiptEntryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the daily receipt get daily receipt entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DailyReceiptGetDailyReceiptEntryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the daily receipt get daily receipt entry params
func (o *DailyReceiptGetDailyReceiptEntryParams) WithTimeout(timeout time.Duration) *DailyReceiptGetDailyReceiptEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the daily receipt get daily receipt entry params
func (o *DailyReceiptGetDailyReceiptEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the daily receipt get daily receipt entry params
func (o *DailyReceiptGetDailyReceiptEntryParams) WithContext(ctx context.Context) *DailyReceiptGetDailyReceiptEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the daily receipt get daily receipt entry params
func (o *DailyReceiptGetDailyReceiptEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the daily receipt get daily receipt entry params
func (o *DailyReceiptGetDailyReceiptEntryParams) WithHTTPClient(client *http.Client) *DailyReceiptGetDailyReceiptEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the daily receipt get daily receipt entry params
func (o *DailyReceiptGetDailyReceiptEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDailyReceiptEntryID adds the dailyReceiptEntryID to the daily receipt get daily receipt entry params
func (o *DailyReceiptGetDailyReceiptEntryParams) WithDailyReceiptEntryID(dailyReceiptEntryID int32) *DailyReceiptGetDailyReceiptEntryParams {
	o.SetDailyReceiptEntryID(dailyReceiptEntryID)
	return o
}

// SetDailyReceiptEntryID adds the dailyReceiptEntryId to the daily receipt get daily receipt entry params
func (o *DailyReceiptGetDailyReceiptEntryParams) SetDailyReceiptEntryID(dailyReceiptEntryID int32) {
	o.DailyReceiptEntryID = dailyReceiptEntryID
}

// WriteToRequest writes these params to a swagger request
func (o *DailyReceiptGetDailyReceiptEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param DailyReceiptEntryID
	if err := r.SetPathParam("DailyReceiptEntryID", swag.FormatInt32(o.DailyReceiptEntryID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
