// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

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

// UploadSampleV3Reader is a Reader for the UploadSampleV3 structure.
type UploadSampleV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadSampleV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadSampleV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadSampleV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadSampleV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUploadSampleV3TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadSampleV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUploadSampleV3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadSampleV3OK creates a UploadSampleV3OK with default headers values
func NewUploadSampleV3OK() *UploadSampleV3OK {
	return &UploadSampleV3OK{}
}

/*UploadSampleV3OK handles this case with default header values.

OK
*/
type UploadSampleV3OK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SamplestoreSampleMetadataResponseV2
}

func (o *UploadSampleV3OK) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3OK  %+v", 200, o.Payload)
}

func (o *UploadSampleV3OK) GetPayload() *models.SamplestoreSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SamplestoreSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSampleV3BadRequest creates a UploadSampleV3BadRequest with default headers values
func NewUploadSampleV3BadRequest() *UploadSampleV3BadRequest {
	return &UploadSampleV3BadRequest{}
}

/*UploadSampleV3BadRequest handles this case with default header values.

Bad Request
*/
type UploadSampleV3BadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SamplestoreSampleMetadataResponseV2
}

func (o *UploadSampleV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3BadRequest  %+v", 400, o.Payload)
}

func (o *UploadSampleV3BadRequest) GetPayload() *models.SamplestoreSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SamplestoreSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSampleV3Forbidden creates a UploadSampleV3Forbidden with default headers values
func NewUploadSampleV3Forbidden() *UploadSampleV3Forbidden {
	return &UploadSampleV3Forbidden{}
}

/*UploadSampleV3Forbidden handles this case with default header values.

Forbidden
*/
type UploadSampleV3Forbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UploadSampleV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3Forbidden  %+v", 403, o.Payload)
}

func (o *UploadSampleV3Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadSampleV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUploadSampleV3TooManyRequests creates a UploadSampleV3TooManyRequests with default headers values
func NewUploadSampleV3TooManyRequests() *UploadSampleV3TooManyRequests {
	return &UploadSampleV3TooManyRequests{}
}

/*UploadSampleV3TooManyRequests handles this case with default header values.

Too Many Requests
*/
type UploadSampleV3TooManyRequests struct {
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

func (o *UploadSampleV3TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3TooManyRequests  %+v", 429, o.Payload)
}

func (o *UploadSampleV3TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadSampleV3TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUploadSampleV3InternalServerError creates a UploadSampleV3InternalServerError with default headers values
func NewUploadSampleV3InternalServerError() *UploadSampleV3InternalServerError {
	return &UploadSampleV3InternalServerError{}
}

/*UploadSampleV3InternalServerError handles this case with default header values.

Internal Server Error
*/
type UploadSampleV3InternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SamplestoreSampleMetadataResponseV2
}

func (o *UploadSampleV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3InternalServerError  %+v", 500, o.Payload)
}

func (o *UploadSampleV3InternalServerError) GetPayload() *models.SamplestoreSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SamplestoreSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSampleV3Default creates a UploadSampleV3Default with default headers values
func NewUploadSampleV3Default(code int) *UploadSampleV3Default {
	return &UploadSampleV3Default{
		_statusCode: code,
	}
}

/*UploadSampleV3Default handles this case with default header values.

OK
*/
type UploadSampleV3Default struct {
	_statusCode int

	Payload *models.SamplestoreSampleMetadataResponseV2
}

// Code gets the status code for the upload sample v3 default response
func (o *UploadSampleV3Default) Code() int {
	return o._statusCode
}

func (o *UploadSampleV3Default) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] UploadSampleV3 default  %+v", o._statusCode, o.Payload)
}

func (o *UploadSampleV3Default) GetPayload() *models.SamplestoreSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SamplestoreSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
