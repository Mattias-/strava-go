// Code generated by go-swagger; DO NOT EDIT.

package segments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// ExploreSegmentsReader is a Reader for the ExploreSegments structure.
type ExploreSegmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExploreSegmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExploreSegmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExploreSegmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExploreSegmentsOK creates a ExploreSegmentsOK with default headers values
func NewExploreSegmentsOK() *ExploreSegmentsOK {
	return &ExploreSegmentsOK{}
}

/*
ExploreSegmentsOK describes a response with status code 200, with default header values.

List of matching segments.
*/
type ExploreSegmentsOK struct {
	Payload *models.ExplorerResponse
}

// IsSuccess returns true when this explore segments o k response has a 2xx status code
func (o *ExploreSegmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this explore segments o k response has a 3xx status code
func (o *ExploreSegmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this explore segments o k response has a 4xx status code
func (o *ExploreSegmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this explore segments o k response has a 5xx status code
func (o *ExploreSegmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this explore segments o k response a status code equal to that given
func (o *ExploreSegmentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExploreSegmentsOK) Error() string {
	return fmt.Sprintf("[GET /segments/explore][%d] exploreSegmentsOK  %+v", 200, o.Payload)
}

func (o *ExploreSegmentsOK) String() string {
	return fmt.Sprintf("[GET /segments/explore][%d] exploreSegmentsOK  %+v", 200, o.Payload)
}

func (o *ExploreSegmentsOK) GetPayload() *models.ExplorerResponse {
	return o.Payload
}

func (o *ExploreSegmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExplorerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExploreSegmentsDefault creates a ExploreSegmentsDefault with default headers values
func NewExploreSegmentsDefault(code int) *ExploreSegmentsDefault {
	return &ExploreSegmentsDefault{
		_statusCode: code,
	}
}

/*
ExploreSegmentsDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type ExploreSegmentsDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the explore segments default response
func (o *ExploreSegmentsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this explore segments default response has a 2xx status code
func (o *ExploreSegmentsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this explore segments default response has a 3xx status code
func (o *ExploreSegmentsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this explore segments default response has a 4xx status code
func (o *ExploreSegmentsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this explore segments default response has a 5xx status code
func (o *ExploreSegmentsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this explore segments default response a status code equal to that given
func (o *ExploreSegmentsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExploreSegmentsDefault) Error() string {
	return fmt.Sprintf("[GET /segments/explore][%d] exploreSegments default  %+v", o._statusCode, o.Payload)
}

func (o *ExploreSegmentsDefault) String() string {
	return fmt.Sprintf("[GET /segments/explore][%d] exploreSegments default  %+v", o._statusCode, o.Payload)
}

func (o *ExploreSegmentsDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *ExploreSegmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
