// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegistrationAWSAccountConsoleURL registration a w s account console URL
//
// swagger:model registration.AWSAccountConsoleURL
type RegistrationAWSAccountConsoleURL struct {

	// account id
	AccountID string `json:"account_id,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this registration a w s account console URL
func (m *RegistrationAWSAccountConsoleURL) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationAWSAccountConsoleURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationAWSAccountConsoleURL) UnmarshalBinary(b []byte) error {
	var res RegistrationAWSAccountConsoleURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
