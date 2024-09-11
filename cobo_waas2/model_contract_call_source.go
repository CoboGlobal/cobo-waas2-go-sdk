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

// ContractCallSource - struct for ContractCallSource
type ContractCallSource struct {
	MpcContractCallSource *MpcContractCallSource
	SafeContractCallSource *SafeContractCallSource
}

// MpcContractCallSourceAsContractCallSource is a convenience function that returns MpcContractCallSource wrapped in ContractCallSource
func MpcContractCallSourceAsContractCallSource(v *MpcContractCallSource) ContractCallSource {
	return ContractCallSource{
		MpcContractCallSource: v,
	}
}

// SafeContractCallSourceAsContractCallSource is a convenience function that returns SafeContractCallSource wrapped in ContractCallSource
func SafeContractCallSourceAsContractCallSource(v *SafeContractCallSource) ContractCallSource {
	return ContractCallSource{
		SafeContractCallSource: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContractCallSource) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Org-Controlled'
	if jsonDict["source_type"] == "Org-Controlled" {
		// try to unmarshal JSON data into MpcContractCallSource
		err = json.Unmarshal(data, &dst.MpcContractCallSource)
		if err == nil {
			return nil // data stored in dst.MpcContractCallSource, return on the first match
		} else {
			dst.MpcContractCallSource = nil
			return fmt.Errorf("failed to unmarshal ContractCallSource as MpcContractCallSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Safe{Wallet}'
	if jsonDict["source_type"] == "Safe{Wallet}" {
		// try to unmarshal JSON data into SafeContractCallSource
		err = json.Unmarshal(data, &dst.SafeContractCallSource)
		if err == nil {
			return nil // data stored in dst.SafeContractCallSource, return on the first match
		} else {
			dst.SafeContractCallSource = nil
			return fmt.Errorf("failed to unmarshal ContractCallSource as SafeContractCallSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'User-Controlled'
	if jsonDict["source_type"] == "User-Controlled" {
		// try to unmarshal JSON data into MpcContractCallSource
		err = json.Unmarshal(data, &dst.MpcContractCallSource)
		if err == nil {
			return nil // data stored in dst.MpcContractCallSource, return on the first match
		} else {
			dst.MpcContractCallSource = nil
			return fmt.Errorf("failed to unmarshal ContractCallSource as MpcContractCallSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MpcContractCallSource'
	if jsonDict["source_type"] == "MpcContractCallSource" {
		// try to unmarshal JSON data into MpcContractCallSource
		err = json.Unmarshal(data, &dst.MpcContractCallSource)
		if err == nil {
			return nil // data stored in dst.MpcContractCallSource, return on the first match
		} else {
			dst.MpcContractCallSource = nil
			return fmt.Errorf("failed to unmarshal ContractCallSource as MpcContractCallSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SafeContractCallSource'
	if jsonDict["source_type"] == "SafeContractCallSource" {
		// try to unmarshal JSON data into SafeContractCallSource
		err = json.Unmarshal(data, &dst.SafeContractCallSource)
		if err == nil {
			return nil // data stored in dst.SafeContractCallSource, return on the first match
		} else {
			dst.SafeContractCallSource = nil
			return fmt.Errorf("failed to unmarshal ContractCallSource as SafeContractCallSource: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContractCallSource) MarshalJSON() ([]byte, error) {
	if src.MpcContractCallSource != nil {
		return json.Marshal(&src.MpcContractCallSource)
	}

	if src.SafeContractCallSource != nil {
		return json.Marshal(&src.SafeContractCallSource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContractCallSource) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MpcContractCallSource != nil {
		return obj.MpcContractCallSource
	}

	if obj.SafeContractCallSource != nil {
		return obj.SafeContractCallSource
	}

	// all schemas are nil
	return nil
}

type NullableContractCallSource struct {
	value *ContractCallSource
	isSet bool
}

func (v NullableContractCallSource) Get() *ContractCallSource {
	return v.value
}

func (v *NullableContractCallSource) Set(val *ContractCallSource) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCallSource) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCallSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCallSource(val *ContractCallSource) *NullableContractCallSource {
	return &NullableContractCallSource{value: val, isSet: true}
}

func (v NullableContractCallSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCallSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


