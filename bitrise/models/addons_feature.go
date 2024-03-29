// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AddonsFeature addons feature
// swagger:model addons.Feature
type AddonsFeature struct {

	// available
	Available bool `json:"available,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// quantity
	Quantity string `json:"quantity,omitempty"`
}

// Validate validates this addons feature
func (m *AddonsFeature) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddonsFeature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddonsFeature) UnmarshalBinary(b []byte) error {
	var res AddonsFeature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
