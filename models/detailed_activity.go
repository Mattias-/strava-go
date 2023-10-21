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
)

// DetailedActivity detailed activity
//
// swagger:model detailedActivity
type DetailedActivity struct {
	SummaryActivity

	// best efforts
	BestEfforts []*DetailedSegmentEffort `json:"best_efforts"`

	// The number of kilocalories consumed during this activity
	Calories float32 `json:"calories,omitempty"`

	// The description of the activity
	Description string `json:"description,omitempty"`

	// The name of the device used to record the activity
	DeviceName string `json:"device_name,omitempty"`

	// The token used to embed a Strava activity
	EmbedToken string `json:"embed_token,omitempty"`

	// gear
	Gear *SummaryGear `json:"gear,omitempty"`

	// laps
	Laps []*Lap `json:"laps"`

	// photos
	Photos *PhotosSummary `json:"photos,omitempty"`

	// segment efforts
	SegmentEfforts []*DetailedSegmentEffort `json:"segment_efforts"`

	// The splits of this activity in metric units (for runs)
	SplitsMetric []*Split `json:"splits_metric"`

	// The splits of this activity in imperial units (for runs)
	SplitsStandard []*Split `json:"splits_standard"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DetailedActivity) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SummaryActivity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SummaryActivity = aO0

	// AO1
	var dataAO1 struct {
		BestEfforts []*DetailedSegmentEffort `json:"best_efforts"`

		Calories float32 `json:"calories,omitempty"`

		Description string `json:"description,omitempty"`

		DeviceName string `json:"device_name,omitempty"`

		EmbedToken string `json:"embed_token,omitempty"`

		Gear *SummaryGear `json:"gear,omitempty"`

		Laps []*Lap `json:"laps"`

		Photos *PhotosSummary `json:"photos,omitempty"`

		SegmentEfforts []*DetailedSegmentEffort `json:"segment_efforts"`

		SplitsMetric []*Split `json:"splits_metric"`

		SplitsStandard []*Split `json:"splits_standard"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BestEfforts = dataAO1.BestEfforts

	m.Calories = dataAO1.Calories

	m.Description = dataAO1.Description

	m.DeviceName = dataAO1.DeviceName

	m.EmbedToken = dataAO1.EmbedToken

	m.Gear = dataAO1.Gear

	m.Laps = dataAO1.Laps

	m.Photos = dataAO1.Photos

	m.SegmentEfforts = dataAO1.SegmentEfforts

	m.SplitsMetric = dataAO1.SplitsMetric

	m.SplitsStandard = dataAO1.SplitsStandard

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DetailedActivity) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SummaryActivity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		BestEfforts []*DetailedSegmentEffort `json:"best_efforts"`

		Calories float32 `json:"calories,omitempty"`

		Description string `json:"description,omitempty"`

		DeviceName string `json:"device_name,omitempty"`

		EmbedToken string `json:"embed_token,omitempty"`

		Gear *SummaryGear `json:"gear,omitempty"`

		Laps []*Lap `json:"laps"`

		Photos *PhotosSummary `json:"photos,omitempty"`

		SegmentEfforts []*DetailedSegmentEffort `json:"segment_efforts"`

		SplitsMetric []*Split `json:"splits_metric"`

		SplitsStandard []*Split `json:"splits_standard"`
	}

	dataAO1.BestEfforts = m.BestEfforts

	dataAO1.Calories = m.Calories

	dataAO1.Description = m.Description

	dataAO1.DeviceName = m.DeviceName

	dataAO1.EmbedToken = m.EmbedToken

	dataAO1.Gear = m.Gear

	dataAO1.Laps = m.Laps

	dataAO1.Photos = m.Photos

	dataAO1.SegmentEfforts = m.SegmentEfforts

	dataAO1.SplitsMetric = m.SplitsMetric

	dataAO1.SplitsStandard = m.SplitsStandard

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this detailed activity
func (m *DetailedActivity) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SummaryActivity
	if err := m.SummaryActivity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBestEfforts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGear(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentEfforts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSplitsMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSplitsStandard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedActivity) validateBestEfforts(formats strfmt.Registry) error {

	if swag.IsZero(m.BestEfforts) { // not required
		return nil
	}

	for i := 0; i < len(m.BestEfforts); i++ {
		if swag.IsZero(m.BestEfforts[i]) { // not required
			continue
		}

		if m.BestEfforts[i] != nil {
			if err := m.BestEfforts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("best_efforts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("best_efforts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedActivity) validateGear(formats strfmt.Registry) error {

	if swag.IsZero(m.Gear) { // not required
		return nil
	}

	if m.Gear != nil {
		if err := m.Gear.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gear")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gear")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedActivity) validateLaps(formats strfmt.Registry) error {

	if swag.IsZero(m.Laps) { // not required
		return nil
	}

	for i := 0; i < len(m.Laps); i++ {
		if swag.IsZero(m.Laps[i]) { // not required
			continue
		}

		if m.Laps[i] != nil {
			if err := m.Laps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("laps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("laps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedActivity) validatePhotos(formats strfmt.Registry) error {

	if swag.IsZero(m.Photos) { // not required
		return nil
	}

	if m.Photos != nil {
		if err := m.Photos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photos")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("photos")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedActivity) validateSegmentEfforts(formats strfmt.Registry) error {

	if swag.IsZero(m.SegmentEfforts) { // not required
		return nil
	}

	for i := 0; i < len(m.SegmentEfforts); i++ {
		if swag.IsZero(m.SegmentEfforts[i]) { // not required
			continue
		}

		if m.SegmentEfforts[i] != nil {
			if err := m.SegmentEfforts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segment_efforts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segment_efforts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedActivity) validateSplitsMetric(formats strfmt.Registry) error {

	if swag.IsZero(m.SplitsMetric) { // not required
		return nil
	}

	for i := 0; i < len(m.SplitsMetric); i++ {
		if swag.IsZero(m.SplitsMetric[i]) { // not required
			continue
		}

		if m.SplitsMetric[i] != nil {
			if err := m.SplitsMetric[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("splits_metric" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("splits_metric" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedActivity) validateSplitsStandard(formats strfmt.Registry) error {

	if swag.IsZero(m.SplitsStandard) { // not required
		return nil
	}

	for i := 0; i < len(m.SplitsStandard); i++ {
		if swag.IsZero(m.SplitsStandard[i]) { // not required
			continue
		}

		if m.SplitsStandard[i] != nil {
			if err := m.SplitsStandard[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("splits_standard" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("splits_standard" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this detailed activity based on the context it is used
func (m *DetailedActivity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SummaryActivity
	if err := m.SummaryActivity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBestEfforts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGear(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLaps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhotos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSegmentEfforts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSplitsMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSplitsStandard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedActivity) contextValidateBestEfforts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BestEfforts); i++ {

		if m.BestEfforts[i] != nil {

			if swag.IsZero(m.BestEfforts[i]) { // not required
				return nil
			}

			if err := m.BestEfforts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("best_efforts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("best_efforts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedActivity) contextValidateGear(ctx context.Context, formats strfmt.Registry) error {

	if m.Gear != nil {

		if swag.IsZero(m.Gear) { // not required
			return nil
		}

		if err := m.Gear.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gear")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gear")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedActivity) contextValidateLaps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Laps); i++ {

		if m.Laps[i] != nil {

			if swag.IsZero(m.Laps[i]) { // not required
				return nil
			}

			if err := m.Laps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("laps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("laps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedActivity) contextValidatePhotos(ctx context.Context, formats strfmt.Registry) error {

	if m.Photos != nil {

		if swag.IsZero(m.Photos) { // not required
			return nil
		}

		if err := m.Photos.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photos")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("photos")
			}
			return err
		}
	}

	return nil
}

func (m *DetailedActivity) contextValidateSegmentEfforts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SegmentEfforts); i++ {

		if m.SegmentEfforts[i] != nil {

			if swag.IsZero(m.SegmentEfforts[i]) { // not required
				return nil
			}

			if err := m.SegmentEfforts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segment_efforts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("segment_efforts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedActivity) contextValidateSplitsMetric(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SplitsMetric); i++ {

		if m.SplitsMetric[i] != nil {

			if swag.IsZero(m.SplitsMetric[i]) { // not required
				return nil
			}

			if err := m.SplitsMetric[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("splits_metric" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("splits_metric" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedActivity) contextValidateSplitsStandard(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SplitsStandard); i++ {

		if m.SplitsStandard[i] != nil {

			if swag.IsZero(m.SplitsStandard[i]) { // not required
				return nil
			}

			if err := m.SplitsStandard[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("splits_standard" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("splits_standard" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedActivity) UnmarshalBinary(b []byte) error {
	var res DetailedActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
