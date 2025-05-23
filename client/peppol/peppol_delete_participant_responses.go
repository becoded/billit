// Code generated by go-swagger; DO NOT EDIT.

package peppol

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PeppolDeleteParticipantReader is a Reader for the PeppolDeleteParticipant structure.
type PeppolDeleteParticipantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PeppolDeleteParticipantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPeppolDeleteParticipantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/peppol/participants] Peppol_DeleteParticipant", response, response.Code())
	}
}

// NewPeppolDeleteParticipantOK creates a PeppolDeleteParticipantOK with default headers values
func NewPeppolDeleteParticipantOK() *PeppolDeleteParticipantOK {
	return &PeppolDeleteParticipantOK{}
}

/*
PeppolDeleteParticipantOK describes a response with status code 200, with default header values.

Send a direct Peppol message. Compatible with receivers that support Invoice5 or Creditnote5, BisV3 Invoice or BisV3 Creditnote
*/
type PeppolDeleteParticipantOK struct {
	Payload int32
}

// IsSuccess returns true when this peppol delete participant o k response has a 2xx status code
func (o *PeppolDeleteParticipantOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this peppol delete participant o k response has a 3xx status code
func (o *PeppolDeleteParticipantOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this peppol delete participant o k response has a 4xx status code
func (o *PeppolDeleteParticipantOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this peppol delete participant o k response has a 5xx status code
func (o *PeppolDeleteParticipantOK) IsServerError() bool {
	return false
}

// IsCode returns true when this peppol delete participant o k response a status code equal to that given
func (o *PeppolDeleteParticipantOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the peppol delete participant o k response
func (o *PeppolDeleteParticipantOK) Code() int {
	return 200
}

func (o *PeppolDeleteParticipantOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/peppol/participants][%d] peppolDeleteParticipantOK %s", 200, payload)
}

func (o *PeppolDeleteParticipantOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/peppol/participants][%d] peppolDeleteParticipantOK %s", 200, payload)
}

func (o *PeppolDeleteParticipantOK) GetPayload() int32 {
	return o.Payload
}

func (o *PeppolDeleteParticipantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
