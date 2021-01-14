// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxIncident falconx incident
//
// swagger:model falconx.Incident
type FalconxIncident struct {

	// details
	Details []string `json:"details"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this falconx incident
func (m *FalconxIncident) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FalconxIncident) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxIncident) UnmarshalBinary(b []byte) error {
	var res FalconxIncident
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
