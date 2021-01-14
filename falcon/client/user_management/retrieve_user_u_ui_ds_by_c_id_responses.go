// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// RetrieveUserUUIDsByCIDReader is a Reader for the RetrieveUserUUIDsByCID structure.
type RetrieveUserUUIDsByCIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveUserUUIDsByCIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveUserUUIDsByCIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveUserUUIDsByCIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRetrieveUserUUIDsByCIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRetrieveUserUUIDsByCIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRetrieveUserUUIDsByCIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRetrieveUserUUIDsByCIDOK creates a RetrieveUserUUIDsByCIDOK with default headers values
func NewRetrieveUserUUIDsByCIDOK() *RetrieveUserUUIDsByCIDOK {
	return &RetrieveUserUUIDsByCIDOK{}
}

/*RetrieveUserUUIDsByCIDOK handles this case with default header values.

OK
*/
type RetrieveUserUUIDsByCIDOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *RetrieveUserUUIDsByCIDOK) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdOK  %+v", 200, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserUUIDsByCIDBadRequest creates a RetrieveUserUUIDsByCIDBadRequest with default headers values
func NewRetrieveUserUUIDsByCIDBadRequest() *RetrieveUserUUIDsByCIDBadRequest {
	return &RetrieveUserUUIDsByCIDBadRequest{}
}

/*RetrieveUserUUIDsByCIDBadRequest handles this case with default header values.

Bad Request
*/
type RetrieveUserUUIDsByCIDBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *RetrieveUserUUIDsByCIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdBadRequest  %+v", 400, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserUUIDsByCIDForbidden creates a RetrieveUserUUIDsByCIDForbidden with default headers values
func NewRetrieveUserUUIDsByCIDForbidden() *RetrieveUserUUIDsByCIDForbidden {
	return &RetrieveUserUUIDsByCIDForbidden{}
}

/*RetrieveUserUUIDsByCIDForbidden handles this case with default header values.

Forbidden
*/
type RetrieveUserUUIDsByCIDForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *RetrieveUserUUIDsByCIDForbidden) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdForbidden  %+v", 403, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDForbidden) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserUUIDsByCIDTooManyRequests creates a RetrieveUserUUIDsByCIDTooManyRequests with default headers values
func NewRetrieveUserUUIDsByCIDTooManyRequests() *RetrieveUserUUIDsByCIDTooManyRequests {
	return &RetrieveUserUUIDsByCIDTooManyRequests{}
}

/*RetrieveUserUUIDsByCIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type RetrieveUserUUIDsByCIDTooManyRequests struct {
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

func (o *RetrieveUserUUIDsByCIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] retrieveUserUUiDsByCIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRetrieveUserUUIDsByCIDDefault creates a RetrieveUserUUIDsByCIDDefault with default headers values
func NewRetrieveUserUUIDsByCIDDefault(code int) *RetrieveUserUUIDsByCIDDefault {
	return &RetrieveUserUUIDsByCIDDefault{
		_statusCode: code,
	}
}

/*RetrieveUserUUIDsByCIDDefault handles this case with default header values.

OK
*/
type RetrieveUserUUIDsByCIDDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the retrieve user u UI ds by c ID default response
func (o *RetrieveUserUUIDsByCIDDefault) Code() int {
	return o._statusCode
}

func (o *RetrieveUserUUIDsByCIDDefault) Error() string {
	return fmt.Sprintf("[GET /users/queries/user-uuids-by-cid/v1][%d] RetrieveUserUUIDsByCID default  %+v", o._statusCode, o.Payload)
}

func (o *RetrieveUserUUIDsByCIDDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *RetrieveUserUUIDsByCIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
