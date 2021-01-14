// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// NewGetCSPMAzureUserScriptsAttachmentParams creates a new GetCSPMAzureUserScriptsAttachmentParams object
// with the default values initialized.
func NewGetCSPMAzureUserScriptsAttachmentParams() *GetCSPMAzureUserScriptsAttachmentParams {
	var ()
	return &GetCSPMAzureUserScriptsAttachmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCSPMAzureUserScriptsAttachmentParamsWithTimeout creates a new GetCSPMAzureUserScriptsAttachmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCSPMAzureUserScriptsAttachmentParamsWithTimeout(timeout time.Duration) *GetCSPMAzureUserScriptsAttachmentParams {
	var ()
	return &GetCSPMAzureUserScriptsAttachmentParams{

		timeout: timeout,
	}
}

// NewGetCSPMAzureUserScriptsAttachmentParamsWithContext creates a new GetCSPMAzureUserScriptsAttachmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCSPMAzureUserScriptsAttachmentParamsWithContext(ctx context.Context) *GetCSPMAzureUserScriptsAttachmentParams {
	var ()
	return &GetCSPMAzureUserScriptsAttachmentParams{

		Context: ctx,
	}
}

// NewGetCSPMAzureUserScriptsAttachmentParamsWithHTTPClient creates a new GetCSPMAzureUserScriptsAttachmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCSPMAzureUserScriptsAttachmentParamsWithHTTPClient(client *http.Client) *GetCSPMAzureUserScriptsAttachmentParams {
	var ()
	return &GetCSPMAzureUserScriptsAttachmentParams{
		HTTPClient: client,
	}
}

/*GetCSPMAzureUserScriptsAttachmentParams contains all the parameters to send to the API endpoint
for the get c s p m azure user scripts attachment operation typically these are written to a http.Request
*/
type GetCSPMAzureUserScriptsAttachmentParams struct {

	/*TenantID
	  Tenant ID to generate script for. Defaults to most recently registered tenant.

	*/
	TenantID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithTimeout(timeout time.Duration) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithContext(ctx context.Context) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithHTTPClient(client *http.Client) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantID adds the tenantID to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) WithTenantID(tenantID *string) *GetCSPMAzureUserScriptsAttachmentParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get c s p m azure user scripts attachment params
func (o *GetCSPMAzureUserScriptsAttachmentParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCSPMAzureUserScriptsAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TenantID != nil {

		// query param tenant-id
		var qrTenantID string
		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {
			if err := r.SetQueryParam("tenant-id", qTenantID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
