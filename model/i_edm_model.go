// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IEdmModel i edm model
//
// swagger:model IEdmModel
type IEdmModel struct {

	// direct value annotations manager
	// Read Only: true
	DirectValueAnnotationsManager IEdmDirectValueAnnotationsManager `json:"DirectValueAnnotationsManager,omitempty"`

	// referenced models
	// Read Only: true
	ReferencedModels []*IEdmModel `json:"ReferencedModels"`

	// schema elements
	// Read Only: true
	SchemaElements []*IEdmSchemaElement `json:"SchemaElements"`

	// vocabulary annotations
	// Read Only: true
	VocabularyAnnotations []*IEdmVocabularyAnnotation `json:"VocabularyAnnotations"`
}

// Validate validates this i edm model
func (m *IEdmModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReferencedModels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemaElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVocabularyAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IEdmModel) validateReferencedModels(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferencedModels) { // not required
		return nil
	}

	for i := 0; i < len(m.ReferencedModels); i++ {
		if swag.IsZero(m.ReferencedModels[i]) { // not required
			continue
		}

		if m.ReferencedModels[i] != nil {
			if err := m.ReferencedModels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ReferencedModels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ReferencedModels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IEdmModel) validateSchemaElements(formats strfmt.Registry) error {
	if swag.IsZero(m.SchemaElements) { // not required
		return nil
	}

	for i := 0; i < len(m.SchemaElements); i++ {
		if swag.IsZero(m.SchemaElements[i]) { // not required
			continue
		}

		if m.SchemaElements[i] != nil {
			if err := m.SchemaElements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SchemaElements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SchemaElements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IEdmModel) validateVocabularyAnnotations(formats strfmt.Registry) error {
	if swag.IsZero(m.VocabularyAnnotations) { // not required
		return nil
	}

	for i := 0; i < len(m.VocabularyAnnotations); i++ {
		if swag.IsZero(m.VocabularyAnnotations[i]) { // not required
			continue
		}

		if m.VocabularyAnnotations[i] != nil {
			if err := m.VocabularyAnnotations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("VocabularyAnnotations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("VocabularyAnnotations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this i edm model based on the context it is used
func (m *IEdmModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReferencedModels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchemaElements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVocabularyAnnotations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IEdmModel) contextValidateReferencedModels(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ReferencedModels", "body", []*IEdmModel(m.ReferencedModels)); err != nil {
		return err
	}

	for i := 0; i < len(m.ReferencedModels); i++ {

		if m.ReferencedModels[i] != nil {

			if swag.IsZero(m.ReferencedModels[i]) { // not required
				return nil
			}

			if err := m.ReferencedModels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ReferencedModels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ReferencedModels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IEdmModel) contextValidateSchemaElements(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "SchemaElements", "body", []*IEdmSchemaElement(m.SchemaElements)); err != nil {
		return err
	}

	for i := 0; i < len(m.SchemaElements); i++ {

		if m.SchemaElements[i] != nil {

			if swag.IsZero(m.SchemaElements[i]) { // not required
				return nil
			}

			if err := m.SchemaElements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SchemaElements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SchemaElements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IEdmModel) contextValidateVocabularyAnnotations(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "VocabularyAnnotations", "body", []*IEdmVocabularyAnnotation(m.VocabularyAnnotations)); err != nil {
		return err
	}

	for i := 0; i < len(m.VocabularyAnnotations); i++ {

		if m.VocabularyAnnotations[i] != nil {

			if swag.IsZero(m.VocabularyAnnotations[i]) { // not required
				return nil
			}

			if err := m.VocabularyAnnotations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("VocabularyAnnotations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("VocabularyAnnotations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IEdmModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IEdmModel) UnmarshalBinary(b []byte) error {
	var res IEdmModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
