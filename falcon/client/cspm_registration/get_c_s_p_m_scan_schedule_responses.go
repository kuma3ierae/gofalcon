// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// GetCSPMScanScheduleReader is a Reader for the GetCSPMScanSchedule structure.
type GetCSPMScanScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMScanScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMScanScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMScanScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMScanScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMScanScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMScanScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMScanScheduleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMScanScheduleOK creates a GetCSPMScanScheduleOK with default headers values
func NewGetCSPMScanScheduleOK() *GetCSPMScanScheduleOK {
	return &GetCSPMScanScheduleOK{}
}

/*GetCSPMScanScheduleOK handles this case with default header values.

OK
*/
type GetCSPMScanScheduleOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationScanScheduleResponseV1
}

func (o *GetCSPMScanScheduleOK) Error() string {
	return fmt.Sprintf("[GET /settings/scan-schedule/v1][%d] getCSPMScanScheduleOK  %+v", 200, o.Payload)
}

func (o *GetCSPMScanScheduleOK) GetPayload() *models.RegistrationScanScheduleResponseV1 {
	return o.Payload
}

func (o *GetCSPMScanScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationScanScheduleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMScanScheduleBadRequest creates a GetCSPMScanScheduleBadRequest with default headers values
func NewGetCSPMScanScheduleBadRequest() *GetCSPMScanScheduleBadRequest {
	return &GetCSPMScanScheduleBadRequest{}
}

/*GetCSPMScanScheduleBadRequest handles this case with default header values.

Bad Request
*/
type GetCSPMScanScheduleBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationScanScheduleResponseV1
}

func (o *GetCSPMScanScheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /settings/scan-schedule/v1][%d] getCSPMScanScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMScanScheduleBadRequest) GetPayload() *models.RegistrationScanScheduleResponseV1 {
	return o.Payload
}

func (o *GetCSPMScanScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationScanScheduleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMScanScheduleForbidden creates a GetCSPMScanScheduleForbidden with default headers values
func NewGetCSPMScanScheduleForbidden() *GetCSPMScanScheduleForbidden {
	return &GetCSPMScanScheduleForbidden{}
}

/*GetCSPMScanScheduleForbidden handles this case with default header values.

Forbidden
*/
type GetCSPMScanScheduleForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetCSPMScanScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /settings/scan-schedule/v1][%d] getCSPMScanScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMScanScheduleForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMScanScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMScanScheduleTooManyRequests creates a GetCSPMScanScheduleTooManyRequests with default headers values
func NewGetCSPMScanScheduleTooManyRequests() *GetCSPMScanScheduleTooManyRequests {
	return &GetCSPMScanScheduleTooManyRequests{}
}

/*GetCSPMScanScheduleTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetCSPMScanScheduleTooManyRequests struct {
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

func (o *GetCSPMScanScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /settings/scan-schedule/v1][%d] getCSPMScanScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMScanScheduleTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMScanScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMScanScheduleInternalServerError creates a GetCSPMScanScheduleInternalServerError with default headers values
func NewGetCSPMScanScheduleInternalServerError() *GetCSPMScanScheduleInternalServerError {
	return &GetCSPMScanScheduleInternalServerError{}
}

/*GetCSPMScanScheduleInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetCSPMScanScheduleInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationScanScheduleResponseV1
}

func (o *GetCSPMScanScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /settings/scan-schedule/v1][%d] getCSPMScanScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMScanScheduleInternalServerError) GetPayload() *models.RegistrationScanScheduleResponseV1 {
	return o.Payload
}

func (o *GetCSPMScanScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationScanScheduleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMScanScheduleDefault creates a GetCSPMScanScheduleDefault with default headers values
func NewGetCSPMScanScheduleDefault(code int) *GetCSPMScanScheduleDefault {
	return &GetCSPMScanScheduleDefault{
		_statusCode: code,
	}
}

/*GetCSPMScanScheduleDefault handles this case with default header values.

OK
*/
type GetCSPMScanScheduleDefault struct {
	_statusCode int

	Payload *models.RegistrationScanScheduleResponseV1
}

// Code gets the status code for the get c s p m scan schedule default response
func (o *GetCSPMScanScheduleDefault) Code() int {
	return o._statusCode
}

func (o *GetCSPMScanScheduleDefault) Error() string {
	return fmt.Sprintf("[GET /settings/scan-schedule/v1][%d] GetCSPMScanSchedule default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMScanScheduleDefault) GetPayload() *models.RegistrationScanScheduleResponseV1 {
	return o.Payload
}

func (o *GetCSPMScanScheduleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationScanScheduleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
