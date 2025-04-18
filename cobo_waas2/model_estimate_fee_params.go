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

// EstimateFeeParams - struct for EstimateFeeParams
type EstimateFeeParams struct {
	EstimateContractCallFeeParams *EstimateContractCallFeeParams
	EstimateTransferFeeParams *EstimateTransferFeeParams
}

// EstimateContractCallFeeParamsAsEstimateFeeParams is a convenience function that returns EstimateContractCallFeeParams wrapped in EstimateFeeParams
func EstimateContractCallFeeParamsAsEstimateFeeParams(v *EstimateContractCallFeeParams) EstimateFeeParams {
	return EstimateFeeParams{
		EstimateContractCallFeeParams: v,
	}
}

// EstimateTransferFeeParamsAsEstimateFeeParams is a convenience function that returns EstimateTransferFeeParams wrapped in EstimateFeeParams
func EstimateTransferFeeParamsAsEstimateFeeParams(v *EstimateTransferFeeParams) EstimateFeeParams {
	return EstimateFeeParams{
		EstimateTransferFeeParams: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EstimateFeeParams) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ContractCall'
	if jsonDict["request_type"] == "ContractCall" {
		// try to unmarshal JSON data into EstimateContractCallFeeParams
		err = json.Unmarshal(data, &dst.EstimateContractCallFeeParams)
		if err == nil {
			return nil // data stored in dst.EstimateContractCallFeeParams, return on the first match
		} else {
			dst.EstimateContractCallFeeParams = nil
			return fmt.Errorf("failed to unmarshal EstimateFeeParams as EstimateContractCallFeeParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Transfer'
	if jsonDict["request_type"] == "Transfer" {
		// try to unmarshal JSON data into EstimateTransferFeeParams
		err = json.Unmarshal(data, &dst.EstimateTransferFeeParams)
		if err == nil {
			return nil // data stored in dst.EstimateTransferFeeParams, return on the first match
		} else {
			dst.EstimateTransferFeeParams = nil
			return fmt.Errorf("failed to unmarshal EstimateFeeParams as EstimateTransferFeeParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EstimateContractCallFeeParams'
	if jsonDict["request_type"] == "EstimateContractCallFeeParams" {
		// try to unmarshal JSON data into EstimateContractCallFeeParams
		err = json.Unmarshal(data, &dst.EstimateContractCallFeeParams)
		if err == nil {
			return nil // data stored in dst.EstimateContractCallFeeParams, return on the first match
		} else {
			dst.EstimateContractCallFeeParams = nil
			return fmt.Errorf("failed to unmarshal EstimateFeeParams as EstimateContractCallFeeParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EstimateTransferFeeParams'
	if jsonDict["request_type"] == "EstimateTransferFeeParams" {
		// try to unmarshal JSON data into EstimateTransferFeeParams
		err = json.Unmarshal(data, &dst.EstimateTransferFeeParams)
		if err == nil {
			return nil // data stored in dst.EstimateTransferFeeParams, return on the first match
		} else {
			dst.EstimateTransferFeeParams = nil
			return fmt.Errorf("failed to unmarshal EstimateFeeParams as EstimateTransferFeeParams: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EstimateFeeParams) MarshalJSON() ([]byte, error) {
	if src.EstimateContractCallFeeParams != nil {
		return json.Marshal(&src.EstimateContractCallFeeParams)
	}

	if src.EstimateTransferFeeParams != nil {
		return json.Marshal(&src.EstimateTransferFeeParams)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EstimateFeeParams) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EstimateContractCallFeeParams != nil {
		return obj.EstimateContractCallFeeParams
	}

	if obj.EstimateTransferFeeParams != nil {
		return obj.EstimateTransferFeeParams
	}

	// all schemas are nil
	return nil
}

type NullableEstimateFeeParams struct {
	value *EstimateFeeParams
	isSet bool
}

func (v NullableEstimateFeeParams) Get() *EstimateFeeParams {
	return v.value
}

func (v *NullableEstimateFeeParams) Set(val *EstimateFeeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateFeeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateFeeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateFeeParams(val *EstimateFeeParams) *NullableEstimateFeeParams {
	return &NullableEstimateFeeParams{value: val, isSet: true}
}

func (v NullableEstimateFeeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateFeeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


