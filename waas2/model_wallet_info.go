/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"fmt"
)

// WalletInfo - struct for WalletInfo
type WalletInfo struct {
	CustodialWalletInfo *CustodialWalletInfo
	ExchangeWalletInfo *ExchangeWalletInfo
	MPCWalletInfo *MPCWalletInfo
	SmartContractWalletInfo *SmartContractWalletInfo
}

// CustodialWalletInfoAsWalletInfo is a convenience function that returns CustodialWalletInfo wrapped in WalletInfo
func CustodialWalletInfoAsWalletInfo(v *CustodialWalletInfo) WalletInfo {
	return WalletInfo{
		CustodialWalletInfo: v,
	}
}

// ExchangeWalletInfoAsWalletInfo is a convenience function that returns ExchangeWalletInfo wrapped in WalletInfo
func ExchangeWalletInfoAsWalletInfo(v *ExchangeWalletInfo) WalletInfo {
	return WalletInfo{
		ExchangeWalletInfo: v,
	}
}

// MPCWalletInfoAsWalletInfo is a convenience function that returns MPCWalletInfo wrapped in WalletInfo
func MPCWalletInfoAsWalletInfo(v *MPCWalletInfo) WalletInfo {
	return WalletInfo{
		MPCWalletInfo: v,
	}
}

// SmartContractWalletInfoAsWalletInfo is a convenience function that returns SmartContractWalletInfo wrapped in WalletInfo
func SmartContractWalletInfoAsWalletInfo(v *SmartContractWalletInfo) WalletInfo {
	return WalletInfo{
		SmartContractWalletInfo: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WalletInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Custodial'
	if jsonDict["wallet_type"] == "Custodial" {
		// try to unmarshal JSON data into CustodialWalletInfo
		err = json.Unmarshal(data, &dst.CustodialWalletInfo)
		if err == nil {
			return nil // data stored in dst.CustodialWalletInfo, return on the first match
		} else {
			dst.CustodialWalletInfo = nil
			return fmt.Errorf("failed to unmarshal WalletInfo as CustodialWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Exchange'
	if jsonDict["wallet_type"] == "Exchange" {
		// try to unmarshal JSON data into ExchangeWalletInfo
		err = json.Unmarshal(data, &dst.ExchangeWalletInfo)
		if err == nil {
			return nil // data stored in dst.ExchangeWalletInfo, return on the first match
		} else {
			dst.ExchangeWalletInfo = nil
			return fmt.Errorf("failed to unmarshal WalletInfo as ExchangeWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MPC'
	if jsonDict["wallet_type"] == "MPC" {
		// try to unmarshal JSON data into MPCWalletInfo
		err = json.Unmarshal(data, &dst.MPCWalletInfo)
		if err == nil {
			return nil // data stored in dst.MPCWalletInfo, return on the first match
		} else {
			dst.MPCWalletInfo = nil
			return fmt.Errorf("failed to unmarshal WalletInfo as MPCWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SmartContract'
	if jsonDict["wallet_type"] == "SmartContract" {
		// try to unmarshal JSON data into SmartContractWalletInfo
		err = json.Unmarshal(data, &dst.SmartContractWalletInfo)
		if err == nil {
			return nil // data stored in dst.SmartContractWalletInfo, return on the first match
		} else {
			dst.SmartContractWalletInfo = nil
			return fmt.Errorf("failed to unmarshal WalletInfo as SmartContractWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CustodialWalletInfo'
	if jsonDict["wallet_type"] == "CustodialWalletInfo" {
		// try to unmarshal JSON data into CustodialWalletInfo
		err = json.Unmarshal(data, &dst.CustodialWalletInfo)
		if err == nil {
			return nil // data stored in dst.CustodialWalletInfo, return on the first match
		} else {
			dst.CustodialWalletInfo = nil
			return fmt.Errorf("failed to unmarshal WalletInfo as CustodialWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeWalletInfo'
	if jsonDict["wallet_type"] == "ExchangeWalletInfo" {
		// try to unmarshal JSON data into ExchangeWalletInfo
		err = json.Unmarshal(data, &dst.ExchangeWalletInfo)
		if err == nil {
			return nil // data stored in dst.ExchangeWalletInfo, return on the first match
		} else {
			dst.ExchangeWalletInfo = nil
			return fmt.Errorf("failed to unmarshal WalletInfo as ExchangeWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MPCWalletInfo'
	if jsonDict["wallet_type"] == "MPCWalletInfo" {
		// try to unmarshal JSON data into MPCWalletInfo
		err = json.Unmarshal(data, &dst.MPCWalletInfo)
		if err == nil {
			return nil // data stored in dst.MPCWalletInfo, return on the first match
		} else {
			dst.MPCWalletInfo = nil
			return fmt.Errorf("failed to unmarshal WalletInfo as MPCWalletInfo: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SmartContractWalletInfo'
	if jsonDict["wallet_type"] == "SmartContractWalletInfo" {
		// try to unmarshal JSON data into SmartContractWalletInfo
		err = json.Unmarshal(data, &dst.SmartContractWalletInfo)
		if err == nil {
			return nil // data stored in dst.SmartContractWalletInfo, return on the first match
		} else {
			dst.SmartContractWalletInfo = nil
			return fmt.Errorf("failed to unmarshal WalletInfo as SmartContractWalletInfo: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WalletInfo) MarshalJSON() ([]byte, error) {
	if src.CustodialWalletInfo != nil {
		return json.Marshal(&src.CustodialWalletInfo)
	}

	if src.ExchangeWalletInfo != nil {
		return json.Marshal(&src.ExchangeWalletInfo)
	}

	if src.MPCWalletInfo != nil {
		return json.Marshal(&src.MPCWalletInfo)
	}

	if src.SmartContractWalletInfo != nil {
		return json.Marshal(&src.SmartContractWalletInfo)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WalletInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CustodialWalletInfo != nil {
		return obj.CustodialWalletInfo
	}

	if obj.ExchangeWalletInfo != nil {
		return obj.ExchangeWalletInfo
	}

	if obj.MPCWalletInfo != nil {
		return obj.MPCWalletInfo
	}

	if obj.SmartContractWalletInfo != nil {
		return obj.SmartContractWalletInfo
	}

	// all schemas are nil
	return nil
}

type NullableWalletInfo struct {
	value *WalletInfo
	isSet bool
}

func (v NullableWalletInfo) Get() *WalletInfo {
	return v.value
}

func (v *NullableWalletInfo) Set(val *WalletInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletInfo(val *WalletInfo) *NullableWalletInfo {
	return &NullableWalletInfo{value: val, isSet: true}
}

func (v NullableWalletInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


