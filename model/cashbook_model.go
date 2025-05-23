// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CashbookModel cashbook model
//
// swagger:model CashbookModel
type CashbookModel struct {

	// cashbook ID
	CashbookID int32 `json:"CashbookID,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"CreationDate,omitempty"`

	// currency
	Currency string `json:"Currency,omitempty"`

	// last entry date
	// Format: date-time
	LastEntryDate *strfmt.DateTime `json:"LastEntryDate,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// saldo
	Saldo float64 `json:"Saldo,omitempty"`
}

// Validate validates this cashbook model
func (m *CashbookModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastEntryDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CashbookModel) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("CreationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CashbookModel) validateLastEntryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastEntryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("LastEntryDate", "body", "date-time", m.LastEntryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cashbook model based on context it is used
func (m *CashbookModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CashbookModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CashbookModel) UnmarshalBinary(b []byte) error {
	var res CashbookModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
