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

// TransferSource - struct for TransferSource
type TransferSource struct {
	CustodialTransferSource *CustodialTransferSource
	CustodialWeb3TransferSource *CustodialWeb3TransferSource
	ExchangeTransferSource *ExchangeTransferSource
	MpcTransferSource *MpcTransferSource
	SafeTransferSource *SafeTransferSource
}

// CustodialTransferSourceAsTransferSource is a convenience function that returns CustodialTransferSource wrapped in TransferSource
func CustodialTransferSourceAsTransferSource(v *CustodialTransferSource) TransferSource {
	return TransferSource{
		CustodialTransferSource: v,
	}
}

// CustodialWeb3TransferSourceAsTransferSource is a convenience function that returns CustodialWeb3TransferSource wrapped in TransferSource
func CustodialWeb3TransferSourceAsTransferSource(v *CustodialWeb3TransferSource) TransferSource {
	return TransferSource{
		CustodialWeb3TransferSource: v,
	}
}

// ExchangeTransferSourceAsTransferSource is a convenience function that returns ExchangeTransferSource wrapped in TransferSource
func ExchangeTransferSourceAsTransferSource(v *ExchangeTransferSource) TransferSource {
	return TransferSource{
		ExchangeTransferSource: v,
	}
}

// MpcTransferSourceAsTransferSource is a convenience function that returns MpcTransferSource wrapped in TransferSource
func MpcTransferSourceAsTransferSource(v *MpcTransferSource) TransferSource {
	return TransferSource{
		MpcTransferSource: v,
	}
}

