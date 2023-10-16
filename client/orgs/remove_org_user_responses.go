// Code generated by go-swagger; DO NOT EDIT.

package orgs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// RemoveOrgUserReader is a Reader for the RemoveOrgUser structure.
type RemoveOrgUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveOrgUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveOrgUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveOrgUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveOrgUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveOrgUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveOrgUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /orgs/{org_id}/users/{user_id}] removeOrgUser", response, response.Code())
	}
}

// NewRemoveOrgUserOK creates a RemoveOrgUserOK with default headers values
func NewRemoveOrgUserOK() *RemoveOrgUserOK {
	return &RemoveOrgUserOK{}
}

/*
RemoveOrgUserOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type RemoveOrgUserOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this remove org user o k response has a 2xx status code
func (o *RemoveOrgUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove org user o k response has a 3xx status code
func (o *RemoveOrgUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user o k response has a 4xx status code
func (o *RemoveOrgUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove org user o k response has a 5xx status code
func (o *RemoveOrgUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove org user o k response a status code equal to that given
func (o *RemoveOrgUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove org user o k response
func (o *RemoveOrgUserOK) Code() int {
	return 200
}

func (o *RemoveOrgUserOK) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserOK  %+v", 200, o.Payload)
}

func (o *RemoveOrgUserOK) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserOK  %+v", 200, o.Payload)
}

func (o *RemoveOrgUserOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserBadRequest creates a RemoveOrgUserBadRequest with default headers values
func NewRemoveOrgUserBadRequest() *RemoveOrgUserBadRequest {
	return &RemoveOrgUserBadRequest{}
}

/*
RemoveOrgUserBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RemoveOrgUserBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove org user bad request response has a 2xx status code
func (o *RemoveOrgUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove org user bad request response has a 3xx status code
func (o *RemoveOrgUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user bad request response has a 4xx status code
func (o *RemoveOrgUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove org user bad request response has a 5xx status code
func (o *RemoveOrgUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove org user bad request response a status code equal to that given
func (o *RemoveOrgUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove org user bad request response
func (o *RemoveOrgUserBadRequest) Code() int {
	return 400
}

func (o *RemoveOrgUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveOrgUserBadRequest) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveOrgUserBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserUnauthorized creates a RemoveOrgUserUnauthorized with default headers values
func NewRemoveOrgUserUnauthorized() *RemoveOrgUserUnauthorized {
	return &RemoveOrgUserUnauthorized{}
}

/*
RemoveOrgUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RemoveOrgUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove org user unauthorized response has a 2xx status code
func (o *RemoveOrgUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove org user unauthorized response has a 3xx status code
func (o *RemoveOrgUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user unauthorized response has a 4xx status code
func (o *RemoveOrgUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove org user unauthorized response has a 5xx status code
func (o *RemoveOrgUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove org user unauthorized response a status code equal to that given
func (o *RemoveOrgUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove org user unauthorized response
func (o *RemoveOrgUserUnauthorized) Code() int {
	return 401
}

func (o *RemoveOrgUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveOrgUserUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveOrgUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForbidden creates a RemoveOrgUserForbidden with default headers values
func NewRemoveOrgUserForbidden() *RemoveOrgUserForbidden {
	return &RemoveOrgUserForbidden{}
}

/*
RemoveOrgUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RemoveOrgUserForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove org user forbidden response has a 2xx status code
func (o *RemoveOrgUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove org user forbidden response has a 3xx status code
func (o *RemoveOrgUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user forbidden response has a 4xx status code
func (o *RemoveOrgUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove org user forbidden response has a 5xx status code
func (o *RemoveOrgUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove org user forbidden response a status code equal to that given
func (o *RemoveOrgUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove org user forbidden response
func (o *RemoveOrgUserForbidden) Code() int {
	return 403
}

func (o *RemoveOrgUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserForbidden  %+v", 403, o.Payload)
}

func (o *RemoveOrgUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserForbidden  %+v", 403, o.Payload)
}

func (o *RemoveOrgUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserInternalServerError creates a RemoveOrgUserInternalServerError with default headers values
func NewRemoveOrgUserInternalServerError() *RemoveOrgUserInternalServerError {
	return &RemoveOrgUserInternalServerError{}
}

/*
RemoveOrgUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RemoveOrgUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this remove org user internal server error response has a 2xx status code
func (o *RemoveOrgUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove org user internal server error response has a 3xx status code
func (o *RemoveOrgUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove org user internal server error response has a 4xx status code
func (o *RemoveOrgUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove org user internal server error response has a 5xx status code
func (o *RemoveOrgUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove org user internal server error response a status code equal to that given
func (o *RemoveOrgUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove org user internal server error response
func (o *RemoveOrgUserInternalServerError) Code() int {
	return 500
}

func (o *RemoveOrgUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveOrgUserInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org_id}/users/{user_id}][%d] removeOrgUserInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveOrgUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}