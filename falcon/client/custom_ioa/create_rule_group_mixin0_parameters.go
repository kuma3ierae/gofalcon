// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// NewCreateRuleGroupMixin0Params creates a new CreateRuleGroupMixin0Params object
// with the default values initialized.
func NewCreateRuleGroupMixin0Params() *CreateRuleGroupMixin0Params {
	var ()
	return &CreateRuleGroupMixin0Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRuleGroupMixin0ParamsWithTimeout creates a new CreateRuleGroupMixin0Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRuleGroupMixin0ParamsWithTimeout(timeout time.Duration) *CreateRuleGroupMixin0Params {
	var ()
	return &CreateRuleGroupMixin0Params{

		timeout: timeout,
	}
}

// NewCreateRuleGroupMixin0ParamsWithContext creates a new CreateRuleGroupMixin0Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRuleGroupMixin0ParamsWithContext(ctx context.Context) *CreateRuleGroupMixin0Params {
	var ()
	return &CreateRuleGroupMixin0Params{

		Context: ctx,
	}
}

// NewCreateRuleGroupMixin0ParamsWithHTTPClient creates a new CreateRuleGroupMixin0Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRuleGroupMixin0ParamsWithHTTPClient(client *http.Client) *CreateRuleGroupMixin0Params {
	var ()
	return &CreateRuleGroupMixin0Params{
		HTTPClient: client,
	}
}

/*CreateRuleGroupMixin0Params contains all the parameters to send to the API endpoint
for the create rule group mixin0 operation typically these are written to a http.Request
*/
type CreateRuleGroupMixin0Params struct {

	/*XCSUSERNAME
	  The user ID

	*/
	XCSUSERNAME string
	/*Body*/
	Body *models.APIRuleGroupCreateRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) WithTimeout(timeout time.Duration) *CreateRuleGroupMixin0Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) WithContext(ctx context.Context) *CreateRuleGroupMixin0Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) WithHTTPClient(client *http.Client) *CreateRuleGroupMixin0Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERNAME adds the xCSUSERNAME to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) WithXCSUSERNAME(xCSUSERNAME string) *CreateRuleGroupMixin0Params {
	o.SetXCSUSERNAME(xCSUSERNAME)
	return o
}

// SetXCSUSERNAME adds the xCSUSERNAME to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) SetXCSUSERNAME(xCSUSERNAME string) {
	o.XCSUSERNAME = xCSUSERNAME
}

// WithBody adds the body to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) WithBody(body *models.APIRuleGroupCreateRequestV1) *CreateRuleGroupMixin0Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create rule group mixin0 params
func (o *CreateRuleGroupMixin0Params) SetBody(body *models.APIRuleGroupCreateRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRuleGroupMixin0Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-CS-USERNAME
	if err := r.SetHeaderParam("X-CS-USERNAME", o.XCSUSERNAME); err != nil {
		return err
	}

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
