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

// CashbookGetCashbooksReader is a Reader for the CashbookGetCashbooks structure.
type CashbookGetCashbooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CashbookGetCashbooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCashbookGetCashbooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/Cashbook/History] Cashbook_GetCashbooks", response, response.Code())
	}
}

// NewCashbookGetCashbooksOK creates a CashbookGetCashbooksOK with default headers values
func NewCashbookGetCashbooksOK() *CashbookGetCashbooksOK {
	return &CashbookGetCashbooksOK{}
}

/*
CashbookGetCashbooksOK describes a response with status code 200, with default header values.

Get an overview of all the closed cash books
*/
type CashbookGetCashbooksOK struct {
	Payload []*model.CashbookModel
}

// IsSuccess returns true when this cashbook get cashbooks o k response has a 2xx status code
func (o *CashbookGetCashbooksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cashbook get cashbooks o k response has a 3xx status code
func (o *CashbookGetCashbooksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cashbook get cashbooks o k response has a 4xx status code
func (o *CashbookGetCashbooksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cashbook get cashbooks o k response has a 5xx status code
func (o *CashbookGetCashbooksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cashbook get cashbooks o k response a status code equal to that given
func (o *CashbookGetCashbooksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cashbook get cashbooks o k response
func (o *CashbookGetCashbooksOK) Code() int {
	return 200
}

func (o *CashbookGetCashbooksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/Cashbook/History][%d] cashbookGetCashbooksOK %s", 200, payload)
}

func (o *CashbookGetCashbooksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/Cashbook/History][%d] cashbookGetCashbooksOK %s", 200, payload)
}

func (o *CashbookGetCashbooksOK) GetPayload() []*model.CashbookModel {
	return o.Payload
}

func (o *CashbookGetCashbooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
