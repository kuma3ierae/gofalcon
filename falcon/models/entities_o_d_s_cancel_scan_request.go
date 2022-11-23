// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EntitiesODSCancelScanRequest entities o d s cancel scan request
//
// swagger:model entities.ODSCancelScanRequest
type EntitiesODSCancelScanRequest struct {

	// ids
	// Required: true
	Ids []string `json:"ids"`
}

// Validate validates this entities o d s cancel scan request
func (m *EntitiesODSCancelScanRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntitiesODSCancelScanRequest) validateIds(formats strfmt.Registry) error {

	if err := validate.Required("ids", "body", m.Ids); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this entities o d s cancel scan request based on context it is used
func (m *EntitiesODSCancelScanRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntitiesODSCancelScanRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitiesODSCancelScanRequest) UnmarshalBinary(b []byte) error {
	var res EntitiesODSCancelScanRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
