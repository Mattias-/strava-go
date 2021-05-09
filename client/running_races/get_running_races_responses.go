// Code generated by go-swagger; DO NOT EDIT.

package running_races

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// GetRunningRacesReader is a Reader for the GetRunningRaces structure.
type GetRunningRacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunningRacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunningRacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRunningRacesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunningRacesOK creates a GetRunningRacesOK with default headers values
func NewGetRunningRacesOK() *GetRunningRacesOK {
	return &GetRunningRacesOK{}
}

/* GetRunningRacesOK describes a response with status code 200, with default header values.

Representation of a list of running race.
*/
type GetRunningRacesOK struct {
	Payload []*models.RunningRace
}

func (o *GetRunningRacesOK) Error() string {
	return fmt.Sprintf("[GET /running_races][%d] getRunningRacesOK  %+v", 200, o.Payload)
}
func (o *GetRunningRacesOK) GetPayload() []*models.RunningRace {
	return o.Payload
}

func (o *GetRunningRacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunningRacesDefault creates a GetRunningRacesDefault with default headers values
func NewGetRunningRacesDefault(code int) *GetRunningRacesDefault {
	return &GetRunningRacesDefault{
		_statusCode: code,
	}
}

/* GetRunningRacesDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetRunningRacesDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get running races default response
func (o *GetRunningRacesDefault) Code() int {
	return o._statusCode
}

func (o *GetRunningRacesDefault) Error() string {
	return fmt.Sprintf("[GET /running_races][%d] getRunningRaces default  %+v", o._statusCode, o.Payload)
}
func (o *GetRunningRacesDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetRunningRacesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
