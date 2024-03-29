// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewAppListByUserParams creates a new AppListByUserParams object
// with the default values initialized.
func NewAppListByUserParams() *AppListByUserParams {
	var ()
	return &AppListByUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppListByUserParamsWithTimeout creates a new AppListByUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppListByUserParamsWithTimeout(timeout time.Duration) *AppListByUserParams {
	var ()
	return &AppListByUserParams{

		timeout: timeout,
	}
}

// NewAppListByUserParamsWithContext creates a new AppListByUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppListByUserParamsWithContext(ctx context.Context) *AppListByUserParams {
	var ()
	return &AppListByUserParams{

		Context: ctx,
	}
}

// NewAppListByUserParamsWithHTTPClient creates a new AppListByUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppListByUserParamsWithHTTPClient(client *http.Client) *AppListByUserParams {
	var ()
	return &AppListByUserParams{
		HTTPClient: client,
	}
}

/*AppListByUserParams contains all the parameters to send to the API endpoint
for the app list by user operation typically these are written to a http.Request
*/
type AppListByUserParams struct {

	/*Limit
	  Max number of elements per page (default: 50)

	*/
	Limit *int64
	/*Next
	  Slug of the first app in the response

	*/
	Next *string
	/*SortBy
	  Order of applications

	*/
	SortBy *string
	/*UserSlug
	  User slug

	*/
	UserSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the app list by user params
func (o *AppListByUserParams) WithTimeout(timeout time.Duration) *AppListByUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app list by user params
func (o *AppListByUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app list by user params
func (o *AppListByUserParams) WithContext(ctx context.Context) *AppListByUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app list by user params
func (o *AppListByUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app list by user params
func (o *AppListByUserParams) WithHTTPClient(client *http.Client) *AppListByUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app list by user params
func (o *AppListByUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the app list by user params
func (o *AppListByUserParams) WithLimit(limit *int64) *AppListByUserParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the app list by user params
func (o *AppListByUserParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the app list by user params
func (o *AppListByUserParams) WithNext(next *string) *AppListByUserParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the app list by user params
func (o *AppListByUserParams) SetNext(next *string) {
	o.Next = next
}

// WithSortBy adds the sortBy to the app list by user params
func (o *AppListByUserParams) WithSortBy(sortBy *string) *AppListByUserParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the app list by user params
func (o *AppListByUserParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithUserSlug adds the userSlug to the app list by user params
func (o *AppListByUserParams) WithUserSlug(userSlug string) *AppListByUserParams {
	o.SetUserSlug(userSlug)
	return o
}

// SetUserSlug adds the userSlug to the app list by user params
func (o *AppListByUserParams) SetUserSlug(userSlug string) {
	o.UserSlug = userSlug
}

// WriteToRequest writes these params to a swagger request
func (o *AppListByUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.SortBy != nil {

		// query param sort_by
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sort_by", qSortBy); err != nil {
				return err
			}
		}

	}

	// path param user-slug
	if err := r.SetPathParam("user-slug", o.UserSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
