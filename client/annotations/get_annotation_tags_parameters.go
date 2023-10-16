// Code generated by go-swagger; DO NOT EDIT.

package annotations

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

// NewGetAnnotationTagsParams creates a new GetAnnotationTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAnnotationTagsParams() *GetAnnotationTagsParams {
	return &GetAnnotationTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnnotationTagsParamsWithTimeout creates a new GetAnnotationTagsParams object
// with the ability to set a timeout on a request.
func NewGetAnnotationTagsParamsWithTimeout(timeout time.Duration) *GetAnnotationTagsParams {
	return &GetAnnotationTagsParams{
		timeout: timeout,
	}
}

// NewGetAnnotationTagsParamsWithContext creates a new GetAnnotationTagsParams object
// with the ability to set a context for a request.
func NewGetAnnotationTagsParamsWithContext(ctx context.Context) *GetAnnotationTagsParams {
	return &GetAnnotationTagsParams{
		Context: ctx,
	}
}

// NewGetAnnotationTagsParamsWithHTTPClient creates a new GetAnnotationTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAnnotationTagsParamsWithHTTPClient(client *http.Client) *GetAnnotationTagsParams {
	return &GetAnnotationTagsParams{
		HTTPClient: client,
	}
}

/*
GetAnnotationTagsParams contains all the parameters to send to the API endpoint

	for the get annotation tags operation.

	Typically these are written to a http.Request.
*/
type GetAnnotationTagsParams struct {

	/* Limit.

	   Max limit for results returned.

	   Default: "100"
	*/
	Limit *string

	/* Tag.

	   Tag is a string that you can use to filter tags.
	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get annotation tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnnotationTagsParams) WithDefaults() *GetAnnotationTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get annotation tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAnnotationTagsParams) SetDefaults() {
	var (
		limitDefault = string("100")
	)

	val := GetAnnotationTagsParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get annotation tags params
func (o *GetAnnotationTagsParams) WithTimeout(timeout time.Duration) *GetAnnotationTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get annotation tags params
func (o *GetAnnotationTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get annotation tags params
func (o *GetAnnotationTagsParams) WithContext(ctx context.Context) *GetAnnotationTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get annotation tags params
func (o *GetAnnotationTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get annotation tags params
func (o *GetAnnotationTagsParams) WithHTTPClient(client *http.Client) *GetAnnotationTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get annotation tags params
func (o *GetAnnotationTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get annotation tags params
func (o *GetAnnotationTagsParams) WithLimit(limit *string) *GetAnnotationTagsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get annotation tags params
func (o *GetAnnotationTagsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithTag adds the tag to the get annotation tags params
func (o *GetAnnotationTagsParams) WithTag(tag *string) *GetAnnotationTagsParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get annotation tags params
func (o *GetAnnotationTagsParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnnotationTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}