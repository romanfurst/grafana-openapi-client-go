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

// GetAlertRulesReader is a Reader for the GetAlertRules structure.
type GetAlertRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/provisioning/alert-rules] GetAlertRules", response, response.Code())
	}
}

// NewGetAlertRulesOK creates a GetAlertRulesOK with default headers values
func NewGetAlertRulesOK() *GetAlertRulesOK {
	return &GetAlertRulesOK{}
}

/*
GetAlertRulesOK describes a response with status code 200, with default header values.

ProvisionedAlertRules
*/
type GetAlertRulesOK struct {
	Payload models.ProvisionedAlertRules
}

// IsSuccess returns true when this get alert rules Ok response has a 2xx status code
func (o *GetAlertRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert rules Ok response has a 3xx status code
func (o *GetAlertRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rules Ok response has a 4xx status code
func (o *GetAlertRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert rules Ok response has a 5xx status code
func (o *GetAlertRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rules Ok response a status code equal to that given
func (o *GetAlertRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert rules Ok response
func (o *GetAlertRulesOK) Code() int {
	return 200
}

func (o *GetAlertRulesOK) Error() string {
	return fmt.Sprintf("[GET /v1/provisioning/alert-rules][%d] getAlertRulesOk  %+v", 200, o.Payload)
}

func (o *GetAlertRulesOK) String() string {
	return fmt.Sprintf("[GET /v1/provisioning/alert-rules][%d] getAlertRulesOk  %+v", 200, o.Payload)
}

func (o *GetAlertRulesOK) GetPayload() models.ProvisionedAlertRules {
	return o.Payload
}

func (o *GetAlertRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
