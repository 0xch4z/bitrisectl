// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0AppFinishParams v0 app finish params
// swagger:model v0.AppFinishParams
type V0AppFinishParams struct {

	// Which config to use `default-android-config`, `default-cordova-config`, `default-fastlane-config`, `default-ionic-config`, `default-ios-config`,`default-macos-config`, `default-react-native-config`, `default-xamarin-config`, `other-config` (default if parameter is not speficied)
	Config string `json:"config,omitempty"`

	// Environment variables for the application workflows, e.g. {"env1":"val1","env2":"val2"}
	Envs interface{} `json:"envs,omitempty"`

	// config specification mode, has to be specified with `manual` value
	Mode string `json:"mode,omitempty"`

	// The slug of the organization, who will be the owner of the application, if it's not specified, then the authenticated user will be the owner
	OrganizationSlug string `json:"organization_slug,omitempty"`

	// The type of your project (`android`, `ios`, `cordova`, `other`, `xamarin`, `macos`, `ionic`, `react-native`, `fastlane`, null)
	ProjectType string `json:"project_type,omitempty"`

	// The id of the stack the application will be built (these can be found in the [system report](https://github.com/bitrise-io/bitrise.io/tree/master/system_reports) file names)
	StackID string `json:"stack_id,omitempty"`
}

// Validate validates this v0 app finish params
func (m *V0AppFinishParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppFinishParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppFinishParams) UnmarshalBinary(b []byte) error {
	var res V0AppFinishParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
