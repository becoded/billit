// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OAuthAccessTokenRequest o auth access token request
//
// swagger:model OAuthAccessTokenRequest
type OAuthAccessTokenRequest struct {

	// client id
	ClientID string `json:"client_id,omitempty"`

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// grant type
	GrantType string `json:"grant_type,omitempty"`

	// redirect uri
	RedirectURI string `json:"redirect_uri,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this o auth access token request
func (m *OAuthAccessTokenRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o auth access token request based on context it is used
func (m *OAuthAccessTokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OAuthAccessTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuthAccessTokenRequest) UnmarshalBinary(b []byte) error {
	var res OAuthAccessTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
