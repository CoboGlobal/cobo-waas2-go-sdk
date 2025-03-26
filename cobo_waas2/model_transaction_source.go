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

// TransactionSource - struct for TransactionSource
type TransactionSource struct {
	TransactionCustodialAssetWalletSource *TransactionCustodialAssetWalletSource
	TransactionCustodialWeb3WalletSource *TransactionCustodialWeb3WalletSource
	TransactionDepositFromAddressSource *TransactionDepositFromAddressSource
	TransactionDepositFromLoopSource *TransactionDepositFromLoopSource
	TransactionDepositFromWalletSource *TransactionDepositFromWalletSource
	TransactionExchangeWalletSource *TransactionExchangeWalletSource
	TransactionMPCWalletSource *TransactionMPCWalletSource
	TransactionSmartContractSafeWalletSource *TransactionSmartContractSafeWalletSource
}

// TransactionCustodialAssetWalletSourceAsTransactionSource is a convenience function that returns TransactionCustodialAssetWalletSource wrapped in TransactionSource
func TransactionCustodialAssetWalletSourceAsTransactionSource(v *TransactionCustodialAssetWalletSource) TransactionSource {
	return TransactionSource{
		TransactionCustodialAssetWalletSource: v,
	}
}

// TransactionCustodialWeb3WalletSourceAsTransactionSource is a convenience function that returns TransactionCustodialWeb3WalletSource wrapped in TransactionSource
func TransactionCustodialWeb3WalletSourceAsTransactionSource(v *TransactionCustodialWeb3WalletSource) TransactionSource {
	return TransactionSource{
		TransactionCustodialWeb3WalletSource: v,
	}
}

// TransactionDepositFromAddressSourceAsTransactionSource is a convenience function that returns TransactionDepositFromAddressSource wrapped in TransactionSource
func TransactionDepositFromAddressSourceAsTransactionSource(v *TransactionDepositFromAddressSource) TransactionSource {
	return TransactionSource{
		TransactionDepositFromAddressSource: v,
	}
}

// TransactionDepositFromLoopSourceAsTransactionSource is a convenience function that returns TransactionDepositFromLoopSource wrapped in TransactionSource
func TransactionDepositFromLoopSourceAsTransactionSource(v *TransactionDepositFromLoopSource) TransactionSource {
	return TransactionSource{
		TransactionDepositFromLoopSource: v,
	}
}

// TransactionDepositFromWalletSourceAsTransactionSource is a convenience function that returns TransactionDepositFromWalletSource wrapped in TransactionSource
func TransactionDepositFromWalletSourceAsTransactionSource(v *TransactionDepositFromWalletSource) TransactionSource {
	return TransactionSource{
		TransactionDepositFromWalletSource: v,
	}
}

// TransactionExchangeWalletSourceAsTransactionSource is a convenience function that returns TransactionExchangeWalletSource wrapped in TransactionSource
func TransactionExchangeWalletSourceAsTransactionSource(v *TransactionExchangeWalletSource) TransactionSource {
	return TransactionSource{
		TransactionExchangeWalletSource: v,
	}
}

// TransactionMPCWalletSourceAsTransactionSource is a convenience function that returns TransactionMPCWalletSource wrapped in TransactionSource
func TransactionMPCWalletSourceAsTransactionSource(v *TransactionMPCWalletSource) TransactionSource {
	return TransactionSource{
		TransactionMPCWalletSource: v,
	}
}

