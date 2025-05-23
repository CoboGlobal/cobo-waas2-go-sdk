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

// TransactionCoboCategory The transaction category defined by Cobo. Possible values include:    - `AutoSweep`: An auto-sweep transaction.   - `AutoFueling`: A transaction where Fee Station pays transaction fees to an address within your MPC Wallets.   - `AutoFuelingRefund`: A refund for an auto-fueling transaction.   - `SafeTxMessage`: A message signing transaction initiated by an MPC wallet to authorize a Smart Contract Wallet (Safe\\{Wallet\\}) transaction.   - `BillPayment`: A transaction to pay Cobo bills through Fee Station.   - `BillRefund`: A refund for a previously made bill payment.   - `CommissionFeeCharge`: A transaction to charge commission fees via Fee Station.   - `CommissionFeeRefund`: A refund of previously charged commission fees. 
type TransactionCoboCategory string

// List of TransactionCoboCategory
const (
	TRANSACTIONCOBOCATEGORY_AUTO_SWEEP TransactionCoboCategory = "AutoSweep"
	TRANSACTIONCOBOCATEGORY_AUTO_FUELING TransactionCoboCategory = "AutoFueling"
	TRANSACTIONCOBOCATEGORY_AUTO_FUELING_REFUND TransactionCoboCategory = "AutoFuelingRefund"
	TRANSACTIONCOBOCATEGORY_SAFE_TX_MESSAGE TransactionCoboCategory = "SafeTxMessage"
	TRANSACTIONCOBOCATEGORY_BILL_PAYMENT TransactionCoboCategory = "BillPayment"
	TRANSACTIONCOBOCATEGORY_BILL_REFUND TransactionCoboCategory = "BillRefund"
	TRANSACTIONCOBOCATEGORY_COMMISSION_FEE_CHARGE TransactionCoboCategory = "CommissionFeeCharge"
	TRANSACTIONCOBOCATEGORY_COMMISSION_FEE_REFUND TransactionCoboCategory = "CommissionFeeRefund"
)

// All allowed values of TransactionCoboCategory enum
var AllowedTransactionCoboCategoryEnumValues = []TransactionCoboCategory{
	"AutoSweep",
	"AutoFueling",
	"AutoFuelingRefund",
	"SafeTxMessage",
	"BillPayment",
	"BillRefund",
	"CommissionFeeCharge",
	"CommissionFeeRefund",
}

func (v *TransactionCoboCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionCoboCategory(value)
	for _, existing := range AllowedTransactionCoboCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = TransactionCoboCategory("unknown")
	return nil
}

// NewTransactionCoboCategoryFromValue returns a pointer to a valid TransactionCoboCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionCoboCategoryFromValue(v string) (*TransactionCoboCategory, error) {
	ev := TransactionCoboCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionCoboCategory: valid values are %v", v, AllowedTransactionCoboCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionCoboCategory) IsValid() bool {
	for _, existing := range AllowedTransactionCoboCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionCoboCategory value
func (v TransactionCoboCategory) Ptr() *TransactionCoboCategory {
	return &v
}

type NullableTransactionCoboCategory struct {
	value *TransactionCoboCategory
	isSet bool
}

func (v NullableTransactionCoboCategory) Get() *TransactionCoboCategory {
	return v.value
}

func (v *NullableTransactionCoboCategory) Set(val *TransactionCoboCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCoboCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCoboCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCoboCategory(val *TransactionCoboCategory) *NullableTransactionCoboCategory {
	return &NullableTransactionCoboCategory{value: val, isSet: true}
}

func (v NullableTransactionCoboCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCoboCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

