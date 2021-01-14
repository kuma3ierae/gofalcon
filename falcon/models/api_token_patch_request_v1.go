// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APITokenPatchRequestV1 api token patch request v1
//
// swagger:model api.tokenPatchRequestV1
type APITokenPatchRequestV1 struct {

	// The token's expiration time (RFC-3339). Null, if the token never expires.
	// Format: date-time
	ExpiresTimestamp strfmt.DateTime `json:"expires_timestamp,omitempty"`

	// The token label.
	Label string `json:"label,omitempty"`

	// Set to true to revoke the token, false to un-revoked it.
	Revoked bool `json:"revoked,omitempty"`
}

// Validate validates this api token patch request v1
func (m *APITokenPatchRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APITokenPatchRequestV1) validateExpiresTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_timestamp", "body", "date-time", m.ExpiresTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APITokenPatchRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APITokenPatchRequestV1) UnmarshalBinary(b []byte) error {
	var res APITokenPatchRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
