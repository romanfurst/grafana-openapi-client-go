// Code generated by go-swagger; DO NOT EDIT.

package library_elements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetLibraryElementByNameReader is a Reader for the GetLibraryElementByName structure.
type GetLibraryElementByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLibraryElementByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLibraryElementByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLibraryElementByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLibraryElementByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLibraryElementByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /library-elements/name/{library_element_name}] getLibraryElementByName", response, response.Code())
	}
}

// NewGetLibraryElementByNameOK creates a GetLibraryElementByNameOK with default headers values
func NewGetLibraryElementByNameOK() *GetLibraryElementByNameOK {
	return &GetLibraryElementByNameOK{}
}

/*
GetLibraryElementByNameOK describes a response with status code 200, with default header values.

(empty)
*/
type GetLibraryElementByNameOK struct {
	Payload *models.LibraryElementArrayResponse
}

// IsSuccess returns true when this get library element by name Ok response has a 2xx status code
func (o *GetLibraryElementByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get library element by name Ok response has a 3xx status code
func (o *GetLibraryElementByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element by name Ok response has a 4xx status code
func (o *GetLibraryElementByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get library element by name Ok response has a 5xx status code
func (o *GetLibraryElementByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get library element by name Ok response a status code equal to that given
func (o *GetLibraryElementByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get library element by name Ok response
func (o *GetLibraryElementByNameOK) Code() int {
	return 200
}

func (o *GetLibraryElementByNameOK) Error() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameOk  %+v", 200, o.Payload)
}

func (o *GetLibraryElementByNameOK) String() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameOk  %+v", 200, o.Payload)
}

func (o *GetLibraryElementByNameOK) GetPayload() *models.LibraryElementArrayResponse {
	return o.Payload
}

func (o *GetLibraryElementByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LibraryElementArrayResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementByNameUnauthorized creates a GetLibraryElementByNameUnauthorized with default headers values
func NewGetLibraryElementByNameUnauthorized() *GetLibraryElementByNameUnauthorized {
	return &GetLibraryElementByNameUnauthorized{}
}

/*
GetLibraryElementByNameUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetLibraryElementByNameUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get library element by name unauthorized response has a 2xx status code
func (o *GetLibraryElementByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get library element by name unauthorized response has a 3xx status code
func (o *GetLibraryElementByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element by name unauthorized response has a 4xx status code
func (o *GetLibraryElementByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get library element by name unauthorized response has a 5xx status code
func (o *GetLibraryElementByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get library element by name unauthorized response a status code equal to that given
func (o *GetLibraryElementByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get library element by name unauthorized response
func (o *GetLibraryElementByNameUnauthorized) Code() int {
	return 401
}

func (o *GetLibraryElementByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLibraryElementByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLibraryElementByNameUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementByNameNotFound creates a GetLibraryElementByNameNotFound with default headers values
func NewGetLibraryElementByNameNotFound() *GetLibraryElementByNameNotFound {
	return &GetLibraryElementByNameNotFound{}
}

/*
GetLibraryElementByNameNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetLibraryElementByNameNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get library element by name not found response has a 2xx status code
func (o *GetLibraryElementByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get library element by name not found response has a 3xx status code
func (o *GetLibraryElementByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element by name not found response has a 4xx status code
func (o *GetLibraryElementByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get library element by name not found response has a 5xx status code
func (o *GetLibraryElementByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get library element by name not found response a status code equal to that given
func (o *GetLibraryElementByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get library element by name not found response
func (o *GetLibraryElementByNameNotFound) Code() int {
	return 404
}

func (o *GetLibraryElementByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetLibraryElementByNameNotFound) String() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameNotFound  %+v", 404, o.Payload)
}

func (o *GetLibraryElementByNameNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementByNameInternalServerError creates a GetLibraryElementByNameInternalServerError with default headers values
func NewGetLibraryElementByNameInternalServerError() *GetLibraryElementByNameInternalServerError {
	return &GetLibraryElementByNameInternalServerError{}
}

/*
GetLibraryElementByNameInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetLibraryElementByNameInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get library element by name internal server error response has a 2xx status code
func (o *GetLibraryElementByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get library element by name internal server error response has a 3xx status code
func (o *GetLibraryElementByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get library element by name internal server error response has a 4xx status code
func (o *GetLibraryElementByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get library element by name internal server error response has a 5xx status code
func (o *GetLibraryElementByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get library element by name internal server error response a status code equal to that given
func (o *GetLibraryElementByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get library element by name internal server error response
func (o *GetLibraryElementByNameInternalServerError) Code() int {
	return 500
}

func (o *GetLibraryElementByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLibraryElementByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLibraryElementByNameInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
