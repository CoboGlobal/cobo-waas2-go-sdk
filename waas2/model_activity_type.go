/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"fmt"
)

// ActivityType The type of activity.
type ActivityType string

// List of ActivityType
const (
	ACTIVITYTYPE_STAKE ActivityType = "Stake"
	ACTIVITYTYPE_UNSTAKE ActivityType = "Unstake"
	ACTIVITYTYPE_WITHDRAW ActivityType = "Withdraw"
	ACTIVITYTYPE_CLAIM ActivityType = "Claim"
	ACTIVITYTYPE_DELEGATE ActivityType = "Delegate"
	ACTIVITYTYPE_UNDELEGATE ActivityType = "Undelegate"
)

// All allowed values of ActivityType enum
var AllowedActivityTypeEnumValues = []ActivityType{
	"Stake",
	"Unstake",
	"Withdraw",
	"Claim",
	"Delegate",
	"Undelegate",
}

func (v *ActivityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActivityType(value)
	for _, existing := range AllowedActivityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = ActivityType("unknown")
    return nil
}

// NewActivityTypeFromValue returns a pointer to a valid ActivityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActivityTypeFromValue(v string) (*ActivityType, error) {
	ev := ActivityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActivityType: valid values are %v", v, AllowedActivityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActivityType) IsValid() bool {
	for _, existing := range AllowedActivityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActivityType value
func (v ActivityType) Ptr() *ActivityType {
	return &v
}

type NullableActivityType struct {
	value *ActivityType
	isSet bool
}

func (v NullableActivityType) Get() *ActivityType {
	return v.value
}

func (v *NullableActivityType) Set(val *ActivityType) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityType) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityType(val *ActivityType) *NullableActivityType {
	return &NullableActivityType{value: val, isSet: true}
}

func (v NullableActivityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

