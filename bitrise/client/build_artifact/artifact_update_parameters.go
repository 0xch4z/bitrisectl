// Code generated by go-swagger; DO NOT EDIT.

package build_artifact

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

	models "github.com/charliekenney23/bitrisectl/var/work/bitrise/models"
)

// NewArtifactUpdateParams creates a new ArtifactUpdateParams object
// with the default values initialized.
func NewArtifactUpdateParams() *ArtifactUpdateParams {
	var ()
	return &ArtifactUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewArtifactUpdateParamsWithTimeout creates a new ArtifactUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewArtifactUpdateParamsWithTimeout(timeout time.Duration) *ArtifactUpdateParams {
	var ()
	return &ArtifactUpdateParams{

		timeout: timeout,
	}
}

// NewArtifactUpdateParamsWithContext creates a new ArtifactUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewArtifactUpdateParamsWithContext(ctx context.Context) *ArtifactUpdateParams {
	var ()
	return &ArtifactUpdateParams{

		Context: ctx,
	}
}

// NewArtifactUpdateParamsWithHTTPClient creates a new ArtifactUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewArtifactUpdateParamsWithHTTPClient(client *http.Client) *ArtifactUpdateParams {
	var ()
	return &ArtifactUpdateParams{
		HTTPClient: client,
	}
}

/*ArtifactUpdateParams contains all the parameters to send to the API endpoint
for the artifact update operation typically these are written to a http.Request
*/
type ArtifactUpdateParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*ArtifactParams
	  Artifact parameters

	*/
	ArtifactParams *models.V0ArtifactUpdateParams
	/*ArtifactSlug
	  Artifact slug

	*/
	ArtifactSlug string
	/*BuildSlug
	  Build slug

	*/
	BuildSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the artifact update params
func (o *ArtifactUpdateParams) WithTimeout(timeout time.Duration) *ArtifactUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the artifact update params
func (o *ArtifactUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the artifact update params
func (o *ArtifactUpdateParams) WithContext(ctx context.Context) *ArtifactUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the artifact update params
func (o *ArtifactUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the artifact update params
func (o *ArtifactUpdateParams) WithHTTPClient(client *http.Client) *ArtifactUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the artifact update params
func (o *ArtifactUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the artifact update params
func (o *ArtifactUpdateParams) WithAppSlug(appSlug string) *ArtifactUpdateParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the artifact update params
func (o *ArtifactUpdateParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithArtifactParams adds the artifactParams to the artifact update params
func (o *ArtifactUpdateParams) WithArtifactParams(artifactParams *models.V0ArtifactUpdateParams) *ArtifactUpdateParams {
	o.SetArtifactParams(artifactParams)
	return o
}

// SetArtifactParams adds the artifactParams to the artifact update params
func (o *ArtifactUpdateParams) SetArtifactParams(artifactParams *models.V0ArtifactUpdateParams) {
	o.ArtifactParams = artifactParams
}

// WithArtifactSlug adds the artifactSlug to the artifact update params
func (o *ArtifactUpdateParams) WithArtifactSlug(artifactSlug string) *ArtifactUpdateParams {
	o.SetArtifactSlug(artifactSlug)
	return o
}

// SetArtifactSlug adds the artifactSlug to the artifact update params
func (o *ArtifactUpdateParams) SetArtifactSlug(artifactSlug string) {
	o.ArtifactSlug = artifactSlug
}

// WithBuildSlug adds the buildSlug to the artifact update params
func (o *ArtifactUpdateParams) WithBuildSlug(buildSlug string) *ArtifactUpdateParams {
	o.SetBuildSlug(buildSlug)
	return o
}

// SetBuildSlug adds the buildSlug to the artifact update params
func (o *ArtifactUpdateParams) SetBuildSlug(buildSlug string) {
	o.BuildSlug = buildSlug
}

// WriteToRequest writes these params to a swagger request
func (o *ArtifactUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if o.ArtifactParams != nil {
		if err := r.SetBodyParam(o.ArtifactParams); err != nil {
			return err
		}
	}

	// path param artifact-slug
	if err := r.SetPathParam("artifact-slug", o.ArtifactSlug); err != nil {
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