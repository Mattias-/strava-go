// Code generated by go-swagger; DO NOT EDIT.

package streams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// GetActivityStreamsReader is a Reader for the GetActivityStreams structure.
type GetActivityStreamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActivityStreamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActivityStreamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetActivityStreamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetActivityStreamsOK creates a GetActivityStreamsOK with default headers values
func NewGetActivityStreamsOK() *GetActivityStreamsOK {
	return &GetActivityStreamsOK{}
}

/* GetActivityStreamsOK describes a response with status code 200, with default header values.

The set of requested streams.
*/
type GetActivityStreamsOK struct {
	Payload *models.StreamSet
}

func (o *GetActivityStreamsOK) Error() string {
	return fmt.Sprintf("[GET /activities/{id}/streams][%d] getActivityStreamsOK  %+v", 200, o.Payload)
}
func (o *GetActivityStreamsOK) GetPayload() *models.StreamSet {
	return o.Payload
}

func (o *GetActivityStreamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StreamSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActivityStreamsDefault creates a GetActivityStreamsDefault with default headers values
func NewGetActivityStreamsDefault(code int) *GetActivityStreamsDefault {
	return &GetActivityStreamsDefault{
		_statusCode: code,
	}
}

/* GetActivityStreamsDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetActivityStreamsDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get activity streams default response
func (o *GetActivityStreamsDefault) Code() int {
	return o._statusCode
}

func (o *GetActivityStreamsDefault) Error() string {
	return fmt.Sprintf("[GET /activities/{id}/streams][%d] getActivityStreams default  %+v", o._statusCode, o.Payload)
}
func (o *GetActivityStreamsDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetActivityStreamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
