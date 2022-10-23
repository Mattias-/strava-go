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
	"github.com/go-openapi/swag"
)

// NewUpdateLoggedInAthleteParams creates a new UpdateLoggedInAthleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLoggedInAthleteParams() *UpdateLoggedInAthleteParams {
	return &UpdateLoggedInAthleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLoggedInAthleteParamsWithTimeout creates a new UpdateLoggedInAthleteParams object
// with the ability to set a timeout on a request.
func NewUpdateLoggedInAthleteParamsWithTimeout(timeout time.Duration) *UpdateLoggedInAthleteParams {
	return &UpdateLoggedInAthleteParams{
		timeout: timeout,
	}
}

// NewUpdateLoggedInAthleteParamsWithContext creates a new UpdateLoggedInAthleteParams object
// with the ability to set a context for a request.
func NewUpdateLoggedInAthleteParamsWithContext(ctx context.Context) *UpdateLoggedInAthleteParams {
	return &UpdateLoggedInAthleteParams{
		Context: ctx,
	}
}

// NewUpdateLoggedInAthleteParamsWithHTTPClient creates a new UpdateLoggedInAthleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLoggedInAthleteParamsWithHTTPClient(client *http.Client) *UpdateLoggedInAthleteParams {
	return &UpdateLoggedInAthleteParams{
		HTTPClient: client,
	}
}

/*
UpdateLoggedInAthleteParams contains all the parameters to send to the API endpoint

	for the update logged in athlete operation.

	Typically these are written to a http.Request.
*/
type UpdateLoggedInAthleteParams struct {

	/* Weight.

	   The weight of the athlete in kilograms.

	   Format: float
	*/
	Weight float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update logged in athlete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLoggedInAthleteParams) WithDefaults() *UpdateLoggedInAthleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update logged in athlete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLoggedInAthleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update logged in athlete params
func (o *UpdateLoggedInAthleteParams) WithTimeout(timeout time.Duration) *UpdateLoggedInAthleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update logged in athlete params
func (o *UpdateLoggedInAthleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update logged in athlete params
func (o *UpdateLoggedInAthleteParams) WithContext(ctx context.Context) *UpdateLoggedInAthleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update logged in athlete params
func (o *UpdateLoggedInAthleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update logged in athlete params
func (o *UpdateLoggedInAthleteParams) WithHTTPClient(client *http.Client) *UpdateLoggedInAthleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update logged in athlete params
func (o *UpdateLoggedInAthleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWeight adds the weight to the update logged in athlete params
func (o *UpdateLoggedInAthleteParams) WithWeight(weight float32) *UpdateLoggedInAthleteParams {
	o.SetWeight(weight)
	return o
}

// SetWeight adds the weight to the update logged in athlete params
func (o *UpdateLoggedInAthleteParams) SetWeight(weight float32) {
	o.Weight = weight
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLoggedInAthleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param weight
	if err := r.SetPathParam("weight", swag.FormatFloat32(o.Weight)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
