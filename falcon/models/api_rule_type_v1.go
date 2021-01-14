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

// APIRuleTypeV1 api rule type v1
//
// swagger:model api.RuleTypeV1
type APIRuleTypeV1 struct {

	// channel
	// Required: true
	Channel *int64 `json:"channel"`

	// disposition map
	// Required: true
	DispositionMap []*DomainDisposition `json:"disposition_map"`

	// fields
	// Required: true
	Fields []*DomainField `json:"fields"`

	// id
	// Required: true
	ID *string `json:"id"`

	// long desc
	// Required: true
	LongDesc *string `json:"long_desc"`

	// name
	// Required: true
	Name *string `json:"name"`

	// platform
	// Required: true
	Platform *string `json:"platform"`

	// released
	// Required: true
	Released *bool `json:"released"`
}

// Validate validates this api rule type v1
func (m *APIRuleTypeV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDispositionMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLongDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleased(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIRuleTypeV1) validateChannel(formats strfmt.Registry) error {

	if err := validate.Required("channel", "body", m.Channel); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleTypeV1) validateDispositionMap(formats strfmt.Registry) error {

	if err := validate.Required("disposition_map", "body", m.DispositionMap); err != nil {
		return err
	}

	for i := 0; i < len(m.DispositionMap); i++ {
		if swag.IsZero(m.DispositionMap[i]) { // not required
			continue
		}

		if m.DispositionMap[i] != nil {
			if err := m.DispositionMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disposition_map" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRuleTypeV1) validateFields(formats strfmt.Registry) error {

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	for i := 0; i < len(m.Fields); i++ {
		if swag.IsZero(m.Fields[i]) { // not required
			continue
		}

		if m.Fields[i] != nil {
			if err := m.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRuleTypeV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleTypeV1) validateLongDesc(formats strfmt.Registry) error {

	if err := validate.Required("long_desc", "body", m.LongDesc); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleTypeV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleTypeV1) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleTypeV1) validateReleased(formats strfmt.Registry) error {

	if err := validate.Required("released", "body", m.Released); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIRuleTypeV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIRuleTypeV1) UnmarshalBinary(b []byte) error {
	var res APIRuleTypeV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
