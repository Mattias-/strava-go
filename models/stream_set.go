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

// StreamSet stream set
//
// swagger:model streamSet
type StreamSet struct {

	// altitude
	Altitude *AltitudeStream `json:"altitude,omitempty"`

	// cadence
	Cadence *CadenceStream `json:"cadence,omitempty"`

	// distance
	Distance *DistanceStream `json:"distance,omitempty"`

	// grade smooth
	GradeSmooth *SmoothGradeStream `json:"grade_smooth,omitempty"`

	// heartrate
	Heartrate *HeartrateStream `json:"heartrate,omitempty"`

	// latlng
	Latlng *LatLngStream `json:"latlng,omitempty"`

	// moving
	Moving *MovingStream `json:"moving,omitempty"`

	// temp
	Temp *TemperatureStream `json:"temp,omitempty"`

	// time
	Time *TimeStream `json:"time,omitempty"`

	// velocity smooth
	VelocitySmooth *SmoothVelocityStream `json:"velocity_smooth,omitempty"`

	// watts
	Watts *PowerStream `json:"watts,omitempty"`
}

// Validate validates this stream set
func (m *StreamSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAltitude(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCadence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGradeSmooth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeartrate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatlng(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoving(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVelocitySmooth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWatts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StreamSet) validateAltitude(formats strfmt.Registry) error {
	if swag.IsZero(m.Altitude) { // not required
		return nil
	}

	if m.Altitude != nil {
		if err := m.Altitude.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("altitude")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("altitude")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateCadence(formats strfmt.Registry) error {
	if swag.IsZero(m.Cadence) { // not required
		return nil
	}

	if m.Cadence != nil {
		if err := m.Cadence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cadence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cadence")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateDistance(formats strfmt.Registry) error {
	if swag.IsZero(m.Distance) { // not required
		return nil
	}

	if m.Distance != nil {
		if err := m.Distance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("distance")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateGradeSmooth(formats strfmt.Registry) error {
	if swag.IsZero(m.GradeSmooth) { // not required
		return nil
	}

	if m.GradeSmooth != nil {
		if err := m.GradeSmooth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grade_smooth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grade_smooth")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateHeartrate(formats strfmt.Registry) error {
	if swag.IsZero(m.Heartrate) { // not required
		return nil
	}

	if m.Heartrate != nil {
		if err := m.Heartrate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heartrate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("heartrate")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateLatlng(formats strfmt.Registry) error {
	if swag.IsZero(m.Latlng) { // not required
		return nil
	}

	if m.Latlng != nil {
		if err := m.Latlng.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latlng")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latlng")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateMoving(formats strfmt.Registry) error {
	if swag.IsZero(m.Moving) { // not required
		return nil
	}

	if m.Moving != nil {
		if err := m.Moving.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moving")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moving")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateTemp(formats strfmt.Registry) error {
	if swag.IsZero(m.Temp) { // not required
		return nil
	}

	if m.Temp != nil {
		if err := m.Temp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("temp")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if m.Time != nil {
		if err := m.Time.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateVelocitySmooth(formats strfmt.Registry) error {
	if swag.IsZero(m.VelocitySmooth) { // not required
		return nil
	}

	if m.VelocitySmooth != nil {
		if err := m.VelocitySmooth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("velocity_smooth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("velocity_smooth")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) validateWatts(formats strfmt.Registry) error {
	if swag.IsZero(m.Watts) { // not required
		return nil
	}

	if m.Watts != nil {
		if err := m.Watts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("watts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("watts")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this stream set based on the context it is used
func (m *StreamSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAltitude(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCadence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDistance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGradeSmooth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeartrate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatlng(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMoving(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVelocitySmooth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWatts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StreamSet) contextValidateAltitude(ctx context.Context, formats strfmt.Registry) error {

	if m.Altitude != nil {
		if err := m.Altitude.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("altitude")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("altitude")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateCadence(ctx context.Context, formats strfmt.Registry) error {

	if m.Cadence != nil {
		if err := m.Cadence.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cadence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cadence")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateDistance(ctx context.Context, formats strfmt.Registry) error {

	if m.Distance != nil {
		if err := m.Distance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("distance")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateGradeSmooth(ctx context.Context, formats strfmt.Registry) error {

	if m.GradeSmooth != nil {
		if err := m.GradeSmooth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grade_smooth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grade_smooth")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateHeartrate(ctx context.Context, formats strfmt.Registry) error {

	if m.Heartrate != nil {
		if err := m.Heartrate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heartrate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("heartrate")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateLatlng(ctx context.Context, formats strfmt.Registry) error {

	if m.Latlng != nil {
		if err := m.Latlng.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latlng")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latlng")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateMoving(ctx context.Context, formats strfmt.Registry) error {

	if m.Moving != nil {
		if err := m.Moving.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moving")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moving")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateTemp(ctx context.Context, formats strfmt.Registry) error {

	if m.Temp != nil {
		if err := m.Temp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("temp")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if m.Time != nil {
		if err := m.Time.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateVelocitySmooth(ctx context.Context, formats strfmt.Registry) error {

	if m.VelocitySmooth != nil {
		if err := m.VelocitySmooth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("velocity_smooth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("velocity_smooth")
			}
			return err
		}
	}

	return nil
}

func (m *StreamSet) contextValidateWatts(ctx context.Context, formats strfmt.Registry) error {

	if m.Watts != nil {
		if err := m.Watts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("watts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("watts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StreamSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StreamSet) UnmarshalBinary(b []byte) error {
	var res StreamSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
