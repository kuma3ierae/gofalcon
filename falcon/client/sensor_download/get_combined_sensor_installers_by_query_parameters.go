// Code generated by go-swagger; DO NOT EDIT.

package sensor_download

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

// NewGetCombinedSensorInstallersByQueryParams creates a new GetCombinedSensorInstallersByQueryParams object
// with the default values initialized.
func NewGetCombinedSensorInstallersByQueryParams() *GetCombinedSensorInstallersByQueryParams {
	var ()
	return &GetCombinedSensorInstallersByQueryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCombinedSensorInstallersByQueryParamsWithTimeout creates a new GetCombinedSensorInstallersByQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCombinedSensorInstallersByQueryParamsWithTimeout(timeout time.Duration) *GetCombinedSensorInstallersByQueryParams {
	var ()
	return &GetCombinedSensorInstallersByQueryParams{

		timeout: timeout,
	}
}

// NewGetCombinedSensorInstallersByQueryParamsWithContext creates a new GetCombinedSensorInstallersByQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCombinedSensorInstallersByQueryParamsWithContext(ctx context.Context) *GetCombinedSensorInstallersByQueryParams {
	var ()
	return &GetCombinedSensorInstallersByQueryParams{

		Context: ctx,
	}
}

// NewGetCombinedSensorInstallersByQueryParamsWithHTTPClient creates a new GetCombinedSensorInstallersByQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCombinedSensorInstallersByQueryParamsWithHTTPClient(client *http.Client) *GetCombinedSensorInstallersByQueryParams {
	var ()
	return &GetCombinedSensorInstallersByQueryParams{
		HTTPClient: client,
	}
}

/*GetCombinedSensorInstallersByQueryParams contains all the parameters to send to the API endpoint
for the get combined sensor installers by query operation typically these are written to a http.Request
*/
type GetCombinedSensorInstallersByQueryParams struct {

	/*Filter
	  Filter items using a query in Falcon Query Language (FQL). An asterisk wildcard * includes all results.

	Common filter options include:
	<ul><li>platform:"windows"</li><li>version:>"5.2"</li></ul>

	*/
	Filter *string
	/*Limit
	  The number of items to return in this response (default: 100, max: 500). Use with the offset parameter to manage pagination of results.

	*/
	Limit *int64
	/*Offset
	  The first item to return, where 0 is the latest item. Use with the limit parameter to manage pagination of results.

	*/
	Offset *int64
	/*Sort
	  Sort items using their properties. Common sort options include:

	<ul><li>version|asc</li><li>release_date|desc</li></ul>

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) WithTimeout(timeout time.Duration) *GetCombinedSensorInstallersByQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) WithContext(ctx context.Context) *GetCombinedSensorInstallersByQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) WithHTTPClient(client *http.Client) *GetCombinedSensorInstallersByQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) WithFilter(filter *string) *GetCombinedSensorInstallersByQueryParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) WithLimit(limit *int64) *GetCombinedSensorInstallersByQueryParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) WithOffset(offset *int64) *GetCombinedSensorInstallersByQueryParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) WithSort(sort *string) *GetCombinedSensorInstallersByQueryParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get combined sensor installers by query params
func (o *GetCombinedSensorInstallersByQueryParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetCombinedSensorInstallersByQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
