// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Route route
//
// swagger:model route
type Route struct {

	// athlete
	Athlete *SummaryAthlete `json:"athlete,omitempty"`

	// The time at which the route was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The description of the route
	Description string `json:"description,omitempty"`

	// The route's distance, in meters
	Distance float32 `json:"distance,omitempty"`

	// The route's elevation gain.
	ElevationGain float32 `json:"elevation_gain,omitempty"`

	// Estimated time in seconds for the authenticated athlete to complete route
	EstimatedMovingTime int64 `json:"estimated_moving_time,omitempty"`

	// The unique identifier of this route
	ID int64 `json:"id,omitempty"`

	// The unique identifier of the route in string format
	IDStr string `json:"id_str,omitempty"`

	// map
	Map *PolylineMap `json:"map,omitempty"`

	// The name of this route
	Name string `json:"name,omitempty"`

	// Whether this route is private
	Private bool `json:"private,omitempty"`

	// The segments traversed by this route
	Segments []*SummarySegment `json:"segments"`

	// Whether this route is starred by the logged-in athlete
	Starred bool `json:"starred,omitempty"`

	// This route's sub-type (1 for road, 2 for mountain bike, 3 for cross, 4 for trail, 5 for mixed)
	SubType int64 `json:"sub_type,omitempty"`

	// An epoch timestamp of when the route was created
	Timestamp int64 `json:"timestamp,omitempty"`

	// This route's type (1 for ride, 2 for runs)
	Type int64 `json:"type,omitempty"`

	// The time at which the route was last updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this route
func (m *Route) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAthlete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Route) validateAthlete(formats strfmt.Registry) error {
	if swag.IsZero(m.Athlete) { // not required
		return nil
	}

	if m.Athlete != nil {
		if err := m.Athlete.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("athlete")
			}
			return err
		}
	}

	return nil
}

func (m *Route) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Route) validateMap(formats strfmt.Registry) error {
	if swag.IsZero(m.Map) { // not required
		return nil
	}

	if m.Map != nil {
		if err := m.Map.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("map")
			}
			return err
		}
	}

	return nil
}

func (m *Route) validateSegments(formats strfmt.Registry) error {
	if swag.IsZero(m.Segments) { // not required
		return nil
	}

	for i := 0; i < len(m.Segments); i++ {
		if swag.IsZero(m.Segments[i]) { // not required
			continue
		}

		if m.Segments[i] != nil {
			if err := m.Segments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Route) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this route based on the context it is used
func (m *Route) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAthlete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSegments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Route) contextValidateAthlete(ctx context.Context, formats strfmt.Registry) error {

	if m.Athlete != nil {

		if swag.IsZero(m.Athlete) { // not required
			return nil
		}

		if err := m.Athlete.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("athlete")
			}
			return err
		}
	}

	return nil
}

func (m *Route) contextValidateMap(ctx context.Context, formats strfmt.Registry) error {

	if m.Map != nil {

		if swag.IsZero(m.Map) { // not required
			return nil
		}

		if err := m.Map.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("map")
			}
			return err
		}
	}

	return nil
}

func (m *Route) contextValidateSegments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Segments); i++ {

		if m.Segments[i] != nil {

			if swag.IsZero(m.Segments[i]) { // not required
				return nil
			}

			if err := m.Segments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Route) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Route) UnmarshalBinary(b []byte) error {
	var res Route
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
