// Code generated by go-swagger; DO NOT EDIT.

package ods

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

// NewQueryMaliciousFilesParams creates a new QueryMaliciousFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryMaliciousFilesParams() *QueryMaliciousFilesParams {
	return &QueryMaliciousFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryMaliciousFilesParamsWithTimeout creates a new QueryMaliciousFilesParams object
// with the ability to set a timeout on a request.
func NewQueryMaliciousFilesParamsWithTimeout(timeout time.Duration) *QueryMaliciousFilesParams {
	return &QueryMaliciousFilesParams{
		timeout: timeout,
	}
}

// NewQueryMaliciousFilesParamsWithContext creates a new QueryMaliciousFilesParams object
// with the ability to set a context for a request.
func NewQueryMaliciousFilesParamsWithContext(ctx context.Context) *QueryMaliciousFilesParams {
	return &QueryMaliciousFilesParams{
		Context: ctx,
	}
}

// NewQueryMaliciousFilesParamsWithHTTPClient creates a new QueryMaliciousFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryMaliciousFilesParamsWithHTTPClient(client *http.Client) *QueryMaliciousFilesParams {
	return &QueryMaliciousFilesParams{
		HTTPClient: client,
	}
}

/*
QueryMaliciousFilesParams contains all the parameters to send to the API endpoint

	for the query malicious files operation.

	Typically these are written to a http.Request.
*/
type QueryMaliciousFilesParams struct {

	/* Filter.

	   A FQL compatible query string. Terms: [id cid scan_id host_id host_scan_id filepath filename hash pattern_id severity quarantined last_updated]
	*/
	Filter string

	/* Limit.

	   The max number of resources to return

	   Default: 500
	*/
	Limit *int64

	/* Offset.

	   Index of the starting resource
	*/
	Offset *int64

	/* Sort.

	   The property to sort on, followed by a |, followed by the sort direction, either "asc" or "desc"

	   Default: "last_updated|desc"
	*/
	Sort string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query malicious files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryMaliciousFilesParams) WithDefaults() *QueryMaliciousFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query malicious files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryMaliciousFilesParams) SetDefaults() {
	var (
		limitDefault = int64(500)

		offsetDefault = int64(0)

		sortDefault = string("last_updated|desc")
	)

	val := QueryMaliciousFilesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		Sort:   sortDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the query malicious files params
func (o *QueryMaliciousFilesParams) WithTimeout(timeout time.Duration) *QueryMaliciousFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query malicious files params
func (o *QueryMaliciousFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query malicious files params
func (o *QueryMaliciousFilesParams) WithContext(ctx context.Context) *QueryMaliciousFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query malicious files params
func (o *QueryMaliciousFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query malicious files params
func (o *QueryMaliciousFilesParams) WithHTTPClient(client *http.Client) *QueryMaliciousFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query malicious files params
func (o *QueryMaliciousFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query malicious files params
func (o *QueryMaliciousFilesParams) WithFilter(filter string) *QueryMaliciousFilesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query malicious files params
func (o *QueryMaliciousFilesParams) SetFilter(filter string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query malicious files params
func (o *QueryMaliciousFilesParams) WithLimit(limit *int64) *QueryMaliciousFilesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query malicious files params
func (o *QueryMaliciousFilesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query malicious files params
func (o *QueryMaliciousFilesParams) WithOffset(offset *int64) *QueryMaliciousFilesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query malicious files params
func (o *QueryMaliciousFilesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query malicious files params
func (o *QueryMaliciousFilesParams) WithSort(sort string) *QueryMaliciousFilesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query malicious files params
func (o *QueryMaliciousFilesParams) SetSort(sort string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryMaliciousFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param filter
	qrFilter := o.Filter
	qFilter := qrFilter

	if err := r.SetQueryParam("filter", qFilter); err != nil {
		return err
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
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	// query param sort
	qrSort := o.Sort
	qSort := qrSort

	if err := r.SetQueryParam("sort", qSort); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
