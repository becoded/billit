// Code generated by go-swagger; DO NOT EDIT.

package accountant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccountantPostConfirmReader is a Reader for the AccountantPostConfirm structure.
type AccountantPostConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountantPostConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountantPostConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/accountant/feeds/{feedName}/{feedItemID}/confirm] Accountant_PostConfirm", response, response.Code())
	}
}

// NewAccountantPostConfirmOK creates a AccountantPostConfirmOK with default headers values
func NewAccountantPostConfirmOK() *AccountantPostConfirmOK {
	return &AccountantPostConfirmOK{}
}

/*
AccountantPostConfirmOK describes a response with status code 200, with default header values.

OK
*/
type AccountantPostConfirmOK struct {
	Payload interface{}
}

// IsSuccess returns true when this accountant post confirm o k response has a 2xx status code
func (o *AccountantPostConfirmOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this accountant post confirm o k response has a 3xx status code
func (o *AccountantPostConfirmOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this accountant post confirm o k response has a 4xx status code
func (o *AccountantPostConfirmOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this accountant post confirm o k response has a 5xx status code
func (o *AccountantPostConfirmOK) IsServerError() bool {
	return false
}

// IsCode returns true when this accountant post confirm o k response a status code equal to that given
func (o *AccountantPostConfirmOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the accountant post confirm o k response
func (o *AccountantPostConfirmOK) Code() int {
	return 200
}

func (o *AccountantPostConfirmOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accountant/feeds/{feedName}/{feedItemID}/confirm][%d] accountantPostConfirmOK %s", 200, payload)
}

func (o *AccountantPostConfirmOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accountant/feeds/{feedName}/{feedItemID}/confirm][%d] accountantPostConfirmOK %s", 200, payload)
}

func (o *AccountantPostConfirmOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AccountantPostConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
