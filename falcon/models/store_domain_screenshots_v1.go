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

// StoreDomainScreenshotsV1 store domain screenshots v1
//
// swagger:model store.domain.ScreenshotsV1
type StoreDomainScreenshotsV1 struct {

	// alt
	Alt string `json:"alt,omitempty"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this store domain screenshots v1
func (m *StoreDomainScreenshotsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoreDomainScreenshotsV1) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this store domain screenshots v1 based on context it is used
func (m *StoreDomainScreenshotsV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StoreDomainScreenshotsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoreDomainScreenshotsV1) UnmarshalBinary(b []byte) error {
	var res StoreDomainScreenshotsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
