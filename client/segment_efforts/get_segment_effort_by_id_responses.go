// Code generated by go-swagger; DO NOT EDIT.

package segment_efforts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// GetSegmentEffortByIDReader is a Reader for the GetSegmentEffortByID structure.
type GetSegmentEffortByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSegmentEffortByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSegmentEffortByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSegmentEffortByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSegmentEffortByIDOK creates a GetSegmentEffortByIDOK with default headers values
func NewGetSegmentEffortByIDOK() *GetSegmentEffortByIDOK {
	return &GetSegmentEffortByIDOK{}
}

/*
GetSegmentEffortByIDOK describes a response with status code 200, with default header values.

Representation of a segment effort.
*/
type GetSegmentEffortByIDOK struct {
	Payload *models.DetailedSegmentEffort
}

// IsSuccess returns true when this get segment effort by Id o k response has a 2xx status code
func (o *GetSegmentEffortByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get segment effort by Id o k response has a 3xx status code
func (o *GetSegmentEffortByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get segment effort by Id o k response has a 4xx status code
func (o *GetSegmentEffortByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get segment effort by Id o k response has a 5xx status code
func (o *GetSegmentEffortByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get segment effort by Id o k response a status code equal to that given
func (o *GetSegmentEffortByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get segment effort by Id o k response
func (o *GetSegmentEffortByIDOK) Code() int {
	return 200
}

func (o *GetSegmentEffortByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /segment_efforts/{id}][%d] getSegmentEffortByIdOK %s", 200, payload)
}

func (o *GetSegmentEffortByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /segment_efforts/{id}][%d] getSegmentEffortByIdOK %s", 200, payload)
}

func (o *GetSegmentEffortByIDOK) GetPayload() *models.DetailedSegmentEffort {
	return o.Payload
}

func (o *GetSegmentEffortByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DetailedSegmentEffort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSegmentEffortByIDDefault creates a GetSegmentEffortByIDDefault with default headers values
func NewGetSegmentEffortByIDDefault(code int) *GetSegmentEffortByIDDefault {
	return &GetSegmentEffortByIDDefault{
		_statusCode: code,
	}
}

/*
GetSegmentEffortByIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetSegmentEffortByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// IsSuccess returns true when this get segment effort by Id default response has a 2xx status code
func (o *GetSegmentEffortByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get segment effort by Id default response has a 3xx status code
func (o *GetSegmentEffortByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get segment effort by Id default response has a 4xx status code
func (o *GetSegmentEffortByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get segment effort by Id default response has a 5xx status code
func (o *GetSegmentEffortByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get segment effort by Id default response a status code equal to that given
func (o *GetSegmentEffortByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get segment effort by Id default response
func (o *GetSegmentEffortByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSegmentEffortByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /segment_efforts/{id}][%d] getSegmentEffortById default %s", o._statusCode, payload)
}

func (o *GetSegmentEffortByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /segment_efforts/{id}][%d] getSegmentEffortById default %s", o._statusCode, payload)
}

func (o *GetSegmentEffortByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetSegmentEffortByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
