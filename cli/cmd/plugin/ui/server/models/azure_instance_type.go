// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureInstanceType azure instance type
//
// swagger:model AzureInstanceType
type AzureInstanceType struct {

	// family
	Family string `json:"family,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size string `json:"size,omitempty"`

	// tier
	Tier string `json:"tier,omitempty"`

	// zones
	Zones []string `json:"zones"`
}

// Validate validates this azure instance type
func (m *AzureInstanceType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure instance type based on context it is used
func (m *AzureInstanceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureInstanceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureInstanceType) UnmarshalBinary(b []byte) error {
	var res AzureInstanceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
