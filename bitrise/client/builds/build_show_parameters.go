// Code generated by go-swagger; DO NOT EDIT.

package builds

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

// NewBuildShowParams creates a new BuildShowParams object
// with the default values initialized.
func NewBuildShowParams() *BuildShowParams {
	var ()
	return &BuildShowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBuildShowParamsWithTimeout creates a new BuildShowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBuildShowParamsWithTimeout(timeout time.Duration) *BuildShowParams {
	var ()
	return &BuildShowParams{

		timeout: timeout,
	}
}

// NewBuildShowParamsWithContext creates a new BuildShowParams object
// with the default values initialized, and the ability to set a context for a request
func NewBuildShowParamsWithContext(ctx context.Context) *BuildShowParams {
	var ()
	return &BuildShowParams{

		Context: ctx,
	}
}

// NewBuildShowParamsWithHTTPClient creates a new BuildShowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBuildShowParamsWithHTTPClient(client *http.Client) *BuildShowParams {
	var ()
	return &BuildShowParams{
		HTTPClient: client,
	}
}

/*BuildShowParams contains all the parameters to send to the API endpoint
for the build show operation typically these are written to a http.Request
*/
type BuildShowParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*BuildSlug
	  Build slug

	*/
	BuildSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the build show params
func (o *BuildShowParams) WithTimeout(timeout time.Duration) *BuildShowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build show params
func (o *BuildShowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build show params
func (o *BuildShowParams) WithContext(ctx context.Context) *BuildShowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build show params
func (o *BuildShowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build show params
func (o *BuildShowParams) WithHTTPClient(client *http.Client) *BuildShowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build show params
func (o *BuildShowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build show params
func (o *BuildShowParams) WithAppSlug(appSlug string) *BuildShowParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build show params
func (o *BuildShowParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithBuildSlug adds the buildSlug to the build show params
func (o *BuildShowParams) WithBuildSlug(buildSlug string) *BuildShowParams {
	o.SetBuildSlug(buildSlug)
	return o
}

// SetBuildSlug adds the buildSlug to the build show params
func (o *BuildShowParams) SetBuildSlug(buildSlug string) {
	o.BuildSlug = buildSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildShowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	// path param build-slug
	if err := r.SetPathParam("build-slug", o.BuildSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}