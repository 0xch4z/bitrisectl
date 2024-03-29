// Code generated by go-swagger; DO NOT EDIT.

package provisioning_profile

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

// NewProvisioningProfileShowParams creates a new ProvisioningProfileShowParams object
// with the default values initialized.
func NewProvisioningProfileShowParams() *ProvisioningProfileShowParams {
	var ()
	return &ProvisioningProfileShowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProvisioningProfileShowParamsWithTimeout creates a new ProvisioningProfileShowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProvisioningProfileShowParamsWithTimeout(timeout time.Duration) *ProvisioningProfileShowParams {
	var ()
	return &ProvisioningProfileShowParams{

		timeout: timeout,
	}
}

// NewProvisioningProfileShowParamsWithContext creates a new ProvisioningProfileShowParams object
// with the default values initialized, and the ability to set a context for a request
func NewProvisioningProfileShowParamsWithContext(ctx context.Context) *ProvisioningProfileShowParams {
	var ()
	return &ProvisioningProfileShowParams{

		Context: ctx,
	}
}

// NewProvisioningProfileShowParamsWithHTTPClient creates a new ProvisioningProfileShowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProvisioningProfileShowParamsWithHTTPClient(client *http.Client) *ProvisioningProfileShowParams {
	var ()
	return &ProvisioningProfileShowParams{
		HTTPClient: client,
	}
}

/*ProvisioningProfileShowParams contains all the parameters to send to the API endpoint
for the provisioning profile show operation typically these are written to a http.Request
*/
type ProvisioningProfileShowParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*ProvisioningProfileSlug
	  Provisioning profile slug

	*/
	ProvisioningProfileSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the provisioning profile show params
func (o *ProvisioningProfileShowParams) WithTimeout(timeout time.Duration) *ProvisioningProfileShowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provisioning profile show params
func (o *ProvisioningProfileShowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provisioning profile show params
func (o *ProvisioningProfileShowParams) WithContext(ctx context.Context) *ProvisioningProfileShowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provisioning profile show params
func (o *ProvisioningProfileShowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provisioning profile show params
func (o *ProvisioningProfileShowParams) WithHTTPClient(client *http.Client) *ProvisioningProfileShowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provisioning profile show params
func (o *ProvisioningProfileShowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the provisioning profile show params
func (o *ProvisioningProfileShowParams) WithAppSlug(appSlug string) *ProvisioningProfileShowParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the provisioning profile show params
func (o *ProvisioningProfileShowParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithProvisioningProfileSlug adds the provisioningProfileSlug to the provisioning profile show params
func (o *ProvisioningProfileShowParams) WithProvisioningProfileSlug(provisioningProfileSlug string) *ProvisioningProfileShowParams {
	o.SetProvisioningProfileSlug(provisioningProfileSlug)
	return o
}

// SetProvisioningProfileSlug adds the provisioningProfileSlug to the provisioning profile show params
func (o *ProvisioningProfileShowParams) SetProvisioningProfileSlug(provisioningProfileSlug string) {
	o.ProvisioningProfileSlug = provisioningProfileSlug
}

// WriteToRequest writes these params to a swagger request
func (o *ProvisioningProfileShowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	// path param provisioning-profile-slug
	if err := r.SetPathParam("provisioning-profile-slug", o.ProvisioningProfileSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
