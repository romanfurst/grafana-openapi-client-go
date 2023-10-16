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

// RouteGetMuteTimingReader is a Reader for the RouteGetMuteTiming structure.
type RouteGetMuteTimingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetMuteTimingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetMuteTimingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRouteGetMuteTimingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/provisioning/mute-timings/{name}] RouteGetMuteTiming", response, response.Code())
	}
}

// NewRouteGetMuteTimingOK creates a RouteGetMuteTimingOK with default headers values
func NewRouteGetMuteTimingOK() *RouteGetMuteTimingOK {
	return &RouteGetMuteTimingOK{}
}

/*
RouteGetMuteTimingOK describes a response with status code 200, with default header values.

MuteTimeInterval
*/
type RouteGetMuteTimingOK struct {
	Payload *models.MuteTimeInterval
}

// IsSuccess returns true when this route get mute timing o k response has a 2xx status code
func (o *RouteGetMuteTimingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this route get mute timing o k response has a 3xx status code
func (o *RouteGetMuteTimingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route get mute timing o k response has a 4xx status code
func (o *RouteGetMuteTimingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this route get mute timing o k response has a 5xx status code
func (o *RouteGetMuteTimingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this route get mute timing o k response a status code equal to that given
func (o *RouteGetMuteTimingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the route get mute timing o k response
func (o *RouteGetMuteTimingOK) Code() int {
	return 200
}

func (o *RouteGetMuteTimingOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/mute-timings/{name}][%d] routeGetMuteTimingOK  %+v", 200, o.Payload)
}

func (o *RouteGetMuteTimingOK) String() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/mute-timings/{name}][%d] routeGetMuteTimingOK  %+v", 200, o.Payload)
}

func (o *RouteGetMuteTimingOK) GetPayload() *models.MuteTimeInterval {
	return o.Payload
}

func (o *RouteGetMuteTimingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MuteTimeInterval)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteGetMuteTimingNotFound creates a RouteGetMuteTimingNotFound with default headers values
func NewRouteGetMuteTimingNotFound() *RouteGetMuteTimingNotFound {
	return &RouteGetMuteTimingNotFound{}
}

/*
RouteGetMuteTimingNotFound describes a response with status code 404, with default header values.

	Not found.
*/
type RouteGetMuteTimingNotFound struct {
}

// IsSuccess returns true when this route get mute timing not found response has a 2xx status code
func (o *RouteGetMuteTimingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this route get mute timing not found response has a 3xx status code
func (o *RouteGetMuteTimingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this route get mute timing not found response has a 4xx status code
func (o *RouteGetMuteTimingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this route get mute timing not found response has a 5xx status code
func (o *RouteGetMuteTimingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this route get mute timing not found response a status code equal to that given
func (o *RouteGetMuteTimingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the route get mute timing not found response
func (o *RouteGetMuteTimingNotFound) Code() int {
	return 404
}

func (o *RouteGetMuteTimingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/mute-timings/{name}][%d] routeGetMuteTimingNotFound ", 404)
}

func (o *RouteGetMuteTimingNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/mute-timings/{name}][%d] routeGetMuteTimingNotFound ", 404)
}

func (o *RouteGetMuteTimingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}