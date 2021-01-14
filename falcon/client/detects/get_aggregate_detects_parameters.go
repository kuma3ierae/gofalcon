// Code generated by go-swagger; DO NOT EDIT.

package detects

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

// NewGetAggregateDetectsParams creates a new GetAggregateDetectsParams object
// with the default values initialized.
func NewGetAggregateDetectsParams() *GetAggregateDetectsParams {
	var ()
	return &GetAggregateDetectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAggregateDetectsParamsWithTimeout creates a new GetAggregateDetectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAggregateDetectsParamsWithTimeout(timeout time.Duration) *GetAggregateDetectsParams {
	var ()
	return &GetAggregateDetectsParams{

		timeout: timeout,
	}
}

// NewGetAggregateDetectsParamsWithContext creates a new GetAggregateDetectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAggregateDetectsParamsWithContext(ctx context.Context) *GetAggregateDetectsParams {
	var ()
	return &GetAggregateDetectsParams{

		Context: ctx,
	}
}

// NewGetAggregateDetectsParamsWithHTTPClient creates a new GetAggregateDetectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAggregateDetectsParamsWithHTTPClient(client *http.Client) *GetAggregateDetectsParams {
	var ()
	return &GetAggregateDetectsParams{
		HTTPClient: client,
	}
}

/*GetAggregateDetectsParams contains all the parameters to send to the API endpoint
for the get aggregate detects operation typically these are written to a http.Request
*/
type GetAggregateDetectsParams struct {

	/*Body
	  Query criteria and settings

	*/
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get aggregate detects params
func (o *GetAggregateDetectsParams) WithTimeout(timeout time.Duration) *GetAggregateDetectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aggregate detects params
func (o *GetAggregateDetectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aggregate detects params
func (o *GetAggregateDetectsParams) WithContext(ctx context.Context) *GetAggregateDetectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aggregate detects params
func (o *GetAggregateDetectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aggregate detects params
func (o *GetAggregateDetectsParams) WithHTTPClient(client *http.Client) *GetAggregateDetectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aggregate detects params
func (o *GetAggregateDetectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get aggregate detects params
func (o *GetAggregateDetectsParams) WithBody(body []*models.MsaAggregateQueryRequest) *GetAggregateDetectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get aggregate detects params
func (o *GetAggregateDetectsParams) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetAggregateDetectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
