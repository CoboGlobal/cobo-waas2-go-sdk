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

// TransactionSource - struct for TransactionSource
type TransactionSource struct {
	BaseTransactionAddress *BaseTransactionAddress
	BaseWalletTransactionAddress *BaseWalletTransactionAddress
	ExchangeWalletTransactionAddress *ExchangeWalletTransactionAddress
	MPCWalletTransactionAddress *MPCWalletTransactionAddress
	SafeTransactionAddress *SafeTransactionAddress
}

// BaseTransactionAddressAsTransactionSource is a convenience function that returns BaseTransactionAddress wrapped in TransactionSource
func BaseTransactionAddressAsTransactionSource(v *BaseTransactionAddress) TransactionSource {
	return TransactionSource{
		BaseTransactionAddress: v,
	}
}

// BaseWalletTransactionAddressAsTransactionSource is a convenience function that returns BaseWalletTransactionAddress wrapped in TransactionSource
func BaseWalletTransactionAddressAsTransactionSource(v *BaseWalletTransactionAddress) TransactionSource {
	return TransactionSource{
		BaseWalletTransactionAddress: v,
	}
}

// ExchangeWalletTransactionAddressAsTransactionSource is a convenience function that returns ExchangeWalletTransactionAddress wrapped in TransactionSource
func ExchangeWalletTransactionAddressAsTransactionSource(v *ExchangeWalletTransactionAddress) TransactionSource {
	return TransactionSource{
		ExchangeWalletTransactionAddress: v,
	}
}

// MPCWalletTransactionAddressAsTransactionSource is a convenience function that returns MPCWalletTransactionAddress wrapped in TransactionSource
func MPCWalletTransactionAddressAsTransactionSource(v *MPCWalletTransactionAddress) TransactionSource {
	return TransactionSource{
		MPCWalletTransactionAddress: v,
	}
}

