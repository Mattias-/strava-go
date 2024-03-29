// Code generated by go-swagger; DO NOT EDIT.

package streams

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

// NewGetSegmentStreamsParams creates a new GetSegmentStreamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSegmentStreamsParams() *GetSegmentStreamsParams {
	return &GetSegmentStreamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSegmentStreamsParamsWithTimeout creates a new GetSegmentStreamsParams object
// with the ability to set a timeout on a request.
func NewGetSegmentStreamsParamsWithTimeout(timeout time.Duration) *GetSegmentStreamsParams {
	return &GetSegmentStreamsParams{
		timeout: timeout,
	}
}

// NewGetSegmentStreamsParamsWithContext creates a new GetSegmentStreamsParams object
// with the ability to set a context for a request.
func NewGetSegmentStreamsParamsWithContext(ctx context.Context) *GetSegmentStreamsParams {
	return &GetSegmentStreamsParams{
		Context: ctx,
	}
}

// NewGetSegmentStreamsParamsWithHTTPClient creates a new GetSegmentStreamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSegmentStreamsParamsWithHTTPClient(client *http.Client) *GetSegmentStreamsParams {
	return &GetSegmentStreamsParams{
		HTTPClient: client,
	}
}

/*
GetSegmentStreamsParams contains all the parameters to send to the API endpoint

	for the get segment streams operation.

	Typically these are written to a http.Request.
*/
type GetSegmentStreamsParams struct {

	/* ID.

	   The identifier of the segment.

	   Format: int64
	*/
	ID int64

	/* KeyByType.

	   Must be true.

	   Default: true
	*/
	KeyByType bool

	/* Keys.

	   The types of streams to return.
	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get segment streams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSegmentStreamsParams) WithDefaults() *GetSegmentStreamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get segment streams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSegmentStreamsParams) SetDefaults() {
	var (
		keyByTypeDefault = bool(true)
	)

	val := GetSegmentStreamsParams{
		KeyByType: keyByTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get segment streams params
func (o *GetSegmentStreamsParams) WithTimeout(timeout time.Duration) *GetSegmentStreamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get segment streams params
func (o *GetSegmentStreamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get segment streams params
func (o *GetSegmentStreamsParams) WithContext(ctx context.Context) *GetSegmentStreamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get segment streams params
func (o *GetSegmentStreamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get segment streams params
func (o *GetSegmentStreamsParams) WithHTTPClient(client *http.Client) *GetSegmentStreamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get segment streams params
func (o *GetSegmentStreamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get segment streams params
func (o *GetSegmentStreamsParams) WithID(id int64) *GetSegmentStreamsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get segment streams params
func (o *GetSegmentStreamsParams) SetID(id int64) {
	o.ID = id
}

// WithKeyByType adds the keyByType to the get segment streams params
func (o *GetSegmentStreamsParams) WithKeyByType(keyByType bool) *GetSegmentStreamsParams {
	o.SetKeyByType(keyByType)
	return o
}

// SetKeyByType adds the keyByType to the get segment streams params
func (o *GetSegmentStreamsParams) SetKeyByType(keyByType bool) {
	o.KeyByType = keyByType
}

// WithKeys adds the keys to the get segment streams params
func (o *GetSegmentStreamsParams) WithKeys(keys []string) *GetSegmentStreamsParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the get segment streams params
func (o *GetSegmentStreamsParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *GetSegmentStreamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// query param key_by_type
	qrKeyByType := o.KeyByType
	qKeyByType := swag.FormatBool(qrKeyByType)
	if qKeyByType != "" {

		if err := r.SetQueryParam("key_by_type", qKeyByType); err != nil {
			return err
		}
	}

	if o.Keys != nil {

		// binding items for keys
		joinedKeys := o.bindParamKeys(reg)

		// query array param keys
		if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetSegmentStreams binds the parameter keys
func (o *GetSegmentStreamsParams) bindParamKeys(formats strfmt.Registry) []string {
	keysIR := o.Keys

	var keysIC []string
	for _, keysIIR := range keysIR { // explode []string

		keysIIV := keysIIR // string as string
		keysIC = append(keysIC, keysIIV)
	}

	// items.CollectionFormat: "csv"
	keysIS := swag.JoinByFormat(keysIC, "csv")

	return keysIS
}
