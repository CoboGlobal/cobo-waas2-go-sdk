/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
	"fmt"
)

// EstimationFee - struct for EstimationFee
type EstimationFee struct {
	EvmEip1559Fee *EvmEip1559Fee
	EvmLegacyFee *EvmLegacyFee
	FixedFee *FixedFee
	UtxoFee *UtxoFee
}

// EvmEip1559FeeAsEstimationFee is a convenience function that returns EvmEip1559Fee wrapped in EstimationFee
func EvmEip1559FeeAsEstimationFee(v *EvmEip1559Fee) EstimationFee {
	return EstimationFee{
		EvmEip1559Fee: v,
	}
}

// EvmLegacyFeeAsEstimationFee is a convenience function that returns EvmLegacyFee wrapped in EstimationFee
func EvmLegacyFeeAsEstimationFee(v *EvmLegacyFee) EstimationFee {
	return EstimationFee{
		EvmLegacyFee: v,
	}
}

// FixedFeeAsEstimationFee is a convenience function that returns FixedFee wrapped in EstimationFee
func FixedFeeAsEstimationFee(v *FixedFee) EstimationFee {
	return EstimationFee{
		FixedFee: v,
	}
}

// UtxoFeeAsEstimationFee is a convenience function that returns UtxoFee wrapped in EstimationFee
func UtxoFeeAsEstimationFee(v *UtxoFee) EstimationFee {
	return EstimationFee{
		UtxoFee: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EstimationFee) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EVM_EIP_1559'
	if jsonDict["fee_type"] == "EVM_EIP_1559" {
		// try to unmarshal JSON data into EvmEip1559Fee
		err = json.Unmarshal(data, &dst.EvmEip1559Fee)
		if err == nil {
			return nil // data stored in dst.EvmEip1559Fee, return on the first match
		} else {
			dst.EvmEip1559Fee = nil
			return fmt.Errorf("failed to unmarshal EstimationFee as EvmEip1559Fee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EVM_Legacy'
	if jsonDict["fee_type"] == "EVM_Legacy" {
		// try to unmarshal JSON data into EvmLegacyFee
		err = json.Unmarshal(data, &dst.EvmLegacyFee)
		if err == nil {
			return nil // data stored in dst.EvmLegacyFee, return on the first match
		} else {
			dst.EvmLegacyFee = nil
			return fmt.Errorf("failed to unmarshal EstimationFee as EvmLegacyFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Fixed'
	if jsonDict["fee_type"] == "Fixed" {
		// try to unmarshal JSON data into FixedFee
		err = json.Unmarshal(data, &dst.FixedFee)
		if err == nil {
			return nil // data stored in dst.FixedFee, return on the first match
		} else {
			dst.FixedFee = nil
			return fmt.Errorf("failed to unmarshal EstimationFee as FixedFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UTXO'
	if jsonDict["fee_type"] == "UTXO" {
		// try to unmarshal JSON data into UtxoFee
		err = json.Unmarshal(data, &dst.UtxoFee)
		if err == nil {
			return nil // data stored in dst.UtxoFee, return on the first match
		} else {
			dst.UtxoFee = nil
			return fmt.Errorf("failed to unmarshal EstimationFee as UtxoFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EvmEip1559Fee'
	if jsonDict["fee_type"] == "EvmEip1559Fee" {
		// try to unmarshal JSON data into EvmEip1559Fee
		err = json.Unmarshal(data, &dst.EvmEip1559Fee)
		if err == nil {
			return nil // data stored in dst.EvmEip1559Fee, return on the first match
		} else {
			dst.EvmEip1559Fee = nil
			return fmt.Errorf("failed to unmarshal EstimationFee as EvmEip1559Fee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EvmLegacyFee'
	if jsonDict["fee_type"] == "EvmLegacyFee" {
		// try to unmarshal JSON data into EvmLegacyFee
		err = json.Unmarshal(data, &dst.EvmLegacyFee)
		if err == nil {
			return nil // data stored in dst.EvmLegacyFee, return on the first match
		} else {
			dst.EvmLegacyFee = nil
			return fmt.Errorf("failed to unmarshal EstimationFee as EvmLegacyFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'FixedFee'
	if jsonDict["fee_type"] == "FixedFee" {
		// try to unmarshal JSON data into FixedFee
		err = json.Unmarshal(data, &dst.FixedFee)
		if err == nil {
			return nil // data stored in dst.FixedFee, return on the first match
		} else {
			dst.FixedFee = nil
			return fmt.Errorf("failed to unmarshal EstimationFee as FixedFee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'UtxoFee'
	if jsonDict["fee_type"] == "UtxoFee" {
		// try to unmarshal JSON data into UtxoFee
		err = json.Unmarshal(data, &dst.UtxoFee)
		if err == nil {
			return nil // data stored in dst.UtxoFee, return on the first match
		} else {
			dst.UtxoFee = nil
			return fmt.Errorf("failed to unmarshal EstimationFee as UtxoFee: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EstimationFee) MarshalJSON() ([]byte, error) {
	if src.EvmEip1559Fee != nil {
		return json.Marshal(&src.EvmEip1559Fee)
	}

	if src.EvmLegacyFee != nil {
		return json.Marshal(&src.EvmLegacyFee)
	}

	if src.FixedFee != nil {
		return json.Marshal(&src.FixedFee)
	}

	if src.UtxoFee != nil {
		return json.Marshal(&src.UtxoFee)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EstimationFee) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EvmEip1559Fee != nil {
		return obj.EvmEip1559Fee
	}

	if obj.EvmLegacyFee != nil {
		return obj.EvmLegacyFee
	}

	if obj.FixedFee != nil {
		return obj.FixedFee
	}

	if obj.UtxoFee != nil {
		return obj.UtxoFee
	}

	// all schemas are nil
	return nil
}

type NullableEstimationFee struct {
	value *EstimationFee
	isSet bool
}

func (v NullableEstimationFee) Get() *EstimationFee {
	return v.value
}

func (v *NullableEstimationFee) Set(val *EstimationFee) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimationFee) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimationFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimationFee(val *EstimationFee) *NullableEstimationFee {
	return &NullableEstimationFee{value: val, isSet: true}
}

func (v NullableEstimationFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimationFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


