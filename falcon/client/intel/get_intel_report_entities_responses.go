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

// GetIntelReportEntitiesReader is a Reader for the GetIntelReportEntities structure.
type GetIntelReportEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntelReportEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntelReportEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetIntelReportEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntelReportEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntelReportEntitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIntelReportEntitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIntelReportEntitiesOK creates a GetIntelReportEntitiesOK with default headers values
func NewGetIntelReportEntitiesOK() *GetIntelReportEntitiesOK {
	return &GetIntelReportEntitiesOK{}
}

/*GetIntelReportEntitiesOK handles this case with default header values.

OK
*/
type GetIntelReportEntitiesOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainNewsResponse
}

func (o *GetIntelReportEntitiesOK) Error() string {
	return fmt.Sprintf("[GET /intel/entities/reports/v1][%d] getIntelReportEntitiesOK  %+v", 200, o.Payload)
}

func (o *GetIntelReportEntitiesOK) GetPayload() *models.DomainNewsResponse {
	return o.Payload
}

func (o *GetIntelReportEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainNewsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntelReportEntitiesForbidden creates a GetIntelReportEntitiesForbidden with default headers values
func NewGetIntelReportEntitiesForbidden() *GetIntelReportEntitiesForbidden {
	return &GetIntelReportEntitiesForbidden{}
}

/*GetIntelReportEntitiesForbidden handles this case with default header values.

Forbidden
*/
type GetIntelReportEntitiesForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetIntelReportEntitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/entities/reports/v1][%d] getIntelReportEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntelReportEntitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntelReportEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelReportEntitiesTooManyRequests creates a GetIntelReportEntitiesTooManyRequests with default headers values
func NewGetIntelReportEntitiesTooManyRequests() *GetIntelReportEntitiesTooManyRequests {
	return &GetIntelReportEntitiesTooManyRequests{}
}

/*GetIntelReportEntitiesTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetIntelReportEntitiesTooManyRequests struct {
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

func (o *GetIntelReportEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/entities/reports/v1][%d] getIntelReportEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntelReportEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntelReportEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelReportEntitiesInternalServerError creates a GetIntelReportEntitiesInternalServerError with default headers values
func NewGetIntelReportEntitiesInternalServerError() *GetIntelReportEntitiesInternalServerError {
	return &GetIntelReportEntitiesInternalServerError{}
}

/*GetIntelReportEntitiesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetIntelReportEntitiesInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetIntelReportEntitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/entities/reports/v1][%d] getIntelReportEntitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntelReportEntitiesInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetIntelReportEntitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntelReportEntitiesDefault creates a GetIntelReportEntitiesDefault with default headers values
func NewGetIntelReportEntitiesDefault(code int) *GetIntelReportEntitiesDefault {
	return &GetIntelReportEntitiesDefault{
		_statusCode: code,
	}
}

/*GetIntelReportEntitiesDefault handles this case with default header values.

OK
*/
type GetIntelReportEntitiesDefault struct {
	_statusCode int

	Payload *models.DomainNewsResponse
}

// Code gets the status code for the get intel report entities default response
func (o *GetIntelReportEntitiesDefault) Code() int {
	return o._statusCode
}

func (o *GetIntelReportEntitiesDefault) Error() string {
	return fmt.Sprintf("[GET /intel/entities/reports/v1][%d] GetIntelReportEntities default  %+v", o._statusCode, o.Payload)
}

func (o *GetIntelReportEntitiesDefault) GetPayload() *models.DomainNewsResponse {
	return o.Payload
}

func (o *GetIntelReportEntitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainNewsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
