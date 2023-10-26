// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceAccountDTO swagger: model
//
// swagger:model ServiceAccountDTO
type ServiceAccountDTO struct {

	// access control
	// Example: {"serviceaccounts:delete":true,"serviceaccounts:read":true,"serviceaccounts:write":true}
	AccessControl map[string]bool `json:"accessControl,omitempty"`

	// avatar Url
	// Example: /avatar/85ec38023d90823d3e5b43ef35646af9
	AvatarURL string `json:"avatarUrl,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is disabled
	// Example: false
	IsDisabled bool `json:"isDisabled,omitempty"`

	// login
	// Example: sa-grafana
	Login string `json:"login,omitempty"`

	// name
	// Example: grafana
	Name string `json:"name,omitempty"`

	// org Id
	// Example: 1
	OrgID int64 `json:"orgId,omitempty"`

	// role
	// Example: Viewer
	Role string `json:"role,omitempty"`

	// tokens
	// Example: 0
	Tokens int64 `json:"tokens,omitempty"`
}

// Validate validates this service account DTO
func (m *ServiceAccountDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service account DTO based on context it is used
func (m *ServiceAccountDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceAccountDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceAccountDTO) UnmarshalBinary(b []byte) error {
	var res ServiceAccountDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}