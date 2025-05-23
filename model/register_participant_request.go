// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegisterParticipantRequest register participant request
//
// swagger:model RegisterParticipantRequest
type RegisterParticipantRequest struct {

	// company ID
	CompanyID string `json:"CompanyID,omitempty"`
}

// Validate validates this register participant request
func (m *RegisterParticipantRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this register participant request based on context it is used
func (m *RegisterParticipantRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterParticipantRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterParticipantRequest) UnmarshalBinary(b []byte) error {
	var res RegisterParticipantRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
