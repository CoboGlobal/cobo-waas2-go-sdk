/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// TSSStatus The status of the request. Possible values include:  - 100 : Undefined. - 110 : Scheduling. - 120 : Initializing. - 130 : Approving. - 140 : Processing. - 160 : Declined. - 170 : Failed. - 180 : Canceled. - 190 : Completed. 
type TSSStatus int32

// List of TSSStatus
const (
	TSSSTATUS_UNDEFINED TSSStatus = 100
	TSSSTATUS_SCHEDULING TSSStatus = 110
	TSSSTATUS_INITIALIZING TSSStatus = 120
	TSSSTATUS_APPROVING TSSStatus = 130
	TSSSTATUS_PROCESSING TSSStatus = 140
	TSSSTATUS_DECLINED TSSStatus = 160
	TSSSTATUS_FAILED TSSStatus = 170
	TSSSTATUS_CANCELED TSSStatus = 180
	TSSSTATUS_COMPLETED TSSStatus = 190
)

// All allowed values of TSSStatus enum
var AllowedTSSStatusEnumValues = []TSSStatus{
	100,
	110,
	120,
	130,
	140,
	160,
	170,
	180,
	190,
}

func (v *TSSStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TSSStatus(value)
	for _, existing := range AllowedTSSStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = TSSStatus(-1)
	return nil
}

// NewTSSStatusFromValue returns a pointer to a valid TSSStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTSSStatusFromValue(v int32) (*TSSStatus, error) {
	ev := TSSStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TSSStatus: valid values are %v", v, AllowedTSSStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TSSStatus) IsValid() bool {
	for _, existing := range AllowedTSSStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TSSStatus value
func (v TSSStatus) Ptr() *TSSStatus {
	return &v
}

type NullableTSSStatus struct {
	value *TSSStatus
	isSet bool
}

func (v NullableTSSStatus) Get() *TSSStatus {
	return v.value
}

func (v *NullableTSSStatus) Set(val *TSSStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSStatus(val *TSSStatus) *NullableTSSStatus {
	return &NullableTSSStatus{value: val, isSet: true}
}

func (v NullableTSSStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

