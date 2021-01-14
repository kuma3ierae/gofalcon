// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewAggregateEventsParams creates a new AggregateEventsParams object
// with the default values initialized.
func NewAggregateEventsParams() *AggregateEventsParams {
	var ()
	return &AggregateEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateEventsParamsWithTimeout creates a new AggregateEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAggregateEventsParamsWithTimeout(timeout time.Duration) *AggregateEventsParams {
	var ()
	return &AggregateEventsParams{

		timeout: timeout,
	}
}

// NewAggregateEventsParamsWithContext creates a new AggregateEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAggregateEventsParamsWithContext(ctx context.Context) *AggregateEventsParams {
	var ()
	return &AggregateEventsParams{

		Context: ctx,
	}
}

// NewAggregateEventsParamsWithHTTPClient creates a new AggregateEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAggregateEventsParamsWithHTTPClient(client *http.Client) *AggregateEventsParams {
	var ()
	return &AggregateEventsParams{
		HTTPClient: client,
	}
}

/*AggregateEventsParams contains all the parameters to send to the API endpoint
for the aggregate events operation typically these are written to a http.Request
*/
type AggregateEventsParams struct {

	/*Body
	  Query criteria and settings

	*/
	Body []*models.FwmgrMsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the aggregate events params
func (o *AggregateEventsParams) WithTimeout(timeout time.Duration) *AggregateEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate events params
func (o *AggregateEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate events params
func (o *AggregateEventsParams) WithContext(ctx context.Context) *AggregateEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate events params
func (o *AggregateEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate events params
func (o *AggregateEventsParams) WithHTTPClient(client *http.Client) *AggregateEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate events params
func (o *AggregateEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the aggregate events params
func (o *AggregateEventsParams) WithBody(body []*models.FwmgrMsaAggregateQueryRequest) *AggregateEventsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the aggregate events params
func (o *AggregateEventsParams) SetBody(body []*models.FwmgrMsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
