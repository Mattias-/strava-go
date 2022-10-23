// Code generated by go-swagger; DO NOT EDIT.

package clubs

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

// NewGetClubActivitiesByIDParams creates a new GetClubActivitiesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClubActivitiesByIDParams() *GetClubActivitiesByIDParams {
	return &GetClubActivitiesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClubActivitiesByIDParamsWithTimeout creates a new GetClubActivitiesByIDParams object
// with the ability to set a timeout on a request.
func NewGetClubActivitiesByIDParamsWithTimeout(timeout time.Duration) *GetClubActivitiesByIDParams {
	return &GetClubActivitiesByIDParams{
		timeout: timeout,
	}
}

// NewGetClubActivitiesByIDParamsWithContext creates a new GetClubActivitiesByIDParams object
// with the ability to set a context for a request.
func NewGetClubActivitiesByIDParamsWithContext(ctx context.Context) *GetClubActivitiesByIDParams {
	return &GetClubActivitiesByIDParams{
		Context: ctx,
	}
}

// NewGetClubActivitiesByIDParamsWithHTTPClient creates a new GetClubActivitiesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClubActivitiesByIDParamsWithHTTPClient(client *http.Client) *GetClubActivitiesByIDParams {
	return &GetClubActivitiesByIDParams{
		HTTPClient: client,
	}
}

/*
GetClubActivitiesByIDParams contains all the parameters to send to the API endpoint

	for the get club activities by Id operation.

	Typically these are written to a http.Request.
*/
type GetClubActivitiesByIDParams struct {

	/* ID.

	   The identifier of the club.

	   Format: int64
	*/
	ID int64

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

// WithDefaults hydrates default values in the get club activities by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClubActivitiesByIDParams) WithDefaults() *GetClubActivitiesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get club activities by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClubActivitiesByIDParams) SetDefaults() {
	var (
		perPageDefault = int64(30)
	)

	val := GetClubActivitiesByIDParams{
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) WithTimeout(timeout time.Duration) *GetClubActivitiesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) WithContext(ctx context.Context) *GetClubActivitiesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) WithHTTPClient(client *http.Client) *GetClubActivitiesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) WithID(id int64) *GetClubActivitiesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) SetID(id int64) {
	o.ID = id
}

// WithPage adds the page to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) WithPage(page *int64) *GetClubActivitiesByIDParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) WithPerPage(perPage *int64) *GetClubActivitiesByIDParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get club activities by Id params
func (o *GetClubActivitiesByIDParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetClubActivitiesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
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
