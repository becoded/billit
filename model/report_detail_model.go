// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReportDetailModel report detail model
//
// swagger:model ReportDetailModel
type ReportDetailModel struct {

	// c s vs
	CSVs []string `json:"CSVs"`

	// ID
	ID int32 `json:"ID,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this report detail model
func (m *ReportDetailModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this report detail model based on context it is used
func (m *ReportDetailModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportDetailModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportDetailModel) UnmarshalBinary(b []byte) error {
	var res ReportDetailModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