// SafeTransactionAddressAsTransactionSource is a convenience function that returns SafeTransactionAddress wrapped in TransactionSource
func SafeTransactionAddressAsTransactionSource(v *SafeTransactionAddress) TransactionSource {
	return TransactionSource{
		SafeTransactionAddress: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionSource) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'CustodialAssetWallet'
	if jsonDict["type"] == "CustodialAssetWallet" {
		// try to unmarshal JSON data into BaseWalletTransactionAddress
		err = json.Unmarshal(data, &dst.BaseWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.BaseWalletTransactionAddress, return on the first match
		} else {
			dst.BaseWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as BaseWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CustodialWeb3Wallet'
	if jsonDict["type"] == "CustodialWeb3Wallet" {
		// try to unmarshal JSON data into BaseWalletTransactionAddress
		err = json.Unmarshal(data, &dst.BaseWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.BaseWalletTransactionAddress, return on the first match
		} else {
			dst.BaseWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as BaseWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeAccount'
	if jsonDict["type"] == "ExchangeAccount" {
		// try to unmarshal JSON data into ExchangeWalletTransactionAddress
		err = json.Unmarshal(data, &dst.ExchangeWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.ExchangeWalletTransactionAddress, return on the first match
		} else {
			dst.ExchangeWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as ExchangeWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExternalAddress'
	if jsonDict["type"] == "ExternalAddress" {
		// try to unmarshal JSON data into BaseTransactionAddress
		err = json.Unmarshal(data, &dst.BaseTransactionAddress)
		if err == nil {
			return nil // data stored in dst.BaseTransactionAddress, return on the first match
		} else {
			dst.BaseTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as BaseTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GasStation'
	if jsonDict["type"] == "GasStation" {
		// try to unmarshal JSON data into BaseWalletTransactionAddress
		err = json.Unmarshal(data, &dst.BaseWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.BaseWalletTransactionAddress, return on the first match
		} else {
			dst.BaseWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as BaseWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MPCClientControlledWallet'
	if jsonDict["type"] == "MPCClientControlledWallet" {
		// try to unmarshal JSON data into MPCWalletTransactionAddress
		err = json.Unmarshal(data, &dst.MPCWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.MPCWalletTransactionAddress, return on the first match
		} else {
			dst.MPCWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as MPCWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MPCUserControlledWallet'
	if jsonDict["type"] == "MPCUserControlledWallet" {
		// try to unmarshal JSON data into MPCWalletTransactionAddress
		err = json.Unmarshal(data, &dst.MPCWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.MPCWalletTransactionAddress, return on the first match
		} else {
			dst.MPCWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as MPCWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SafeContractWallet'
	if jsonDict["type"] == "SafeContractWallet" {
		// try to unmarshal JSON data into SafeWalletTransactionAddress
		err = json.Unmarshal(data, &dst.SafeWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.SafeWalletTransactionAddress, return on the first match
		} else {
			dst.SafeWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as SafeWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BaseTransactionAddress'
	if jsonDict["type"] == "BaseTransactionAddress" {
		// try to unmarshal JSON data into BaseTransactionAddress
		err = json.Unmarshal(data, &dst.BaseTransactionAddress)
		if err == nil {
			return nil // data stored in dst.BaseTransactionAddress, return on the first match
		} else {
			dst.BaseTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as BaseTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BaseWalletTransactionAddress'
	if jsonDict["type"] == "BaseWalletTransactionAddress" {
		// try to unmarshal JSON data into BaseWalletTransactionAddress
		err = json.Unmarshal(data, &dst.BaseWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.BaseWalletTransactionAddress, return on the first match
		} else {
			dst.BaseWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as BaseWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeWalletTransactionAddress'
	if jsonDict["type"] == "ExchangeWalletTransactionAddress" {
		// try to unmarshal JSON data into ExchangeWalletTransactionAddress
		err = json.Unmarshal(data, &dst.ExchangeWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.ExchangeWalletTransactionAddress, return on the first match
		} else {
			dst.ExchangeWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as ExchangeWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MPCWalletTransactionAddress'
	if jsonDict["type"] == "MPCWalletTransactionAddress" {
		// try to unmarshal JSON data into MPCWalletTransactionAddress
		err = json.Unmarshal(data, &dst.MPCWalletTransactionAddress)
		if err == nil {
			return nil // data stored in dst.MPCWalletTransactionAddress, return on the first match
		} else {
			dst.MPCWalletTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as MPCWalletTransactionAddress: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SafeTransactionAddress'
	if jsonDict["type"] == "SafeTransactionAddress" {
		// try to unmarshal JSON data into SafeTransactionAddress
		err = json.Unmarshal(data, &dst.SafeTransactionAddress)
		if err == nil {
			return nil // data stored in dst.SafeTransactionAddress, return on the first match
		} else {
			dst.SafeTransactionAddress = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as SafeTransactionAddress: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionSource) MarshalJSON() ([]byte, error) {
	if src.BaseTransactionAddress != nil {
		return json.Marshal(&src.BaseTransactionAddress)
	}

	if src.BaseWalletTransactionAddress != nil {
		return json.Marshal(&src.BaseWalletTransactionAddress)
	}

	if src.ExchangeWalletTransactionAddress != nil {
		return json.Marshal(&src.ExchangeWalletTransactionAddress)
	}

	if src.MPCWalletTransactionAddress != nil {
		return json.Marshal(&src.MPCWalletTransactionAddress)
	}

	if src.SafeTransactionAddress != nil {
		return json.Marshal(&src.SafeTransactionAddress)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionSource) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BaseTransactionAddress != nil {
		return obj.BaseTransactionAddress
	}

	if obj.BaseWalletTransactionAddress != nil {
		return obj.BaseWalletTransactionAddress
	}

	if obj.ExchangeWalletTransactionAddress != nil {
		return obj.ExchangeWalletTransactionAddress
	}

	if obj.MPCWalletTransactionAddress != nil {
		return obj.MPCWalletTransactionAddress
	}

	if obj.SafeTransactionAddress != nil {
		return obj.SafeTransactionAddress
	}

	// all schemas are nil
	return nil
}

type NullableTransactionSource struct {
	value *TransactionSource
	isSet bool
}

func (v NullableTransactionSource) Get() *TransactionSource {
	return v.value
}

func (v *NullableTransactionSource) Set(val *TransactionSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionSource(val *TransactionSource) *NullableTransactionSource {
	return &NullableTransactionSource{value: val, isSet: true}
}

func (v NullableTransactionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

