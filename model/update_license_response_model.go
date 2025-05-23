// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateLicenseResponseModel update license response model
//
// swagger:model UpdateLicenseResponseModel
type UpdateLicenseResponseModel struct {

	// company ID
	CompanyID int32 `json:"CompanyID,omitempty"`
}

// Validate validates this update license response model
func (m *UpdateLicenseResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update license response model based on context it is used
func (m *UpdateLicenseResponseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateLicenseResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateLicenseResponseModel) UnmarshalBinary(b []byte) error {
	var res UpdateLicenseResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
