// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDailyReceiptEntryResponseModel get daily receipt entry response model
//
// swagger:model GetDailyReceiptEntryResponseModel
type GetDailyReceiptEntryResponseModel struct {

	// cashbook entry ID
	CashbookEntryID int32 `json:"CashbookEntryID,omitempty"`

	// cashbook ID
	CashbookID int32 `json:"CashbookID,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// orderid
	Orderid int32 `json:"Orderid,omitempty"`

	// total excl
	TotalExcl float64 `json:"TotalExcl,omitempty"`

	// total incl
	TotalIncl float64 `json:"TotalIncl,omitempty"`

	// total v a t
	TotalVAT float64 `json:"TotalVAT,omitempty"`

	// v a t
	VAT float64 `json:"VAT,omitempty"`
}

// Validate validates this get daily receipt entry response model
func (m *GetDailyReceiptEntryResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get daily receipt entry response model based on context it is used
func (m *GetDailyReceiptEntryResponseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetDailyReceiptEntryResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDailyReceiptEntryResponseModel) UnmarshalBinary(b []byte) error {
	var res GetDailyReceiptEntryResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
