// Code generated by go-swagger; DO NOT EDIT.

package installation_tokens

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

// NewTokensDeleteParams creates a new TokensDeleteParams object
// with the default values initialized.
func NewTokensDeleteParams() *TokensDeleteParams {
	var ()
	return &TokensDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokensDeleteParamsWithTimeout creates a new TokensDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokensDeleteParamsWithTimeout(timeout time.Duration) *TokensDeleteParams {
	var ()
	return &TokensDeleteParams{

		timeout: timeout,
	}
}

// NewTokensDeleteParamsWithContext creates a new TokensDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokensDeleteParamsWithContext(ctx context.Context) *TokensDeleteParams {
	var ()
	return &TokensDeleteParams{

		Context: ctx,
	}
}

// NewTokensDeleteParamsWithHTTPClient creates a new TokensDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokensDeleteParamsWithHTTPClient(client *http.Client) *TokensDeleteParams {
	var ()
	return &TokensDeleteParams{
		HTTPClient: client,
	}
}

/*TokensDeleteParams contains all the parameters to send to the API endpoint
for the tokens delete operation typically these are written to a http.Request
*/
type TokensDeleteParams struct {

	/*Ids
	  The token ids to delete.

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tokens delete params
func (o *TokensDeleteParams) WithTimeout(timeout time.Duration) *TokensDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tokens delete params
func (o *TokensDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tokens delete params
func (o *TokensDeleteParams) WithContext(ctx context.Context) *TokensDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tokens delete params
func (o *TokensDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tokens delete params
func (o *TokensDeleteParams) WithHTTPClient(client *http.Client) *TokensDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tokens delete params
func (o *TokensDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the tokens delete params
func (o *TokensDeleteParams) WithIds(ids []string) *TokensDeleteParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the tokens delete params
func (o *TokensDeleteParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *TokensDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "csv")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
