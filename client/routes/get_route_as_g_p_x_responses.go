// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// GetRouteAsGPXReader is a Reader for the GetRouteAsGPX structure.
type GetRouteAsGPXReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRouteAsGPXReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRouteAsGPXOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRouteAsGPXDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRouteAsGPXOK creates a GetRouteAsGPXOK with default headers values
func NewGetRouteAsGPXOK() *GetRouteAsGPXOK {
	return &GetRouteAsGPXOK{}
}

/* GetRouteAsGPXOK describes a response with status code 200, with default header values.

A GPX file with the route.
*/
type GetRouteAsGPXOK struct {
}

func (o *GetRouteAsGPXOK) Error() string {
	return fmt.Sprintf("[GET /routes/{id}/export_gpx][%d] getRouteAsGPXOK ", 200)
}

func (o *GetRouteAsGPXOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRouteAsGPXDefault creates a GetRouteAsGPXDefault with default headers values
func NewGetRouteAsGPXDefault(code int) *GetRouteAsGPXDefault {
	return &GetRouteAsGPXDefault{
		_statusCode: code,
	}
}

/* GetRouteAsGPXDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetRouteAsGPXDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get route as g p x default response
func (o *GetRouteAsGPXDefault) Code() int {
	return o._statusCode
}

func (o *GetRouteAsGPXDefault) Error() string {
	return fmt.Sprintf("[GET /routes/{id}/export_gpx][%d] getRouteAsGPX default  %+v", o._statusCode, o.Payload)
}
func (o *GetRouteAsGPXDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetRouteAsGPXDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
