// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SummarySegment summary segment
//
// swagger:model summarySegment
type SummarySegment struct {

	// activity type
	// Enum: [Ride Run]
	ActivityType string `json:"activity_type,omitempty"`

	// athlete pr effort
	AthletePrEffort *SummarySegmentEffort `json:"athlete_pr_effort,omitempty"`

	// athlete segment stats
	AthleteSegmentStats *SummaryPRSegmentEffort `json:"athlete_segment_stats,omitempty"`

	// The segment's average grade, in percents
	AverageGrade float32 `json:"average_grade,omitempty"`

	// The segments's city.
	City string `json:"city,omitempty"`

	// The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category.
	ClimbCategory int64 `json:"climb_category,omitempty"`

	// The segment's country.
	Country string `json:"country,omitempty"`

	// The segment's distance, in meters
	Distance float32 `json:"distance,omitempty"`

	// The segments's highest elevation, in meters
	ElevationHigh float32 `json:"elevation_high,omitempty"`

	// The segments's lowest elevation, in meters
	ElevationLow float32 `json:"elevation_low,omitempty"`

	// end latlng
	EndLatlng LatLng `json:"end_latlng,omitempty"`

	// The unique identifier of this segment
	ID int64 `json:"id,omitempty"`

	// The segments's maximum grade, in percents
	MaximumGrade float32 `json:"maximum_grade,omitempty"`

	// The name of this segment
	Name string `json:"name,omitempty"`

	// Whether this segment is private.
	Private bool `json:"private,omitempty"`

	// start latlng
	StartLatlng LatLng `json:"start_latlng,omitempty"`

	// The segments's state or geographical region.
	State string `json:"state,omitempty"`
}

// Validate validates this summary segment
func (m *SummarySegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAthletePrEffort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAthleteSegmentStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndLatlng(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartLatlng(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var summarySegmentTypeActivityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Ride","Run"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		summarySegmentTypeActivityTypePropEnum = append(summarySegmentTypeActivityTypePropEnum, v)
	}
}

const (

	// SummarySegmentActivityTypeRide captures enum value "Ride"
	SummarySegmentActivityTypeRide string = "Ride"

	// SummarySegmentActivityTypeRun captures enum value "Run"
	SummarySegmentActivityTypeRun string = "Run"
)

// prop value enum
func (m *SummarySegment) validateActivityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, summarySegmentTypeActivityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SummarySegment) validateActivityType(formats strfmt.Registry) error {
	if swag.IsZero(m.ActivityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateActivityTypeEnum("activity_type", "body", m.ActivityType); err != nil {
		return err
	}

	return nil
}

func (m *SummarySegment) validateAthletePrEffort(formats strfmt.Registry) error {
	if swag.IsZero(m.AthletePrEffort) { // not required
		return nil
	}

	if m.AthletePrEffort != nil {
		if err := m.AthletePrEffort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete_pr_effort")
			}
			return err
		}
	}

	return nil
}

func (m *SummarySegment) validateAthleteSegmentStats(formats strfmt.Registry) error {
	if swag.IsZero(m.AthleteSegmentStats) { // not required
		return nil
	}

	if m.AthleteSegmentStats != nil {
		if err := m.AthleteSegmentStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete_segment_stats")
			}
			return err
		}
	}

	return nil
}

func (m *SummarySegment) validateEndLatlng(formats strfmt.Registry) error {
	if swag.IsZero(m.EndLatlng) { // not required
		return nil
	}

	if err := m.EndLatlng.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("end_latlng")
		}
		return err
	}

	return nil
}

func (m *SummarySegment) validateStartLatlng(formats strfmt.Registry) error {
	if swag.IsZero(m.StartLatlng) { // not required
		return nil
	}

	if err := m.StartLatlng.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("start_latlng")
		}
		return err
	}

	return nil
}

// ContextValidate validate this summary segment based on the context it is used
func (m *SummarySegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAthletePrEffort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAthleteSegmentStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndLatlng(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartLatlng(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummarySegment) contextValidateAthletePrEffort(ctx context.Context, formats strfmt.Registry) error {

	if m.AthletePrEffort != nil {
		if err := m.AthletePrEffort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete_pr_effort")
			}
			return err
		}
	}

	return nil
}

func (m *SummarySegment) contextValidateAthleteSegmentStats(ctx context.Context, formats strfmt.Registry) error {

	if m.AthleteSegmentStats != nil {
		if err := m.AthleteSegmentStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete_segment_stats")
			}
			return err
		}
	}

	return nil
}

func (m *SummarySegment) contextValidateEndLatlng(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EndLatlng.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("end_latlng")
		}
		return err
	}

	return nil
}

func (m *SummarySegment) contextValidateStartLatlng(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StartLatlng.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("start_latlng")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummarySegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummarySegment) UnmarshalBinary(b []byte) error {
	var res SummarySegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
