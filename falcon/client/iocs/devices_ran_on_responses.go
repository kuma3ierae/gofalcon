// Code generated by go-swagger; DO NOT EDIT.

package iocs

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

// DevicesRanOnReader is a Reader for the DevicesRanOn structure.
type DevicesRanOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesRanOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesRanOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDevicesRanOnForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDevicesRanOnTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDevicesRanOnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDevicesRanOnOK creates a DevicesRanOnOK with default headers values
func NewDevicesRanOnOK() *DevicesRanOnOK {
	return &DevicesRanOnOK{}
}

/*DevicesRanOnOK handles this case with default header values.

OK
*/
type DevicesRanOnOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIMsaReplyDevicesRanOn
}

func (o *DevicesRanOnOK) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnOK  %+v", 200, o.Payload)
}

func (o *DevicesRanOnOK) GetPayload() *models.APIMsaReplyDevicesRanOn {
	return o.Payload
}

func (o *DevicesRanOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIMsaReplyDevicesRanOn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDevicesRanOnForbidden creates a DevicesRanOnForbidden with default headers values
func NewDevicesRanOnForbidden() *DevicesRanOnForbidden {
	return &DevicesRanOnForbidden{}
}

/*DevicesRanOnForbidden handles this case with default header values.

Forbidden
*/
type DevicesRanOnForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DevicesRanOnForbidden) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnForbidden  %+v", 403, o.Payload)
}

func (o *DevicesRanOnForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DevicesRanOnForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDevicesRanOnTooManyRequests creates a DevicesRanOnTooManyRequests with default headers values
func NewDevicesRanOnTooManyRequests() *DevicesRanOnTooManyRequests {
	return &DevicesRanOnTooManyRequests{}
}

/*DevicesRanOnTooManyRequests handles this case with default header values.

Too Many Requests
*/
type DevicesRanOnTooManyRequests struct {
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

func (o *DevicesRanOnTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] devicesRanOnTooManyRequests  %+v", 429, o.Payload)
}

func (o *DevicesRanOnTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DevicesRanOnTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDevicesRanOnDefault creates a DevicesRanOnDefault with default headers values
func NewDevicesRanOnDefault(code int) *DevicesRanOnDefault {
	return &DevicesRanOnDefault{
		_statusCode: code,
	}
}

/*DevicesRanOnDefault handles this case with default header values.

OK
*/
type DevicesRanOnDefault struct {
	_statusCode int

	Payload *models.APIMsaReplyDevicesRanOn
}

// Code gets the status code for the devices ran on default response
func (o *DevicesRanOnDefault) Code() int {
	return o._statusCode
}

func (o *DevicesRanOnDefault) Error() string {
	return fmt.Sprintf("[GET /indicators/queries/devices/v1][%d] DevicesRanOn default  %+v", o._statusCode, o.Payload)
}

func (o *DevicesRanOnDefault) GetPayload() *models.APIMsaReplyDevicesRanOn {
	return o.Payload
}

func (o *DevicesRanOnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMsaReplyDevicesRanOn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
