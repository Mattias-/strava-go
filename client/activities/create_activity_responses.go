// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// CreateActivityReader is a Reader for the CreateActivity structure.
type CreateActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateActivityCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateActivityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateActivityCreated creates a CreateActivityCreated with default headers values
func NewCreateActivityCreated() *CreateActivityCreated {
	return &CreateActivityCreated{}
}

/* CreateActivityCreated describes a response with status code 201, with default header values.

The activity's detailed representation.
*/
type CreateActivityCreated struct {
	Payload *models.DetailedActivity
}

func (o *CreateActivityCreated) Error() string {
	return fmt.Sprintf("[POST /activities][%d] createActivityCreated  %+v", 201, o.Payload)
}
func (o *CreateActivityCreated) GetPayload() *models.DetailedActivity {
	return o.Payload
}

func (o *CreateActivityCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DetailedActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActivityDefault creates a CreateActivityDefault with default headers values
func NewCreateActivityDefault(code int) *CreateActivityDefault {
	return &CreateActivityDefault{
		_statusCode: code,
	}
}

/* CreateActivityDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type CreateActivityDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the create activity default response
func (o *CreateActivityDefault) Code() int {
	return o._statusCode
}

func (o *CreateActivityDefault) Error() string {
	return fmt.Sprintf("[POST /activities][%d] createActivity default  %+v", o._statusCode, o.Payload)
}
func (o *CreateActivityDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *CreateActivityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
