// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// AppActionsReader is a Reader for the AppActions structure.
type AppActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppActionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppActionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAppActionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAppActionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAppActionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppActionsOK creates a AppActionsOK with default headers values
func NewAppActionsOK() *AppActionsOK {
	return &AppActionsOK{}
}

/*
AppActionsOK describes a response with status code 200, with default header values.

OK
*/
type AppActionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.StoreDomainAppDetailsResponse
}

// IsSuccess returns true when this app actions o k response has a 2xx status code
func (o *AppActionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this app actions o k response has a 3xx status code
func (o *AppActionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app actions o k response has a 4xx status code
func (o *AppActionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this app actions o k response has a 5xx status code
func (o *AppActionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this app actions o k response a status code equal to that given
func (o *AppActionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *AppActionsOK) Error() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] appActionsOK  %+v", 200, o.Payload)
}

func (o *AppActionsOK) String() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] appActionsOK  %+v", 200, o.Payload)
}

func (o *AppActionsOK) GetPayload() *models.StoreDomainAppDetailsResponse {
	return o.Payload
}

func (o *AppActionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.StoreDomainAppDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppActionsBadRequest creates a AppActionsBadRequest with default headers values
func NewAppActionsBadRequest() *AppActionsBadRequest {
	return &AppActionsBadRequest{}
}

/*
AppActionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AppActionsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.StoreDomainAppDetailsResponse
}

// IsSuccess returns true when this app actions bad request response has a 2xx status code
func (o *AppActionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app actions bad request response has a 3xx status code
func (o *AppActionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app actions bad request response has a 4xx status code
func (o *AppActionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this app actions bad request response has a 5xx status code
func (o *AppActionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this app actions bad request response a status code equal to that given
func (o *AppActionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AppActionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] appActionsBadRequest  %+v", 400, o.Payload)
}

func (o *AppActionsBadRequest) String() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] appActionsBadRequest  %+v", 400, o.Payload)
}

func (o *AppActionsBadRequest) GetPayload() *models.StoreDomainAppDetailsResponse {
	return o.Payload
}

func (o *AppActionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.StoreDomainAppDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppActionsForbidden creates a AppActionsForbidden with default headers values
func NewAppActionsForbidden() *AppActionsForbidden {
	return &AppActionsForbidden{}
}

/*
AppActionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AppActionsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this app actions forbidden response has a 2xx status code
func (o *AppActionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app actions forbidden response has a 3xx status code
func (o *AppActionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app actions forbidden response has a 4xx status code
func (o *AppActionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this app actions forbidden response has a 5xx status code
func (o *AppActionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this app actions forbidden response a status code equal to that given
func (o *AppActionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AppActionsForbidden) Error() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] appActionsForbidden  %+v", 403, o.Payload)
}

func (o *AppActionsForbidden) String() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] appActionsForbidden  %+v", 403, o.Payload)
}

func (o *AppActionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AppActionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppActionsTooManyRequests creates a AppActionsTooManyRequests with default headers values
func NewAppActionsTooManyRequests() *AppActionsTooManyRequests {
	return &AppActionsTooManyRequests{}
}

/*
AppActionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AppActionsTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this app actions too many requests response has a 2xx status code
func (o *AppActionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this app actions too many requests response has a 3xx status code
func (o *AppActionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this app actions too many requests response has a 4xx status code
func (o *AppActionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this app actions too many requests response has a 5xx status code
func (o *AppActionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this app actions too many requests response a status code equal to that given
func (o *AppActionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *AppActionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] appActionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AppActionsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] appActionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AppActionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AppActionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppActionsDefault creates a AppActionsDefault with default headers values
func NewAppActionsDefault(code int) *AppActionsDefault {
	return &AppActionsDefault{
		_statusCode: code,
	}
}

/*
AppActionsDefault describes a response with status code -1, with default header values.

OK
*/
type AppActionsDefault struct {
	_statusCode int

	Payload *models.StoreDomainAppDetailsResponse
}

// Code gets the status code for the app actions default response
func (o *AppActionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this app actions default response has a 2xx status code
func (o *AppActionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this app actions default response has a 3xx status code
func (o *AppActionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this app actions default response has a 4xx status code
func (o *AppActionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this app actions default response has a 5xx status code
func (o *AppActionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this app actions default response a status code equal to that given
func (o *AppActionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AppActionsDefault) Error() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] AppActions default  %+v", o._statusCode, o.Payload)
}

func (o *AppActionsDefault) String() string {
	return fmt.Sprintf("[POST /store/entities/app-actions/v1][%d] AppActions default  %+v", o._statusCode, o.Payload)
}

func (o *AppActionsDefault) GetPayload() *models.StoreDomainAppDetailsResponse {
	return o.Payload
}

func (o *AppActionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoreDomainAppDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
