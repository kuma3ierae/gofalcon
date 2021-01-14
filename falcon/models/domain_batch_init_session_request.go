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

// DomainBatchInitSessionRequest domain batch init session request
//
// swagger:model domain.BatchInitSessionRequest
type DomainBatchInitSessionRequest struct {

	// existing batch id
	// Required: true
	ExistingBatchID *string `json:"existing_batch_id"`

	// host ids
	// Required: true
	HostIds []string `json:"host_ids"`

	// queue offline
	// Required: true
	QueueOffline *bool `json:"queue_offline"`
}

// Validate validates this domain batch init session request
func (m *DomainBatchInitSessionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExistingBatchID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueOffline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainBatchInitSessionRequest) validateExistingBatchID(formats strfmt.Registry) error {

	if err := validate.Required("existing_batch_id", "body", m.ExistingBatchID); err != nil {
		return err
	}

	return nil
}

func (m *DomainBatchInitSessionRequest) validateHostIds(formats strfmt.Registry) error {

	if err := validate.Required("host_ids", "body", m.HostIds); err != nil {
		return err
	}

	return nil
}

func (m *DomainBatchInitSessionRequest) validateQueueOffline(formats strfmt.Registry) error {

	if err := validate.Required("queue_offline", "body", m.QueueOffline); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainBatchInitSessionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainBatchInitSessionRequest) UnmarshalBinary(b []byte) error {
	var res DomainBatchInitSessionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
