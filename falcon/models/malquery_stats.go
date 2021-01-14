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

// MalqueryStats malquery stats
//
// swagger:model malquery.Stats
type MalqueryStats struct {

	// Number of clean samples
	// Required: true
	CleanCount *int32 `json:"clean_count"`

	// Number of malicious samples
	// Required: true
	MalwareCount *int32 `json:"malware_count"`

	// Number of potentially unwanted samples such as adware
	// Required: true
	PuaCount *int32 `json:"pua_count"`

	// Total number of samples
	// Required: true
	TotalCount *int32 `json:"total_count"`

	// Number of unknown (which could not be classified) samples
	// Required: true
	UnknownCount *int32 `json:"unknown_count"`
}

// Validate validates this malquery stats
func (m *MalqueryStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCleanCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMalwareCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePuaCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnknownCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MalqueryStats) validateCleanCount(formats strfmt.Registry) error {

	if err := validate.Required("clean_count", "body", m.CleanCount); err != nil {
		return err
	}

	return nil
}

func (m *MalqueryStats) validateMalwareCount(formats strfmt.Registry) error {

	if err := validate.Required("malware_count", "body", m.MalwareCount); err != nil {
		return err
	}

	return nil
}

func (m *MalqueryStats) validatePuaCount(formats strfmt.Registry) error {

	if err := validate.Required("pua_count", "body", m.PuaCount); err != nil {
		return err
	}

	return nil
}

func (m *MalqueryStats) validateTotalCount(formats strfmt.Registry) error {

	if err := validate.Required("total_count", "body", m.TotalCount); err != nil {
		return err
	}

	return nil
}

func (m *MalqueryStats) validateUnknownCount(formats strfmt.Registry) error {

	if err := validate.Required("unknown_count", "body", m.UnknownCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MalqueryStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MalqueryStats) UnmarshalBinary(b []byte) error {
	var res MalqueryStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
