// Code generated by go-swagger; DO NOT EDIT.

package ml_exclusions

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

// UpdateMLExclusionsV1Reader is a Reader for the UpdateMLExclusionsV1 structure.
type UpdateMLExclusionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMLExclusionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMLExclusionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMLExclusionsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMLExclusionsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateMLExclusionsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMLExclusionsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateMLExclusionsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMLExclusionsV1OK creates a UpdateMLExclusionsV1OK with default headers values
func NewUpdateMLExclusionsV1OK() *UpdateMLExclusionsV1OK {
	return &UpdateMLExclusionsV1OK{}
}

/*UpdateMLExclusionsV1OK handles this case with default header values.

OK
*/
type UpdateMLExclusionsV1OK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesMlExclusionRespV1
}

func (o *UpdateMLExclusionsV1OK) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ml-exclusions/v1][%d] updateMLExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *UpdateMLExclusionsV1OK) GetPayload() *models.ResponsesMlExclusionRespV1 {
	return o.Payload
}

func (o *UpdateMLExclusionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesMlExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMLExclusionsV1BadRequest creates a UpdateMLExclusionsV1BadRequest with default headers values
func NewUpdateMLExclusionsV1BadRequest() *UpdateMLExclusionsV1BadRequest {
	return &UpdateMLExclusionsV1BadRequest{}
}

/*UpdateMLExclusionsV1BadRequest handles this case with default header values.

Bad Request
*/
type UpdateMLExclusionsV1BadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesMlExclusionRespV1
}

func (o *UpdateMLExclusionsV1BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ml-exclusions/v1][%d] updateMLExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateMLExclusionsV1BadRequest) GetPayload() *models.ResponsesMlExclusionRespV1 {
	return o.Payload
}

func (o *UpdateMLExclusionsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesMlExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMLExclusionsV1Forbidden creates a UpdateMLExclusionsV1Forbidden with default headers values
func NewUpdateMLExclusionsV1Forbidden() *UpdateMLExclusionsV1Forbidden {
	return &UpdateMLExclusionsV1Forbidden{}
}

/*UpdateMLExclusionsV1Forbidden handles this case with default header values.

Forbidden
*/
type UpdateMLExclusionsV1Forbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *UpdateMLExclusionsV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ml-exclusions/v1][%d] updateMLExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateMLExclusionsV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateMLExclusionsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateMLExclusionsV1TooManyRequests creates a UpdateMLExclusionsV1TooManyRequests with default headers values
func NewUpdateMLExclusionsV1TooManyRequests() *UpdateMLExclusionsV1TooManyRequests {
	return &UpdateMLExclusionsV1TooManyRequests{}
}

/*UpdateMLExclusionsV1TooManyRequests handles this case with default header values.

Too Many Requests
*/
type UpdateMLExclusionsV1TooManyRequests struct {
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

func (o *UpdateMLExclusionsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ml-exclusions/v1][%d] updateMLExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateMLExclusionsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateMLExclusionsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateMLExclusionsV1InternalServerError creates a UpdateMLExclusionsV1InternalServerError with default headers values
func NewUpdateMLExclusionsV1InternalServerError() *UpdateMLExclusionsV1InternalServerError {
	return &UpdateMLExclusionsV1InternalServerError{}
}

/*UpdateMLExclusionsV1InternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateMLExclusionsV1InternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesMlExclusionRespV1
}

func (o *UpdateMLExclusionsV1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ml-exclusions/v1][%d] updateMLExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMLExclusionsV1InternalServerError) GetPayload() *models.ResponsesMlExclusionRespV1 {
	return o.Payload
}

func (o *UpdateMLExclusionsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesMlExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMLExclusionsV1Default creates a UpdateMLExclusionsV1Default with default headers values
func NewUpdateMLExclusionsV1Default(code int) *UpdateMLExclusionsV1Default {
	return &UpdateMLExclusionsV1Default{
		_statusCode: code,
	}
}

/*UpdateMLExclusionsV1Default handles this case with default header values.

OK
*/
type UpdateMLExclusionsV1Default struct {
	_statusCode int

	Payload *models.ResponsesMlExclusionRespV1
}

// Code gets the status code for the update m l exclusions v1 default response
func (o *UpdateMLExclusionsV1Default) Code() int {
	return o._statusCode
}

func (o *UpdateMLExclusionsV1Default) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/ml-exclusions/v1][%d] updateMLExclusionsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMLExclusionsV1Default) GetPayload() *models.ResponsesMlExclusionRespV1 {
	return o.Payload
}

func (o *UpdateMLExclusionsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesMlExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
