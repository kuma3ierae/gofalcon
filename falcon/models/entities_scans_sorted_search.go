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

// EntitiesScansSortedSearch entities scans sorted search
//
// swagger:model entities.ScansSortedSearch
type EntitiesScansSortedSearch struct {

	// filter
	// Required: true
	Filter *string `json:"filter"`

	// sort
	// Required: true
	Sort *string `json:"sort"`
}

// Validate validates this entities scans sorted search
func (m *EntitiesScansSortedSearch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntitiesScansSortedSearch) validateFilter(formats strfmt.Registry) error {

	if err := validate.Required("filter", "body", m.Filter); err != nil {
		return err
	}

	return nil
}

func (m *EntitiesScansSortedSearch) validateSort(formats strfmt.Registry) error {

	if err := validate.Required("sort", "body", m.Sort); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this entities scans sorted search based on context it is used
func (m *EntitiesScansSortedSearch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntitiesScansSortedSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitiesScansSortedSearch) UnmarshalBinary(b []byte) error {
	var res EntitiesScansSortedSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
