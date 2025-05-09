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

// SmartContractWalletInfo - struct for SmartContractWalletInfo
type SmartContractWalletInfo struct {
	SafeWallet *SafeWallet
}

// SafeWalletAsSmartContractWalletInfo is a convenience function that returns SafeWallet wrapped in SmartContractWalletInfo
func SafeWalletAsSmartContractWalletInfo(v *SafeWallet) SmartContractWalletInfo {
	return SmartContractWalletInfo{
		SafeWallet: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SmartContractWalletInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Safe{Wallet}'
	if jsonDict["smart_contract_wallet_type"] == "Safe{Wallet}" {
		// try to unmarshal JSON data into SafeWallet
		err = json.Unmarshal(data, &dst.SafeWallet)
		if err == nil {
			return nil // data stored in dst.SafeWallet, return on the first match
		} else {
			dst.SafeWallet = nil
			return fmt.Errorf("failed to unmarshal SmartContractWalletInfo as SafeWallet: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SafeWallet'
	if jsonDict["smart_contract_wallet_type"] == "SafeWallet" {
		// try to unmarshal JSON data into SafeWallet
		err = json.Unmarshal(data, &dst.SafeWallet)
		if err == nil {
			return nil // data stored in dst.SafeWallet, return on the first match
		} else {
			dst.SafeWallet = nil
			return fmt.Errorf("failed to unmarshal SmartContractWalletInfo as SafeWallet: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SmartContractWalletInfo) MarshalJSON() ([]byte, error) {
	if src.SafeWallet != nil {
		return json.Marshal(&src.SafeWallet)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SmartContractWalletInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SafeWallet != nil {
		return obj.SafeWallet
	}

	// all schemas are nil
	return nil
}

type NullableSmartContractWalletInfo struct {
	value *SmartContractWalletInfo
	isSet bool
}

func (v NullableSmartContractWalletInfo) Get() *SmartContractWalletInfo {
	return v.value
}

func (v *NullableSmartContractWalletInfo) Set(val *SmartContractWalletInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractWalletInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractWalletInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractWalletInfo(val *SmartContractWalletInfo) *NullableSmartContractWalletInfo {
	return &NullableSmartContractWalletInfo{value: val, isSet: true}
}

func (v NullableSmartContractWalletInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractWalletInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


