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

// Zones zones
//
// swagger:model zones
type Zones struct {

	// heart rate
	HeartRate *HeartRateZoneRanges `json:"heart_rate,omitempty"`

	// power
	Power *PowerZoneRanges `json:"power,omitempty"`
}

// Validate validates this zones
func (m *Zones) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeartRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePower(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zones) validateHeartRate(formats strfmt.Registry) error {
	if swag.IsZero(m.HeartRate) { // not required
		return nil
	}

	if m.HeartRate != nil {
		if err := m.HeartRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heart_rate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("heart_rate")
			}
			return err
		}
	}

	return nil
}

func (m *Zones) validatePower(formats strfmt.Registry) error {
	if swag.IsZero(m.Power) { // not required
		return nil
	}

	if m.Power != nil {
		if err := m.Power.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("power")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this zones based on the context it is used
func (m *Zones) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeartRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePower(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zones) contextValidateHeartRate(ctx context.Context, formats strfmt.Registry) error {

	if m.HeartRate != nil {
		if err := m.HeartRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heart_rate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("heart_rate")
			}
			return err
		}
	}

	return nil
}

func (m *Zones) contextValidatePower(ctx context.Context, formats strfmt.Registry) error {

	if m.Power != nil {
		if err := m.Power.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("power")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Zones) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Zones) UnmarshalBinary(b []byte) error {
	var res Zones
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
