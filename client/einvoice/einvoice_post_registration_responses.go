// Code generated by go-swagger; DO NOT EDIT.

package einvoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/becoded/billit/model"
)

// EinvoicePostRegistrationReader is a Reader for the EinvoicePostRegistration structure.
type EinvoicePostRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EinvoicePostRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEinvoicePostRegistrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/einvoices/registrations] Einvoice_PostRegistration", response, response.Code())
	}
}

// NewEinvoicePostRegistrationOK creates a EinvoicePostRegistrationOK with default headers values
func NewEinvoicePostRegistrationOK() *EinvoicePostRegistrationOK {
	return &EinvoicePostRegistrationOK{}
}

/*
EinvoicePostRegistrationOK describes a response with status code 200, with default header values.

Register a new entity
*/
type EinvoicePostRegistrationOK struct {
	Payload *model.AccountRegistrationResponse
}

// IsSuccess returns true when this einvoice post registration o k response has a 2xx status code
func (o *EinvoicePostRegistrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this einvoice post registration o k response has a 3xx status code
func (o *EinvoicePostRegistrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this einvoice post registration o k response has a 4xx status code
func (o *EinvoicePostRegistrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this einvoice post registration o k response has a 5xx status code
func (o *EinvoicePostRegistrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this einvoice post registration o k response a status code equal to that given
func (o *EinvoicePostRegistrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the einvoice post registration o k response
func (o *EinvoicePostRegistrationOK) Code() int {
	return 200
}

func (o *EinvoicePostRegistrationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/einvoices/registrations][%d] einvoicePostRegistrationOK %s", 200, payload)
}

func (o *EinvoicePostRegistrationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/einvoices/registrations][%d] einvoicePostRegistrationOK %s", 200, payload)
}

func (o *EinvoicePostRegistrationOK) GetPayload() *model.AccountRegistrationResponse {
	return o.Payload
}

func (o *EinvoicePostRegistrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.AccountRegistrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
