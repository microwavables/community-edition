// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureVirtualMachine azure virtual machine
//
// swagger:model AzureVirtualMachine
type AzureVirtualMachine struct {

	// name
	Name string `json:"name,omitempty"`

	// os info
	OsInfo *OSInfo `json:"osInfo,omitempty"`
}

// Validate validates this azure virtual machine
func (m *AzureVirtualMachine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureVirtualMachine) validateOsInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.OsInfo) { // not required
		return nil
	}

	if m.OsInfo != nil {
		if err := m.OsInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("osInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this azure virtual machine based on the context it is used
func (m *AzureVirtualMachine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureVirtualMachine) contextValidateOsInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.OsInfo != nil {
		if err := m.OsInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("osInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("osInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureVirtualMachine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureVirtualMachine) UnmarshalBinary(b []byte) error {
	var res AzureVirtualMachine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
