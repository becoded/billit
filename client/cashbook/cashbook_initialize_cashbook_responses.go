// Code generated by go-swagger; DO NOT EDIT.

package cashbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CashbookInitializeCashbookReader is a Reader for the CashbookInitializeCashbook structure.
type CashbookInitializeCashbookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CashbookInitializeCashbookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCashbookInitializeCashbookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/Cashbook/Initialize] Cashbook_InitializeCashbook", response, response.Code())
	}
}

// NewCashbookInitializeCashbookOK creates a CashbookInitializeCashbookOK with default headers values
func NewCashbookInitializeCashbookOK() *CashbookInitializeCashbookOK {
	return &CashbookInitializeCashbookOK{}
}

/*
CashbookInitializeCashbookOK describes a response with status code 200, with default header values.

Initialize the cash book for first use
*/
type CashbookInitializeCashbookOK struct {
}

// IsSuccess returns true when this cashbook initialize cashbook o k response has a 2xx status code
func (o *CashbookInitializeCashbookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cashbook initialize cashbook o k response has a 3xx status code
func (o *CashbookInitializeCashbookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cashbook initialize cashbook o k response has a 4xx status code
func (o *CashbookInitializeCashbookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cashbook initialize cashbook o k response has a 5xx status code
func (o *CashbookInitializeCashbookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cashbook initialize cashbook o k response a status code equal to that given
func (o *CashbookInitializeCashbookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cashbook initialize cashbook o k response
func (o *CashbookInitializeCashbookOK) Code() int {
	return 200
}

func (o *CashbookInitializeCashbookOK) Error() string {
	return fmt.Sprintf("[POST /v1/Cashbook/Initialize][%d] cashbookInitializeCashbookOK", 200)
}

func (o *CashbookInitializeCashbookOK) String() string {
	return fmt.Sprintf("[POST /v1/Cashbook/Initialize][%d] cashbookInitializeCashbookOK", 200)
}

func (o *CashbookInitializeCashbookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
