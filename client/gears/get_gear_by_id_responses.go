// Code generated by go-swagger; DO NOT EDIT.

package gears

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// GetGearByIDReader is a Reader for the GetGearByID structure.
type GetGearByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGearByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGearByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetGearByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGearByIDOK creates a GetGearByIDOK with default headers values
func NewGetGearByIDOK() *GetGearByIDOK {
	return &GetGearByIDOK{}
}

/* GetGearByIDOK describes a response with status code 200, with default header values.

A representation of the gear.
*/
type GetGearByIDOK struct {
	Payload *models.DetailedGear
}

func (o *GetGearByIDOK) Error() string {
	return fmt.Sprintf("[GET /gear/{id}][%d] getGearByIdOK  %+v", 200, o.Payload)
}
func (o *GetGearByIDOK) GetPayload() *models.DetailedGear {
	return o.Payload
}

func (o *GetGearByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DetailedGear)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGearByIDDefault creates a GetGearByIDDefault with default headers values
func NewGetGearByIDDefault(code int) *GetGearByIDDefault {
	return &GetGearByIDDefault{
		_statusCode: code,
	}
}

/* GetGearByIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetGearByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get gear by Id default response
func (o *GetGearByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetGearByIDDefault) Error() string {
	return fmt.Sprintf("[GET /gear/{id}][%d] getGearById default  %+v", o._statusCode, o.Payload)
}
func (o *GetGearByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetGearByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
