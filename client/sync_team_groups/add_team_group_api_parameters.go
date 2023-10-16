// Code generated by go-swagger; DO NOT EDIT.

package sync_team_groups

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

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewAddTeamGroupAPIParams creates a new AddTeamGroupAPIParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddTeamGroupAPIParams() *AddTeamGroupAPIParams {
	return &AddTeamGroupAPIParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddTeamGroupAPIParamsWithTimeout creates a new AddTeamGroupAPIParams object
// with the ability to set a timeout on a request.
func NewAddTeamGroupAPIParamsWithTimeout(timeout time.Duration) *AddTeamGroupAPIParams {
	return &AddTeamGroupAPIParams{
		timeout: timeout,
	}
}

// NewAddTeamGroupAPIParamsWithContext creates a new AddTeamGroupAPIParams object
// with the ability to set a context for a request.
func NewAddTeamGroupAPIParamsWithContext(ctx context.Context) *AddTeamGroupAPIParams {
	return &AddTeamGroupAPIParams{
		Context: ctx,
	}
}

// NewAddTeamGroupAPIParamsWithHTTPClient creates a new AddTeamGroupAPIParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddTeamGroupAPIParamsWithHTTPClient(client *http.Client) *AddTeamGroupAPIParams {
	return &AddTeamGroupAPIParams{
		HTTPClient: client,
	}
}

/*
AddTeamGroupAPIParams contains all the parameters to send to the API endpoint

	for the add team group Api operation.

	Typically these are written to a http.Request.
*/
type AddTeamGroupAPIParams struct {

	// Body.
	Body *models.TeamGroupMapping

	// TeamID.
	//
	// Format: int64
	TeamID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add team group Api params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTeamGroupAPIParams) WithDefaults() *AddTeamGroupAPIParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add team group Api params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTeamGroupAPIParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add team group Api params
func (o *AddTeamGroupAPIParams) WithTimeout(timeout time.Duration) *AddTeamGroupAPIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add team group Api params
func (o *AddTeamGroupAPIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add team group Api params
func (o *AddTeamGroupAPIParams) WithContext(ctx context.Context) *AddTeamGroupAPIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add team group Api params
func (o *AddTeamGroupAPIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add team group Api params
func (o *AddTeamGroupAPIParams) WithHTTPClient(client *http.Client) *AddTeamGroupAPIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add team group Api params
func (o *AddTeamGroupAPIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add team group Api params
func (o *AddTeamGroupAPIParams) WithBody(body *models.TeamGroupMapping) *AddTeamGroupAPIParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add team group Api params
func (o *AddTeamGroupAPIParams) SetBody(body *models.TeamGroupMapping) {
	o.Body = body
}

// WithTeamID adds the teamID to the add team group Api params
func (o *AddTeamGroupAPIParams) WithTeamID(teamID int64) *AddTeamGroupAPIParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the add team group Api params
func (o *AddTeamGroupAPIParams) SetTeamID(teamID int64) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *AddTeamGroupAPIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param teamId
	if err := r.SetPathParam("teamId", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}