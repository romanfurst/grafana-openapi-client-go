// Code generated by go-swagger; DO NOT EDIT.

package migrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new migrations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for migrations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCloudMigrationToken(opts ...ClientOption) (*CreateCloudMigrationTokenOK, error)
	CreateCloudMigrationTokenWithParams(params *CreateCloudMigrationTokenParams, opts ...ClientOption) (*CreateCloudMigrationTokenOK, error)

	CreateMigration(body *models.CloudMigrationRequest, opts ...ClientOption) (*CreateMigrationOK, error)
	CreateMigrationWithParams(params *CreateMigrationParams, opts ...ClientOption) (*CreateMigrationOK, error)

	DeleteCloudMigration(id int64, opts ...ClientOption) error
	DeleteCloudMigrationWithParams(params *DeleteCloudMigrationParams, opts ...ClientOption) error

	GetCloudMigration(id int64, opts ...ClientOption) (*GetCloudMigrationOK, error)
	GetCloudMigrationWithParams(params *GetCloudMigrationParams, opts ...ClientOption) (*GetCloudMigrationOK, error)

	GetCloudMigrationRun(runID int64, id int64, opts ...ClientOption) (*GetCloudMigrationRunOK, error)
	GetCloudMigrationRunWithParams(params *GetCloudMigrationRunParams, opts ...ClientOption) (*GetCloudMigrationRunOK, error)

	GetCloudMigrationRunList(id int64, opts ...ClientOption) (*GetCloudMigrationRunListOK, error)
	GetCloudMigrationRunListWithParams(params *GetCloudMigrationRunListParams, opts ...ClientOption) (*GetCloudMigrationRunListOK, error)

	GetMigrationList(opts ...ClientOption) (*GetMigrationListOK, error)
	GetMigrationListWithParams(params *GetMigrationListParams, opts ...ClientOption) (*GetMigrationListOK, error)

	RunCloudMigration(id int64, opts ...ClientOption) (*RunCloudMigrationOK, error)
	RunCloudMigrationWithParams(params *RunCloudMigrationParams, opts ...ClientOption) (*RunCloudMigrationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateCloudMigrationToken creates gcom access token
*/
func (a *Client) CreateCloudMigrationToken(opts ...ClientOption) (*CreateCloudMigrationTokenOK, error) {
	params := NewCreateCloudMigrationTokenParams()
	return a.CreateCloudMigrationTokenWithParams(params, opts...)
}

func (a *Client) CreateCloudMigrationTokenWithParams(params *CreateCloudMigrationTokenParams, opts ...ClientOption) (*CreateCloudMigrationTokenOK, error) {
	if params == nil {
		params = NewCreateCloudMigrationTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCloudMigrationToken",
		Method:             "POST",
		PathPattern:        "/cloudmigration/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCloudMigrationTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCloudMigrationTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCloudMigrationToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateMigration creates a migration
*/
func (a *Client) CreateMigration(body *models.CloudMigrationRequest, opts ...ClientOption) (*CreateMigrationOK, error) {
	params := NewCreateMigrationParams().WithBody(body)
	return a.CreateMigrationWithParams(params, opts...)
}

func (a *Client) CreateMigrationWithParams(params *CreateMigrationParams, opts ...ClientOption) (*CreateMigrationOK, error) {
	if params == nil {
		params = NewCreateMigrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMigration",
		Method:             "POST",
		PathPattern:        "/cloudmigration/migration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateMigrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMigrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMigration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteCloudMigration deletes a migration
*/
func (a *Client) DeleteCloudMigration(id int64, opts ...ClientOption) error {
	params := NewDeleteCloudMigrationParams().WithID(id)
	return a.DeleteCloudMigrationWithParams(params, opts...)
}

func (a *Client) DeleteCloudMigrationWithParams(params *DeleteCloudMigrationParams, opts ...ClientOption) error {
	if params == nil {
		params = NewDeleteCloudMigrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCloudMigration",
		Method:             "DELETE",
		PathPattern:        "/cloudmigration/migration/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCloudMigrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetCloudMigration gets a cloud migration

It returns migrations that has been created.
*/
func (a *Client) GetCloudMigration(id int64, opts ...ClientOption) (*GetCloudMigrationOK, error) {
	params := NewGetCloudMigrationParams().WithID(id)
	return a.GetCloudMigrationWithParams(params, opts...)
}

func (a *Client) GetCloudMigrationWithParams(params *GetCloudMigrationParams, opts ...ClientOption) (*GetCloudMigrationOK, error) {
	if params == nil {
		params = NewGetCloudMigrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCloudMigration",
		Method:             "GET",
		PathPattern:        "/cloudmigration/migration/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCloudMigrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCloudMigrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCloudMigration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCloudMigrationRun gets the result of a single migration run
*/
func (a *Client) GetCloudMigrationRun(runID int64, id int64, opts ...ClientOption) (*GetCloudMigrationRunOK, error) {
	params := NewGetCloudMigrationRunParams().WithID(id).WithRunID(runID)
	return a.GetCloudMigrationRunWithParams(params, opts...)
}

func (a *Client) GetCloudMigrationRunWithParams(params *GetCloudMigrationRunParams, opts ...ClientOption) (*GetCloudMigrationRunOK, error) {
	if params == nil {
		params = NewGetCloudMigrationRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCloudMigrationRun",
		Method:             "GET",
		PathPattern:        "/cloudmigration/migration/{id}/run/{runID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCloudMigrationRunReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCloudMigrationRunOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCloudMigrationRun: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCloudMigrationRunList gets a list of migration runs for a migration
*/
func (a *Client) GetCloudMigrationRunList(id int64, opts ...ClientOption) (*GetCloudMigrationRunListOK, error) {
	params := NewGetCloudMigrationRunListParams().WithID(id)
	return a.GetCloudMigrationRunListWithParams(params, opts...)
}

func (a *Client) GetCloudMigrationRunListWithParams(params *GetCloudMigrationRunListParams, opts ...ClientOption) (*GetCloudMigrationRunListOK, error) {
	if params == nil {
		params = NewGetCloudMigrationRunListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCloudMigrationRunList",
		Method:             "GET",
		PathPattern:        "/cloudmigration/migration/{id}/run",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCloudMigrationRunListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCloudMigrationRunListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCloudMigrationRunList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMigrationList gets a list of all cloud migrations
*/
func (a *Client) GetMigrationList(opts ...ClientOption) (*GetMigrationListOK, error) {
	params := NewGetMigrationListParams()
	return a.GetMigrationListWithParams(params, opts...)
}

func (a *Client) GetMigrationListWithParams(params *GetMigrationListParams, opts ...ClientOption) (*GetMigrationListOK, error) {
	if params == nil {
		params = NewGetMigrationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMigrationList",
		Method:             "GET",
		PathPattern:        "/cloudmigration/migration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMigrationListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMigrationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMigrationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RunCloudMigration triggers the run of a migration to the grafana cloud

It returns migrations that has been created.
*/
func (a *Client) RunCloudMigration(id int64, opts ...ClientOption) (*RunCloudMigrationOK, error) {
	params := NewRunCloudMigrationParams().WithID(id)
	return a.RunCloudMigrationWithParams(params, opts...)
}

func (a *Client) RunCloudMigrationWithParams(params *RunCloudMigrationParams, opts ...ClientOption) (*RunCloudMigrationOK, error) {
	if params == nil {
		params = NewRunCloudMigrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "runCloudMigration",
		Method:             "POST",
		PathPattern:        "/cloudmigration/migration/{id}/run",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunCloudMigrationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunCloudMigrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for runCloudMigration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// WithAuthInfo changes the transport on the client
func WithAuthInfo(authInfo runtime.ClientAuthInfoWriter) ClientOption {
	return func(op *runtime.ClientOperation) {
		op.AuthInfo = authInfo
	}
}
