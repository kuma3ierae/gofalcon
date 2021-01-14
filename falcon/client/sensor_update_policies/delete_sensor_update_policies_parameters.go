// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteSensorUpdatePoliciesParams creates a new DeleteSensorUpdatePoliciesParams object
// with the default values initialized.
func NewDeleteSensorUpdatePoliciesParams() *DeleteSensorUpdatePoliciesParams {
	var ()
	return &DeleteSensorUpdatePoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSensorUpdatePoliciesParamsWithTimeout creates a new DeleteSensorUpdatePoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSensorUpdatePoliciesParamsWithTimeout(timeout time.Duration) *DeleteSensorUpdatePoliciesParams {
	var ()
	return &DeleteSensorUpdatePoliciesParams{

		timeout: timeout,
	}
}

// NewDeleteSensorUpdatePoliciesParamsWithContext creates a new DeleteSensorUpdatePoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSensorUpdatePoliciesParamsWithContext(ctx context.Context) *DeleteSensorUpdatePoliciesParams {
	var ()
	return &DeleteSensorUpdatePoliciesParams{

		Context: ctx,
	}
}

// NewDeleteSensorUpdatePoliciesParamsWithHTTPClient creates a new DeleteSensorUpdatePoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSensorUpdatePoliciesParamsWithHTTPClient(client *http.Client) *DeleteSensorUpdatePoliciesParams {
	var ()
	return &DeleteSensorUpdatePoliciesParams{
		HTTPClient: client,
	}
}

/*DeleteSensorUpdatePoliciesParams contains all the parameters to send to the API endpoint
for the delete sensor update policies operation typically these are written to a http.Request
*/
type DeleteSensorUpdatePoliciesParams struct {

	/*Ids
	  The IDs of the Sensor Update Policies to delete

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sensor update policies params
func (o *DeleteSensorUpdatePoliciesParams) WithTimeout(timeout time.Duration) *DeleteSensorUpdatePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sensor update policies params
func (o *DeleteSensorUpdatePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sensor update policies params
func (o *DeleteSensorUpdatePoliciesParams) WithContext(ctx context.Context) *DeleteSensorUpdatePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sensor update policies params
func (o *DeleteSensorUpdatePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sensor update policies params
func (o *DeleteSensorUpdatePoliciesParams) WithHTTPClient(client *http.Client) *DeleteSensorUpdatePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sensor update policies params
func (o *DeleteSensorUpdatePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the delete sensor update policies params
func (o *DeleteSensorUpdatePoliciesParams) WithIds(ids []string) *DeleteSensorUpdatePoliciesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete sensor update policies params
func (o *DeleteSensorUpdatePoliciesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSensorUpdatePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "multi")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
