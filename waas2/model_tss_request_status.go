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

// TSSRequestStatus The TSS request status. Possible values include: - `PendingKeyHolderConfirmation`: The action done to the TSS request is currently pending enough key share holders to approve.  - `KeyHolderConfirmationFailed`: Key share holders failed to approve the the action to be done to the TSS request.  - `KeyGenerating`: The key share is currently being generated for the action to be done to the TSS request.  - `KeyGeneratingFailed`: The key share generation process has failed for the action to be done to the TSS request.  - `Success`: The action done to the TSS request has been completed successfully. If you see this status while running [Cancel TSS request](http://localhost:3000/v2/api-references/wallets--mpc-wallets/cancel-tss-request), this mean the specified TSS request has been successfully canceled. 
type TSSRequestStatus string

// List of TSSRequestStatus
const (
	TSSREQUESTSTATUS_PENDING_KEY_HOLDER_CONFIRMATION TSSRequestStatus = "PendingKeyHolderConfirmation"
	TSSREQUESTSTATUS_KEY_HOLDER_CONFIRMATION_FAILED TSSRequestStatus = "KeyHolderConfirmationFailed"
	TSSREQUESTSTATUS_KEY_GENERATING TSSRequestStatus = "KeyGenerating"
	TSSREQUESTSTATUS_KEY_GENERATING_FAILED TSSRequestStatus = "KeyGeneratingFailed"
	TSSREQUESTSTATUS_SUCCESS TSSRequestStatus = "Success"
)

// All allowed values of TSSRequestStatus enum
var AllowedTSSRequestStatusEnumValues = []TSSRequestStatus{
	"PendingKeyHolderConfirmation",
	"KeyHolderConfirmationFailed",
	"KeyGenerating",
	"KeyGeneratingFailed",
	"Success",
}

func (v *TSSRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TSSRequestStatus(value)
	for _, existing := range AllowedTSSRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = TSSRequestStatus("unknown")
    return nil
}

// NewTSSRequestStatusFromValue returns a pointer to a valid TSSRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTSSRequestStatusFromValue(v string) (*TSSRequestStatus, error) {
	ev := TSSRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TSSRequestStatus: valid values are %v", v, AllowedTSSRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TSSRequestStatus) IsValid() bool {
	for _, existing := range AllowedTSSRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TSSRequestStatus value
func (v TSSRequestStatus) Ptr() *TSSRequestStatus {
	return &v
}

type NullableTSSRequestStatus struct {
	value *TSSRequestStatus
	isSet bool
}

func (v NullableTSSRequestStatus) Get() *TSSRequestStatus {
	return v.value
}

func (v *NullableTSSRequestStatus) Set(val *TSSRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSRequestStatus(val *TSSRequestStatus) *NullableTSSRequestStatus {
	return &NullableTSSRequestStatus{value: val, isSet: true}
}

func (v NullableTSSRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

