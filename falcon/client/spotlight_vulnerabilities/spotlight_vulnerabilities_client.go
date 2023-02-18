// Code generated by go-swagger; DO NOT EDIT.

package spotlight_vulnerabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new spotlight vulnerabilities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for spotlight vulnerabilities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CombinedQueryVulnerabilities(params *CombinedQueryVulnerabilitiesParams, opts ...ClientOption) (*CombinedQueryVulnerabilitiesOK, error)

	GetRemediationsV2(params *GetRemediationsV2Params, opts ...ClientOption) (*GetRemediationsV2OK, error)

	GetVulnerabilities(params *GetVulnerabilitiesParams, opts ...ClientOption) (*GetVulnerabilitiesOK, error)

	QueryVulnerabilities(params *QueryVulnerabilitiesParams, opts ...ClientOption) (*QueryVulnerabilitiesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CombinedQueryVulnerabilities searches for vulnerabilities in your environment by providing an f q l filter and paging details returns a set of vulnerability entities which match the filter criteria
*/
func (a *Client) CombinedQueryVulnerabilities(params *CombinedQueryVulnerabilitiesParams, opts ...ClientOption) (*CombinedQueryVulnerabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCombinedQueryVulnerabilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "combinedQueryVulnerabilities",
		Method:             "GET",
		PathPattern:        "/spotlight/combined/vulnerabilities/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CombinedQueryVulnerabilitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CombinedQueryVulnerabilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for combinedQueryVulnerabilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRemediationsV2 gets details on remediation by providing one or more i ds
*/
func (a *Client) GetRemediationsV2(params *GetRemediationsV2Params, opts ...ClientOption) (*GetRemediationsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRemediationsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRemediationsV2",
		Method:             "GET",
		PathPattern:        "/spotlight/entities/remediations/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRemediationsV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRemediationsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRemediationsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVulnerabilities gets details on vulnerabilities by providing one or more i ds
*/
func (a *Client) GetVulnerabilities(params *GetVulnerabilitiesParams, opts ...ClientOption) (*GetVulnerabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVulnerabilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVulnerabilities",
		Method:             "GET",
		PathPattern:        "/spotlight/entities/vulnerabilities/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVulnerabilitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVulnerabilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVulnerabilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryVulnerabilities searches for vulnerabilities in your environment by providing an f q l filter and paging details returns a set of vulnerability i ds which match the filter criteria
*/
func (a *Client) QueryVulnerabilities(params *QueryVulnerabilitiesParams, opts ...ClientOption) (*QueryVulnerabilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryVulnerabilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryVulnerabilities",
		Method:             "GET",
		PathPattern:        "/spotlight/queries/vulnerabilities/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryVulnerabilitiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryVulnerabilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryVulnerabilitiesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
