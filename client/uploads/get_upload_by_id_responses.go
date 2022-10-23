// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Mattias-/strava-go/models"
)

// GetUploadByIDReader is a Reader for the GetUploadByID structure.
type GetUploadByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUploadByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUploadByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUploadByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUploadByIDOK creates a GetUploadByIDOK with default headers values
func NewGetUploadByIDOK() *GetUploadByIDOK {
	return &GetUploadByIDOK{}
}

/*
GetUploadByIDOK describes a response with status code 200, with default header values.

Representation of the upload.
*/
type GetUploadByIDOK struct {
	Payload *models.Upload
}

// IsSuccess returns true when this get upload by Id o k response has a 2xx status code
func (o *GetUploadByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upload by Id o k response has a 3xx status code
func (o *GetUploadByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upload by Id o k response has a 4xx status code
func (o *GetUploadByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upload by Id o k response has a 5xx status code
func (o *GetUploadByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upload by Id o k response a status code equal to that given
func (o *GetUploadByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUploadByIDOK) Error() string {
	return fmt.Sprintf("[GET /uploads/{uploadId}][%d] getUploadByIdOK  %+v", 200, o.Payload)
}

func (o *GetUploadByIDOK) String() string {
	return fmt.Sprintf("[GET /uploads/{uploadId}][%d] getUploadByIdOK  %+v", 200, o.Payload)
}

func (o *GetUploadByIDOK) GetPayload() *models.Upload {
	return o.Payload
}

func (o *GetUploadByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Upload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUploadByIDDefault creates a GetUploadByIDDefault with default headers values
func NewGetUploadByIDDefault(code int) *GetUploadByIDDefault {
	return &GetUploadByIDDefault{
		_statusCode: code,
	}
}

/*
GetUploadByIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetUploadByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get upload by Id default response
func (o *GetUploadByIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get upload by Id default response has a 2xx status code
func (o *GetUploadByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get upload by Id default response has a 3xx status code
func (o *GetUploadByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get upload by Id default response has a 4xx status code
func (o *GetUploadByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get upload by Id default response has a 5xx status code
func (o *GetUploadByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get upload by Id default response a status code equal to that given
func (o *GetUploadByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetUploadByIDDefault) Error() string {
	return fmt.Sprintf("[GET /uploads/{uploadId}][%d] getUploadById default  %+v", o._statusCode, o.Payload)
}

func (o *GetUploadByIDDefault) String() string {
	return fmt.Sprintf("[GET /uploads/{uploadId}][%d] getUploadById default  %+v", o._statusCode, o.Payload)
}

func (o *GetUploadByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetUploadByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
