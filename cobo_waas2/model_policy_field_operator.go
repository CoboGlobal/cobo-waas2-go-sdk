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

// PolicyFieldOperator The policy field operator. Possible values include:    - `>`: Greater than.   - `>=`: Greater than or equal.   - `<`: Less than.   - `<=`: Less than or equal.   - `=`: Equal. 
type PolicyFieldOperator string

// List of PolicyFieldOperator
const (
	POLICYFIELDOPERATOR_GREATER_THAN PolicyFieldOperator = ">"
	POLICYFIELDOPERATOR_GREATER_THAN_OR_EQUAL_TO PolicyFieldOperator = ">="
	POLICYFIELDOPERATOR_LESS_THAN PolicyFieldOperator = "<"
	POLICYFIELDOPERATOR_LESS_THAN_OR_EQUAL_TO PolicyFieldOperator = "<="
	POLICYFIELDOPERATOR_EQUAL PolicyFieldOperator = "="
)

// All allowed values of PolicyFieldOperator enum
var AllowedPolicyFieldOperatorEnumValues = []PolicyFieldOperator{
	">",
	">=",
	"<",
	"<=",
	"=",
}

func (v *PolicyFieldOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyFieldOperator(value)
	for _, existing := range AllowedPolicyFieldOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = PolicyFieldOperator("unknown")
	return nil
}

// NewPolicyFieldOperatorFromValue returns a pointer to a valid PolicyFieldOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyFieldOperatorFromValue(v string) (*PolicyFieldOperator, error) {
	ev := PolicyFieldOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyFieldOperator: valid values are %v", v, AllowedPolicyFieldOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyFieldOperator) IsValid() bool {
	for _, existing := range AllowedPolicyFieldOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyFieldOperator value
func (v PolicyFieldOperator) Ptr() *PolicyFieldOperator {
	return &v
}

type NullablePolicyFieldOperator struct {
	value *PolicyFieldOperator
	isSet bool
}

func (v NullablePolicyFieldOperator) Get() *PolicyFieldOperator {
	return v.value
}

func (v *NullablePolicyFieldOperator) Set(val *PolicyFieldOperator) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyFieldOperator) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyFieldOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyFieldOperator(val *PolicyFieldOperator) *NullablePolicyFieldOperator {
	return &NullablePolicyFieldOperator{value: val, isSet: true}
}

func (v NullablePolicyFieldOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyFieldOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

