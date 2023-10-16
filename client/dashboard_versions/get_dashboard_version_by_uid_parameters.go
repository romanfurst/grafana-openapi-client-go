// Code generated by go-swagger; DO NOT EDIT.

package dashboard_versions

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

// NewGetDashboardVersionByUIDParams creates a new GetDashboardVersionByUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDashboardVersionByUIDParams() *GetDashboardVersionByUIDParams {
	return &GetDashboardVersionByUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardVersionByUIDParamsWithTimeout creates a new GetDashboardVersionByUIDParams object
// with the ability to set a timeout on a request.
func NewGetDashboardVersionByUIDParamsWithTimeout(timeout time.Duration) *GetDashboardVersionByUIDParams {
	return &GetDashboardVersionByUIDParams{
		timeout: timeout,
	}
}

// NewGetDashboardVersionByUIDParamsWithContext creates a new GetDashboardVersionByUIDParams object
// with the ability to set a context for a request.
func NewGetDashboardVersionByUIDParamsWithContext(ctx context.Context) *GetDashboardVersionByUIDParams {
	return &GetDashboardVersionByUIDParams{
		Context: ctx,
	}
}

// NewGetDashboardVersionByUIDParamsWithHTTPClient creates a new GetDashboardVersionByUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDashboardVersionByUIDParamsWithHTTPClient(client *http.Client) *GetDashboardVersionByUIDParams {
	return &GetDashboardVersionByUIDParams{
		HTTPClient: client,
	}
}

/*
GetDashboardVersionByUIDParams contains all the parameters to send to the API endpoint

	for the get dashboard version by UID operation.

	Typically these are written to a http.Request.
*/
type GetDashboardVersionByUIDParams struct {

	// DashboardVersionID.
	//
	// Format: int64
	DashboardVersionID int64

	// UID.
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dashboard version by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardVersionByUIDParams) WithDefaults() *GetDashboardVersionByUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dashboard version by UID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDashboardVersionByUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) WithTimeout(timeout time.Duration) *GetDashboardVersionByUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) WithContext(ctx context.Context) *GetDashboardVersionByUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) WithHTTPClient(client *http.Client) *GetDashboardVersionByUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardVersionID adds the dashboardVersionID to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) WithDashboardVersionID(dashboardVersionID int64) *GetDashboardVersionByUIDParams {
	o.SetDashboardVersionID(dashboardVersionID)
	return o
}

// SetDashboardVersionID adds the dashboardVersionId to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) SetDashboardVersionID(dashboardVersionID int64) {
	o.DashboardVersionID = dashboardVersionID
}

// WithUID adds the uid to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) WithUID(uid string) *GetDashboardVersionByUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the get dashboard version by UID params
func (o *GetDashboardVersionByUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardVersionByUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param DashboardVersionID
	if err := r.SetPathParam("DashboardVersionID", swag.FormatInt64(o.DashboardVersionID)); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}