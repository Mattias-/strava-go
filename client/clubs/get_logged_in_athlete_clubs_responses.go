// Code generated by go-swagger; DO NOT EDIT.

package clubs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// GetLoggedInAthleteClubsReader is a Reader for the GetLoggedInAthleteClubs structure.
type GetLoggedInAthleteClubsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoggedInAthleteClubsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoggedInAthleteClubsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLoggedInAthleteClubsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLoggedInAthleteClubsOK creates a GetLoggedInAthleteClubsOK with default headers values
func NewGetLoggedInAthleteClubsOK() *GetLoggedInAthleteClubsOK {
	return &GetLoggedInAthleteClubsOK{}
}

/*
GetLoggedInAthleteClubsOK describes a response with status code 200, with default header values.

A list of summary club representations.
*/
type GetLoggedInAthleteClubsOK struct {
	Payload []*models.SummaryClub
}

// IsSuccess returns true when this get logged in athlete clubs o k response has a 2xx status code
func (o *GetLoggedInAthleteClubsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logged in athlete clubs o k response has a 3xx status code
func (o *GetLoggedInAthleteClubsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logged in athlete clubs o k response has a 4xx status code
func (o *GetLoggedInAthleteClubsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logged in athlete clubs o k response has a 5xx status code
func (o *GetLoggedInAthleteClubsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logged in athlete clubs o k response a status code equal to that given
func (o *GetLoggedInAthleteClubsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get logged in athlete clubs o k response
func (o *GetLoggedInAthleteClubsOK) Code() int {
	return 200
}

func (o *GetLoggedInAthleteClubsOK) Error() string {
	return fmt.Sprintf("[GET /athlete/clubs][%d] getLoggedInAthleteClubsOK  %+v", 200, o.Payload)
}

func (o *GetLoggedInAthleteClubsOK) String() string {
	return fmt.Sprintf("[GET /athlete/clubs][%d] getLoggedInAthleteClubsOK  %+v", 200, o.Payload)
}

func (o *GetLoggedInAthleteClubsOK) GetPayload() []*models.SummaryClub {
	return o.Payload
}

func (o *GetLoggedInAthleteClubsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoggedInAthleteClubsDefault creates a GetLoggedInAthleteClubsDefault with default headers values
func NewGetLoggedInAthleteClubsDefault(code int) *GetLoggedInAthleteClubsDefault {
	return &GetLoggedInAthleteClubsDefault{
		_statusCode: code,
	}
}

/*
GetLoggedInAthleteClubsDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetLoggedInAthleteClubsDefault struct {
	_statusCode int

	Payload *models.Fault
}

// IsSuccess returns true when this get logged in athlete clubs default response has a 2xx status code
func (o *GetLoggedInAthleteClubsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get logged in athlete clubs default response has a 3xx status code
func (o *GetLoggedInAthleteClubsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get logged in athlete clubs default response has a 4xx status code
func (o *GetLoggedInAthleteClubsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get logged in athlete clubs default response has a 5xx status code
func (o *GetLoggedInAthleteClubsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get logged in athlete clubs default response a status code equal to that given
func (o *GetLoggedInAthleteClubsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get logged in athlete clubs default response
func (o *GetLoggedInAthleteClubsDefault) Code() int {
	return o._statusCode
}

func (o *GetLoggedInAthleteClubsDefault) Error() string {
	return fmt.Sprintf("[GET /athlete/clubs][%d] getLoggedInAthleteClubs default  %+v", o._statusCode, o.Payload)
}

func (o *GetLoggedInAthleteClubsDefault) String() string {
	return fmt.Sprintf("[GET /athlete/clubs][%d] getLoggedInAthleteClubs default  %+v", o._statusCode, o.Payload)
}

func (o *GetLoggedInAthleteClubsDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetLoggedInAthleteClubsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
