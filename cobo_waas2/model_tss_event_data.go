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

// TSSEventData - struct for TSSEventData
type TSSEventData struct {
	TSSKeyGenEventData *TSSKeyGenEventData
	TSSKeyReshareEventData *TSSKeyReshareEventData
	TSSKeyShareSignEventData *TSSKeyShareSignEventData
	TSSKeySignEventData *TSSKeySignEventData
}

// TSSKeyGenEventDataAsTSSEventData is a convenience function that returns TSSKeyGenEventData wrapped in TSSEventData
func TSSKeyGenEventDataAsTSSEventData(v *TSSKeyGenEventData) TSSEventData {
	return TSSEventData{
		TSSKeyGenEventData: v,
	}
}

// TSSKeyReshareEventDataAsTSSEventData is a convenience function that returns TSSKeyReshareEventData wrapped in TSSEventData
func TSSKeyReshareEventDataAsTSSEventData(v *TSSKeyReshareEventData) TSSEventData {
	return TSSEventData{
		TSSKeyReshareEventData: v,
	}
}

// TSSKeyShareSignEventDataAsTSSEventData is a convenience function that returns TSSKeyShareSignEventData wrapped in TSSEventData
func TSSKeyShareSignEventDataAsTSSEventData(v *TSSKeyShareSignEventData) TSSEventData {
	return TSSEventData{
		TSSKeyShareSignEventData: v,
	}
}

// TSSKeySignEventDataAsTSSEventData is a convenience function that returns TSSKeySignEventData wrapped in TSSEventData
func TSSKeySignEventDataAsTSSEventData(v *TSSKeySignEventData) TSSEventData {
	return TSSEventData{
		TSSKeySignEventData: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TSSEventData) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'KeyGen'
	if jsonDict["data_type"] == "KeyGen" {
		// try to unmarshal JSON data into TSSKeyGenEventData
		err = json.Unmarshal(data, &dst.TSSKeyGenEventData)
		if err == nil {
			return nil // data stored in dst.TSSKeyGenEventData, return on the first match
		} else {
			dst.TSSKeyGenEventData = nil
			return fmt.Errorf("failed to unmarshal TSSEventData as TSSKeyGenEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'KeyReshare'
	if jsonDict["data_type"] == "KeyReshare" {
		// try to unmarshal JSON data into TSSKeyReshareEventData
		err = json.Unmarshal(data, &dst.TSSKeyReshareEventData)
		if err == nil {
			return nil // data stored in dst.TSSKeyReshareEventData, return on the first match
		} else {
			dst.TSSKeyReshareEventData = nil
			return fmt.Errorf("failed to unmarshal TSSEventData as TSSKeyReshareEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'KeyShareSign'
	if jsonDict["data_type"] == "KeyShareSign" {
		// try to unmarshal JSON data into TSSKeyShareSignEventData
		err = json.Unmarshal(data, &dst.TSSKeyShareSignEventData)
		if err == nil {
			return nil // data stored in dst.TSSKeyShareSignEventData, return on the first match
		} else {
			dst.TSSKeyShareSignEventData = nil
			return fmt.Errorf("failed to unmarshal TSSEventData as TSSKeyShareSignEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'KeySign'
	if jsonDict["data_type"] == "KeySign" {
		// try to unmarshal JSON data into TSSKeySignEventData
		err = json.Unmarshal(data, &dst.TSSKeySignEventData)
		if err == nil {
			return nil // data stored in dst.TSSKeySignEventData, return on the first match
		} else {
			dst.TSSKeySignEventData = nil
			return fmt.Errorf("failed to unmarshal TSSEventData as TSSKeySignEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TSSKeyGenEventData'
	if jsonDict["data_type"] == "TSSKeyGenEventData" {
		// try to unmarshal JSON data into TSSKeyGenEventData
		err = json.Unmarshal(data, &dst.TSSKeyGenEventData)
		if err == nil {
			return nil // data stored in dst.TSSKeyGenEventData, return on the first match
		} else {
			dst.TSSKeyGenEventData = nil
			return fmt.Errorf("failed to unmarshal TSSEventData as TSSKeyGenEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TSSKeyReshareEventData'
	if jsonDict["data_type"] == "TSSKeyReshareEventData" {
		// try to unmarshal JSON data into TSSKeyReshareEventData
		err = json.Unmarshal(data, &dst.TSSKeyReshareEventData)
		if err == nil {
			return nil // data stored in dst.TSSKeyReshareEventData, return on the first match
		} else {
			dst.TSSKeyReshareEventData = nil
			return fmt.Errorf("failed to unmarshal TSSEventData as TSSKeyReshareEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TSSKeyShareSignEventData'
	if jsonDict["data_type"] == "TSSKeyShareSignEventData" {
		// try to unmarshal JSON data into TSSKeyShareSignEventData
		err = json.Unmarshal(data, &dst.TSSKeyShareSignEventData)
		if err == nil {
			return nil // data stored in dst.TSSKeyShareSignEventData, return on the first match
		} else {
			dst.TSSKeyShareSignEventData = nil
			return fmt.Errorf("failed to unmarshal TSSEventData as TSSKeyShareSignEventData: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TSSKeySignEventData'
	if jsonDict["data_type"] == "TSSKeySignEventData" {
		// try to unmarshal JSON data into TSSKeySignEventData
		err = json.Unmarshal(data, &dst.TSSKeySignEventData)
		if err == nil {
			return nil // data stored in dst.TSSKeySignEventData, return on the first match
		} else {
			dst.TSSKeySignEventData = nil
			return fmt.Errorf("failed to unmarshal TSSEventData as TSSKeySignEventData: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TSSEventData) MarshalJSON() ([]byte, error) {
	if src.TSSKeyGenEventData != nil {
		return json.Marshal(&src.TSSKeyGenEventData)
	}

	if src.TSSKeyReshareEventData != nil {
		return json.Marshal(&src.TSSKeyReshareEventData)
	}

	if src.TSSKeyShareSignEventData != nil {
		return json.Marshal(&src.TSSKeyShareSignEventData)
	}

	if src.TSSKeySignEventData != nil {
		return json.Marshal(&src.TSSKeySignEventData)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TSSEventData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TSSKeyGenEventData != nil {
		return obj.TSSKeyGenEventData
	}

	if obj.TSSKeyReshareEventData != nil {
		return obj.TSSKeyReshareEventData
	}

	if obj.TSSKeyShareSignEventData != nil {
		return obj.TSSKeyShareSignEventData
	}

	if obj.TSSKeySignEventData != nil {
		return obj.TSSKeySignEventData
	}

	// all schemas are nil
	return nil
}

type NullableTSSEventData struct {
	value *TSSEventData
	isSet bool
}

func (v NullableTSSEventData) Get() *TSSEventData {
	return v.value
}

func (v *NullableTSSEventData) Set(val *TSSEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSEventData(val *TSSEventData) *NullableTSSEventData {
	return &NullableTSSEventData{value: val, isSet: true}
}

func (v NullableTSSEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


