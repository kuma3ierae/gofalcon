// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MalqueryRequestMetaInfo malquery request meta info
//
// swagger:model malquery.RequestMetaInfo
type MalqueryRequestMetaInfo struct {

	// Options for a hunt or exact search request
	Options *MalqueryExternalHuntOptions `json:"options,omitempty"`

	// pagination
	Pagination *MsaPaging `json:"pagination,omitempty"`

	// Patterns to search for
	Patterns []*MalquerySearchParameter `json:"patterns"`

	// powered by
	PoweredBy string `json:"powered_by,omitempty"`

	// Elapsed time since the request started in seconds
	QueryTime float64 `json:"query_time,omitempty"`

	// Request ID returned after creating a hunt or exact search
	Reqid string `json:"reqid,omitempty"`

	// Request type. Possible values: hunt, search
	Reqtype string `json:"reqtype,omitempty"`

	// Sample ID
	Sample string `json:"sample,omitempty"`

	// Result statistics
	Stats *MalqueryStats `json:"stats,omitempty"`

	// Request status. Possible values: inprogress, failed, done
	Status string `json:"status,omitempty"`

	// trace id
	// Required: true
	TraceID *string `json:"trace_id"`

	// writes
	Writes *MsaResources `json:"writes,omitempty"`

	// YARA rule to be monitored
	YaraRule string `json:"yara_rule,omitempty"`
}

// Validate validates this malquery request meta info
func (m *MalqueryRequestMetaInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatterns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrites(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MalqueryRequestMetaInfo) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *MalqueryRequestMetaInfo) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *MalqueryRequestMetaInfo) validatePatterns(formats strfmt.Registry) error {

	if swag.IsZero(m.Patterns) { // not required
		return nil
	}

	for i := 0; i < len(m.Patterns); i++ {
		if swag.IsZero(m.Patterns[i]) { // not required
			continue
		}

		if m.Patterns[i] != nil {
			if err := m.Patterns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MalqueryRequestMetaInfo) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *MalqueryRequestMetaInfo) validateTraceID(formats strfmt.Registry) error {

	if err := validate.Required("trace_id", "body", m.TraceID); err != nil {
		return err
	}

	return nil
}

func (m *MalqueryRequestMetaInfo) validateWrites(formats strfmt.Registry) error {

	if swag.IsZero(m.Writes) { // not required
		return nil
	}

	if m.Writes != nil {
		if err := m.Writes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MalqueryRequestMetaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MalqueryRequestMetaInfo) UnmarshalBinary(b []byte) error {
	var res MalqueryRequestMetaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
