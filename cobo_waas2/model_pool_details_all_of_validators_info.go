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

// PoolDetailsAllOfValidatorsInfo - struct for PoolDetailsAllOfValidatorsInfo
type PoolDetailsAllOfValidatorsInfo struct {
	BabylonValidator *BabylonValidator
}

// BabylonValidatorAsPoolDetailsAllOfValidatorsInfo is a convenience function that returns BabylonValidator wrapped in PoolDetailsAllOfValidatorsInfo
func BabylonValidatorAsPoolDetailsAllOfValidatorsInfo(v *BabylonValidator) PoolDetailsAllOfValidatorsInfo {
	return PoolDetailsAllOfValidatorsInfo{
		BabylonValidator: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PoolDetailsAllOfValidatorsInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Babylon'
	if jsonDict["pool_type"] == "Babylon" {
		// try to unmarshal JSON data into BabylonValidator
		err = json.Unmarshal(data, &dst.BabylonValidator)
		if err == nil {
			return nil // data stored in dst.BabylonValidator, return on the first match
		} else {
			dst.BabylonValidator = nil
			return fmt.Errorf("failed to unmarshal PoolDetailsAllOfValidatorsInfo as BabylonValidator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BabylonValidator'
	if jsonDict["pool_type"] == "BabylonValidator" {
		// try to unmarshal JSON data into BabylonValidator
		err = json.Unmarshal(data, &dst.BabylonValidator)
		if err == nil {
			return nil // data stored in dst.BabylonValidator, return on the first match
		} else {
			dst.BabylonValidator = nil
			return fmt.Errorf("failed to unmarshal PoolDetailsAllOfValidatorsInfo as BabylonValidator: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PoolDetailsAllOfValidatorsInfo) MarshalJSON() ([]byte, error) {
	if src.BabylonValidator != nil {
		return json.Marshal(&src.BabylonValidator)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PoolDetailsAllOfValidatorsInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BabylonValidator != nil {
		return obj.BabylonValidator
	}

	// all schemas are nil
	return nil
}

type NullablePoolDetailsAllOfValidatorsInfo struct {
	value *PoolDetailsAllOfValidatorsInfo
	isSet bool
}

func (v NullablePoolDetailsAllOfValidatorsInfo) Get() *PoolDetailsAllOfValidatorsInfo {
	return v.value
}

func (v *NullablePoolDetailsAllOfValidatorsInfo) Set(val *PoolDetailsAllOfValidatorsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDetailsAllOfValidatorsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDetailsAllOfValidatorsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDetailsAllOfValidatorsInfo(val *PoolDetailsAllOfValidatorsInfo) *NullablePoolDetailsAllOfValidatorsInfo {
	return &NullablePoolDetailsAllOfValidatorsInfo{value: val, isSet: true}
}

func (v NullablePoolDetailsAllOfValidatorsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDetailsAllOfValidatorsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