// TransactionSmartContractSafeWalletSourceAsTransactionSource is a convenience function that returns TransactionSmartContractSafeWalletSource wrapped in TransactionSource
func TransactionSmartContractSafeWalletSourceAsTransactionSource(v *TransactionSmartContractSafeWalletSource) TransactionSource {
	return TransactionSource{
		TransactionSmartContractSafeWalletSource: v,
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

	// check if the discriminator value is 'Asset'
	if jsonDict["source_type"] == "Asset" {
		// try to unmarshal JSON data into TransactionCustodialAssetWalletSource
		err = json.Unmarshal(data, &dst.TransactionCustodialAssetWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionCustodialAssetWalletSource, return on the first match
		} else {
			dst.TransactionCustodialAssetWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionCustodialAssetWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DepositFromAddress'
	if jsonDict["source_type"] == "DepositFromAddress" {
		// try to unmarshal JSON data into TransactionDepositFromAddressSource
		err = json.Unmarshal(data, &dst.TransactionDepositFromAddressSource)
		if err == nil {
			return nil // data stored in dst.TransactionDepositFromAddressSource, return on the first match
		} else {
			dst.TransactionDepositFromAddressSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionDepositFromAddressSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DepositFromLoop'
	if jsonDict["source_type"] == "DepositFromLoop" {
		// try to unmarshal JSON data into TransactionDepositFromLoopSource
		err = json.Unmarshal(data, &dst.TransactionDepositFromLoopSource)
		if err == nil {
			return nil // data stored in dst.TransactionDepositFromLoopSource, return on the first match
		} else {
			dst.TransactionDepositFromLoopSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionDepositFromLoopSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DepositFromWallet'
	if jsonDict["source_type"] == "DepositFromWallet" {
		// try to unmarshal JSON data into TransactionDepositFromWalletSource
		err = json.Unmarshal(data, &dst.TransactionDepositFromWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionDepositFromWalletSource, return on the first match
		} else {
			dst.TransactionDepositFromWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionDepositFromWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Main'
	if jsonDict["source_type"] == "Main" {
		// try to unmarshal JSON data into TransactionExchangeWalletSource
		err = json.Unmarshal(data, &dst.TransactionExchangeWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionExchangeWalletSource, return on the first match
		} else {
			dst.TransactionExchangeWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionExchangeWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Org-Controlled'
	if jsonDict["source_type"] == "Org-Controlled" {
		// try to unmarshal JSON data into TransactionMPCWalletSource
		err = json.Unmarshal(data, &dst.TransactionMPCWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionMPCWalletSource, return on the first match
		} else {
			dst.TransactionMPCWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionMPCWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Safe{Wallet}'
	if jsonDict["source_type"] == "Safe{Wallet}" {
		// try to unmarshal JSON data into TransactionSmartContractSafeWalletSource
		err = json.Unmarshal(data, &dst.TransactionSmartContractSafeWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionSmartContractSafeWalletSource, return on the first match
		} else {
			dst.TransactionSmartContractSafeWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionSmartContractSafeWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Sub'
	if jsonDict["source_type"] == "Sub" {
		// try to unmarshal JSON data into TransactionExchangeWalletSource
		err = json.Unmarshal(data, &dst.TransactionExchangeWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionExchangeWalletSource, return on the first match
		} else {
			dst.TransactionExchangeWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionExchangeWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'User-Controlled'
	if jsonDict["source_type"] == "User-Controlled" {
		// try to unmarshal JSON data into TransactionMPCWalletSource
		err = json.Unmarshal(data, &dst.TransactionMPCWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionMPCWalletSource, return on the first match
		} else {
			dst.TransactionMPCWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionMPCWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Web3'
	if jsonDict["source_type"] == "Web3" {
		// try to unmarshal JSON data into TransactionCustodialWeb3WalletSource
		err = json.Unmarshal(data, &dst.TransactionCustodialWeb3WalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionCustodialWeb3WalletSource, return on the first match
		} else {
			dst.TransactionCustodialWeb3WalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionCustodialWeb3WalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionCustodialAssetWalletSource'
	if jsonDict["source_type"] == "TransactionCustodialAssetWalletSource" {
		// try to unmarshal JSON data into TransactionCustodialAssetWalletSource
		err = json.Unmarshal(data, &dst.TransactionCustodialAssetWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionCustodialAssetWalletSource, return on the first match
		} else {
			dst.TransactionCustodialAssetWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionCustodialAssetWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionCustodialWeb3WalletSource'
	if jsonDict["source_type"] == "TransactionCustodialWeb3WalletSource" {
		// try to unmarshal JSON data into TransactionCustodialWeb3WalletSource
		err = json.Unmarshal(data, &dst.TransactionCustodialWeb3WalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionCustodialWeb3WalletSource, return on the first match
		} else {
			dst.TransactionCustodialWeb3WalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionCustodialWeb3WalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionDepositFromAddressSource'
	if jsonDict["source_type"] == "TransactionDepositFromAddressSource" {
		// try to unmarshal JSON data into TransactionDepositFromAddressSource
		err = json.Unmarshal(data, &dst.TransactionDepositFromAddressSource)
		if err == nil {
			return nil // data stored in dst.TransactionDepositFromAddressSource, return on the first match
		} else {
			dst.TransactionDepositFromAddressSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionDepositFromAddressSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionDepositFromLoopSource'
	if jsonDict["source_type"] == "TransactionDepositFromLoopSource" {
		// try to unmarshal JSON data into TransactionDepositFromLoopSource
		err = json.Unmarshal(data, &dst.TransactionDepositFromLoopSource)
		if err == nil {
			return nil // data stored in dst.TransactionDepositFromLoopSource, return on the first match
		} else {
			dst.TransactionDepositFromLoopSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionDepositFromLoopSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionDepositFromWalletSource'
	if jsonDict["source_type"] == "TransactionDepositFromWalletSource" {
		// try to unmarshal JSON data into TransactionDepositFromWalletSource
		err = json.Unmarshal(data, &dst.TransactionDepositFromWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionDepositFromWalletSource, return on the first match
		} else {
			dst.TransactionDepositFromWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionDepositFromWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionExchangeWalletSource'
	if jsonDict["source_type"] == "TransactionExchangeWalletSource" {
		// try to unmarshal JSON data into TransactionExchangeWalletSource
		err = json.Unmarshal(data, &dst.TransactionExchangeWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionExchangeWalletSource, return on the first match
		} else {
			dst.TransactionExchangeWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionExchangeWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionMPCWalletSource'
	if jsonDict["source_type"] == "TransactionMPCWalletSource" {
		// try to unmarshal JSON data into TransactionMPCWalletSource
		err = json.Unmarshal(data, &dst.TransactionMPCWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionMPCWalletSource, return on the first match
		} else {
			dst.TransactionMPCWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionMPCWalletSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionSmartContractSafeWalletSource'
	if jsonDict["source_type"] == "TransactionSmartContractSafeWalletSource" {
		// try to unmarshal JSON data into TransactionSmartContractSafeWalletSource
		err = json.Unmarshal(data, &dst.TransactionSmartContractSafeWalletSource)
		if err == nil {
			return nil // data stored in dst.TransactionSmartContractSafeWalletSource, return on the first match
		} else {
			dst.TransactionSmartContractSafeWalletSource = nil
			return fmt.Errorf("failed to unmarshal TransactionSource as TransactionSmartContractSafeWalletSource: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionSource) MarshalJSON() ([]byte, error) {
	if src.TransactionCustodialAssetWalletSource != nil {
		return json.Marshal(&src.TransactionCustodialAssetWalletSource)
	}

	if src.TransactionCustodialWeb3WalletSource != nil {
		return json.Marshal(&src.TransactionCustodialWeb3WalletSource)
	}

	if src.TransactionDepositFromAddressSource != nil {
		return json.Marshal(&src.TransactionDepositFromAddressSource)
	}

	if src.TransactionDepositFromLoopSource != nil {
		return json.Marshal(&src.TransactionDepositFromLoopSource)
	}

	if src.TransactionDepositFromWalletSource != nil {
		return json.Marshal(&src.TransactionDepositFromWalletSource)
	}

	if src.TransactionExchangeWalletSource != nil {
		return json.Marshal(&src.TransactionExchangeWalletSource)
	}

	if src.TransactionMPCWalletSource != nil {
		return json.Marshal(&src.TransactionMPCWalletSource)
	}

	if src.TransactionSmartContractSafeWalletSource != nil {
		return json.Marshal(&src.TransactionSmartContractSafeWalletSource)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionSource) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TransactionCustodialAssetWalletSource != nil {
		return obj.TransactionCustodialAssetWalletSource
	}

	if obj.TransactionCustodialWeb3WalletSource != nil {
		return obj.TransactionCustodialWeb3WalletSource
	}

	if obj.TransactionDepositFromAddressSource != nil {
		return obj.TransactionDepositFromAddressSource
	}

	if obj.TransactionDepositFromLoopSource != nil {
		return obj.TransactionDepositFromLoopSource
	}

	if obj.TransactionDepositFromWalletSource != nil {
		return obj.TransactionDepositFromWalletSource
	}

	if obj.TransactionExchangeWalletSource != nil {
		return obj.TransactionExchangeWalletSource
	}

	if obj.TransactionMPCWalletSource != nil {
		return obj.TransactionMPCWalletSource
	}

	if obj.TransactionSmartContractSafeWalletSource != nil {
		return obj.TransactionSmartContractSafeWalletSource
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


