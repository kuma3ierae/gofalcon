// Code generated by go-swagger; DO NOT EDIT.

package event_streams

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
)

// NewListAvailableStreamsOAuth2Params creates a new ListAvailableStreamsOAuth2Params object
// with the default values initialized.
func NewListAvailableStreamsOAuth2Params() *ListAvailableStreamsOAuth2Params {
	var ()
	return &ListAvailableStreamsOAuth2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAvailableStreamsOAuth2ParamsWithTimeout creates a new ListAvailableStreamsOAuth2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAvailableStreamsOAuth2ParamsWithTimeout(timeout time.Duration) *ListAvailableStreamsOAuth2Params {
	var ()
	return &ListAvailableStreamsOAuth2Params{

		timeout: timeout,
	}
}

// NewListAvailableStreamsOAuth2ParamsWithContext creates a new ListAvailableStreamsOAuth2Params object
// with the default values initialized, and the ability to set a context for a request
func NewListAvailableStreamsOAuth2ParamsWithContext(ctx context.Context) *ListAvailableStreamsOAuth2Params {
	var ()
	return &ListAvailableStreamsOAuth2Params{

		Context: ctx,
	}
}

// NewListAvailableStreamsOAuth2ParamsWithHTTPClient creates a new ListAvailableStreamsOAuth2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAvailableStreamsOAuth2ParamsWithHTTPClient(client *http.Client) *ListAvailableStreamsOAuth2Params {
	var ()
	return &ListAvailableStreamsOAuth2Params{
		HTTPClient: client,
	}
}

/*ListAvailableStreamsOAuth2Params contains all the parameters to send to the API endpoint
for the list available streams o auth2 operation typically these are written to a http.Request
*/
type ListAvailableStreamsOAuth2Params struct {

	/*AppID
	  Label that identifies your connection. Max: 32 alphanumeric characters (a-z, A-Z, 0-9).

	*/
	AppID string
	/*Format
	  Format for streaming events. Valid values: json, flatjson

	*/
	Format *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) WithTimeout(timeout time.Duration) *ListAvailableStreamsOAuth2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) WithContext(ctx context.Context) *ListAvailableStreamsOAuth2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) WithHTTPClient(client *http.Client) *ListAvailableStreamsOAuth2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) WithAppID(appID string) *ListAvailableStreamsOAuth2Params {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) SetAppID(appID string) {
	o.AppID = appID
}

// WithFormat adds the format to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) WithFormat(format *string) *ListAvailableStreamsOAuth2Params {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the list available streams o auth2 params
func (o *ListAvailableStreamsOAuth2Params) SetFormat(format *string) {
	o.Format = format
}

// WriteToRequest writes these params to a swagger request
func (o *ListAvailableStreamsOAuth2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param appId
	qrAppID := o.AppID
	qAppID := qrAppID
	if qAppID != "" {
		if err := r.SetQueryParam("appId", qAppID); err != nil {
			return err
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
