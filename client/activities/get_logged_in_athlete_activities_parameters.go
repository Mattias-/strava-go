// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetLoggedInAthleteActivitiesParams creates a new GetLoggedInAthleteActivitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoggedInAthleteActivitiesParams() *GetLoggedInAthleteActivitiesParams {
	return &GetLoggedInAthleteActivitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoggedInAthleteActivitiesParamsWithTimeout creates a new GetLoggedInAthleteActivitiesParams object
// with the ability to set a timeout on a request.
func NewGetLoggedInAthleteActivitiesParamsWithTimeout(timeout time.Duration) *GetLoggedInAthleteActivitiesParams {
	return &GetLoggedInAthleteActivitiesParams{
		timeout: timeout,
	}
}

// NewGetLoggedInAthleteActivitiesParamsWithContext creates a new GetLoggedInAthleteActivitiesParams object
// with the ability to set a context for a request.
func NewGetLoggedInAthleteActivitiesParamsWithContext(ctx context.Context) *GetLoggedInAthleteActivitiesParams {
	return &GetLoggedInAthleteActivitiesParams{
		Context: ctx,
	}
}

// NewGetLoggedInAthleteActivitiesParamsWithHTTPClient creates a new GetLoggedInAthleteActivitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoggedInAthleteActivitiesParamsWithHTTPClient(client *http.Client) *GetLoggedInAthleteActivitiesParams {
	return &GetLoggedInAthleteActivitiesParams{
		HTTPClient: client,
	}
}

/*
GetLoggedInAthleteActivitiesParams contains all the parameters to send to the API endpoint

	for the get logged in athlete activities operation.

	Typically these are written to a http.Request.
*/
type GetLoggedInAthleteActivitiesParams struct {

	/* After.

	   An epoch timestamp to use for filtering activities that have taken place after a certain time.
	*/
	After *int64

	/* Before.

	   An epoch timestamp to use for filtering activities that have taken place before a certain time.
	*/
	Before *int64

	/* Page.

	   Page number. Defaults to 1.
	*/
	Page *int64

	/* PerPage.

	   Number of items per page. Defaults to 30.

	   Default: 30
	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logged in athlete activities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggedInAthleteActivitiesParams) WithDefaults() *GetLoggedInAthleteActivitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logged in athlete activities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggedInAthleteActivitiesParams) SetDefaults() {
	var (
		perPageDefault = int64(30)
	)

	val := GetLoggedInAthleteActivitiesParams{
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) WithTimeout(timeout time.Duration) *GetLoggedInAthleteActivitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) WithContext(ctx context.Context) *GetLoggedInAthleteActivitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) WithHTTPClient(client *http.Client) *GetLoggedInAthleteActivitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) WithAfter(after *int64) *GetLoggedInAthleteActivitiesParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) SetAfter(after *int64) {
	o.After = after
}

// WithBefore adds the before to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) WithBefore(before *int64) *GetLoggedInAthleteActivitiesParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) SetBefore(before *int64) {
	o.Before = before
}

// WithPage adds the page to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) WithPage(page *int64) *GetLoggedInAthleteActivitiesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) WithPerPage(perPage *int64) *GetLoggedInAthleteActivitiesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get logged in athlete activities params
func (o *GetLoggedInAthleteActivitiesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggedInAthleteActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter int64

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := swag.FormatInt64(qrAfter)
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Before != nil {

		// query param before
		var qrBefore int64

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := swag.FormatInt64(qrBefore)
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
