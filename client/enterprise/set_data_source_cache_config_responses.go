// Code generated by go-swagger; DO NOT EDIT.

package enterprise

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// SetDataSourceCacheConfigReader is a Reader for the SetDataSourceCacheConfig structure.
type SetDataSourceCacheConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDataSourceCacheConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDataSourceCacheConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetDataSourceCacheConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetDataSourceCacheConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /datasources/{dataSourceUID}/cache] setDataSourceCacheConfig", response, response.Code())
	}
}

// NewSetDataSourceCacheConfigOK creates a SetDataSourceCacheConfigOK with default headers values
func NewSetDataSourceCacheConfigOK() *SetDataSourceCacheConfigOK {
	return &SetDataSourceCacheConfigOK{}
}

/*
SetDataSourceCacheConfigOK describes a response with status code 200, with default header values.

CacheConfigResponse
*/
type SetDataSourceCacheConfigOK struct {
	Payload *models.CacheConfigResponse
}

// IsSuccess returns true when this set data source cache config Ok response has a 2xx status code
func (o *SetDataSourceCacheConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set data source cache config Ok response has a 3xx status code
func (o *SetDataSourceCacheConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set data source cache config Ok response has a 4xx status code
func (o *SetDataSourceCacheConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set data source cache config Ok response has a 5xx status code
func (o *SetDataSourceCacheConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set data source cache config Ok response a status code equal to that given
func (o *SetDataSourceCacheConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set data source cache config Ok response
func (o *SetDataSourceCacheConfigOK) Code() int {
	return 200
}

func (o *SetDataSourceCacheConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /datasources/{dataSourceUID}/cache][%d] setDataSourceCacheConfigOk %s", 200, payload)
}

func (o *SetDataSourceCacheConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /datasources/{dataSourceUID}/cache][%d] setDataSourceCacheConfigOk %s", 200, payload)
}

func (o *SetDataSourceCacheConfigOK) GetPayload() *models.CacheConfigResponse {
	return o.Payload
}

func (o *SetDataSourceCacheConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CacheConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDataSourceCacheConfigBadRequest creates a SetDataSourceCacheConfigBadRequest with default headers values
func NewSetDataSourceCacheConfigBadRequest() *SetDataSourceCacheConfigBadRequest {
	return &SetDataSourceCacheConfigBadRequest{}
}

/*
SetDataSourceCacheConfigBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type SetDataSourceCacheConfigBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set data source cache config bad request response has a 2xx status code
func (o *SetDataSourceCacheConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set data source cache config bad request response has a 3xx status code
func (o *SetDataSourceCacheConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set data source cache config bad request response has a 4xx status code
func (o *SetDataSourceCacheConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this set data source cache config bad request response has a 5xx status code
func (o *SetDataSourceCacheConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this set data source cache config bad request response a status code equal to that given
func (o *SetDataSourceCacheConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the set data source cache config bad request response
func (o *SetDataSourceCacheConfigBadRequest) Code() int {
	return 400
}

func (o *SetDataSourceCacheConfigBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /datasources/{dataSourceUID}/cache][%d] setDataSourceCacheConfigBadRequest %s", 400, payload)
}

func (o *SetDataSourceCacheConfigBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /datasources/{dataSourceUID}/cache][%d] setDataSourceCacheConfigBadRequest %s", 400, payload)
}

func (o *SetDataSourceCacheConfigBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetDataSourceCacheConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDataSourceCacheConfigInternalServerError creates a SetDataSourceCacheConfigInternalServerError with default headers values
func NewSetDataSourceCacheConfigInternalServerError() *SetDataSourceCacheConfigInternalServerError {
	return &SetDataSourceCacheConfigInternalServerError{}
}

/*
SetDataSourceCacheConfigInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SetDataSourceCacheConfigInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set data source cache config internal server error response has a 2xx status code
func (o *SetDataSourceCacheConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set data source cache config internal server error response has a 3xx status code
func (o *SetDataSourceCacheConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set data source cache config internal server error response has a 4xx status code
func (o *SetDataSourceCacheConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set data source cache config internal server error response has a 5xx status code
func (o *SetDataSourceCacheConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set data source cache config internal server error response a status code equal to that given
func (o *SetDataSourceCacheConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set data source cache config internal server error response
func (o *SetDataSourceCacheConfigInternalServerError) Code() int {
	return 500
}

func (o *SetDataSourceCacheConfigInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /datasources/{dataSourceUID}/cache][%d] setDataSourceCacheConfigInternalServerError %s", 500, payload)
}

func (o *SetDataSourceCacheConfigInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /datasources/{dataSourceUID}/cache][%d] setDataSourceCacheConfigInternalServerError %s", 500, payload)
}

func (o *SetDataSourceCacheConfigInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetDataSourceCacheConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
