// Code generated by go-swagger; DO NOT EDIT.

package athletes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// UpdateLoggedInAthleteReader is a Reader for the UpdateLoggedInAthlete structure.
type UpdateLoggedInAthleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLoggedInAthleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLoggedInAthleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateLoggedInAthleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateLoggedInAthleteOK creates a UpdateLoggedInAthleteOK with default headers values
func NewUpdateLoggedInAthleteOK() *UpdateLoggedInAthleteOK {
	return &UpdateLoggedInAthleteOK{}
}

/*
UpdateLoggedInAthleteOK describes a response with status code 200, with default header values.

Profile information for the authenticated athlete.
*/
type UpdateLoggedInAthleteOK struct {
	Payload *models.DetailedAthlete
}

// IsSuccess returns true when this update logged in athlete o k response has a 2xx status code
func (o *UpdateLoggedInAthleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update logged in athlete o k response has a 3xx status code
func (o *UpdateLoggedInAthleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update logged in athlete o k response has a 4xx status code
func (o *UpdateLoggedInAthleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update logged in athlete o k response has a 5xx status code
func (o *UpdateLoggedInAthleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update logged in athlete o k response a status code equal to that given
func (o *UpdateLoggedInAthleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update logged in athlete o k response
func (o *UpdateLoggedInAthleteOK) Code() int {
	return 200
}

func (o *UpdateLoggedInAthleteOK) Error() string {
	return fmt.Sprintf("[PUT /athlete][%d] updateLoggedInAthleteOK  %+v", 200, o.Payload)
}

func (o *UpdateLoggedInAthleteOK) String() string {
	return fmt.Sprintf("[PUT /athlete][%d] updateLoggedInAthleteOK  %+v", 200, o.Payload)
}

func (o *UpdateLoggedInAthleteOK) GetPayload() *models.DetailedAthlete {
	return o.Payload
}

func (o *UpdateLoggedInAthleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DetailedAthlete)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLoggedInAthleteDefault creates a UpdateLoggedInAthleteDefault with default headers values
func NewUpdateLoggedInAthleteDefault(code int) *UpdateLoggedInAthleteDefault {
	return &UpdateLoggedInAthleteDefault{
		_statusCode: code,
	}
}

/*
UpdateLoggedInAthleteDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type UpdateLoggedInAthleteDefault struct {
	_statusCode int

	Payload *models.Fault
}

// IsSuccess returns true when this update logged in athlete default response has a 2xx status code
func (o *UpdateLoggedInAthleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update logged in athlete default response has a 3xx status code
func (o *UpdateLoggedInAthleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update logged in athlete default response has a 4xx status code
func (o *UpdateLoggedInAthleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update logged in athlete default response has a 5xx status code
func (o *UpdateLoggedInAthleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update logged in athlete default response a status code equal to that given
func (o *UpdateLoggedInAthleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update logged in athlete default response
func (o *UpdateLoggedInAthleteDefault) Code() int {
	return o._statusCode
}

func (o *UpdateLoggedInAthleteDefault) Error() string {
	return fmt.Sprintf("[PUT /athlete][%d] updateLoggedInAthlete default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateLoggedInAthleteDefault) String() string {
	return fmt.Sprintf("[PUT /athlete][%d] updateLoggedInAthlete default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateLoggedInAthleteDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *UpdateLoggedInAthleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
