// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Lap lap
//
// swagger:model lap
type Lap struct {

	// activity
	Activity *MetaActivity `json:"activity,omitempty"`

	// athlete
	Athlete *MetaAthlete `json:"athlete,omitempty"`

	// The lap's average cadence
	AverageCadence float32 `json:"average_cadence,omitempty"`

	// The lap's average speed
	AverageSpeed float32 `json:"average_speed,omitempty"`

	// The lap's distance, in meters
	Distance float32 `json:"distance,omitempty"`

	// The lap's elapsed time, in seconds
	ElapsedTime int64 `json:"elapsed_time,omitempty"`

	// The end index of this effort in its activity's stream
	EndIndex int64 `json:"end_index,omitempty"`

	// The unique identifier of this lap
	ID int64 `json:"id,omitempty"`

	// The index of this lap in the activity it belongs to
	LapIndex int64 `json:"lap_index,omitempty"`

	// The maximum speed of this lat, in meters per second
	MaxSpeed float32 `json:"max_speed,omitempty"`

	// The lap's moving time, in seconds
	MovingTime int64 `json:"moving_time,omitempty"`

	// The name of the lap
	Name string `json:"name,omitempty"`

	// The athlete's pace zone during this lap
	PaceZone int64 `json:"pace_zone,omitempty"`

	// split
	Split int64 `json:"split,omitempty"`

	// The time at which the lap was started.
	// Format: date-time
	StartDate strfmt.DateTime `json:"start_date,omitempty"`

	// The time at which the lap was started in the local timezone.
	// Format: date-time
	StartDateLocal strfmt.DateTime `json:"start_date_local,omitempty"`

	// The start index of this effort in its activity's stream
	StartIndex int64 `json:"start_index,omitempty"`

	// The elevation gain of this lap, in meters
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`
}

// Validate validates this lap
func (m *Lap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAthlete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateLocal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Lap) validateActivity(formats strfmt.Registry) error {
	if swag.IsZero(m.Activity) { // not required
		return nil
	}

	if m.Activity != nil {
		if err := m.Activity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity")
			}
			return err
		}
	}

	return nil
}

func (m *Lap) validateAthlete(formats strfmt.Registry) error {
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

func (m *Lap) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Lap) validateStartDateLocal(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDateLocal) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date_local", "body", "date-time", m.StartDateLocal.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this lap based on the context it is used
func (m *Lap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAthlete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Lap) contextValidateActivity(ctx context.Context, formats strfmt.Registry) error {

	if m.Activity != nil {

		if swag.IsZero(m.Activity) { // not required
			return nil
		}

		if err := m.Activity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity")
			}
			return err
		}
	}

	return nil
}

func (m *Lap) contextValidateAthlete(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Lap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Lap) UnmarshalBinary(b []byte) error {
	var res Lap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
