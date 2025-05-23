// Code generated by go-swagger; DO NOT EDIT.

package document

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

// DocumentGetDocumentsReader is a Reader for the DocumentGetDocuments structure.
type DocumentGetDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentGetDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentGetDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/documents] Document_GetDocuments", response, response.Code())
	}
}

// NewDocumentGetDocumentsOK creates a DocumentGetDocumentsOK with default headers values
func NewDocumentGetDocumentsOK() *DocumentGetDocumentsOK {
	return &DocumentGetDocumentsOK{}
}

/*
DocumentGetDocumentsOK describes a response with status code 200, with default header values.

Get a list of documents
*/
type DocumentGetDocumentsOK struct {
	Payload []*model.DocumentAPIView
}

// IsSuccess returns true when this document get documents o k response has a 2xx status code
func (o *DocumentGetDocumentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this document get documents o k response has a 3xx status code
func (o *DocumentGetDocumentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this document get documents o k response has a 4xx status code
func (o *DocumentGetDocumentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this document get documents o k response has a 5xx status code
func (o *DocumentGetDocumentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this document get documents o k response a status code equal to that given
func (o *DocumentGetDocumentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the document get documents o k response
func (o *DocumentGetDocumentsOK) Code() int {
	return 200
}

func (o *DocumentGetDocumentsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/documents][%d] documentGetDocumentsOK %s", 200, payload)
}

func (o *DocumentGetDocumentsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/documents][%d] documentGetDocumentsOK %s", 200, payload)
}

func (o *DocumentGetDocumentsOK) GetPayload() []*model.DocumentAPIView {
	return o.Payload
}

func (o *DocumentGetDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
