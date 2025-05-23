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

// Document document
//
// swagger:model Document
type Document struct {

	// company ID
	CompanyID int32 `json:"CompanyID,omitempty"`

	// creation date
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"CreationDate,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// document date
	// Format: date-time
	DocumentDate *strfmt.DateTime `json:"DocumentDate,omitempty"`

	// document ID
	DocumentID int32 `json:"DocumentID,omitempty"`

	// file
	File *File `json:"File,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// tags
	Tags []string `json:"Tags"`
}

// Validate validates this document
func (m *Document) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocumentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("CreationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateDocumentDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DocumentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("DocumentDate", "body", "date-time", m.DocumentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Document) validateFile(formats strfmt.Registry) error {
	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {
		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("File")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("File")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this document based on the context it is used
func (m *Document) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Document) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

	if m.File != nil {

		if swag.IsZero(m.File) { // not required
			return nil
		}

		if err := m.File.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("File")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("File")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Document) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Document) UnmarshalBinary(b []byte) error {
	var res Document
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
