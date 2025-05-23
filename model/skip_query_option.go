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

// SkipQueryOption skip query option
//
// swagger:model SkipQueryOption
type SkipQueryOption struct {

	// context
	// Read Only: true
	Context *ODataQueryContext `json:"Context,omitempty"`

	// raw value
	// Read Only: true
	RawValue string `json:"RawValue,omitempty"`

	// validator
	Validator SkipQueryValidator `json:"Validator,omitempty"`

	// value
	// Read Only: true
	Value int32 `json:"Value,omitempty"`
}

// Validate validates this skip query option
func (m *SkipQueryOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SkipQueryOption) validateContext(formats strfmt.Registry) error {
	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Context")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this skip query option based on the context it is used
func (m *SkipQueryOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRawValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SkipQueryOption) contextValidateContext(ctx context.Context, formats strfmt.Registry) error {

	if m.Context != nil {

		if swag.IsZero(m.Context) { // not required
			return nil
		}

		if err := m.Context.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Context")
			}
			return err
		}
	}

	return nil
}

func (m *SkipQueryOption) contextValidateRawValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "RawValue", "body", string(m.RawValue)); err != nil {
		return err
	}

	return nil
}

func (m *SkipQueryOption) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "Value", "body", int32(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SkipQueryOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SkipQueryOption) UnmarshalBinary(b []byte) error {
	var res SkipQueryOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
