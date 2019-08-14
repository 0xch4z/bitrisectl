// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0AppUploadParams v0 app upload params
// swagger:model v0.AppUploadParams
type V0AppUploadParams struct {

	// The slug of the owner of the repository at the git provider
	GitOwner string `json:"git_owner,omitempty"`

	// The slug of the repository at the git provider
	GitRepoSlug string `json:"git_repo_slug,omitempty"`

	// If `true` then the repository visibility setting will be public, in case of `false` it will be private
	IsPublic bool `json:"is_public,omitempty"`

	// The git provider you are using, it can be `github`, `bitbucket`, `gitlab`, `gitlab-self-hosted` or `custom`
	Provider string `json:"provider,omitempty"`

	// The URL of your repository
	RepoURL string `json:"repo_url,omitempty"`

	// It has to be provided by legacy reasons and has to have the `git` value
	Type string `json:"type,omitempty"`
}

// Validate validates this v0 app upload params
func (m *V0AppUploadParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppUploadParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppUploadParams) UnmarshalBinary(b []byte) error {
	var res V0AppUploadParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}