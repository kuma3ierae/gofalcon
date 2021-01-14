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
	"github.com/go-openapi/swag"
)

// NewQueryPolicyRulesParams creates a new QueryPolicyRulesParams object
// with the default values initialized.
func NewQueryPolicyRulesParams() *QueryPolicyRulesParams {
	var ()
	return &QueryPolicyRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryPolicyRulesParamsWithTimeout creates a new QueryPolicyRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryPolicyRulesParamsWithTimeout(timeout time.Duration) *QueryPolicyRulesParams {
	var ()
	return &QueryPolicyRulesParams{

		timeout: timeout,
	}
}

// NewQueryPolicyRulesParamsWithContext creates a new QueryPolicyRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryPolicyRulesParamsWithContext(ctx context.Context) *QueryPolicyRulesParams {
	var ()
	return &QueryPolicyRulesParams{

		Context: ctx,
	}
}

// NewQueryPolicyRulesParamsWithHTTPClient creates a new QueryPolicyRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryPolicyRulesParamsWithHTTPClient(client *http.Client) *QueryPolicyRulesParams {
	var ()
	return &QueryPolicyRulesParams{
		HTTPClient: client,
	}
}

/*QueryPolicyRulesParams contains all the parameters to send to the API endpoint
for the query policy rules operation typically these are written to a http.Request
*/
type QueryPolicyRulesParams struct {

	/*Filter
	  FQL query specifying the filter parameters. Filter term criteria: enabled, platform, name, description, etc TODO. Filter range criteria: created_on, modified_on; use any common date format, such as '2010-05-15T14:55:21.892315096Z'.

	*/
	Filter *string
	/*ID
	  The ID of the policy container within which to query

	*/
	ID *string
	/*Limit
	  Number of ids to return.

	*/
	Limit *int64
	/*Offset
	  Starting index of overall result set from which to return ids.

	*/
	Offset *string
	/*Q
	  Match query criteria, which includes all the filter string fields, plus TODO

	*/
	Q *string
	/*Sort
	  Possible order by fields:

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query policy rules params
func (o *QueryPolicyRulesParams) WithTimeout(timeout time.Duration) *QueryPolicyRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query policy rules params
func (o *QueryPolicyRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query policy rules params
func (o *QueryPolicyRulesParams) WithContext(ctx context.Context) *QueryPolicyRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query policy rules params
func (o *QueryPolicyRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query policy rules params
func (o *QueryPolicyRulesParams) WithHTTPClient(client *http.Client) *QueryPolicyRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query policy rules params
func (o *QueryPolicyRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query policy rules params
func (o *QueryPolicyRulesParams) WithFilter(filter *string) *QueryPolicyRulesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query policy rules params
func (o *QueryPolicyRulesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the query policy rules params
func (o *QueryPolicyRulesParams) WithID(id *string) *QueryPolicyRulesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the query policy rules params
func (o *QueryPolicyRulesParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the query policy rules params
func (o *QueryPolicyRulesParams) WithLimit(limit *int64) *QueryPolicyRulesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query policy rules params
func (o *QueryPolicyRulesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query policy rules params
func (o *QueryPolicyRulesParams) WithOffset(offset *string) *QueryPolicyRulesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query policy rules params
func (o *QueryPolicyRulesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithQ adds the q to the query policy rules params
func (o *QueryPolicyRulesParams) WithQ(q *string) *QueryPolicyRulesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the query policy rules params
func (o *QueryPolicyRulesParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the query policy rules params
func (o *QueryPolicyRulesParams) WithSort(sort *string) *QueryPolicyRulesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query policy rules params
func (o *QueryPolicyRulesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryPolicyRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
