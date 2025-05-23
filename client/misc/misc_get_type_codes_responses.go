// Code generated by go-swagger; DO NOT EDIT.

package misc

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

// MiscGetTypeCodesReader is a Reader for the MiscGetTypeCodes structure.
type MiscGetTypeCodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MiscGetTypeCodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMiscGetTypeCodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/misc/typecodes/{TypeCodeType}] Misc_GetTypeCodes", response, response.Code())
	}
}

// NewMiscGetTypeCodesOK creates a MiscGetTypeCodesOK with default headers values
func NewMiscGetTypeCodesOK() *MiscGetTypeCodesOK {
	return &MiscGetTypeCodesOK{}
}

/*
MiscGetTypeCodesOK describes a response with status code 200, with default header values.

Get all translations of a TypeCode
*/
type MiscGetTypeCodesOK struct {
	Payload *model.TypeCodeResult
}

// IsSuccess returns true when this misc get type codes o k response has a 2xx status code
func (o *MiscGetTypeCodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this misc get type codes o k response has a 3xx status code
func (o *MiscGetTypeCodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this misc get type codes o k response has a 4xx status code
func (o *MiscGetTypeCodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this misc get type codes o k response has a 5xx status code
func (o *MiscGetTypeCodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this misc get type codes o k response a status code equal to that given
func (o *MiscGetTypeCodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the misc get type codes o k response
func (o *MiscGetTypeCodesOK) Code() int {
	return 200
}

func (o *MiscGetTypeCodesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/misc/typecodes/{TypeCodeType}][%d] miscGetTypeCodesOK %s", 200, payload)
}

func (o *MiscGetTypeCodesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/misc/typecodes/{TypeCodeType}][%d] miscGetTypeCodesOK %s", 200, payload)
}

func (o *MiscGetTypeCodesOK) GetPayload() *model.TypeCodeResult {
	return o.Payload
}

func (o *MiscGetTypeCodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.TypeCodeResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
