// Code generated by go-swagger; DO NOT EDIT.

package order

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

// OrderGetOrderReader is a Reader for the OrderGetOrder structure.
type OrderGetOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderGetOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderGetOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/orders/{orderID}] Order_GetOrder", response, response.Code())
	}
}

// NewOrderGetOrderOK creates a OrderGetOrderOK with default headers values
func NewOrderGetOrderOK() *OrderGetOrderOK {
	return &OrderGetOrderOK{}
}

/*
OrderGetOrderOK describes a response with status code 200, with default header values.

Get an order
*/
type OrderGetOrderOK struct {
	Payload *model.Order
}

// IsSuccess returns true when this order get order o k response has a 2xx status code
func (o *OrderGetOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this order get order o k response has a 3xx status code
func (o *OrderGetOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this order get order o k response has a 4xx status code
func (o *OrderGetOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this order get order o k response has a 5xx status code
func (o *OrderGetOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this order get order o k response a status code equal to that given
func (o *OrderGetOrderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the order get order o k response
func (o *OrderGetOrderOK) Code() int {
	return 200
}

func (o *OrderGetOrderOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orders/{orderID}][%d] orderGetOrderOK %s", 200, payload)
}

func (o *OrderGetOrderOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orders/{orderID}][%d] orderGetOrderOK %s", 200, payload)
}

func (o *OrderGetOrderOK) GetPayload() *model.Order {
	return o.Payload
}

func (o *OrderGetOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
