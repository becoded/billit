// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ToProcessSaveRequest to process save request
//
// swagger:model ToProcessSaveRequest
type ToProcessSaveRequest struct {

	// file
	File *File `json:"File,omitempty"`

	// file source
	FileSource *File `json:"FileSource,omitempty"`

	// Optional
	MetaData *UploadMetadata `json:"MetaData,omitempty"`

	// split
	Split bool `json:"Split,omitempty"`
}

// Validate validates this to process save request
func (m *ToProcessSaveRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ToProcessSaveRequest) validateFile(formats strfmt.Registry) error {
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

func (m *ToProcessSaveRequest) validateFileSource(formats strfmt.Registry) error {
	if swag.IsZero(m.FileSource) { // not required
		return nil
	}

	if m.FileSource != nil {
		if err := m.FileSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FileSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FileSource")
			}
			return err
		}
	}

	return nil
}

func (m *ToProcessSaveRequest) validateMetaData(formats strfmt.Registry) error {
	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	if m.MetaData != nil {
		if err := m.MetaData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MetaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MetaData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this to process save request based on the context it is used
func (m *ToProcessSaveRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetaData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ToProcessSaveRequest) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ToProcessSaveRequest) contextValidateFileSource(ctx context.Context, formats strfmt.Registry) error {

	if m.FileSource != nil {

		if swag.IsZero(m.FileSource) { // not required
			return nil
		}

		if err := m.FileSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FileSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FileSource")
			}
			return err
		}
	}

	return nil
}

func (m *ToProcessSaveRequest) contextValidateMetaData(ctx context.Context, formats strfmt.Registry) error {

	if m.MetaData != nil {

		if swag.IsZero(m.MetaData) { // not required
			return nil
		}

		if err := m.MetaData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MetaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MetaData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ToProcessSaveRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ToProcessSaveRequest) UnmarshalBinary(b []byte) error {
	var res ToProcessSaveRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
