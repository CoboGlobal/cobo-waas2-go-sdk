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

// MessageSignSource - struct for MessageSignSource
type MessageSignSource struct {
	MpcMessageSignSource *MpcMessageSignSource
}

// MpcMessageSignSourceAsMessageSignSource is a convenience function that returns MpcMessageSignSource wrapped in MessageSignSource
func MpcMessageSignSourceAsMessageSignSource(v *MpcMessageSignSource) MessageSignSource {
	return MessageSignSource{
		MpcMessageSignSource: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MessageSignSource) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Org-Controlled'
	if jsonDict["source_type"] == "Org-Controlled" {
		// try to unmarshal JSON data into MpcMessageSignSource
		err = json.Unmarshal(data, &dst.MpcMessageSignSource)
		if err == nil {
			return nil // data stored in dst.MpcMessageSignSource, return on the first match
		} else {
			dst.MpcMessageSignSource = nil
			return fmt.Errorf("failed to unmarshal MessageSignSource as MpcMessageSignSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'User-Controlled'
	if jsonDict["source_type"] == "User-Controlled" {
		// try to unmarshal JSON data into MpcMessageSignSource
		err = json.Unmarshal(data, &dst.MpcMessageSignSource)
		if err == nil {
			return nil // data stored in dst.MpcMessageSignSource, return on the first match
		} else {
			dst.MpcMessageSignSource = nil
			return fmt.Errorf("failed to unmarshal MessageSignSource as MpcMessageSignSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MpcMessageSignSource'
	if jsonDict["source_type"] == "MpcMessageSignSource" {
		// try to unmarshal JSON data into MpcMessageSignSource
		err = json.Unmarshal(data, &dst.MpcMessageSignSource)
		if err == nil {
			return nil // data stored in dst.MpcMessageSignSource, return on the first match
		} else {
			dst.MpcMessageSignSource = nil
			return fmt.Errorf("failed to unmarshal MessageSignSource as MpcMessageSignSource: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MessageSignSource) MarshalJSON() ([]byte, error) {
	if src.MpcMessageSignSource != nil {
		return json.Marshal(&src.MpcMessageSignSource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MessageSignSource) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MpcMessageSignSource != nil {
		return obj.MpcMessageSignSource
	}

	// all schemas are nil
	return nil
}

type NullableMessageSignSource struct {
	value *MessageSignSource
	isSet bool
}

func (v NullableMessageSignSource) Get() *MessageSignSource {
	return v.value
}

func (v *NullableMessageSignSource) Set(val *MessageSignSource) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSignSource) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSignSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSignSource(val *MessageSignSource) *NullableMessageSignSource {
	return &NullableMessageSignSource{value: val, isSet: true}
}

func (v NullableMessageSignSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSignSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


