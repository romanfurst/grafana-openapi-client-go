// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// NoticeSeverity NoticeSeverity is a type for the Severity property of a Notice.
//
// swagger:model NoticeSeverity
type NoticeSeverity int64

// Validate validates this notice severity
func (m NoticeSeverity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notice severity based on context it is used
func (m NoticeSeverity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}