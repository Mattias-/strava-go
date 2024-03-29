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

// NewGetActivityByIDParams creates a new GetActivityByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetActivityByIDParams() *GetActivityByIDParams {
	return &GetActivityByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetActivityByIDParamsWithTimeout creates a new GetActivityByIDParams object
// with the ability to set a timeout on a request.
func NewGetActivityByIDParamsWithTimeout(timeout time.Duration) *GetActivityByIDParams {
	return &GetActivityByIDParams{
		timeout: timeout,
	}
}

// NewGetActivityByIDParamsWithContext creates a new GetActivityByIDParams object
// with the ability to set a context for a request.
func NewGetActivityByIDParamsWithContext(ctx context.Context) *GetActivityByIDParams {
	return &GetActivityByIDParams{
		Context: ctx,
	}
}

// NewGetActivityByIDParamsWithHTTPClient creates a new GetActivityByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetActivityByIDParamsWithHTTPClient(client *http.Client) *GetActivityByIDParams {
	return &GetActivityByIDParams{
		HTTPClient: client,
	}
}

/*
GetActivityByIDParams contains all the parameters to send to the API endpoint

	for the get activity by Id operation.

	Typically these are written to a http.Request.
*/
type GetActivityByIDParams struct {

	/* ID.

	   The identifier of the activity.

	   Format: int64
	*/
	ID int64

	/* IncludeAllEfforts.

	   To include all segments efforts.
	*/
	IncludeAllEfforts *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get activity by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActivityByIDParams) WithDefaults() *GetActivityByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get activity by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActivityByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get activity by Id params
func (o *GetActivityByIDParams) WithTimeout(timeout time.Duration) *GetActivityByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get activity by Id params
func (o *GetActivityByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get activity by Id params
func (o *GetActivityByIDParams) WithContext(ctx context.Context) *GetActivityByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get activity by Id params
func (o *GetActivityByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get activity by Id params
func (o *GetActivityByIDParams) WithHTTPClient(client *http.Client) *GetActivityByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get activity by Id params
func (o *GetActivityByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get activity by Id params
func (o *GetActivityByIDParams) WithID(id int64) *GetActivityByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get activity by Id params
func (o *GetActivityByIDParams) SetID(id int64) {
	o.ID = id
}

// WithIncludeAllEfforts adds the includeAllEfforts to the get activity by Id params
func (o *GetActivityByIDParams) WithIncludeAllEfforts(includeAllEfforts *bool) *GetActivityByIDParams {
	o.SetIncludeAllEfforts(includeAllEfforts)
	return o
}

// SetIncludeAllEfforts adds the includeAllEfforts to the get activity by Id params
func (o *GetActivityByIDParams) SetIncludeAllEfforts(includeAllEfforts *bool) {
	o.IncludeAllEfforts = includeAllEfforts
}

// WriteToRequest writes these params to a swagger request
func (o *GetActivityByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.IncludeAllEfforts != nil {

		// query param include_all_efforts
		var qrIncludeAllEfforts bool

		if o.IncludeAllEfforts != nil {
			qrIncludeAllEfforts = *o.IncludeAllEfforts
		}
		qIncludeAllEfforts := swag.FormatBool(qrIncludeAllEfforts)
		if qIncludeAllEfforts != "" {

			if err := r.SetQueryParam("include_all_efforts", qIncludeAllEfforts); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
