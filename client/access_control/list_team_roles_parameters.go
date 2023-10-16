// Code generated by go-swagger; DO NOT EDIT.

package access_control

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
	"github.com/go-openapi/swag"
)

// NewListTeamRolesParams creates a new ListTeamRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTeamRolesParams() *ListTeamRolesParams {
	return &ListTeamRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTeamRolesParamsWithTimeout creates a new ListTeamRolesParams object
// with the ability to set a timeout on a request.
func NewListTeamRolesParamsWithTimeout(timeout time.Duration) *ListTeamRolesParams {
	return &ListTeamRolesParams{
		timeout: timeout,
	}
}

// NewListTeamRolesParamsWithContext creates a new ListTeamRolesParams object
// with the ability to set a context for a request.
func NewListTeamRolesParamsWithContext(ctx context.Context) *ListTeamRolesParams {
	return &ListTeamRolesParams{
		Context: ctx,
	}
}

// NewListTeamRolesParamsWithHTTPClient creates a new ListTeamRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTeamRolesParamsWithHTTPClient(client *http.Client) *ListTeamRolesParams {
	return &ListTeamRolesParams{
		HTTPClient: client,
	}
}

/*
ListTeamRolesParams contains all the parameters to send to the API endpoint

	for the list team roles operation.

	Typically these are written to a http.Request.
*/
type ListTeamRolesParams struct {

	// TeamID.
	//
	// Format: int64
	TeamID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list team roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTeamRolesParams) WithDefaults() *ListTeamRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list team roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTeamRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list team roles params
func (o *ListTeamRolesParams) WithTimeout(timeout time.Duration) *ListTeamRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list team roles params
func (o *ListTeamRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list team roles params
func (o *ListTeamRolesParams) WithContext(ctx context.Context) *ListTeamRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list team roles params
func (o *ListTeamRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list team roles params
func (o *ListTeamRolesParams) WithHTTPClient(client *http.Client) *ListTeamRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list team roles params
func (o *ListTeamRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the list team roles params
func (o *ListTeamRolesParams) WithTeamID(teamID int64) *ListTeamRolesParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the list team roles params
func (o *ListTeamRolesParams) SetTeamID(teamID int64) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *ListTeamRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param teamId
	if err := r.SetPathParam("teamId", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}