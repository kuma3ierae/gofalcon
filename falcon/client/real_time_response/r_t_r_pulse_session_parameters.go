// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// NewRTRPulseSessionParams creates a new RTRPulseSessionParams object
// with the default values initialized.
func NewRTRPulseSessionParams() *RTRPulseSessionParams {
	var ()
	return &RTRPulseSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRTRPulseSessionParamsWithTimeout creates a new RTRPulseSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRTRPulseSessionParamsWithTimeout(timeout time.Duration) *RTRPulseSessionParams {
	var ()
	return &RTRPulseSessionParams{

		timeout: timeout,
	}
}

// NewRTRPulseSessionParamsWithContext creates a new RTRPulseSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRTRPulseSessionParamsWithContext(ctx context.Context) *RTRPulseSessionParams {
	var ()
	return &RTRPulseSessionParams{

		Context: ctx,
	}
}

// NewRTRPulseSessionParamsWithHTTPClient creates a new RTRPulseSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRTRPulseSessionParamsWithHTTPClient(client *http.Client) *RTRPulseSessionParams {
	var ()
	return &RTRPulseSessionParams{
		HTTPClient: client,
	}
}

/*RTRPulseSessionParams contains all the parameters to send to the API endpoint
for the r t r pulse session operation typically these are written to a http.Request
*/
type RTRPulseSessionParams struct {

	/*Body
	 **`device_id`** The host agent ID to refresh the RTR session on.  RTR will retrieve an existing session for the calling user on this host

	 */
	Body *models.DomainInitRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the r t r pulse session params
func (o *RTRPulseSessionParams) WithTimeout(timeout time.Duration) *RTRPulseSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r pulse session params
func (o *RTRPulseSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r pulse session params
func (o *RTRPulseSessionParams) WithContext(ctx context.Context) *RTRPulseSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r pulse session params
func (o *RTRPulseSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r pulse session params
func (o *RTRPulseSessionParams) WithHTTPClient(client *http.Client) *RTRPulseSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r pulse session params
func (o *RTRPulseSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the r t r pulse session params
func (o *RTRPulseSessionParams) WithBody(body *models.DomainInitRequest) *RTRPulseSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the r t r pulse session params
func (o *RTRPulseSessionParams) SetBody(body *models.DomainInitRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RTRPulseSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
