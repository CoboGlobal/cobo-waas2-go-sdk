/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// TransactionRequestFee - struct for TransactionRequestFee
type TransactionRequestFee struct {
	TransactionRequestEvmEip1559Fee *TransactionRequestEvmEip1559Fee
	TransactionRequestEvmLegacyFee *TransactionRequestEvmLegacyFee
	TransactionRequestFixedFee *TransactionRequestFixedFee
	TransactionRequestUtxoFee *TransactionRequestUtxoFee
}

// TransactionRequestEvmEip1559FeeAsTransactionRequestFee is a convenience function that returns TransactionRequestEvmEip1559Fee wrapped in TransactionRequestFee
func TransactionRequestEvmEip1559FeeAsTransactionRequestFee(v *TransactionRequestEvmEip1559Fee) TransactionRequestFee {
	return TransactionRequestFee{
		TransactionRequestEvmEip1559Fee: v,
	}
}

// TransactionRequestEvmLegacyFeeAsTransactionRequestFee is a convenience function that returns TransactionRequestEvmLegacyFee wrapped in TransactionRequestFee
func TransactionRequestEvmLegacyFeeAsTransactionRequestFee(v *TransactionRequestEvmLegacyFee) TransactionRequestFee {
	return TransactionRequestFee{
		TransactionRequestEvmLegacyFee: v,
	}
}

// TransactionRequestFixedFeeAsTransactionRequestFee is a convenience function that returns TransactionRequestFixedFee wrapped in TransactionRequestFee
func TransactionRequestFixedFeeAsTransactionRequestFee(v *TransactionRequestFixedFee) TransactionRequestFee {
	return TransactionRequestFee{
		TransactionRequestFixedFee: v,
	}
}

