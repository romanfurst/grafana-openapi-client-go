// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetTemplateParams creates a new GetTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTemplateParams() *GetTemplateParams {
	return &GetTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTemplateParamsWithTimeout creates a new GetTemplateParams object
// with the ability to set a timeout on a request.
func NewGetTemplateParamsWithTimeout(timeout time.Duration) *GetTemplateParams {
	return &GetTemplateParams{
		timeout: timeout,
	}
}

// NewGetTemplateParamsWithContext creates a new GetTemplateParams object
// with the ability to set a context for a request.
func NewGetTemplateParamsWithContext(ctx context.Context) *GetTemplateParams {
	return &GetTemplateParams{
		Context: ctx,
	}
}

// NewGetTemplateParamsWithHTTPClient creates a new GetTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTemplateParamsWithHTTPClient(client *http.Client) *GetTemplateParams {
	return &GetTemplateParams{
		HTTPClient: client,
	}
}

/*
GetTemplateParams contains all the parameters to send to the API endpoint

	for the get template operation.

	Typically these are written to a http.Request.
*/
type GetTemplateParams struct {

	/* Name.

	   Template Name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTemplateParams) WithDefaults() *GetTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get template params
func (o *GetTemplateParams) WithTimeout(timeout time.Duration) *GetTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get template params
func (o *GetTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get template params
func (o *GetTemplateParams) WithContext(ctx context.Context) *GetTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get template params
func (o *GetTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get template params
func (o *GetTemplateParams) WithHTTPClient(client *http.Client) *GetTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get template params
func (o *GetTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get template params
func (o *GetTemplateParams) WithName(name string) *GetTemplateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get template params
func (o *GetTemplateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}