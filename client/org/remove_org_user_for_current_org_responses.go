// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// RemoveOrgUserForCurrentOrgReader is a Reader for the RemoveOrgUserForCurrentOrg structure.
type RemoveOrgUserForCurrentOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveOrgUserForCurrentOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveOrgUserForCurrentOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveOrgUserForCurrentOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveOrgUserForCurrentOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveOrgUserForCurrentOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveOrgUserForCurrentOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /org/users/{user_id}] removeOrgUserForCurrentOrg", response, response.Code())
	}
}

// NewRemoveOrgUserForCurrentOrgOK creates a RemoveOrgUserForCurrentOrgOK with default headers values
func NewRemoveOrgUserForCurrentOrgOK() *RemoveOrgUserForCurrentOrgOK {
	return &RemoveOrgUserForCurrentOrgOK{}
}

/*
RemoveOrgUserForCurrentOrgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type RemoveOrgUserForCurrentOrgOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this remove org user for current org o k response has a 2xx status code
func (o *RemoveOrgUserForCurrentOrgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove org user for current org o k response has a 3xx status code
func (o *RemoveOrgUserForCurrentOrgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user for current org o k response has a 4xx status code
func (o *RemoveOrgUserForCurrentOrgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove org user for current org o k response has a 5xx status code
func (o *RemoveOrgUserForCurrentOrgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove org user for current org o k response a status code equal to that given
func (o *RemoveOrgUserForCurrentOrgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove org user for current org o k response
func (o *RemoveOrgUserForCurrentOrgOK) Code() int {
	return 200
}

func (o *RemoveOrgUserForCurrentOrgOK) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgOK  %+v", 200, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgOK) String() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgOK  %+v", 200, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForCurrentOrgBadRequest creates a RemoveOrgUserForCurrentOrgBadRequest with default headers values
func NewRemoveOrgUserForCurrentOrgBadRequest() *RemoveOrgUserForCurrentOrgBadRequest {
	return &RemoveOrgUserForCurrentOrgBadRequest{}
}

/*
RemoveOrgUserForCurrentOrgBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RemoveOrgUserForCurrentOrgBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove org user for current org bad request response has a 2xx status code
func (o *RemoveOrgUserForCurrentOrgBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove org user for current org bad request response has a 3xx status code
func (o *RemoveOrgUserForCurrentOrgBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user for current org bad request response has a 4xx status code
func (o *RemoveOrgUserForCurrentOrgBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove org user for current org bad request response has a 5xx status code
func (o *RemoveOrgUserForCurrentOrgBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove org user for current org bad request response a status code equal to that given
func (o *RemoveOrgUserForCurrentOrgBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove org user for current org bad request response
func (o *RemoveOrgUserForCurrentOrgBadRequest) Code() int {
	return 400
}

func (o *RemoveOrgUserForCurrentOrgBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgBadRequest) String() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForCurrentOrgUnauthorized creates a RemoveOrgUserForCurrentOrgUnauthorized with default headers values
func NewRemoveOrgUserForCurrentOrgUnauthorized() *RemoveOrgUserForCurrentOrgUnauthorized {
	return &RemoveOrgUserForCurrentOrgUnauthorized{}
}

/*
RemoveOrgUserForCurrentOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RemoveOrgUserForCurrentOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove org user for current org unauthorized response has a 2xx status code
func (o *RemoveOrgUserForCurrentOrgUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove org user for current org unauthorized response has a 3xx status code
func (o *RemoveOrgUserForCurrentOrgUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user for current org unauthorized response has a 4xx status code
func (o *RemoveOrgUserForCurrentOrgUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove org user for current org unauthorized response has a 5xx status code
func (o *RemoveOrgUserForCurrentOrgUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove org user for current org unauthorized response a status code equal to that given
func (o *RemoveOrgUserForCurrentOrgUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove org user for current org unauthorized response
func (o *RemoveOrgUserForCurrentOrgUnauthorized) Code() int {
	return 401
}

func (o *RemoveOrgUserForCurrentOrgUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForCurrentOrgForbidden creates a RemoveOrgUserForCurrentOrgForbidden with default headers values
func NewRemoveOrgUserForCurrentOrgForbidden() *RemoveOrgUserForCurrentOrgForbidden {
	return &RemoveOrgUserForCurrentOrgForbidden{}
}

/*
RemoveOrgUserForCurrentOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RemoveOrgUserForCurrentOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove org user for current org forbidden response has a 2xx status code
func (o *RemoveOrgUserForCurrentOrgForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove org user for current org forbidden response has a 3xx status code
func (o *RemoveOrgUserForCurrentOrgForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user for current org forbidden response has a 4xx status code
func (o *RemoveOrgUserForCurrentOrgForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove org user for current org forbidden response has a 5xx status code
func (o *RemoveOrgUserForCurrentOrgForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove org user for current org forbidden response a status code equal to that given
func (o *RemoveOrgUserForCurrentOrgForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove org user for current org forbidden response
func (o *RemoveOrgUserForCurrentOrgForbidden) Code() int {
	return 403
}

func (o *RemoveOrgUserForCurrentOrgForbidden) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgForbidden  %+v", 403, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgForbidden) String() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgForbidden  %+v", 403, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForCurrentOrgInternalServerError creates a RemoveOrgUserForCurrentOrgInternalServerError with default headers values
func NewRemoveOrgUserForCurrentOrgInternalServerError() *RemoveOrgUserForCurrentOrgInternalServerError {
	return &RemoveOrgUserForCurrentOrgInternalServerError{}
}

/*
RemoveOrgUserForCurrentOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RemoveOrgUserForCurrentOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove org user for current org internal server error response has a 2xx status code
func (o *RemoveOrgUserForCurrentOrgInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove org user for current org internal server error response has a 3xx status code
func (o *RemoveOrgUserForCurrentOrgInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user for current org internal server error response has a 4xx status code
func (o *RemoveOrgUserForCurrentOrgInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove org user for current org internal server error response has a 5xx status code
func (o *RemoveOrgUserForCurrentOrgInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove org user for current org internal server error response a status code equal to that given
func (o *RemoveOrgUserForCurrentOrgInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove org user for current org internal server error response
func (o *RemoveOrgUserForCurrentOrgInternalServerError) Code() int {
	return 500
}

func (o *RemoveOrgUserForCurrentOrgInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveOrgUserForCurrentOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}