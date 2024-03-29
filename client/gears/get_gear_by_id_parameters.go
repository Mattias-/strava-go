// Code generated by go-swagger; DO NOT EDIT.

package gears

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
)

// NewGetGearByIDParams creates a new GetGearByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGearByIDParams() *GetGearByIDParams {
	return &GetGearByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGearByIDParamsWithTimeout creates a new GetGearByIDParams object
// with the ability to set a timeout on a request.
func NewGetGearByIDParamsWithTimeout(timeout time.Duration) *GetGearByIDParams {
	return &GetGearByIDParams{
		timeout: timeout,
	}
}

// NewGetGearByIDParamsWithContext creates a new GetGearByIDParams object
// with the ability to set a context for a request.
func NewGetGearByIDParamsWithContext(ctx context.Context) *GetGearByIDParams {
	return &GetGearByIDParams{
		Context: ctx,
	}
}

// NewGetGearByIDParamsWithHTTPClient creates a new GetGearByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGearByIDParamsWithHTTPClient(client *http.Client) *GetGearByIDParams {
	return &GetGearByIDParams{
		HTTPClient: client,
	}
}

/*
GetGearByIDParams contains all the parameters to send to the API endpoint

	for the get gear by Id operation.

	Typically these are written to a http.Request.
*/
type GetGearByIDParams struct {

	/* ID.

	   The identifier of the gear.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gear by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGearByIDParams) WithDefaults() *GetGearByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gear by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGearByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gear by Id params
func (o *GetGearByIDParams) WithTimeout(timeout time.Duration) *GetGearByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gear by Id params
func (o *GetGearByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gear by Id params
func (o *GetGearByIDParams) WithContext(ctx context.Context) *GetGearByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gear by Id params
func (o *GetGearByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gear by Id params
func (o *GetGearByIDParams) WithHTTPClient(client *http.Client) *GetGearByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gear by Id params
func (o *GetGearByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get gear by Id params
func (o *GetGearByIDParams) WithID(id string) *GetGearByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get gear by Id params
func (o *GetGearByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGearByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
