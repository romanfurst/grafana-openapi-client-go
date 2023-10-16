// Code generated by go-swagger; DO NOT EDIT.

package folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command
// To edit this file, edit the custom template in templates/clientClient.gotmpl
// More info on custom templates can be found in the README or here: https://github.com/go-swagger/go-swagger/blob/master/docs/generate/templates.md
// The template itself can be found here: https://github.com/go-swagger/go-swagger/blob/master/generator/templates/client/client.gotmpl

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new folders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for folders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateFolder(params *CreateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFolderOK, error)

	DeleteFolder(params *DeleteFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFolderOK, error)

	GetFolderByID(params *GetFolderByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderByIDOK, error)

	GetFolderByUID(params *GetFolderByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderByUIDOK, error)

	GetFolderDescendantCounts(params *GetFolderDescendantCountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderDescendantCountsOK, error)

	GetFolders(params *GetFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFoldersOK, error)

	MoveFolder(params *MoveFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MoveFolderOK, error)

	UpdateFolder(params *UpdateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFolderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateFolder creates folder

If nested folders are enabled then it additionally expects the parent folder UID.
*/
func (a *Client) CreateFolder(params *CreateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFolder",
		Method:             "POST",
		PathPattern:        "/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	DeleteFolder deletes folder

	Deletes an existing folder identified by UID along with all dashboards (and their alerts) stored in the folder. This operation cannot be reverted.

If nested folders are enabled then it also deletes all the subfolders.
*/
func (a *Client) DeleteFolder(params *DeleteFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteFolder",
		Method:             "DELETE",
		PathPattern:        "/folders/{folder_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetFolderByID gets folder by id

	Returns the folder identified by id. This is deprecated.

Please refer to [updated API](#/folders/getFolderByUID) instead
*/
func (a *Client) GetFolderByID(params *GetFolderByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFolderByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolderByID",
		Method:             "GET",
		PathPattern:        "/folders/id/{folder_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFolderByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFolderByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolderByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFolderByUID gets folder by uid
*/
func (a *Client) GetFolderByUID(params *GetFolderByUIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderByUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFolderByUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolderByUID",
		Method:             "GET",
		PathPattern:        "/folders/{folder_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFolderByUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFolderByUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolderByUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFolderDescendantCounts gets the count of each descendant of a folder by kind the folder is identified by UID
*/
func (a *Client) GetFolderDescendantCounts(params *GetFolderDescendantCountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFolderDescendantCountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFolderDescendantCountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolderDescendantCounts",
		Method:             "GET",
		PathPattern:        "/folders/{folder_uid}/counts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFolderDescendantCountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFolderDescendantCountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolderDescendantCounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetFolders gets all folders

	Returns all folders that the authenticated user has permission to view.

If nested folders are enabled, it expects an additional query parameter with the parent folder UID
and returns the immediate subfolders that the authenticated user has permission to view.
If the parameter is not supplied then it returns immediate subfolders under the root
that the authenticated user has permission to view.
*/
func (a *Client) GetFolders(params *GetFoldersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFoldersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFolders",
		Method:             "GET",
		PathPattern:        "/folders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFoldersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFoldersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFolders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
MoveFolder moves folder
*/
func (a *Client) MoveFolder(params *MoveFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MoveFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "moveFolder",
		Method:             "POST",
		PathPattern:        "/folders/{folder_uid}/move",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &MoveFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MoveFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for moveFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateFolder updates folder
*/
func (a *Client) UpdateFolder(params *UpdateFolderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateFolderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFolderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateFolder",
		Method:             "PUT",
		PathPattern:        "/folders/{folder_uid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateFolderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateFolderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFolder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}