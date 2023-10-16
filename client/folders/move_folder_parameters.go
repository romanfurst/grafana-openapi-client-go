// Code generated by go-swagger; DO NOT EDIT.

package folders

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

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewMoveFolderParams creates a new MoveFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMoveFolderParams() *MoveFolderParams {
	return &MoveFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMoveFolderParamsWithTimeout creates a new MoveFolderParams object
// with the ability to set a timeout on a request.
func NewMoveFolderParamsWithTimeout(timeout time.Duration) *MoveFolderParams {
	return &MoveFolderParams{
		timeout: timeout,
	}
}

// NewMoveFolderParamsWithContext creates a new MoveFolderParams object
// with the ability to set a context for a request.
func NewMoveFolderParamsWithContext(ctx context.Context) *MoveFolderParams {
	return &MoveFolderParams{
		Context: ctx,
	}
}

// NewMoveFolderParamsWithHTTPClient creates a new MoveFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewMoveFolderParamsWithHTTPClient(client *http.Client) *MoveFolderParams {
	return &MoveFolderParams{
		HTTPClient: client,
	}
}

/*
MoveFolderParams contains all the parameters to send to the API endpoint

	for the move folder operation.

	Typically these are written to a http.Request.
*/
type MoveFolderParams struct {

	// Body.
	Body *models.MoveFolderCommand

	// FolderUID.
	FolderUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the move folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MoveFolderParams) WithDefaults() *MoveFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the move folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MoveFolderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the move folder params
func (o *MoveFolderParams) WithTimeout(timeout time.Duration) *MoveFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move folder params
func (o *MoveFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move folder params
func (o *MoveFolderParams) WithContext(ctx context.Context) *MoveFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move folder params
func (o *MoveFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move folder params
func (o *MoveFolderParams) WithHTTPClient(client *http.Client) *MoveFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move folder params
func (o *MoveFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the move folder params
func (o *MoveFolderParams) WithBody(body *models.MoveFolderCommand) *MoveFolderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move folder params
func (o *MoveFolderParams) SetBody(body *models.MoveFolderCommand) {
	o.Body = body
}

// WithFolderUID adds the folderUID to the move folder params
func (o *MoveFolderParams) WithFolderUID(folderUID string) *MoveFolderParams {
	o.SetFolderUID(folderUID)
	return o
}

// SetFolderUID adds the folderUid to the move folder params
func (o *MoveFolderParams) SetFolderUID(folderUID string) {
	o.FolderUID = folderUID
}

// WriteToRequest writes these params to a swagger request
func (o *MoveFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param folder_uid
	if err := r.SetPathParam("folder_uid", o.FolderUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}