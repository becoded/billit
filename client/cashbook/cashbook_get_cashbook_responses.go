// Code generated by go-swagger; DO NOT EDIT.

package cashbook

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

// CashbookGetCashbookReader is a Reader for the CashbookGetCashbook structure.
type CashbookGetCashbookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CashbookGetCashbookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCashbookGetCashbookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/Cashbook] Cashbook_GetCashbook", response, response.Code())
	}
}

// NewCashbookGetCashbookOK creates a CashbookGetCashbookOK with default headers values
func NewCashbookGetCashbookOK() *CashbookGetCashbookOK {
	return &CashbookGetCashbookOK{}
}

/*
CashbookGetCashbookOK describes a response with status code 200, with default header values.

Get information about the current active cash book
*/
type CashbookGetCashbookOK struct {
	Payload *model.CashbookModel
}

// IsSuccess returns true when this cashbook get cashbook o k response has a 2xx status code
func (o *CashbookGetCashbookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cashbook get cashbook o k response has a 3xx status code
func (o *CashbookGetCashbookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cashbook get cashbook o k response has a 4xx status code
func (o *CashbookGetCashbookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cashbook get cashbook o k response has a 5xx status code
func (o *CashbookGetCashbookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cashbook get cashbook o k response a status code equal to that given
func (o *CashbookGetCashbookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cashbook get cashbook o k response
func (o *CashbookGetCashbookOK) Code() int {
	return 200
}

func (o *CashbookGetCashbookOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/Cashbook][%d] cashbookGetCashbookOK %s", 200, payload)
}

func (o *CashbookGetCashbookOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/Cashbook][%d] cashbookGetCashbookOK %s", 200, payload)
}

func (o *CashbookGetCashbookOK) GetPayload() *model.CashbookModel {
	return o.Payload
}

func (o *CashbookGetCashbookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CashbookModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
