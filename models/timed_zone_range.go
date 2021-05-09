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

// TimedZoneRange A union type representing the time spent in a given zone.
//
// swagger:model timedZoneRange
type TimedZoneRange struct {
	ZoneRange

	// The number of seconds spent in this zone
	Time int64 `json:"time,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TimedZoneRange) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ZoneRange
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ZoneRange = aO0

	// AO1
	var dataAO1 struct {
		Time int64 `json:"time,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Time = dataAO1.Time

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TimedZoneRange) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ZoneRange)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Time int64 `json:"time,omitempty"`
	}

	dataAO1.Time = m.Time

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this timed zone range
func (m *TimedZoneRange) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ZoneRange
	if err := m.ZoneRange.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this timed zone range based on the context it is used
func (m *TimedZoneRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ZoneRange
	if err := m.ZoneRange.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TimedZoneRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimedZoneRange) UnmarshalBinary(b []byte) error {
	var res TimedZoneRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
