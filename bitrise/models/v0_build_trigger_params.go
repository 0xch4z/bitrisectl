// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V0BuildTriggerParams v0 build trigger params
// swagger:model v0.BuildTriggerParams
type V0BuildTriggerParams struct {

	// The public part of the SSH key you would like to use
	BuildParams *V0BuildTriggerParamsBuildParams `json:"build_params,omitempty"`

	// hook info
	HookInfo *V0BuildTriggerParamsHookInfo `json:"hook_info,omitempty"`
}

// Validate validates this v0 build trigger params
func (m *V0BuildTriggerParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHookInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0BuildTriggerParams) validateBuildParams(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildParams) { // not required
		return nil
	}

	if m.BuildParams != nil {
		if err := m.BuildParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build_params")
			}
			return err
		}
	}

	return nil
}

func (m *V0BuildTriggerParams) validateHookInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.HookInfo) { // not required
		return nil
	}

	if m.HookInfo != nil {
		if err := m.HookInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hook_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V0BuildTriggerParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0BuildTriggerParams) UnmarshalBinary(b []byte) error {
	var res V0BuildTriggerParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
