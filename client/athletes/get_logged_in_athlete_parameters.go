// Code generated by go-swagger; DO NOT EDIT.

package athletes

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

// NewGetLoggedInAthleteParams creates a new GetLoggedInAthleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoggedInAthleteParams() *GetLoggedInAthleteParams {
	return &GetLoggedInAthleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoggedInAthleteParamsWithTimeout creates a new GetLoggedInAthleteParams object
// with the ability to set a timeout on a request.
func NewGetLoggedInAthleteParamsWithTimeout(timeout time.Duration) *GetLoggedInAthleteParams {
	return &GetLoggedInAthleteParams{
		timeout: timeout,
	}
}

// NewGetLoggedInAthleteParamsWithContext creates a new GetLoggedInAthleteParams object
// with the ability to set a context for a request.
func NewGetLoggedInAthleteParamsWithContext(ctx context.Context) *GetLoggedInAthleteParams {
	return &GetLoggedInAthleteParams{
		Context: ctx,
	}
}

// NewGetLoggedInAthleteParamsWithHTTPClient creates a new GetLoggedInAthleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoggedInAthleteParamsWithHTTPClient(client *http.Client) *GetLoggedInAthleteParams {
	return &GetLoggedInAthleteParams{
		HTTPClient: client,
	}
}

/*
GetLoggedInAthleteParams contains all the parameters to send to the API endpoint

	for the get logged in athlete operation.

	Typically these are written to a http.Request.
*/
type GetLoggedInAthleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logged in athlete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggedInAthleteParams) WithDefaults() *GetLoggedInAthleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logged in athlete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggedInAthleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get logged in athlete params
func (o *GetLoggedInAthleteParams) WithTimeout(timeout time.Duration) *GetLoggedInAthleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logged in athlete params
func (o *GetLoggedInAthleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logged in athlete params
func (o *GetLoggedInAthleteParams) WithContext(ctx context.Context) *GetLoggedInAthleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logged in athlete params
func (o *GetLoggedInAthleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logged in athlete params
func (o *GetLoggedInAthleteParams) WithHTTPClient(client *http.Client) *GetLoggedInAthleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logged in athlete params
func (o *GetLoggedInAthleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggedInAthleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
