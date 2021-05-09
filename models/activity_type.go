// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ActivityType An enumeration of the types an activity may have.
//
// swagger:model activityType
type ActivityType string

func NewActivityType(value ActivityType) *ActivityType {
	v := value
	return &v
}

const (

	// ActivityTypeAlpineSki captures enum value "AlpineSki"
	ActivityTypeAlpineSki ActivityType = "AlpineSki"

	// ActivityTypeBackcountrySki captures enum value "BackcountrySki"
	ActivityTypeBackcountrySki ActivityType = "BackcountrySki"

	// ActivityTypeCanoeing captures enum value "Canoeing"
	ActivityTypeCanoeing ActivityType = "Canoeing"

	// ActivityTypeCrossfit captures enum value "Crossfit"
	ActivityTypeCrossfit ActivityType = "Crossfit"

	// ActivityTypeEBikeRide captures enum value "EBikeRide"
	ActivityTypeEBikeRide ActivityType = "EBikeRide"

	// ActivityTypeElliptical captures enum value "Elliptical"
	ActivityTypeElliptical ActivityType = "Elliptical"

	// ActivityTypeGolf captures enum value "Golf"
	ActivityTypeGolf ActivityType = "Golf"

	// ActivityTypeHandcycle captures enum value "Handcycle"
	ActivityTypeHandcycle ActivityType = "Handcycle"

	// ActivityTypeHike captures enum value "Hike"
	ActivityTypeHike ActivityType = "Hike"

	// ActivityTypeIceSkate captures enum value "IceSkate"
	ActivityTypeIceSkate ActivityType = "IceSkate"

	// ActivityTypeInlineSkate captures enum value "InlineSkate"
	ActivityTypeInlineSkate ActivityType = "InlineSkate"

	// ActivityTypeKayaking captures enum value "Kayaking"
	ActivityTypeKayaking ActivityType = "Kayaking"

	// ActivityTypeKitesurf captures enum value "Kitesurf"
	ActivityTypeKitesurf ActivityType = "Kitesurf"

	// ActivityTypeNordicSki captures enum value "NordicSki"
	ActivityTypeNordicSki ActivityType = "NordicSki"

	// ActivityTypeRide captures enum value "Ride"
	ActivityTypeRide ActivityType = "Ride"

	// ActivityTypeRockClimbing captures enum value "RockClimbing"
	ActivityTypeRockClimbing ActivityType = "RockClimbing"

	// ActivityTypeRollerSki captures enum value "RollerSki"
	ActivityTypeRollerSki ActivityType = "RollerSki"

	// ActivityTypeRowing captures enum value "Rowing"
	ActivityTypeRowing ActivityType = "Rowing"

	// ActivityTypeRun captures enum value "Run"
	ActivityTypeRun ActivityType = "Run"

	// ActivityTypeSail captures enum value "Sail"
	ActivityTypeSail ActivityType = "Sail"

	// ActivityTypeSkateboard captures enum value "Skateboard"
	ActivityTypeSkateboard ActivityType = "Skateboard"

	// ActivityTypeSnowboard captures enum value "Snowboard"
	ActivityTypeSnowboard ActivityType = "Snowboard"

	// ActivityTypeSnowshoe captures enum value "Snowshoe"
	ActivityTypeSnowshoe ActivityType = "Snowshoe"

	// ActivityTypeSoccer captures enum value "Soccer"
	ActivityTypeSoccer ActivityType = "Soccer"

	// ActivityTypeStairStepper captures enum value "StairStepper"
	ActivityTypeStairStepper ActivityType = "StairStepper"

	// ActivityTypeStandUpPaddling captures enum value "StandUpPaddling"
	ActivityTypeStandUpPaddling ActivityType = "StandUpPaddling"

	// ActivityTypeSurfing captures enum value "Surfing"
	ActivityTypeSurfing ActivityType = "Surfing"

	// ActivityTypeSwim captures enum value "Swim"
	ActivityTypeSwim ActivityType = "Swim"

	// ActivityTypeVelomobile captures enum value "Velomobile"
	ActivityTypeVelomobile ActivityType = "Velomobile"

	// ActivityTypeVirtualRide captures enum value "VirtualRide"
	ActivityTypeVirtualRide ActivityType = "VirtualRide"

	// ActivityTypeVirtualRun captures enum value "VirtualRun"
	ActivityTypeVirtualRun ActivityType = "VirtualRun"

	// ActivityTypeWalk captures enum value "Walk"
	ActivityTypeWalk ActivityType = "Walk"

	// ActivityTypeWeightTraining captures enum value "WeightTraining"
	ActivityTypeWeightTraining ActivityType = "WeightTraining"

	// ActivityTypeWheelchair captures enum value "Wheelchair"
	ActivityTypeWheelchair ActivityType = "Wheelchair"

	// ActivityTypeWindsurf captures enum value "Windsurf"
	ActivityTypeWindsurf ActivityType = "Windsurf"

	// ActivityTypeWorkout captures enum value "Workout"
	ActivityTypeWorkout ActivityType = "Workout"

	// ActivityTypeYoga captures enum value "Yoga"
	ActivityTypeYoga ActivityType = "Yoga"
)

// for schema
var activityTypeEnum []interface{}

func init() {
	var res []ActivityType
	if err := json.Unmarshal([]byte(`["AlpineSki","BackcountrySki","Canoeing","Crossfit","EBikeRide","Elliptical","Golf","Handcycle","Hike","IceSkate","InlineSkate","Kayaking","Kitesurf","NordicSki","Ride","RockClimbing","RollerSki","Rowing","Run","Sail","Skateboard","Snowboard","Snowshoe","Soccer","StairStepper","StandUpPaddling","Surfing","Swim","Velomobile","VirtualRide","VirtualRun","Walk","WeightTraining","Wheelchair","Windsurf","Workout","Yoga"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activityTypeEnum = append(activityTypeEnum, v)
	}
}

func (m ActivityType) validateActivityTypeEnum(path, location string, value ActivityType) error {
	if err := validate.EnumCase(path, location, value, activityTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this activity type
func (m ActivityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActivityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this activity type based on context it is used
func (m ActivityType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
