// Code generated by go-swagger; DO NOT EDIT.

package iocs

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

// NewQueryIOCsParams creates a new QueryIOCsParams object
// with the default values initialized.
func NewQueryIOCsParams() *QueryIOCsParams {
	var ()
	return &QueryIOCsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryIOCsParamsWithTimeout creates a new QueryIOCsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryIOCsParamsWithTimeout(timeout time.Duration) *QueryIOCsParams {
	var ()
	return &QueryIOCsParams{

		timeout: timeout,
	}
}

// NewQueryIOCsParamsWithContext creates a new QueryIOCsParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryIOCsParamsWithContext(ctx context.Context) *QueryIOCsParams {
	var ()
	return &QueryIOCsParams{

		Context: ctx,
	}
}

// NewQueryIOCsParamsWithHTTPClient creates a new QueryIOCsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryIOCsParamsWithHTTPClient(client *http.Client) *QueryIOCsParams {
	var ()
	return &QueryIOCsParams{
		HTTPClient: client,
	}
}

/*QueryIOCsParams contains all the parameters to send to the API endpoint
for the query i o cs operation typically these are written to a http.Request
*/
type QueryIOCsParams struct {

	/*CreatedBy
	  created_by

	*/
	CreatedBy *string
	/*DeletedBy
	  The user or API client who deleted the custom IOC

	*/
	DeletedBy *string
	/*FromExpirationTimestamp
	  Find custom IOCs created after this time (RFC-3339 timestamp)

	*/
	FromExpirationTimestamp *string
	/*IncludeDeleted

	true: Include deleted IOCs

	false: Don't include deleted IOCs (default)


	*/
	IncludeDeleted *string
	/*Policies
	  \ndetect: Find custom IOCs that produce notifications\n\nnone: Find custom IOCs the particular indicator has been detected on a host. This is equivalent to turning the indicator off.


	*/
	Policies *string
	/*ShareLevels
	  The level at which the indicator will be shared. Currently only red share level (not shared) is supported, indicating that the IOC isn't shared with other FH customers.

	*/
	ShareLevels *string
	/*Sources
	  The source where this indicator originated. This can be used for tracking where this indicator was defined. Limit 200 characters.

	*/
	Sources *string
	/*ToExpirationTimestamp
	  Find custom IOCs created before this time (RFC-3339 timestamp)

	*/
	ToExpirationTimestamp *string
	/*Types

	The type of the indicator. Valid types include:

	sha256: A hex-encoded sha256 hash string. Length - min: 64, max: 64.

	md5: A hex-encoded md5 hash string. Length - min 32, max: 32.

	domain: A domain name. Length - min: 1, max: 200.

	ipv4: An IPv4 address. Must be a valid IP address.

	ipv6: An IPv6 address. Must be a valid IP address.


	*/
	Types *string
	/*Values
	  The string representation of the indicator

	*/
	Values *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query i o cs params
func (o *QueryIOCsParams) WithTimeout(timeout time.Duration) *QueryIOCsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query i o cs params
func (o *QueryIOCsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query i o cs params
func (o *QueryIOCsParams) WithContext(ctx context.Context) *QueryIOCsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query i o cs params
func (o *QueryIOCsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query i o cs params
func (o *QueryIOCsParams) WithHTTPClient(client *http.Client) *QueryIOCsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query i o cs params
func (o *QueryIOCsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedBy adds the createdBy to the query i o cs params
func (o *QueryIOCsParams) WithCreatedBy(createdBy *string) *QueryIOCsParams {
	o.SetCreatedBy(createdBy)
	return o
}

// SetCreatedBy adds the createdBy to the query i o cs params
func (o *QueryIOCsParams) SetCreatedBy(createdBy *string) {
	o.CreatedBy = createdBy
}

// WithDeletedBy adds the deletedBy to the query i o cs params
func (o *QueryIOCsParams) WithDeletedBy(deletedBy *string) *QueryIOCsParams {
	o.SetDeletedBy(deletedBy)
	return o
}

// SetDeletedBy adds the deletedBy to the query i o cs params
func (o *QueryIOCsParams) SetDeletedBy(deletedBy *string) {
	o.DeletedBy = deletedBy
}

// WithFromExpirationTimestamp adds the fromExpirationTimestamp to the query i o cs params
func (o *QueryIOCsParams) WithFromExpirationTimestamp(fromExpirationTimestamp *string) *QueryIOCsParams {
	o.SetFromExpirationTimestamp(fromExpirationTimestamp)
	return o
}

// SetFromExpirationTimestamp adds the fromExpirationTimestamp to the query i o cs params
func (o *QueryIOCsParams) SetFromExpirationTimestamp(fromExpirationTimestamp *string) {
	o.FromExpirationTimestamp = fromExpirationTimestamp
}

// WithIncludeDeleted adds the includeDeleted to the query i o cs params
func (o *QueryIOCsParams) WithIncludeDeleted(includeDeleted *string) *QueryIOCsParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the query i o cs params
func (o *QueryIOCsParams) SetIncludeDeleted(includeDeleted *string) {
	o.IncludeDeleted = includeDeleted
}

// WithPolicies adds the policies to the query i o cs params
func (o *QueryIOCsParams) WithPolicies(policies *string) *QueryIOCsParams {
	o.SetPolicies(policies)
	return o
}

// SetPolicies adds the policies to the query i o cs params
func (o *QueryIOCsParams) SetPolicies(policies *string) {
	o.Policies = policies
}

// WithShareLevels adds the shareLevels to the query i o cs params
func (o *QueryIOCsParams) WithShareLevels(shareLevels *string) *QueryIOCsParams {
	o.SetShareLevels(shareLevels)
	return o
}

// SetShareLevels adds the shareLevels to the query i o cs params
func (o *QueryIOCsParams) SetShareLevels(shareLevels *string) {
	o.ShareLevels = shareLevels
}

// WithSources adds the sources to the query i o cs params
func (o *QueryIOCsParams) WithSources(sources *string) *QueryIOCsParams {
	o.SetSources(sources)
	return o
}

// SetSources adds the sources to the query i o cs params
func (o *QueryIOCsParams) SetSources(sources *string) {
	o.Sources = sources
}

// WithToExpirationTimestamp adds the toExpirationTimestamp to the query i o cs params
func (o *QueryIOCsParams) WithToExpirationTimestamp(toExpirationTimestamp *string) *QueryIOCsParams {
	o.SetToExpirationTimestamp(toExpirationTimestamp)
	return o
}

// SetToExpirationTimestamp adds the toExpirationTimestamp to the query i o cs params
func (o *QueryIOCsParams) SetToExpirationTimestamp(toExpirationTimestamp *string) {
	o.ToExpirationTimestamp = toExpirationTimestamp
}

// WithTypes adds the types to the query i o cs params
func (o *QueryIOCsParams) WithTypes(types *string) *QueryIOCsParams {
	o.SetTypes(types)
	return o
}

// SetTypes adds the types to the query i o cs params
func (o *QueryIOCsParams) SetTypes(types *string) {
	o.Types = types
}

// WithValues adds the values to the query i o cs params
func (o *QueryIOCsParams) WithValues(values *string) *QueryIOCsParams {
	o.SetValues(values)
	return o
}

// SetValues adds the values to the query i o cs params
func (o *QueryIOCsParams) SetValues(values *string) {
	o.Values = values
}

// WriteToRequest writes these params to a swagger request
func (o *QueryIOCsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatedBy != nil {

		// query param created_by
		var qrCreatedBy string
		if o.CreatedBy != nil {
			qrCreatedBy = *o.CreatedBy
		}
		qCreatedBy := qrCreatedBy
		if qCreatedBy != "" {
			if err := r.SetQueryParam("created_by", qCreatedBy); err != nil {
				return err
			}
		}

	}

	if o.DeletedBy != nil {

		// query param deleted_by
		var qrDeletedBy string
		if o.DeletedBy != nil {
			qrDeletedBy = *o.DeletedBy
		}
		qDeletedBy := qrDeletedBy
		if qDeletedBy != "" {
			if err := r.SetQueryParam("deleted_by", qDeletedBy); err != nil {
				return err
			}
		}

	}

	if o.FromExpirationTimestamp != nil {

		// query param from.expiration_timestamp
		var qrFromExpirationTimestamp string
		if o.FromExpirationTimestamp != nil {
			qrFromExpirationTimestamp = *o.FromExpirationTimestamp
		}
		qFromExpirationTimestamp := qrFromExpirationTimestamp
		if qFromExpirationTimestamp != "" {
			if err := r.SetQueryParam("from.expiration_timestamp", qFromExpirationTimestamp); err != nil {
				return err
			}
		}

	}

	if o.IncludeDeleted != nil {

		// query param include_deleted
		var qrIncludeDeleted string
		if o.IncludeDeleted != nil {
			qrIncludeDeleted = *o.IncludeDeleted
		}
		qIncludeDeleted := qrIncludeDeleted
		if qIncludeDeleted != "" {
			if err := r.SetQueryParam("include_deleted", qIncludeDeleted); err != nil {
				return err
			}
		}

	}

	if o.Policies != nil {

		// query param policies
		var qrPolicies string
		if o.Policies != nil {
			qrPolicies = *o.Policies
		}
		qPolicies := qrPolicies
		if qPolicies != "" {
			if err := r.SetQueryParam("policies", qPolicies); err != nil {
				return err
			}
		}

	}

	if o.ShareLevels != nil {

		// query param share_levels
		var qrShareLevels string
		if o.ShareLevels != nil {
			qrShareLevels = *o.ShareLevels
		}
		qShareLevels := qrShareLevels
		if qShareLevels != "" {
			if err := r.SetQueryParam("share_levels", qShareLevels); err != nil {
				return err
			}
		}

	}

	if o.Sources != nil {

		// query param sources
		var qrSources string
		if o.Sources != nil {
			qrSources = *o.Sources
		}
		qSources := qrSources
		if qSources != "" {
			if err := r.SetQueryParam("sources", qSources); err != nil {
				return err
			}
		}

	}

	if o.ToExpirationTimestamp != nil {

		// query param to.expiration_timestamp
		var qrToExpirationTimestamp string
		if o.ToExpirationTimestamp != nil {
			qrToExpirationTimestamp = *o.ToExpirationTimestamp
		}
		qToExpirationTimestamp := qrToExpirationTimestamp
		if qToExpirationTimestamp != "" {
			if err := r.SetQueryParam("to.expiration_timestamp", qToExpirationTimestamp); err != nil {
				return err
			}
		}

	}

	if o.Types != nil {

		// query param types
		var qrTypes string
		if o.Types != nil {
			qrTypes = *o.Types
		}
		qTypes := qrTypes
		if qTypes != "" {
			if err := r.SetQueryParam("types", qTypes); err != nil {
				return err
			}
		}

	}

	if o.Values != nil {

		// query param values
		var qrValues string
		if o.Values != nil {
			qrValues = *o.Values
		}
		qValues := qrValues
		if qValues != "" {
			if err := r.SetQueryParam("values", qValues); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
