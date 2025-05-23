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

// TransactionStatus The transaction status. For more details including sub-statuses, please refer to [Transaction statuses and sub-statuses](https://www.cobo.com/developers/v2/guides/transactions/status). 
type TransactionStatus string

// List of TransactionStatus
const (
	TRANSACTIONSTATUS_SUBMITTED TransactionStatus = "Submitted"
	TRANSACTIONSTATUS_PENDING_SCREENING TransactionStatus = "PendingScreening"
	TRANSACTIONSTATUS_PENDING_AUTHORIZATION TransactionStatus = "PendingAuthorization"
	TRANSACTIONSTATUS_PENDING_SIGNATURE TransactionStatus = "PendingSignature"
	TRANSACTIONSTATUS_BROADCASTING TransactionStatus = "Broadcasting"
	TRANSACTIONSTATUS_CONFIRMING TransactionStatus = "Confirming"
	TRANSACTIONSTATUS_COMPLETED TransactionStatus = "Completed"
	TRANSACTIONSTATUS_FAILED TransactionStatus = "Failed"
	TRANSACTIONSTATUS_REJECTED TransactionStatus = "Rejected"
	TRANSACTIONSTATUS_PENDING TransactionStatus = "Pending"
)

// All allowed values of TransactionStatus enum
var AllowedTransactionStatusEnumValues = []TransactionStatus{
	"Submitted",
	"PendingScreening",
	"PendingAuthorization",
	"PendingSignature",
	"Broadcasting",
	"Confirming",
	"Completed",
	"Failed",
	"Rejected",
	"Pending",
}

func (v *TransactionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionStatus(value)
	for _, existing := range AllowedTransactionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = TransactionStatus("unknown")
	return nil
}

// NewTransactionStatusFromValue returns a pointer to a valid TransactionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionStatusFromValue(v string) (*TransactionStatus, error) {
	ev := TransactionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionStatus: valid values are %v", v, AllowedTransactionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionStatus) IsValid() bool {
	for _, existing := range AllowedTransactionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionStatus value
func (v TransactionStatus) Ptr() *TransactionStatus {
	return &v
}

type NullableTransactionStatus struct {
	value *TransactionStatus
	isSet bool
}

func (v NullableTransactionStatus) Get() *TransactionStatus {
	return v.value
}

func (v *NullableTransactionStatus) Set(val *TransactionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStatus(val *TransactionStatus) *NullableTransactionStatus {
	return &NullableTransactionStatus{value: val, isSet: true}
}

func (v NullableTransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

