// Code generated by go-swagger; DO NOT EDIT.

package addons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAddonListByOrganizationParams creates a new AddonListByOrganizationParams object
// with the default values initialized.
func NewAddonListByOrganizationParams() *AddonListByOrganizationParams {
	var ()
	return &AddonListByOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddonListByOrganizationParamsWithTimeout creates a new AddonListByOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddonListByOrganizationParamsWithTimeout(timeout time.Duration) *AddonListByOrganizationParams {
	var ()
	return &AddonListByOrganizationParams{

		timeout: timeout,
	}
}

// NewAddonListByOrganizationParamsWithContext creates a new AddonListByOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddonListByOrganizationParamsWithContext(ctx context.Context) *AddonListByOrganizationParams {
	var ()
	return &AddonListByOrganizationParams{

		Context: ctx,
	}
}

// NewAddonListByOrganizationParamsWithHTTPClient creates a new AddonListByOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddonListByOrganizationParamsWithHTTPClient(client *http.Client) *AddonListByOrganizationParams {
	var ()
	return &AddonListByOrganizationParams{
		HTTPClient: client,
	}
}

/*AddonListByOrganizationParams contains all the parameters to send to the API endpoint
for the addon list by organization operation typically these are written to a http.Request
*/
type AddonListByOrganizationParams struct {

	/*OrganizationSlug
	  Organization slug

	*/
	OrganizationSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the addon list by organization params
func (o *AddonListByOrganizationParams) WithTimeout(timeout time.Duration) *AddonListByOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the addon list by organization params
func (o *AddonListByOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the addon list by organization params
func (o *AddonListByOrganizationParams) WithContext(ctx context.Context) *AddonListByOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the addon list by organization params
func (o *AddonListByOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the addon list by organization params
func (o *AddonListByOrganizationParams) WithHTTPClient(client *http.Client) *AddonListByOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the addon list by organization params
func (o *AddonListByOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationSlug adds the organizationSlug to the addon list by organization params
func (o *AddonListByOrganizationParams) WithOrganizationSlug(organizationSlug string) *AddonListByOrganizationParams {
	o.SetOrganizationSlug(organizationSlug)
	return o
}

// SetOrganizationSlug adds the organizationSlug to the addon list by organization params
func (o *AddonListByOrganizationParams) SetOrganizationSlug(organizationSlug string) {
	o.OrganizationSlug = organizationSlug
}

// WriteToRequest writes these params to a swagger request
func (o *AddonListByOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization-slug
	if err := r.SetPathParam("organization-slug", o.OrganizationSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
