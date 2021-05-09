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

// ExplorerSegment explorer segment
//
// swagger:model explorerSegment
type ExplorerSegment struct {

	// The segment's average grade, in percents
	AvgGrade float32 `json:"avg_grade,omitempty"`

	// The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category. If climb_category = 5, climb_category_desc = HC. If climb_category = 2, climb_category_desc = 3.
	// Maximum: 5
	// Minimum: 0
	ClimbCategory *int64 `json:"climb_category,omitempty"`

	// The description for the category of the climb
	// Enum: [NC 4 3 2 1 HC]
	ClimbCategoryDesc string `json:"climb_category_desc,omitempty"`

	// The segment's distance, in meters
	Distance float32 `json:"distance,omitempty"`

	// The segments's evelation difference, in meters
	ElevDifference float32 `json:"elev_difference,omitempty"`

	// end latlng
	EndLatlng LatLng `json:"end_latlng,omitempty"`

	// The unique identifier of this segment
	ID int64 `json:"id,omitempty"`

	// The name of this segment
	Name string `json:"name,omitempty"`

	// The polyline of the segment
	Points string `json:"points,omitempty"`

	// start latlng
	StartLatlng LatLng `json:"start_latlng,omitempty"`
}

// Validate validates this explorer segment
func (m *ExplorerSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClimbCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClimbCategoryDesc(formats); err != nil {
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

func (m *ExplorerSegment) validateClimbCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.ClimbCategory) { // not required
		return nil
	}

	if err := validate.MinimumInt("climb_category", "body", *m.ClimbCategory, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("climb_category", "body", *m.ClimbCategory, 5, false); err != nil {
		return err
	}

	return nil
}

var explorerSegmentTypeClimbCategoryDescPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NC","4","3","2","1","HC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		explorerSegmentTypeClimbCategoryDescPropEnum = append(explorerSegmentTypeClimbCategoryDescPropEnum, v)
	}
}

const (

	// ExplorerSegmentClimbCategoryDescNC captures enum value "NC"
	ExplorerSegmentClimbCategoryDescNC string = "NC"

	// ExplorerSegmentClimbCategoryDescNr4 captures enum value "4"
	ExplorerSegmentClimbCategoryDescNr4 string = "4"

	// ExplorerSegmentClimbCategoryDescNr3 captures enum value "3"
	ExplorerSegmentClimbCategoryDescNr3 string = "3"

	// ExplorerSegmentClimbCategoryDescNr2 captures enum value "2"
	ExplorerSegmentClimbCategoryDescNr2 string = "2"

	// ExplorerSegmentClimbCategoryDescNr1 captures enum value "1"
	ExplorerSegmentClimbCategoryDescNr1 string = "1"

	// ExplorerSegmentClimbCategoryDescHC captures enum value "HC"
	ExplorerSegmentClimbCategoryDescHC string = "HC"
)

// prop value enum
func (m *ExplorerSegment) validateClimbCategoryDescEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, explorerSegmentTypeClimbCategoryDescPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExplorerSegment) validateClimbCategoryDesc(formats strfmt.Registry) error {
	if swag.IsZero(m.ClimbCategoryDesc) { // not required
		return nil
	}

	// value enum
	if err := m.validateClimbCategoryDescEnum("climb_category_desc", "body", m.ClimbCategoryDesc); err != nil {
		return err
	}

	return nil
}

func (m *ExplorerSegment) validateEndLatlng(formats strfmt.Registry) error {
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

func (m *ExplorerSegment) validateStartLatlng(formats strfmt.Registry) error {
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

// ContextValidate validate this explorer segment based on the context it is used
func (m *ExplorerSegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *ExplorerSegment) contextValidateEndLatlng(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EndLatlng.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("end_latlng")
		}
		return err
	}

	return nil
}

func (m *ExplorerSegment) contextValidateStartLatlng(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StartLatlng.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("start_latlng")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExplorerSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExplorerSegment) UnmarshalBinary(b []byte) error {
	var res ExplorerSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