// TransactionRequestUtxoFeeAsTransactionRequestFee is a convenience function that returns TransactionRequestUtxoFee wrapped in TransactionRequestFee
func TransactionRequestUtxoFeeAsTransactionRequestFee(v *TransactionRequestUtxoFee) TransactionRequestFee {
	return TransactionRequestFee{
		TransactionRequestUtxoFee: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionRequestFee) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EVM_EIP_1559'
	if jsonDict["fee_type"] == "EVM_EIP_1559" {
		// try to unmarshal JSON data into TransactionRequestEvmEip1559Fee
		err = json.Unmarshal(data, &dst.TransactionRequestEvmEip1559Fee)
		if err == nil {
			return nil // data stored in dst.TransactionRequestEvmEip1559Fee, return on the first match
		} else {
			dst.TransactionRequestEvmEip1559Fee = nil
			return fmt.Errorf("failed to unmarshal TransactionRequestFee as TransactionRequestEvmEip1559Fee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EVM_Legacy'
	if jsonDict["fee_type"] == "EVM_Legacy" {
		// try to unmarshal JSON data into TransactionRequestEvmLegacyFee
		err = json.Unmarshal(data, &dst.TransactionRequestEvmLegacyFee)
		if err == nil {
			return nil // data stored in dst.TransactionRequestEvmLegacyFee, return on the first match
		} else {
			dst.TransactionRequestEvmLegacyFee = nil
			return fmt.Errorf("failed to unmarshal TransactionRequestFee as TransactionRequestEvmLegacyFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Fixed'
	if jsonDict["fee_type"] == "Fixed" {
		// try to unmarshal JSON data into TransactionRequestFixedFee
		err = json.Unmarshal(data, &dst.TransactionRequestFixedFee)
		if err == nil {
			return nil // data stored in dst.TransactionRequestFixedFee, return on the first match
		} else {
			dst.TransactionRequestFixedFee = nil
			return fmt.Errorf("failed to unmarshal TransactionRequestFee as TransactionRequestFixedFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UTXO'
	if jsonDict["fee_type"] == "UTXO" {
		// try to unmarshal JSON data into TransactionRequestUtxoFee
		err = json.Unmarshal(data, &dst.TransactionRequestUtxoFee)
		if err == nil {
			return nil // data stored in dst.TransactionRequestUtxoFee, return on the first match
		} else {
			dst.TransactionRequestUtxoFee = nil
			return fmt.Errorf("failed to unmarshal TransactionRequestFee as TransactionRequestUtxoFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionRequestEvmEip1559Fee'
	if jsonDict["fee_type"] == "TransactionRequestEvmEip1559Fee" {
		// try to unmarshal JSON data into TransactionRequestEvmEip1559Fee
		err = json.Unmarshal(data, &dst.TransactionRequestEvmEip1559Fee)
		if err == nil {
			return nil // data stored in dst.TransactionRequestEvmEip1559Fee, return on the first match
		} else {
			dst.TransactionRequestEvmEip1559Fee = nil
			return fmt.Errorf("failed to unmarshal TransactionRequestFee as TransactionRequestEvmEip1559Fee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionRequestEvmLegacyFee'
	if jsonDict["fee_type"] == "TransactionRequestEvmLegacyFee" {
		// try to unmarshal JSON data into TransactionRequestEvmLegacyFee
		err = json.Unmarshal(data, &dst.TransactionRequestEvmLegacyFee)
		if err == nil {
			return nil // data stored in dst.TransactionRequestEvmLegacyFee, return on the first match
		} else {
			dst.TransactionRequestEvmLegacyFee = nil
			return fmt.Errorf("failed to unmarshal TransactionRequestFee as TransactionRequestEvmLegacyFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionRequestFixedFee'
	if jsonDict["fee_type"] == "TransactionRequestFixedFee" {
		// try to unmarshal JSON data into TransactionRequestFixedFee
		err = json.Unmarshal(data, &dst.TransactionRequestFixedFee)
		if err == nil {
			return nil // data stored in dst.TransactionRequestFixedFee, return on the first match
		} else {
			dst.TransactionRequestFixedFee = nil
			return fmt.Errorf("failed to unmarshal TransactionRequestFee as TransactionRequestFixedFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionRequestUtxoFee'
	if jsonDict["fee_type"] == "TransactionRequestUtxoFee" {
		// try to unmarshal JSON data into TransactionRequestUtxoFee
		err = json.Unmarshal(data, &dst.TransactionRequestUtxoFee)
		if err == nil {
			return nil // data stored in dst.TransactionRequestUtxoFee, return on the first match
		} else {
			dst.TransactionRequestUtxoFee = nil
			return fmt.Errorf("failed to unmarshal TransactionRequestFee as TransactionRequestUtxoFee: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionRequestFee) MarshalJSON() ([]byte, error) {
	if src.TransactionRequestEvmEip1559Fee != nil {
		return json.Marshal(&src.TransactionRequestEvmEip1559Fee)
	}

	if src.TransactionRequestEvmLegacyFee != nil {
		return json.Marshal(&src.TransactionRequestEvmLegacyFee)
	}

	if src.TransactionRequestFixedFee != nil {
		return json.Marshal(&src.TransactionRequestFixedFee)
	}

	if src.TransactionRequestUtxoFee != nil {
		return json.Marshal(&src.TransactionRequestUtxoFee)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionRequestFee) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TransactionRequestEvmEip1559Fee != nil {
		return obj.TransactionRequestEvmEip1559Fee
	}

	if obj.TransactionRequestEvmLegacyFee != nil {
		return obj.TransactionRequestEvmLegacyFee
	}

	if obj.TransactionRequestFixedFee != nil {
		return obj.TransactionRequestFixedFee
	}

	if obj.TransactionRequestUtxoFee != nil {
		return obj.TransactionRequestUtxoFee
	}

	// all schemas are nil
	return nil
}

type NullableTransactionRequestFee struct {
	value *TransactionRequestFee
	isSet bool
}

func (v NullableTransactionRequestFee) Get() *TransactionRequestFee {
	return v.value
}

func (v *NullableTransactionRequestFee) Set(val *TransactionRequestFee) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequestFee) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequestFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequestFee(val *TransactionRequestFee) *NullableTransactionRequestFee {
	return &NullableTransactionRequestFee{value: val, isSet: true}
}

func (v NullableTransactionRequestFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequestFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


