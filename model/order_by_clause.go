// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderByClause order by clause
//
// swagger:model OrderByClause
type OrderByClause struct {

	// direction
	// Read Only: true
	// Enum: ["Ascending","Descending"]
	Direction string `json:"Direction,omitempty"`

	// expression
	// Read Only: true
	Expression *SingleValueNode `json:"Expression,omitempty"`

	// item type
	// Read Only: true
	ItemType *IEdmTypeReference `json:"ItemType,omitempty"`

	// range variable
	// Read Only: true
	RangeVariable *RangeVariable `json:"RangeVariable,omitempty"`

	// then by
	// Read Only: true
	ThenBy *OrderByClause `json:"ThenBy,omitempty"`
}

// Validate validates this order by clause
func (m *OrderByClause) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRangeVariable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThenBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var orderByClauseTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ascending","Descending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderByClauseTypeDirectionPropEnum = append(orderByClauseTypeDirectionPropEnum, v)
	}
}

const (

	// OrderByClauseDirectionAscending captures enum value "Ascending"
	OrderByClauseDirectionAscending string = "Ascending"

	// OrderByClauseDirectionDescending captures enum value "Descending"
	OrderByClauseDirectionDescending string = "Descending"
)

// prop value enum
func (m *OrderByClause) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderByClauseTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrderByClause) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *OrderByClause) validateExpression(formats strfmt.Registry) error {
	if swag.IsZero(m.Expression) { // not required
		return nil
	}

	if m.Expression != nil {
		if err := m.Expression.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Expression")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Expression")
			}
			return err
		}
	}

	return nil
}

func (m *OrderByClause) validateItemType(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemType) { // not required
		return nil
	}

	if m.ItemType != nil {
		if err := m.ItemType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemType")
			}
			return err
		}
	}

	return nil
}

func (m *OrderByClause) validateRangeVariable(formats strfmt.Registry) error {
	if swag.IsZero(m.RangeVariable) { // not required
		return nil
	}

	if m.RangeVariable != nil {
		if err := m.RangeVariable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RangeVariable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RangeVariable")
			}
			return err
		}
	}

	return nil
}

func (m *OrderByClause) validateThenBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ThenBy) { // not required
		return nil
	}

	if m.ThenBy != nil {
		if err := m.ThenBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ThenBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ThenBy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order by clause based on the context it is used
func (m *OrderByClause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpression(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRangeVariable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThenBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderByClause) contextValidateDirection(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "Direction", "body", string(m.Direction)); err != nil {
		return err
	}

	return nil
}

func (m *OrderByClause) contextValidateExpression(ctx context.Context, formats strfmt.Registry) error {

	if m.Expression != nil {

		if swag.IsZero(m.Expression) { // not required
			return nil
		}

		if err := m.Expression.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Expression")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Expression")
			}
			return err
		}
	}

	return nil
}

func (m *OrderByClause) contextValidateItemType(ctx context.Context, formats strfmt.Registry) error {

	if m.ItemType != nil {

		if swag.IsZero(m.ItemType) { // not required
			return nil
		}

		if err := m.ItemType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemType")
			}
			return err
		}
	}

	return nil
}

func (m *OrderByClause) contextValidateRangeVariable(ctx context.Context, formats strfmt.Registry) error {

	if m.RangeVariable != nil {

		if swag.IsZero(m.RangeVariable) { // not required
			return nil
		}

		if err := m.RangeVariable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RangeVariable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("RangeVariable")
			}
			return err
		}
	}

	return nil
}

func (m *OrderByClause) contextValidateThenBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ThenBy != nil {

		if swag.IsZero(m.ThenBy) { // not required
			return nil
		}

		if err := m.ThenBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ThenBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ThenBy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderByClause) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderByClause) UnmarshalBinary(b []byte) error {
	var res OrderByClause
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
