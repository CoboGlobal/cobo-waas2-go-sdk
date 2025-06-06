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

// TransactionInitiatorType The transaction initiator type. Possible values include:   - `API`: An API initiator, who initiates the transaction by using the WaaS API.   - `Web`: An web initiator, who initiates the transaction from Cobo Portal.   - `App`: An App initiator, who initiates the transaction from Cobo Portal Apps.   - `External`: An external initiator, who initiates the transaction outside Cobo. 
type TransactionInitiatorType string

// List of TransactionInitiatorType
const (
	TRANSACTIONINITIATORTYPE_API TransactionInitiatorType = "API"
	TRANSACTIONINITIATORTYPE_WEB TransactionInitiatorType = "Web"
	TRANSACTIONINITIATORTYPE_APP TransactionInitiatorType = "App"
	TRANSACTIONINITIATORTYPE_EXTERNAL TransactionInitiatorType = "External"
)

// All allowed values of TransactionInitiatorType enum
var AllowedTransactionInitiatorTypeEnumValues = []TransactionInitiatorType{
	"API",
	"Web",
	"App",
	"External",
}

func (v *TransactionInitiatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionInitiatorType(value)
	for _, existing := range AllowedTransactionInitiatorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = TransactionInitiatorType("unknown")
	return nil
}

// NewTransactionInitiatorTypeFromValue returns a pointer to a valid TransactionInitiatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionInitiatorTypeFromValue(v string) (*TransactionInitiatorType, error) {
	ev := TransactionInitiatorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionInitiatorType: valid values are %v", v, AllowedTransactionInitiatorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionInitiatorType) IsValid() bool {
	for _, existing := range AllowedTransactionInitiatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionInitiatorType value
func (v TransactionInitiatorType) Ptr() *TransactionInitiatorType {
	return &v
}

type NullableTransactionInitiatorType struct {
	value *TransactionInitiatorType
	isSet bool
}

func (v NullableTransactionInitiatorType) Get() *TransactionInitiatorType {
	return v.value
}

func (v *NullableTransactionInitiatorType) Set(val *TransactionInitiatorType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInitiatorType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInitiatorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInitiatorType(val *TransactionInitiatorType) *NullableTransactionInitiatorType {
	return &NullableTransactionInitiatorType{value: val, isSet: true}
}

func (v NullableTransactionInitiatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInitiatorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

