// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// RoutePutMuteTimingReader is a Reader for the RoutePutMuteTiming structure.
type RoutePutMuteTimingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoutePutMuteTimingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoutePutMuteTimingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoutePutMuteTimingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/v1/provisioning/mute-timings/{name}] RoutePutMuteTiming", response, response.Code())
	}
}

// NewRoutePutMuteTimingOK creates a RoutePutMuteTimingOK with default headers values
func NewRoutePutMuteTimingOK() *RoutePutMuteTimingOK {
	return &RoutePutMuteTimingOK{}
}

/*
RoutePutMuteTimingOK describes a response with status code 200, with default header values.

MuteTimeInterval
*/
type RoutePutMuteTimingOK struct {
	Payload *models.MuteTimeInterval
}

// IsSuccess returns true when this route put mute timing o k response has a 2xx status code
func (o *RoutePutMuteTimingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this route put mute timing o k response has a 3xx status code
func (o *RoutePutMuteTimingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route put mute timing o k response has a 4xx status code
func (o *RoutePutMuteTimingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this route put mute timing o k response has a 5xx status code
func (o *RoutePutMuteTimingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this route put mute timing o k response a status code equal to that given
func (o *RoutePutMuteTimingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the route put mute timing o k response
func (o *RoutePutMuteTimingOK) Code() int {
	return 200
}

func (o *RoutePutMuteTimingOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/mute-timings/{name}][%d] routePutMuteTimingOK  %+v", 200, o.Payload)
}

func (o *RoutePutMuteTimingOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/mute-timings/{name}][%d] routePutMuteTimingOK  %+v", 200, o.Payload)
}

func (o *RoutePutMuteTimingOK) GetPayload() *models.MuteTimeInterval {
	return o.Payload
}

func (o *RoutePutMuteTimingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MuteTimeInterval)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoutePutMuteTimingBadRequest creates a RoutePutMuteTimingBadRequest with default headers values
func NewRoutePutMuteTimingBadRequest() *RoutePutMuteTimingBadRequest {
	return &RoutePutMuteTimingBadRequest{}
}

/*
RoutePutMuteTimingBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RoutePutMuteTimingBadRequest struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this route put mute timing bad request response has a 2xx status code
func (o *RoutePutMuteTimingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this route put mute timing bad request response has a 3xx status code
func (o *RoutePutMuteTimingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route put mute timing bad request response has a 4xx status code
func (o *RoutePutMuteTimingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this route put mute timing bad request response has a 5xx status code
func (o *RoutePutMuteTimingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this route put mute timing bad request response a status code equal to that given
func (o *RoutePutMuteTimingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the route put mute timing bad request response
func (o *RoutePutMuteTimingBadRequest) Code() int {
	return 400
}

func (o *RoutePutMuteTimingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/mute-timings/{name}][%d] routePutMuteTimingBadRequest  %+v", 400, o.Payload)
}

func (o *RoutePutMuteTimingBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/mute-timings/{name}][%d] routePutMuteTimingBadRequest  %+v", 400, o.Payload)
}

func (o *RoutePutMuteTimingBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RoutePutMuteTimingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}