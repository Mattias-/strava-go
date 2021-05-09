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

// GetClubMembersByIDReader is a Reader for the GetClubMembersByID structure.
type GetClubMembersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClubMembersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClubMembersByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClubMembersByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClubMembersByIDOK creates a GetClubMembersByIDOK with default headers values
func NewGetClubMembersByIDOK() *GetClubMembersByIDOK {
	return &GetClubMembersByIDOK{}
}

/* GetClubMembersByIDOK describes a response with status code 200, with default header values.

A list of summary athlete representations.
*/
type GetClubMembersByIDOK struct {
	Payload []*models.SummaryAthlete
}

func (o *GetClubMembersByIDOK) Error() string {
	return fmt.Sprintf("[GET /clubs/{id}/members][%d] getClubMembersByIdOK  %+v", 200, o.Payload)
}
func (o *GetClubMembersByIDOK) GetPayload() []*models.SummaryAthlete {
	return o.Payload
}

func (o *GetClubMembersByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClubMembersByIDDefault creates a GetClubMembersByIDDefault with default headers values
func NewGetClubMembersByIDDefault(code int) *GetClubMembersByIDDefault {
	return &GetClubMembersByIDDefault{
		_statusCode: code,
	}
}

/* GetClubMembersByIDDefault describes a response with status code -1, with default header values.

Unexpected error.
*/
type GetClubMembersByIDDefault struct {
	_statusCode int

	Payload *models.Fault
}

// Code gets the status code for the get club members by Id default response
func (o *GetClubMembersByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClubMembersByIDDefault) Error() string {
	return fmt.Sprintf("[GET /clubs/{id}/members][%d] getClubMembersById default  %+v", o._statusCode, o.Payload)
}
func (o *GetClubMembersByIDDefault) GetPayload() *models.Fault {
	return o.Payload
}

func (o *GetClubMembersByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
