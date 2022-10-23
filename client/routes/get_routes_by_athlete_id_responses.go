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

// GetRoutesByAthleteIDReader is a Reader for the GetRoutesByAthleteID structure.
type GetRoutesByAthleteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutesByAthleteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutesByAthleteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRoutesByAthleteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoutesByAthleteIDOK creates a GetRoutesByAthleteIDOK with default headers values
func NewGetRoutesByAthleteIDOK() *GetRoutesByAthleteIDOK {
	return &GetRoutesByAthleteIDOK{}
}

/*
GetRoutesByAthleteIDOK describes a response with status code 200, with default header values.

A representation of the route.
*/
type GetRoutesByAthleteIDOK struct {
	Payload []*models.Route
}

// IsSuccess returns true when this get routes by athlete Id o k response has a 2xx status code
func (o *GetRoutesByAthleteIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routes by athlete Id o k response has a 3xx status code
func (o *GetRoutesByAthleteIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routes by athlete Id o k response has a 4xx status code
func (o *GetRoutesByAthleteIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routes by athlete Id o k response has a 5xx status code
func (o *GetRoutesByAthleteIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routes by athlete Id o k response a status code equal to that given
func (o *GetRoutesByAthleteIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutesByAthleteIDOK) Error() string {
	return fmt.Sprintf("[GET /athletes/{id}/routes][%d] getRoutesByAthleteIdOK  %+v", 200, o.Payload)
}

func (o *GetRoutesByAthleteIDOK) String() string {
	return fmt.Sprintf("[GET /athletes/{id}/routes][%d] getRoutesByAthleteIdOK  %+v", 200, o.Payload)
}

func (o *GetRoutesByAthleteIDOK) GetPayload() []*models.Route {
	return o.Payload
}

func (o *GetRoutesByAthleteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutesByAthleteIDDefault creates a GetRoutesByAthleteIDDefault with default headers values
func NewGetRoutesByAthleteIDDefault(code int) *GetRoutesByAthleteIDDefault {
	return &GetRoutesByAthleteIDDefault{
		_statusCode: code,
	}
}

/*
GetRoutesByAthleteIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetRoutesByAthleteIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get routes by athlete Id default response
func (o *GetRoutesByAthleteIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get routes by athlete Id default response has a 2xx status code
func (o *GetRoutesByAthleteIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get routes by athlete Id default response has a 3xx status code
func (o *GetRoutesByAthleteIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get routes by athlete Id default response has a 4xx status code
func (o *GetRoutesByAthleteIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get routes by athlete Id default response has a 5xx status code
func (o *GetRoutesByAthleteIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get routes by athlete Id default response a status code equal to that given
func (o *GetRoutesByAthleteIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetRoutesByAthleteIDDefault) Error() string {
	return fmt.Sprintf("[GET /athletes/{id}/routes][%d] getRoutesByAthleteId default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoutesByAthleteIDDefault) String() string {
	return fmt.Sprintf("[GET /athletes/{id}/routes][%d] getRoutesByAthleteId default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoutesByAthleteIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetRoutesByAthleteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
