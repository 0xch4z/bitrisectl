// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0SSHKeyRespModel v0 SSH key resp model
// swagger:model v0.SSHKeyRespModel
type V0SSHKeyRespModel struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this v0 SSH key resp model
func (m *V0SSHKeyRespModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0SSHKeyRespModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0SSHKeyRespModel) UnmarshalBinary(b []byte) error {
	var res V0SSHKeyRespModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
