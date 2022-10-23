// Code generated by go-swagger; DO NOT EDIT.

package segments

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

// NewGetLoggedInAthleteStarredSegmentsParams creates a new GetLoggedInAthleteStarredSegmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoggedInAthleteStarredSegmentsParams() *GetLoggedInAthleteStarredSegmentsParams {
	return &GetLoggedInAthleteStarredSegmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoggedInAthleteStarredSegmentsParamsWithTimeout creates a new GetLoggedInAthleteStarredSegmentsParams object
// with the ability to set a timeout on a request.
func NewGetLoggedInAthleteStarredSegmentsParamsWithTimeout(timeout time.Duration) *GetLoggedInAthleteStarredSegmentsParams {
	return &GetLoggedInAthleteStarredSegmentsParams{
		timeout: timeout,
	}
}

// NewGetLoggedInAthleteStarredSegmentsParamsWithContext creates a new GetLoggedInAthleteStarredSegmentsParams object
// with the ability to set a context for a request.
func NewGetLoggedInAthleteStarredSegmentsParamsWithContext(ctx context.Context) *GetLoggedInAthleteStarredSegmentsParams {
	return &GetLoggedInAthleteStarredSegmentsParams{
		Context: ctx,
	}
}

// NewGetLoggedInAthleteStarredSegmentsParamsWithHTTPClient creates a new GetLoggedInAthleteStarredSegmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoggedInAthleteStarredSegmentsParamsWithHTTPClient(client *http.Client) *GetLoggedInAthleteStarredSegmentsParams {
	return &GetLoggedInAthleteStarredSegmentsParams{
		HTTPClient: client,
	}
}

/*
GetLoggedInAthleteStarredSegmentsParams contains all the parameters to send to the API endpoint

	for the get logged in athlete starred segments operation.

	Typically these are written to a http.Request.
*/
type GetLoggedInAthleteStarredSegmentsParams struct {

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

// WithDefaults hydrates default values in the get logged in athlete starred segments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggedInAthleteStarredSegmentsParams) WithDefaults() *GetLoggedInAthleteStarredSegmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logged in athlete starred segments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggedInAthleteStarredSegmentsParams) SetDefaults() {
	var (
		perPageDefault = int64(30)
	)

	val := GetLoggedInAthleteStarredSegmentsParams{
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) WithTimeout(timeout time.Duration) *GetLoggedInAthleteStarredSegmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) WithContext(ctx context.Context) *GetLoggedInAthleteStarredSegmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) WithHTTPClient(client *http.Client) *GetLoggedInAthleteStarredSegmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) WithPage(page *int64) *GetLoggedInAthleteStarredSegmentsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) WithPerPage(perPage *int64) *GetLoggedInAthleteStarredSegmentsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get logged in athlete starred segments params
func (o *GetLoggedInAthleteStarredSegmentsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggedInAthleteStarredSegmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