// SafeTransferSourceAsTransferSource is a convenience function that returns SafeTransferSource wrapped in TransferSource
func SafeTransferSourceAsTransferSource(v *SafeTransferSource) TransferSource {
	return TransferSource{
		SafeTransferSource: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransferSource) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Asset'
	if jsonDict["source_type"] == "Asset" {
		// try to unmarshal JSON data into CustodialTransferSource
		err = json.Unmarshal(data, &dst.CustodialTransferSource)
		if err == nil {
			return nil // data stored in dst.CustodialTransferSource, return on the first match
		} else {
			dst.CustodialTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as CustodialTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Main'
	if jsonDict["source_type"] == "Main" {
		// try to unmarshal JSON data into ExchangeTransferSource
		err = json.Unmarshal(data, &dst.ExchangeTransferSource)
		if err == nil {
			return nil // data stored in dst.ExchangeTransferSource, return on the first match
		} else {
			dst.ExchangeTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as ExchangeTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Org-Controlled'
	if jsonDict["source_type"] == "Org-Controlled" {
		// try to unmarshal JSON data into MpcTransferSource
		err = json.Unmarshal(data, &dst.MpcTransferSource)
		if err == nil {
			return nil // data stored in dst.MpcTransferSource, return on the first match
		} else {
			dst.MpcTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as MpcTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Safe{Wallet}'
	if jsonDict["source_type"] == "Safe{Wallet}" {
		// try to unmarshal JSON data into SafeTransferSource
		err = json.Unmarshal(data, &dst.SafeTransferSource)
		if err == nil {
			return nil // data stored in dst.SafeTransferSource, return on the first match
		} else {
			dst.SafeTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as SafeTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Sub'
	if jsonDict["source_type"] == "Sub" {
		// try to unmarshal JSON data into ExchangeTransferSource
		err = json.Unmarshal(data, &dst.ExchangeTransferSource)
		if err == nil {
			return nil // data stored in dst.ExchangeTransferSource, return on the first match
		} else {
			dst.ExchangeTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as ExchangeTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'User-Controlled'
	if jsonDict["source_type"] == "User-Controlled" {
		// try to unmarshal JSON data into MpcTransferSource
		err = json.Unmarshal(data, &dst.MpcTransferSource)
		if err == nil {
			return nil // data stored in dst.MpcTransferSource, return on the first match
		} else {
			dst.MpcTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as MpcTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Web3'
	if jsonDict["source_type"] == "Web3" {
		// try to unmarshal JSON data into CustodialWeb3TransferSource
		err = json.Unmarshal(data, &dst.CustodialWeb3TransferSource)
		if err == nil {
			return nil // data stored in dst.CustodialWeb3TransferSource, return on the first match
		} else {
			dst.CustodialWeb3TransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as CustodialWeb3TransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CustodialTransferSource'
	if jsonDict["source_type"] == "CustodialTransferSource" {
		// try to unmarshal JSON data into CustodialTransferSource
		err = json.Unmarshal(data, &dst.CustodialTransferSource)
		if err == nil {
			return nil // data stored in dst.CustodialTransferSource, return on the first match
		} else {
			dst.CustodialTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as CustodialTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CustodialWeb3TransferSource'
	if jsonDict["source_type"] == "CustodialWeb3TransferSource" {
		// try to unmarshal JSON data into CustodialWeb3TransferSource
		err = json.Unmarshal(data, &dst.CustodialWeb3TransferSource)
		if err == nil {
			return nil // data stored in dst.CustodialWeb3TransferSource, return on the first match
		} else {
			dst.CustodialWeb3TransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as CustodialWeb3TransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeTransferSource'
	if jsonDict["source_type"] == "ExchangeTransferSource" {
		// try to unmarshal JSON data into ExchangeTransferSource
		err = json.Unmarshal(data, &dst.ExchangeTransferSource)
		if err == nil {
			return nil // data stored in dst.ExchangeTransferSource, return on the first match
		} else {
			dst.ExchangeTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as ExchangeTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MpcTransferSource'
	if jsonDict["source_type"] == "MpcTransferSource" {
		// try to unmarshal JSON data into MpcTransferSource
		err = json.Unmarshal(data, &dst.MpcTransferSource)
		if err == nil {
			return nil // data stored in dst.MpcTransferSource, return on the first match
		} else {
			dst.MpcTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as MpcTransferSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SafeTransferSource'
	if jsonDict["source_type"] == "SafeTransferSource" {
		// try to unmarshal JSON data into SafeTransferSource
		err = json.Unmarshal(data, &dst.SafeTransferSource)
		if err == nil {
			return nil // data stored in dst.SafeTransferSource, return on the first match
		} else {
			dst.SafeTransferSource = nil
			return fmt.Errorf("failed to unmarshal TransferSource as SafeTransferSource: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransferSource) MarshalJSON() ([]byte, error) {
	if src.CustodialTransferSource != nil {
		return json.Marshal(&src.CustodialTransferSource)
	}

	if src.CustodialWeb3TransferSource != nil {
		return json.Marshal(&src.CustodialWeb3TransferSource)
	}

	if src.ExchangeTransferSource != nil {
		return json.Marshal(&src.ExchangeTransferSource)
	}

	if src.MpcTransferSource != nil {
		return json.Marshal(&src.MpcTransferSource)
	}

	if src.SafeTransferSource != nil {
		return json.Marshal(&src.SafeTransferSource)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransferSource) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CustodialTransferSource != nil {
		return obj.CustodialTransferSource
	}

	if obj.CustodialWeb3TransferSource != nil {
		return obj.CustodialWeb3TransferSource
	}

	if obj.ExchangeTransferSource != nil {
		return obj.ExchangeTransferSource
	}

	if obj.MpcTransferSource != nil {
		return obj.MpcTransferSource
	}

	if obj.SafeTransferSource != nil {
		return obj.SafeTransferSource
	}

	// all schemas are nil
	return nil
}

type NullableTransferSource struct {
	value *TransferSource
	isSet bool
}

func (v NullableTransferSource) Get() *TransferSource {
	return v.value
}

func (v *NullableTransferSource) Set(val *TransferSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferSource(val *TransferSource) *NullableTransferSource {
	return &NullableTransferSource{value: val, isSet: true}
}

func (v NullableTransferSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


