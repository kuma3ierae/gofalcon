// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// QueryIntelIndicatorEntitiesReader is a Reader for the QueryIntelIndicatorEntities structure.
type QueryIntelIndicatorEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryIntelIndicatorEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryIntelIndicatorEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryIntelIndicatorEntitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryIntelIndicatorEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryIntelIndicatorEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryIntelIndicatorEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryIntelIndicatorEntitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryIntelIndicatorEntitiesOK creates a QueryIntelIndicatorEntitiesOK with default headers values
func NewQueryIntelIndicatorEntitiesOK() *QueryIntelIndicatorEntitiesOK {
	return &QueryIntelIndicatorEntitiesOK{}
}

/*QueryIntelIndicatorEntitiesOK handles this case with default header values.

OK
*/
type QueryIntelIndicatorEntitiesOK struct {
	/*Provides next page pagination URL. Available only if sorting was done using using _marker field, which is the default one.
	 */
	NextPage string
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainPublicIndicatorsV3Response
}

func (o *QueryIntelIndicatorEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /intel/combined/indicators/v1][%d] queryIntelIndicatorEntitiesOK  %+v", 200, o.Payload)
}

func (o *QueryIntelIndicatorEntitiesOK) GetPayload() *models.DomainPublicIndicatorsV3Response {
	return o.Payload
}

func (o *QueryIntelIndicatorEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Next-Page
	o.NextPage = response.GetHeader("Next-Page")

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.DomainPublicIndicatorsV3Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelIndicatorEntitiesBadRequest creates a QueryIntelIndicatorEntitiesBadRequest with default headers values
func NewQueryIntelIndicatorEntitiesBadRequest() *QueryIntelIndicatorEntitiesBadRequest {
	return &QueryIntelIndicatorEntitiesBadRequest{}
}

/*QueryIntelIndicatorEntitiesBadRequest handles this case with default header values.

Bad Request
*/
type QueryIntelIndicatorEntitiesBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryIntelIndicatorEntitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /intel/combined/indicators/v1][%d] queryIntelIndicatorEntitiesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryIntelIndicatorEntitiesBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryIntelIndicatorEntitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelIndicatorEntitiesForbidden creates a QueryIntelIndicatorEntitiesForbidden with default headers values
func NewQueryIntelIndicatorEntitiesForbidden() *QueryIntelIndicatorEntitiesForbidden {
	return &QueryIntelIndicatorEntitiesForbidden{}
}

/*QueryIntelIndicatorEntitiesForbidden handles this case with default header values.

Forbidden
*/
type QueryIntelIndicatorEntitiesForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryIntelIndicatorEntitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/combined/indicators/v1][%d] queryIntelIndicatorEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *QueryIntelIndicatorEntitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIntelIndicatorEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelIndicatorEntitiesTooManyRequests creates a QueryIntelIndicatorEntitiesTooManyRequests with default headers values
func NewQueryIntelIndicatorEntitiesTooManyRequests() *QueryIntelIndicatorEntitiesTooManyRequests {
	return &QueryIntelIndicatorEntitiesTooManyRequests{}
}

/*QueryIntelIndicatorEntitiesTooManyRequests handles this case with default header values.

Too Many Requests
*/
type QueryIntelIndicatorEntitiesTooManyRequests struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
	/*Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryIntelIndicatorEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/combined/indicators/v1][%d] queryIntelIndicatorEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryIntelIndicatorEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIntelIndicatorEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	// response header X-RateLimit-RetryAfter
	xRateLimitRetryAfter, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-RetryAfter"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", response.GetHeader("X-RateLimit-RetryAfter"))
	}
	o.XRateLimitRetryAfter = xRateLimitRetryAfter

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelIndicatorEntitiesInternalServerError creates a QueryIntelIndicatorEntitiesInternalServerError with default headers values
func NewQueryIntelIndicatorEntitiesInternalServerError() *QueryIntelIndicatorEntitiesInternalServerError {
	return &QueryIntelIndicatorEntitiesInternalServerError{}
}

/*QueryIntelIndicatorEntitiesInternalServerError handles this case with default header values.

Internal Server Error
*/
type QueryIntelIndicatorEntitiesInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryIntelIndicatorEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/combined/indicators/v1][%d] queryIntelIndicatorEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryIntelIndicatorEntitiesInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryIntelIndicatorEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelIndicatorEntitiesDefault creates a QueryIntelIndicatorEntitiesDefault with default headers values
func NewQueryIntelIndicatorEntitiesDefault(code int) *QueryIntelIndicatorEntitiesDefault {
	return &QueryIntelIndicatorEntitiesDefault{
		_statusCode: code,
	}
}

/*QueryIntelIndicatorEntitiesDefault handles this case with default header values.

OK
*/
type QueryIntelIndicatorEntitiesDefault struct {
	_statusCode int

	/*Provides next page pagination URL. Available only if sorting was done using using _marker field, which is the default one.
	 */
	NextPage string

	Payload *models.DomainPublicIndicatorsV3Response
}

// Code gets the status code for the query intel indicator entities default response
func (o *QueryIntelIndicatorEntitiesDefault) Code() int {
	return o._statusCode
}

func (o *QueryIntelIndicatorEntitiesDefault) Error() string {
	return fmt.Sprintf("[GET /intel/combined/indicators/v1][%d] QueryIntelIndicatorEntities default  %+v", o._statusCode, o.Payload)
}

func (o *QueryIntelIndicatorEntitiesDefault) GetPayload() *models.DomainPublicIndicatorsV3Response {
	return o.Payload
}

func (o *QueryIntelIndicatorEntitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Next-Page
	o.NextPage = response.GetHeader("Next-Page")

	o.Payload = new(models.DomainPublicIndicatorsV3Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
