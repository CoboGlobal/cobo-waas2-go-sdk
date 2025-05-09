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

// CreateUnstakeActivityExtra - Additional protocol-specific information required for the unstaking operation. The required fields vary depending on the staking protocol.
type CreateUnstakeActivityExtra struct {
	EthUnstakeExtra *EthUnstakeExtra
}

// EthUnstakeExtraAsCreateUnstakeActivityExtra is a convenience function that returns EthUnstakeExtra wrapped in CreateUnstakeActivityExtra
func EthUnstakeExtraAsCreateUnstakeActivityExtra(v *EthUnstakeExtra) CreateUnstakeActivityExtra {
	return CreateUnstakeActivityExtra{
		EthUnstakeExtra: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateUnstakeActivityExtra) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ETHBeacon'
	if jsonDict["pool_type"] == "ETHBeacon" {
		// try to unmarshal JSON data into EthUnstakeExtra
		err = json.Unmarshal(data, &dst.EthUnstakeExtra)
		if err == nil {
			return nil // data stored in dst.EthUnstakeExtra, return on the first match
		} else {
			dst.EthUnstakeExtra = nil
			return fmt.Errorf("failed to unmarshal CreateUnstakeActivityExtra as EthUnstakeExtra: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EthUnstakeExtra'
	if jsonDict["pool_type"] == "EthUnstakeExtra" {
		// try to unmarshal JSON data into EthUnstakeExtra
		err = json.Unmarshal(data, &dst.EthUnstakeExtra)
		if err == nil {
			return nil // data stored in dst.EthUnstakeExtra, return on the first match
		} else {
			dst.EthUnstakeExtra = nil
			return fmt.Errorf("failed to unmarshal CreateUnstakeActivityExtra as EthUnstakeExtra: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateUnstakeActivityExtra) MarshalJSON() ([]byte, error) {
	if src.EthUnstakeExtra != nil {
		return json.Marshal(&src.EthUnstakeExtra)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateUnstakeActivityExtra) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EthUnstakeExtra != nil {
		return obj.EthUnstakeExtra
	}

	// all schemas are nil
	return nil
}

type NullableCreateUnstakeActivityExtra struct {
	value *CreateUnstakeActivityExtra
	isSet bool
}

func (v NullableCreateUnstakeActivityExtra) Get() *CreateUnstakeActivityExtra {
	return v.value
}

func (v *NullableCreateUnstakeActivityExtra) Set(val *CreateUnstakeActivityExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUnstakeActivityExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUnstakeActivityExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUnstakeActivityExtra(val *CreateUnstakeActivityExtra) *NullableCreateUnstakeActivityExtra {
	return &NullableCreateUnstakeActivityExtra{value: val, isSet: true}
}

func (v NullableCreateUnstakeActivityExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUnstakeActivityExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


