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

// ApprovalResult The single approval result. Possible values include:    - `1`: The transaction is pending approval.   - `2`: The transaction is approved.   - `3`: The transaction is rejected. 
type ApprovalResult int32

// List of ApprovalResult
const (
	APPROVALRESULT_PENDING ApprovalResult = 1
	APPROVALRESULT_APPROVED ApprovalResult = 2
	APPROVALRESULT_REJECTED ApprovalResult = 3
)

// All allowed values of ApprovalResult enum
var AllowedApprovalResultEnumValues = []ApprovalResult{
	1,
	2,
	3,
}

func (v *ApprovalResult) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApprovalResult(value)
	for _, existing := range AllowedApprovalResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = ApprovalResult(-1)
	return nil
}

// NewApprovalResultFromValue returns a pointer to a valid ApprovalResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApprovalResultFromValue(v int32) (*ApprovalResult, error) {
	ev := ApprovalResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApprovalResult: valid values are %v", v, AllowedApprovalResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApprovalResult) IsValid() bool {
	for _, existing := range AllowedApprovalResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApprovalResult value
func (v ApprovalResult) Ptr() *ApprovalResult {
	return &v
}

type NullableApprovalResult struct {
	value *ApprovalResult
	isSet bool
}

func (v NullableApprovalResult) Get() *ApprovalResult {
	return v.value
}

func (v *NullableApprovalResult) Set(val *ApprovalResult) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalResult) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalResult(val *ApprovalResult) *NullableApprovalResult {
	return &NullableApprovalResult{value: val, isSet: true}
}

func (v NullableApprovalResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

