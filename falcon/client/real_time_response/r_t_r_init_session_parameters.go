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

// NewRTRInitSessionParams creates a new RTRInitSessionParams object
// with the default values initialized.
func NewRTRInitSessionParams() *RTRInitSessionParams {
	var ()
	return &RTRInitSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRTRInitSessionParamsWithTimeout creates a new RTRInitSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRTRInitSessionParamsWithTimeout(timeout time.Duration) *RTRInitSessionParams {
	var ()
	return &RTRInitSessionParams{

		timeout: timeout,
	}
}

// NewRTRInitSessionParamsWithContext creates a new RTRInitSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewRTRInitSessionParamsWithContext(ctx context.Context) *RTRInitSessionParams {
	var ()
	return &RTRInitSessionParams{

		Context: ctx,
	}
}

// NewRTRInitSessionParamsWithHTTPClient creates a new RTRInitSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRTRInitSessionParamsWithHTTPClient(client *http.Client) *RTRInitSessionParams {
	var ()
	return &RTRInitSessionParams{
		HTTPClient: client,
	}
}

/*RTRInitSessionParams contains all the parameters to send to the API endpoint
for the r t r init session operation typically these are written to a http.Request
*/
type RTRInitSessionParams struct {

	/*Body
	 **`device_id`** The host agent ID to initialize the RTR session on.  RTR will retrieve an existing session for the calling user on this host

	 */
	Body *models.DomainInitRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the r t r init session params
func (o *RTRInitSessionParams) WithTimeout(timeout time.Duration) *RTRInitSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r init session params
func (o *RTRInitSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r init session params
func (o *RTRInitSessionParams) WithContext(ctx context.Context) *RTRInitSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r init session params
func (o *RTRInitSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r init session params
func (o *RTRInitSessionParams) WithHTTPClient(client *http.Client) *RTRInitSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r init session params
func (o *RTRInitSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the r t r init session params
func (o *RTRInitSessionParams) WithBody(body *models.DomainInitRequest) *RTRInitSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the r t r init session params
func (o *RTRInitSessionParams) SetBody(body *models.DomainInitRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RTRInitSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
