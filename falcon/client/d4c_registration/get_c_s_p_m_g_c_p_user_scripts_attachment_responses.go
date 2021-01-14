// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// GetCSPMGCPUserScriptsAttachmentReader is a Reader for the GetCSPMGCPUserScriptsAttachment structure.
type GetCSPMGCPUserScriptsAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMGCPUserScriptsAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMGCPUserScriptsAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMGCPUserScriptsAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMGCPUserScriptsAttachmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMGCPUserScriptsAttachmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMGCPUserScriptsAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCSPMGCPUserScriptsAttachmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCSPMGCPUserScriptsAttachmentOK creates a GetCSPMGCPUserScriptsAttachmentOK with default headers values
func NewGetCSPMGCPUserScriptsAttachmentOK() *GetCSPMGCPUserScriptsAttachmentOK {
	return &GetCSPMGCPUserScriptsAttachmentOK{}
}

/*GetCSPMGCPUserScriptsAttachmentOK handles this case with default header values.

OK
*/
type GetCSPMGCPUserScriptsAttachmentOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

func (o *GetCSPMGCPUserScriptsAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentOK) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPUserScriptsAttachmentBadRequest creates a GetCSPMGCPUserScriptsAttachmentBadRequest with default headers values
func NewGetCSPMGCPUserScriptsAttachmentBadRequest() *GetCSPMGCPUserScriptsAttachmentBadRequest {
	return &GetCSPMGCPUserScriptsAttachmentBadRequest{}
}

/*GetCSPMGCPUserScriptsAttachmentBadRequest handles this case with default header values.

Bad Request
*/
type GetCSPMGCPUserScriptsAttachmentBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPUserScriptsAttachmentForbidden creates a GetCSPMGCPUserScriptsAttachmentForbidden with default headers values
func NewGetCSPMGCPUserScriptsAttachmentForbidden() *GetCSPMGCPUserScriptsAttachmentForbidden {
	return &GetCSPMGCPUserScriptsAttachmentForbidden{}
}

/*GetCSPMGCPUserScriptsAttachmentForbidden handles this case with default header values.

Forbidden
*/
type GetCSPMGCPUserScriptsAttachmentForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetCSPMGCPUserScriptsAttachmentForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMGCPUserScriptsAttachmentTooManyRequests creates a GetCSPMGCPUserScriptsAttachmentTooManyRequests with default headers values
func NewGetCSPMGCPUserScriptsAttachmentTooManyRequests() *GetCSPMGCPUserScriptsAttachmentTooManyRequests {
	return &GetCSPMGCPUserScriptsAttachmentTooManyRequests{}
}

/*GetCSPMGCPUserScriptsAttachmentTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetCSPMGCPUserScriptsAttachmentTooManyRequests struct {
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

func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMGCPUserScriptsAttachmentInternalServerError creates a GetCSPMGCPUserScriptsAttachmentInternalServerError with default headers values
func NewGetCSPMGCPUserScriptsAttachmentInternalServerError() *GetCSPMGCPUserScriptsAttachmentInternalServerError {
	return &GetCSPMGCPUserScriptsAttachmentInternalServerError{}
}

/*GetCSPMGCPUserScriptsAttachmentInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetCSPMGCPUserScriptsAttachmentInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] getCSPMGCPUserScriptsAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMGCPUserScriptsAttachmentDefault creates a GetCSPMGCPUserScriptsAttachmentDefault with default headers values
func NewGetCSPMGCPUserScriptsAttachmentDefault(code int) *GetCSPMGCPUserScriptsAttachmentDefault {
	return &GetCSPMGCPUserScriptsAttachmentDefault{
		_statusCode: code,
	}
}

/*GetCSPMGCPUserScriptsAttachmentDefault handles this case with default header values.

OK
*/
type GetCSPMGCPUserScriptsAttachmentDefault struct {
	_statusCode int

	Payload *models.RegistrationGCPProvisionGetUserScriptResponseV1
}

// Code gets the status code for the get c s p m g c p user scripts attachment default response
func (o *GetCSPMGCPUserScriptsAttachmentDefault) Code() int {
	return o._statusCode
}

func (o *GetCSPMGCPUserScriptsAttachmentDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-gcp/entities/user-scripts-download/v1][%d] GetCSPMGCPUserScriptsAttachment default  %+v", o._statusCode, o.Payload)
}

func (o *GetCSPMGCPUserScriptsAttachmentDefault) GetPayload() *models.RegistrationGCPProvisionGetUserScriptResponseV1 {
	return o.Payload
}

func (o *GetCSPMGCPUserScriptsAttachmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationGCPProvisionGetUserScriptResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
