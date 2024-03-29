// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClubActivity club activity
//
// swagger:model clubActivity
type ClubActivity struct {

	// athlete
	Athlete *MetaAthlete `json:"athlete,omitempty"`

	// The activity's distance, in meters
	Distance float32 `json:"distance,omitempty"`

	// The activity's elapsed time, in seconds
	ElapsedTime int64 `json:"elapsed_time,omitempty"`

	// The activity's moving time, in seconds
	MovingTime int64 `json:"moving_time,omitempty"`

	// The name of the activity
	Name string `json:"name,omitempty"`

	// sport type
	SportType SportType `json:"sport_type,omitempty"`

	// The activity's total elevation gain.
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`

	// type
	Type ActivityType `json:"type,omitempty"`

	// The activity's workout type
	WorkoutType int64 `json:"workout_type,omitempty"`
}

// Validate validates this club activity
func (m *ClubActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAthlete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSportType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClubActivity) validateAthlete(formats strfmt.Registry) error {
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

func (m *ClubActivity) validateSportType(formats strfmt.Registry) error {
	if swag.IsZero(m.SportType) { // not required
		return nil
	}

	if err := m.SportType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sport_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sport_type")
		}
		return err
	}

	return nil
}

func (m *ClubActivity) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this club activity based on the context it is used
func (m *ClubActivity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAthlete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSportType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClubActivity) contextValidateAthlete(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClubActivity) contextValidateSportType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SportType) { // not required
		return nil
	}

	if err := m.SportType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sport_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sport_type")
		}
		return err
	}

	return nil
}

func (m *ClubActivity) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClubActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClubActivity) UnmarshalBinary(b []byte) error {
	var res ClubActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
