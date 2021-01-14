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

// NewTokensReadParams creates a new TokensReadParams object
// with the default values initialized.
func NewTokensReadParams() *TokensReadParams {
	var ()
	return &TokensReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTokensReadParamsWithTimeout creates a new TokensReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTokensReadParamsWithTimeout(timeout time.Duration) *TokensReadParams {
	var ()
	return &TokensReadParams{

		timeout: timeout,
	}
}

// NewTokensReadParamsWithContext creates a new TokensReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewTokensReadParamsWithContext(ctx context.Context) *TokensReadParams {
	var ()
	return &TokensReadParams{

		Context: ctx,
	}
}

// NewTokensReadParamsWithHTTPClient creates a new TokensReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTokensReadParamsWithHTTPClient(client *http.Client) *TokensReadParams {
	var ()
	return &TokensReadParams{
		HTTPClient: client,
	}
}

/*TokensReadParams contains all the parameters to send to the API endpoint
for the tokens read operation typically these are written to a http.Request
*/
type TokensReadParams struct {

	/*Ids
	  IDs of tokens to retrieve details for

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tokens read params
func (o *TokensReadParams) WithTimeout(timeout time.Duration) *TokensReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tokens read params
func (o *TokensReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tokens read params
func (o *TokensReadParams) WithContext(ctx context.Context) *TokensReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tokens read params
func (o *TokensReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tokens read params
func (o *TokensReadParams) WithHTTPClient(client *http.Client) *TokensReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tokens read params
func (o *TokensReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the tokens read params
func (o *TokensReadParams) WithIds(ids []string) *TokensReadParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the tokens read params
func (o *TokensReadParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *TokensReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
