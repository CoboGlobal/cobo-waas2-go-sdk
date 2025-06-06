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

// GuardPubkeyStatus The status of a Cobo Guard public key binding. Possible values include:    - `New`: The binding is created.   - `ChangeNew`: A new binding is created    - `WaitSelfConfirm`: The binding is waiting for user confirmation on the old Cobo Guard.   - `WaitConfirm`: The binding is waiting for admin confirmation.   - `WaitActive`: The binding is waiting to become active.   - `Active`: The binding has come into effect.   - `Freeze`: The binding is frozen.   - `Invalid`: The binding is invalid. 
type GuardPubkeyStatus string

// List of GuardPubkeyStatus
const (
	GUARDPUBKEYSTATUS_NEW GuardPubkeyStatus = "New"
	GUARDPUBKEYSTATUS_CHANGE_NEW GuardPubkeyStatus = "ChangeNew"
	GUARDPUBKEYSTATUS_WAIT_SELF_CONFIRM GuardPubkeyStatus = "WaitSelfConfirm"
	GUARDPUBKEYSTATUS_WAIT_CONFIRM GuardPubkeyStatus = "WaitConfirm"
	GUARDPUBKEYSTATUS_WAIT_ACTIVE GuardPubkeyStatus = "WaitActive"
	GUARDPUBKEYSTATUS_ACTIVE GuardPubkeyStatus = "Active"
	GUARDPUBKEYSTATUS_FREEZE GuardPubkeyStatus = "Freeze"
	GUARDPUBKEYSTATUS_INVALID GuardPubkeyStatus = "Invalid"
)

// All allowed values of GuardPubkeyStatus enum
var AllowedGuardPubkeyStatusEnumValues = []GuardPubkeyStatus{
	"New",
	"ChangeNew",
	"WaitSelfConfirm",
	"WaitConfirm",
	"WaitActive",
	"Active",
	"Freeze",
	"Invalid",
}

func (v *GuardPubkeyStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GuardPubkeyStatus(value)
	for _, existing := range AllowedGuardPubkeyStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = GuardPubkeyStatus("unknown")
	return nil
}

// NewGuardPubkeyStatusFromValue returns a pointer to a valid GuardPubkeyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGuardPubkeyStatusFromValue(v string) (*GuardPubkeyStatus, error) {
	ev := GuardPubkeyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GuardPubkeyStatus: valid values are %v", v, AllowedGuardPubkeyStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GuardPubkeyStatus) IsValid() bool {
	for _, existing := range AllowedGuardPubkeyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuardPubkeyStatus value
func (v GuardPubkeyStatus) Ptr() *GuardPubkeyStatus {
	return &v
}

type NullableGuardPubkeyStatus struct {
	value *GuardPubkeyStatus
	isSet bool
}

func (v NullableGuardPubkeyStatus) Get() *GuardPubkeyStatus {
	return v.value
}

func (v *NullableGuardPubkeyStatus) Set(val *GuardPubkeyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGuardPubkeyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGuardPubkeyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuardPubkeyStatus(val *GuardPubkeyStatus) *NullableGuardPubkeyStatus {
	return &NullableGuardPubkeyStatus{value: val, isSet: true}
}

func (v NullableGuardPubkeyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuardPubkeyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

