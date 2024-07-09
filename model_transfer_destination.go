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

// TransferDestination - struct for TransferDestination
type TransferDestination struct {
	AddressTransferDestination *AddressTransferDestination
	ExchangeTransferDestination *ExchangeTransferDestination
}

// AddressTransferDestinationAsTransferDestination is a convenience function that returns AddressTransferDestination wrapped in TransferDestination
func AddressTransferDestinationAsTransferDestination(v *AddressTransferDestination) TransferDestination {
	return TransferDestination{
		AddressTransferDestination: v,
	}
}

// ExchangeTransferDestinationAsTransferDestination is a convenience function that returns ExchangeTransferDestination wrapped in TransferDestination
func ExchangeTransferDestinationAsTransferDestination(v *ExchangeTransferDestination) TransferDestination {
	return TransferDestination{
		ExchangeTransferDestination: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransferDestination) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Address'
	if jsonDict["destination_type"] == "Address" {
		// try to unmarshal JSON data into AddressTransferDestination
		err = json.Unmarshal(data, &dst.AddressTransferDestination)
		if err == nil {
			return nil // data stored in dst.AddressTransferDestination, return on the first match
		} else {
			dst.AddressTransferDestination = nil
			return fmt.Errorf("failed to unmarshal TransferDestination as AddressTransferDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeWallet'
	if jsonDict["destination_type"] == "ExchangeWallet" {
		// try to unmarshal JSON data into ExchangeTransferDestination
		err = json.Unmarshal(data, &dst.ExchangeTransferDestination)
		if err == nil {
			return nil // data stored in dst.ExchangeTransferDestination, return on the first match
		} else {
			dst.ExchangeTransferDestination = nil
			return fmt.Errorf("failed to unmarshal TransferDestination as ExchangeTransferDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AddressTransferDestination'
	if jsonDict["destination_type"] == "AddressTransferDestination" {
		// try to unmarshal JSON data into AddressTransferDestination
		err = json.Unmarshal(data, &dst.AddressTransferDestination)
		if err == nil {
			return nil // data stored in dst.AddressTransferDestination, return on the first match
		} else {
			dst.AddressTransferDestination = nil
			return fmt.Errorf("failed to unmarshal TransferDestination as AddressTransferDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeTransferDestination'
	if jsonDict["destination_type"] == "ExchangeTransferDestination" {
		// try to unmarshal JSON data into ExchangeTransferDestination
		err = json.Unmarshal(data, &dst.ExchangeTransferDestination)
		if err == nil {
			return nil // data stored in dst.ExchangeTransferDestination, return on the first match
		} else {
			dst.ExchangeTransferDestination = nil
			return fmt.Errorf("failed to unmarshal TransferDestination as ExchangeTransferDestination: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransferDestination) MarshalJSON() ([]byte, error) {
	if src.AddressTransferDestination != nil {
		return json.Marshal(&src.AddressTransferDestination)
	}

	if src.ExchangeTransferDestination != nil {
		return json.Marshal(&src.ExchangeTransferDestination)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransferDestination) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AddressTransferDestination != nil {
		return obj.AddressTransferDestination
	}

	if obj.ExchangeTransferDestination != nil {
		return obj.ExchangeTransferDestination
	}

	// all schemas are nil
	return nil
}

type NullableTransferDestination struct {
	value *TransferDestination
	isSet bool
}

func (v NullableTransferDestination) Get() *TransferDestination {
	return v.value
}

func (v *NullableTransferDestination) Set(val *TransferDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDestination(val *TransferDestination) *NullableTransferDestination {
	return &NullableTransferDestination{value: val, isSet: true}
}

func (v NullableTransferDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


