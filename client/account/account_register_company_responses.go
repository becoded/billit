// Code generated by go-swagger; DO NOT EDIT.

package account

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

// AccountRegisterCompanyReader is a Reader for the AccountRegisterCompany structure.
type AccountRegisterCompanyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountRegisterCompanyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountRegisterCompanyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/account/registercompany] Account_RegisterCompany", response, response.Code())
	}
}

// NewAccountRegisterCompanyOK creates a AccountRegisterCompanyOK with default headers values
func NewAccountRegisterCompanyOK() *AccountRegisterCompanyOK {
	return &AccountRegisterCompanyOK{}
}

/*
AccountRegisterCompanyOK describes a response with status code 200, with default header values.

Registers a company
*/
type AccountRegisterCompanyOK struct {
	Payload *model.RegisterAccountResponse
}

// IsSuccess returns true when this account register company o k response has a 2xx status code
func (o *AccountRegisterCompanyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account register company o k response has a 3xx status code
func (o *AccountRegisterCompanyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account register company o k response has a 4xx status code
func (o *AccountRegisterCompanyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account register company o k response has a 5xx status code
func (o *AccountRegisterCompanyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account register company o k response a status code equal to that given
func (o *AccountRegisterCompanyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account register company o k response
func (o *AccountRegisterCompanyOK) Code() int {
	return 200
}

func (o *AccountRegisterCompanyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/account/registercompany][%d] accountRegisterCompanyOK %s", 200, payload)
}

func (o *AccountRegisterCompanyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/account/registercompany][%d] accountRegisterCompanyOK %s", 200, payload)
}

func (o *AccountRegisterCompanyOK) GetPayload() *model.RegisterAccountResponse {
	return o.Payload
}

func (o *AccountRegisterCompanyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.RegisterAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
