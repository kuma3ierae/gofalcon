// Code generated by go-swagger; DO NOT EDIT.

package malquery

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

// GetMalQueryRequestV1Reader is a Reader for the GetMalQueryRequestV1 structure.
type GetMalQueryRequestV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMalQueryRequestV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMalQueryRequestV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMalQueryRequestV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMalQueryRequestV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMalQueryRequestV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMalQueryRequestV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMalQueryRequestV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMalQueryRequestV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMalQueryRequestV1OK creates a GetMalQueryRequestV1OK with default headers values
func NewGetMalQueryRequestV1OK() *GetMalQueryRequestV1OK {
	return &GetMalQueryRequestV1OK{}
}

/*GetMalQueryRequestV1OK handles this case with default header values.

OK
*/
type GetMalQueryRequestV1OK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRequestResponse
}

func (o *GetMalQueryRequestV1OK) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1OK  %+v", 200, o.Payload)
}

func (o *GetMalQueryRequestV1OK) GetPayload() *models.MalqueryRequestResponse {
	return o.Payload
}

func (o *GetMalQueryRequestV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryRequestV1BadRequest creates a GetMalQueryRequestV1BadRequest with default headers values
func NewGetMalQueryRequestV1BadRequest() *GetMalQueryRequestV1BadRequest {
	return &GetMalQueryRequestV1BadRequest{}
}

/*GetMalQueryRequestV1BadRequest handles this case with default header values.

Bad Request
*/
type GetMalQueryRequestV1BadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRequestResponse
}

func (o *GetMalQueryRequestV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMalQueryRequestV1BadRequest) GetPayload() *models.MalqueryRequestResponse {
	return o.Payload
}

func (o *GetMalQueryRequestV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryRequestV1Unauthorized creates a GetMalQueryRequestV1Unauthorized with default headers values
func NewGetMalQueryRequestV1Unauthorized() *GetMalQueryRequestV1Unauthorized {
	return &GetMalQueryRequestV1Unauthorized{}
}

/*GetMalQueryRequestV1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetMalQueryRequestV1Unauthorized struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetMalQueryRequestV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetMalQueryRequestV1Unauthorized) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryRequestV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryRequestV1Forbidden creates a GetMalQueryRequestV1Forbidden with default headers values
func NewGetMalQueryRequestV1Forbidden() *GetMalQueryRequestV1Forbidden {
	return &GetMalQueryRequestV1Forbidden{}
}

/*GetMalQueryRequestV1Forbidden handles this case with default header values.

Forbidden
*/
type GetMalQueryRequestV1Forbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetMalQueryRequestV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMalQueryRequestV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryRequestV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryRequestV1TooManyRequests creates a GetMalQueryRequestV1TooManyRequests with default headers values
func NewGetMalQueryRequestV1TooManyRequests() *GetMalQueryRequestV1TooManyRequests {
	return &GetMalQueryRequestV1TooManyRequests{}
}

/*GetMalQueryRequestV1TooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetMalQueryRequestV1TooManyRequests struct {
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

func (o *GetMalQueryRequestV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMalQueryRequestV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMalQueryRequestV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryRequestV1InternalServerError creates a GetMalQueryRequestV1InternalServerError with default headers values
func NewGetMalQueryRequestV1InternalServerError() *GetMalQueryRequestV1InternalServerError {
	return &GetMalQueryRequestV1InternalServerError{}
}

/*GetMalQueryRequestV1InternalServerError handles this case with default header values.

Internal Server Error
*/
type GetMalQueryRequestV1InternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRequestResponse
}

func (o *GetMalQueryRequestV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] getMalQueryRequestV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMalQueryRequestV1InternalServerError) GetPayload() *models.MalqueryRequestResponse {
	return o.Payload
}

func (o *GetMalQueryRequestV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryRequestV1Default creates a GetMalQueryRequestV1Default with default headers values
func NewGetMalQueryRequestV1Default(code int) *GetMalQueryRequestV1Default {
	return &GetMalQueryRequestV1Default{
		_statusCode: code,
	}
}

/*GetMalQueryRequestV1Default handles this case with default header values.

OK
*/
type GetMalQueryRequestV1Default struct {
	_statusCode int

	Payload *models.MalqueryRequestResponse
}

// Code gets the status code for the get mal query request v1 default response
func (o *GetMalQueryRequestV1Default) Code() int {
	return o._statusCode
}

func (o *GetMalQueryRequestV1Default) Error() string {
	return fmt.Sprintf("[GET /malquery/entities/requests/v1][%d] GetMalQueryRequestV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetMalQueryRequestV1Default) GetPayload() *models.MalqueryRequestResponse {
	return o.Payload
}

func (o *GetMalQueryRequestV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MalqueryRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
