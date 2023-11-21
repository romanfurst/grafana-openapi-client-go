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

// PostAlertRuleReader is a Reader for the PostAlertRule structure.
type PostAlertRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAlertRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAlertRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAlertRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/provisioning/alert-rules] PostAlertRule", response, response.Code())
	}
}

// NewPostAlertRuleCreated creates a PostAlertRuleCreated with default headers values
func NewPostAlertRuleCreated() *PostAlertRuleCreated {
	return &PostAlertRuleCreated{}
}

/*
PostAlertRuleCreated describes a response with status code 201, with default header values.

ProvisionedAlertRule
*/
type PostAlertRuleCreated struct {
	Payload *models.ProvisionedAlertRule
}

// IsSuccess returns true when this post alert rule created response has a 2xx status code
func (o *PostAlertRuleCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post alert rule created response has a 3xx status code
func (o *PostAlertRuleCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post alert rule created response has a 4xx status code
func (o *PostAlertRuleCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post alert rule created response has a 5xx status code
func (o *PostAlertRuleCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post alert rule created response a status code equal to that given
func (o *PostAlertRuleCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post alert rule created response
func (o *PostAlertRuleCreated) Code() int {
	return 201
}

func (o *PostAlertRuleCreated) Error() string {
	return fmt.Sprintf("[POST /v1/provisioning/alert-rules][%d] postAlertRuleCreated  %+v", 201, o.Payload)
}

func (o *PostAlertRuleCreated) String() string {
	return fmt.Sprintf("[POST /v1/provisioning/alert-rules][%d] postAlertRuleCreated  %+v", 201, o.Payload)
}

func (o *PostAlertRuleCreated) GetPayload() *models.ProvisionedAlertRule {
	return o.Payload
}

func (o *PostAlertRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvisionedAlertRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertRuleBadRequest creates a PostAlertRuleBadRequest with default headers values
func NewPostAlertRuleBadRequest() *PostAlertRuleBadRequest {
	return &PostAlertRuleBadRequest{}
}

/*
PostAlertRuleBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type PostAlertRuleBadRequest struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this post alert rule bad request response has a 2xx status code
func (o *PostAlertRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post alert rule bad request response has a 3xx status code
func (o *PostAlertRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post alert rule bad request response has a 4xx status code
func (o *PostAlertRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post alert rule bad request response has a 5xx status code
func (o *PostAlertRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post alert rule bad request response a status code equal to that given
func (o *PostAlertRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post alert rule bad request response
func (o *PostAlertRuleBadRequest) Code() int {
	return 400
}

func (o *PostAlertRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/provisioning/alert-rules][%d] postAlertRuleBadRequest  %+v", 400, o.Payload)
}

func (o *PostAlertRuleBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/provisioning/alert-rules][%d] postAlertRuleBadRequest  %+v", 400, o.Payload)
}

func (o *PostAlertRuleBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *PostAlertRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
