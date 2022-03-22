// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Policy The model for policy
//
// swagger:model Policy
type Policy struct {

	// The policy id
	ID int64 `json:"id,omitempty"`

	// The policy name
	Name string `json:"name,omitempty"`
}

// Validate validates this policy
func (m *Policy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy based on context it is used
func (m *Policy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policy) UnmarshalBinary(b []byte) error {
	var res Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}