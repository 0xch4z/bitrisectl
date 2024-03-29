// Code generated by go-swagger; DO NOT EDIT.

package generic_project_file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGenericProjectFileListParams creates a new GenericProjectFileListParams object
// with the default values initialized.
func NewGenericProjectFileListParams() *GenericProjectFileListParams {
	var ()
	return &GenericProjectFileListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGenericProjectFileListParamsWithTimeout creates a new GenericProjectFileListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGenericProjectFileListParamsWithTimeout(timeout time.Duration) *GenericProjectFileListParams {
	var ()
	return &GenericProjectFileListParams{

		timeout: timeout,
	}
}

// NewGenericProjectFileListParamsWithContext creates a new GenericProjectFileListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGenericProjectFileListParamsWithContext(ctx context.Context) *GenericProjectFileListParams {
	var ()
	return &GenericProjectFileListParams{

		Context: ctx,
	}
}

// NewGenericProjectFileListParamsWithHTTPClient creates a new GenericProjectFileListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGenericProjectFileListParamsWithHTTPClient(client *http.Client) *GenericProjectFileListParams {
	var ()
	return &GenericProjectFileListParams{
		HTTPClient: client,
	}
}

/*GenericProjectFileListParams contains all the parameters to send to the API endpoint
for the generic project file list operation typically these are written to a http.Request
*/
type GenericProjectFileListParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*Limit
	  Max number of build certificates per page is 50.

	*/
	Limit *int64
	/*Next
	  Slug of the first generic project file in the response

	*/
	Next *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the generic project file list params
func (o *GenericProjectFileListParams) WithTimeout(timeout time.Duration) *GenericProjectFileListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generic project file list params
func (o *GenericProjectFileListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generic project file list params
func (o *GenericProjectFileListParams) WithContext(ctx context.Context) *GenericProjectFileListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generic project file list params
func (o *GenericProjectFileListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generic project file list params
func (o *GenericProjectFileListParams) WithHTTPClient(client *http.Client) *GenericProjectFileListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generic project file list params
func (o *GenericProjectFileListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the generic project file list params
func (o *GenericProjectFileListParams) WithAppSlug(appSlug string) *GenericProjectFileListParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the generic project file list params
func (o *GenericProjectFileListParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithLimit adds the limit to the generic project file list params
func (o *GenericProjectFileListParams) WithLimit(limit *int64) *GenericProjectFileListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the generic project file list params
func (o *GenericProjectFileListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the generic project file list params
func (o *GenericProjectFileListParams) WithNext(next *string) *GenericProjectFileListParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the generic project file list params
func (o *GenericProjectFileListParams) SetNext(next *string) {
	o.Next = next
}

// WriteToRequest writes these params to a swagger request
func (o *GenericProjectFileListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Next != nil {

		// query param next
		var qrNext string
		if o.Next != nil {
			qrNext = *o.Next
		}
		qNext := qrNext
		if qNext != "" {
			if err := r.SetQueryParam("next", qNext); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
