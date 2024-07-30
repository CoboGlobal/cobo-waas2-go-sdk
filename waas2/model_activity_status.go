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

// ActivityStatus The status of activity.
type ActivityStatus string

// List of ActivityStatus
const (
	ACTIVITYSTATUS_SUCCESS ActivityStatus = "Success"
	ACTIVITYSTATUS_PROCESSING ActivityStatus = "Processing"
	ACTIVITYSTATUS_FAILED ActivityStatus = "Failed"
)

// All allowed values of ActivityStatus enum
var AllowedActivityStatusEnumValues = []ActivityStatus{
	"Success",
	"Processing",
	"Failed",
}

func (v *ActivityStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActivityStatus(value)
	for _, existing := range AllowedActivityStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = ActivityStatus("unknown")
    return nil
}

// NewActivityStatusFromValue returns a pointer to a valid ActivityStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActivityStatusFromValue(v string) (*ActivityStatus, error) {
	ev := ActivityStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActivityStatus: valid values are %v", v, AllowedActivityStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActivityStatus) IsValid() bool {
	for _, existing := range AllowedActivityStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActivityStatus value
func (v ActivityStatus) Ptr() *ActivityStatus {
	return &v
}

type NullableActivityStatus struct {
	value *ActivityStatus
	isSet bool
}

func (v NullableActivityStatus) Get() *ActivityStatus {
	return v.value
}

func (v *NullableActivityStatus) Set(val *ActivityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityStatus(val *ActivityStatus) *NullableActivityStatus {
	return &NullableActivityStatus{value: val, isSet: true}
}

func (v NullableActivityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
