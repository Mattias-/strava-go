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

// NewGetZonesByActivityIDParams creates a new GetZonesByActivityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetZonesByActivityIDParams() *GetZonesByActivityIDParams {
	return &GetZonesByActivityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetZonesByActivityIDParamsWithTimeout creates a new GetZonesByActivityIDParams object
// with the ability to set a timeout on a request.
func NewGetZonesByActivityIDParamsWithTimeout(timeout time.Duration) *GetZonesByActivityIDParams {
	return &GetZonesByActivityIDParams{
		timeout: timeout,
	}
}

// NewGetZonesByActivityIDParamsWithContext creates a new GetZonesByActivityIDParams object
// with the ability to set a context for a request.
func NewGetZonesByActivityIDParamsWithContext(ctx context.Context) *GetZonesByActivityIDParams {
	return &GetZonesByActivityIDParams{
		Context: ctx,
	}
}

// NewGetZonesByActivityIDParamsWithHTTPClient creates a new GetZonesByActivityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetZonesByActivityIDParamsWithHTTPClient(client *http.Client) *GetZonesByActivityIDParams {
	return &GetZonesByActivityIDParams{
		HTTPClient: client,
	}
}

/*
GetZonesByActivityIDParams contains all the parameters to send to the API endpoint

	for the get zones by activity Id operation.

	Typically these are written to a http.Request.
*/
type GetZonesByActivityIDParams struct {

	/* ID.

	   The identifier of the activity.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get zones by activity Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZonesByActivityIDParams) WithDefaults() *GetZonesByActivityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get zones by activity Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZonesByActivityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get zones by activity Id params
func (o *GetZonesByActivityIDParams) WithTimeout(timeout time.Duration) *GetZonesByActivityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get zones by activity Id params
func (o *GetZonesByActivityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get zones by activity Id params
func (o *GetZonesByActivityIDParams) WithContext(ctx context.Context) *GetZonesByActivityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get zones by activity Id params
func (o *GetZonesByActivityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get zones by activity Id params
func (o *GetZonesByActivityIDParams) WithHTTPClient(client *http.Client) *GetZonesByActivityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get zones by activity Id params
func (o *GetZonesByActivityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get zones by activity Id params
func (o *GetZonesByActivityIDParams) WithID(id int64) *GetZonesByActivityIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get zones by activity Id params
func (o *GetZonesByActivityIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetZonesByActivityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
