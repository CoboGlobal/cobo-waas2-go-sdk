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

// TransactionDestination - struct for TransactionDestination
type TransactionDestination struct {
	TransactionDepositToAddressDestination *TransactionDepositToAddressDestination
	TransactionDepositToWalletDestination *TransactionDepositToWalletDestination
	TransactionEvmContractDestination *TransactionEvmContractDestination
	TransactionMessageSignEIP191Destination *TransactionMessageSignEIP191Destination
	TransactionMessageSignEIP712Destination *TransactionMessageSignEIP712Destination
	TransactionRawMessageSignDestination *TransactionRawMessageSignDestination
	TransactionSolContractDestination *TransactionSolContractDestination
	TransactionTransferToAddressDestination *TransactionTransferToAddressDestination
	TransactionTransferToWalletDestination *TransactionTransferToWalletDestination
}

// TransactionDepositToAddressDestinationAsTransactionDestination is a convenience function that returns TransactionDepositToAddressDestination wrapped in TransactionDestination
func TransactionDepositToAddressDestinationAsTransactionDestination(v *TransactionDepositToAddressDestination) TransactionDestination {
	return TransactionDestination{
		TransactionDepositToAddressDestination: v,
	}
}

// TransactionDepositToWalletDestinationAsTransactionDestination is a convenience function that returns TransactionDepositToWalletDestination wrapped in TransactionDestination
func TransactionDepositToWalletDestinationAsTransactionDestination(v *TransactionDepositToWalletDestination) TransactionDestination {
	return TransactionDestination{
		TransactionDepositToWalletDestination: v,
	}
}

// TransactionEvmContractDestinationAsTransactionDestination is a convenience function that returns TransactionEvmContractDestination wrapped in TransactionDestination
func TransactionEvmContractDestinationAsTransactionDestination(v *TransactionEvmContractDestination) TransactionDestination {
	return TransactionDestination{
		TransactionEvmContractDestination: v,
	}
}

// TransactionMessageSignEIP191DestinationAsTransactionDestination is a convenience function that returns TransactionMessageSignEIP191Destination wrapped in TransactionDestination
func TransactionMessageSignEIP191DestinationAsTransactionDestination(v *TransactionMessageSignEIP191Destination) TransactionDestination {
	return TransactionDestination{
		TransactionMessageSignEIP191Destination: v,
	}
}

// TransactionMessageSignEIP712DestinationAsTransactionDestination is a convenience function that returns TransactionMessageSignEIP712Destination wrapped in TransactionDestination
func TransactionMessageSignEIP712DestinationAsTransactionDestination(v *TransactionMessageSignEIP712Destination) TransactionDestination {
	return TransactionDestination{
		TransactionMessageSignEIP712Destination: v,
	}
}

// TransactionRawMessageSignDestinationAsTransactionDestination is a convenience function that returns TransactionRawMessageSignDestination wrapped in TransactionDestination
func TransactionRawMessageSignDestinationAsTransactionDestination(v *TransactionRawMessageSignDestination) TransactionDestination {
	return TransactionDestination{
		TransactionRawMessageSignDestination: v,
	}
}

// TransactionSolContractDestinationAsTransactionDestination is a convenience function that returns TransactionSolContractDestination wrapped in TransactionDestination
func TransactionSolContractDestinationAsTransactionDestination(v *TransactionSolContractDestination) TransactionDestination {
	return TransactionDestination{
		TransactionSolContractDestination: v,
	}
}

// TransactionTransferToAddressDestinationAsTransactionDestination is a convenience function that returns TransactionTransferToAddressDestination wrapped in TransactionDestination
func TransactionTransferToAddressDestinationAsTransactionDestination(v *TransactionTransferToAddressDestination) TransactionDestination {
	return TransactionDestination{
		TransactionTransferToAddressDestination: v,
	}
}

// TransactionTransferToWalletDestinationAsTransactionDestination is a convenience function that returns TransactionTransferToWalletDestination wrapped in TransactionDestination
func TransactionTransferToWalletDestinationAsTransactionDestination(v *TransactionTransferToWalletDestination) TransactionDestination {
	return TransactionDestination{
		TransactionTransferToWalletDestination: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransactionDestination) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Address'
	if jsonDict["destination_type"] == "Address" {
		// try to unmarshal JSON data into TransactionTransferToAddressDestination
		err = json.Unmarshal(data, &dst.TransactionTransferToAddressDestination)
		if err == nil {
			return nil // data stored in dst.TransactionTransferToAddressDestination, return on the first match
		} else {
			dst.TransactionTransferToAddressDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionTransferToAddressDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DepositToAddress'
	if jsonDict["destination_type"] == "DepositToAddress" {
		// try to unmarshal JSON data into TransactionDepositToAddressDestination
		err = json.Unmarshal(data, &dst.TransactionDepositToAddressDestination)
		if err == nil {
			return nil // data stored in dst.TransactionDepositToAddressDestination, return on the first match
		} else {
			dst.TransactionDepositToAddressDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionDepositToAddressDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DepositToWallet'
	if jsonDict["destination_type"] == "DepositToWallet" {
		// try to unmarshal JSON data into TransactionDepositToWalletDestination
		err = json.Unmarshal(data, &dst.TransactionDepositToWalletDestination)
		if err == nil {
			return nil // data stored in dst.TransactionDepositToWalletDestination, return on the first match
		} else {
			dst.TransactionDepositToWalletDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionDepositToWalletDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EVM_Contract'
	if jsonDict["destination_type"] == "EVM_Contract" {
		// try to unmarshal JSON data into TransactionEvmContractDestination
		err = json.Unmarshal(data, &dst.TransactionEvmContractDestination)
		if err == nil {
			return nil // data stored in dst.TransactionEvmContractDestination, return on the first match
		} else {
			dst.TransactionEvmContractDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionEvmContractDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EVM_EIP_191_Signature'
	if jsonDict["destination_type"] == "EVM_EIP_191_Signature" {
		// try to unmarshal JSON data into TransactionMessageSignEIP191Destination
		err = json.Unmarshal(data, &dst.TransactionMessageSignEIP191Destination)
		if err == nil {
			return nil // data stored in dst.TransactionMessageSignEIP191Destination, return on the first match
		} else {
			dst.TransactionMessageSignEIP191Destination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionMessageSignEIP191Destination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EVM_EIP_712_Signature'
	if jsonDict["destination_type"] == "EVM_EIP_712_Signature" {
		// try to unmarshal JSON data into TransactionMessageSignEIP712Destination
		err = json.Unmarshal(data, &dst.TransactionMessageSignEIP712Destination)
		if err == nil {
			return nil // data stored in dst.TransactionMessageSignEIP712Destination, return on the first match
		} else {
			dst.TransactionMessageSignEIP712Destination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionMessageSignEIP712Destination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ExchangeWallet'
	if jsonDict["destination_type"] == "ExchangeWallet" {
		// try to unmarshal JSON data into TransactionTransferToWalletDestination
		err = json.Unmarshal(data, &dst.TransactionTransferToWalletDestination)
		if err == nil {
			return nil // data stored in dst.TransactionTransferToWalletDestination, return on the first match
		} else {
			dst.TransactionTransferToWalletDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionTransferToWalletDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Raw_Message_Signature'
	if jsonDict["destination_type"] == "Raw_Message_Signature" {
		// try to unmarshal JSON data into TransactionRawMessageSignDestination
		err = json.Unmarshal(data, &dst.TransactionRawMessageSignDestination)
		if err == nil {
			return nil // data stored in dst.TransactionRawMessageSignDestination, return on the first match
		} else {
			dst.TransactionRawMessageSignDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionRawMessageSignDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SOL_Contract'
	if jsonDict["destination_type"] == "SOL_Contract" {
		// try to unmarshal JSON data into TransactionSolContractDestination
		err = json.Unmarshal(data, &dst.TransactionSolContractDestination)
		if err == nil {
			return nil // data stored in dst.TransactionSolContractDestination, return on the first match
		} else {
			dst.TransactionSolContractDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionSolContractDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionDepositToAddressDestination'
	if jsonDict["destination_type"] == "TransactionDepositToAddressDestination" {
		// try to unmarshal JSON data into TransactionDepositToAddressDestination
		err = json.Unmarshal(data, &dst.TransactionDepositToAddressDestination)
		if err == nil {
			return nil // data stored in dst.TransactionDepositToAddressDestination, return on the first match
		} else {
			dst.TransactionDepositToAddressDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionDepositToAddressDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionDepositToWalletDestination'
	if jsonDict["destination_type"] == "TransactionDepositToWalletDestination" {
		// try to unmarshal JSON data into TransactionDepositToWalletDestination
		err = json.Unmarshal(data, &dst.TransactionDepositToWalletDestination)
		if err == nil {
			return nil // data stored in dst.TransactionDepositToWalletDestination, return on the first match
		} else {
			dst.TransactionDepositToWalletDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionDepositToWalletDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionEvmContractDestination'
	if jsonDict["destination_type"] == "TransactionEvmContractDestination" {
		// try to unmarshal JSON data into TransactionEvmContractDestination
		err = json.Unmarshal(data, &dst.TransactionEvmContractDestination)
		if err == nil {
			return nil // data stored in dst.TransactionEvmContractDestination, return on the first match
		} else {
			dst.TransactionEvmContractDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionEvmContractDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionMessageSignEIP191Destination'
	if jsonDict["destination_type"] == "TransactionMessageSignEIP191Destination" {
		// try to unmarshal JSON data into TransactionMessageSignEIP191Destination
		err = json.Unmarshal(data, &dst.TransactionMessageSignEIP191Destination)
		if err == nil {
			return nil // data stored in dst.TransactionMessageSignEIP191Destination, return on the first match
		} else {
			dst.TransactionMessageSignEIP191Destination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionMessageSignEIP191Destination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionMessageSignEIP712Destination'
	if jsonDict["destination_type"] == "TransactionMessageSignEIP712Destination" {
		// try to unmarshal JSON data into TransactionMessageSignEIP712Destination
		err = json.Unmarshal(data, &dst.TransactionMessageSignEIP712Destination)
		if err == nil {
			return nil // data stored in dst.TransactionMessageSignEIP712Destination, return on the first match
		} else {
			dst.TransactionMessageSignEIP712Destination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionMessageSignEIP712Destination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionRawMessageSignDestination'
	if jsonDict["destination_type"] == "TransactionRawMessageSignDestination" {
		// try to unmarshal JSON data into TransactionRawMessageSignDestination
		err = json.Unmarshal(data, &dst.TransactionRawMessageSignDestination)
		if err == nil {
			return nil // data stored in dst.TransactionRawMessageSignDestination, return on the first match
		} else {
			dst.TransactionRawMessageSignDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionRawMessageSignDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionSolContractDestination'
	if jsonDict["destination_type"] == "TransactionSolContractDestination" {
		// try to unmarshal JSON data into TransactionSolContractDestination
		err = json.Unmarshal(data, &dst.TransactionSolContractDestination)
		if err == nil {
			return nil // data stored in dst.TransactionSolContractDestination, return on the first match
		} else {
			dst.TransactionSolContractDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionSolContractDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionTransferToAddressDestination'
	if jsonDict["destination_type"] == "TransactionTransferToAddressDestination" {
		// try to unmarshal JSON data into TransactionTransferToAddressDestination
		err = json.Unmarshal(data, &dst.TransactionTransferToAddressDestination)
		if err == nil {
			return nil // data stored in dst.TransactionTransferToAddressDestination, return on the first match
		} else {
			dst.TransactionTransferToAddressDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionTransferToAddressDestination: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TransactionTransferToWalletDestination'
	if jsonDict["destination_type"] == "TransactionTransferToWalletDestination" {
		// try to unmarshal JSON data into TransactionTransferToWalletDestination
		err = json.Unmarshal(data, &dst.TransactionTransferToWalletDestination)
		if err == nil {
			return nil // data stored in dst.TransactionTransferToWalletDestination, return on the first match
		} else {
			dst.TransactionTransferToWalletDestination = nil
			return fmt.Errorf("failed to unmarshal TransactionDestination as TransactionTransferToWalletDestination: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransactionDestination) MarshalJSON() ([]byte, error) {
	if src.TransactionDepositToAddressDestination != nil {
		return json.Marshal(&src.TransactionDepositToAddressDestination)
	}

	if src.TransactionDepositToWalletDestination != nil {
		return json.Marshal(&src.TransactionDepositToWalletDestination)
	}

	if src.TransactionEvmContractDestination != nil {
		return json.Marshal(&src.TransactionEvmContractDestination)
	}

	if src.TransactionMessageSignEIP191Destination != nil {
		return json.Marshal(&src.TransactionMessageSignEIP191Destination)
	}

	if src.TransactionMessageSignEIP712Destination != nil {
		return json.Marshal(&src.TransactionMessageSignEIP712Destination)
	}

	if src.TransactionRawMessageSignDestination != nil {
		return json.Marshal(&src.TransactionRawMessageSignDestination)
	}

	if src.TransactionSolContractDestination != nil {
		return json.Marshal(&src.TransactionSolContractDestination)
	}

	if src.TransactionTransferToAddressDestination != nil {
		return json.Marshal(&src.TransactionTransferToAddressDestination)
	}

	if src.TransactionTransferToWalletDestination != nil {
		return json.Marshal(&src.TransactionTransferToWalletDestination)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransactionDestination) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TransactionDepositToAddressDestination != nil {
		return obj.TransactionDepositToAddressDestination
	}

	if obj.TransactionDepositToWalletDestination != nil {
		return obj.TransactionDepositToWalletDestination
	}

	if obj.TransactionEvmContractDestination != nil {
		return obj.TransactionEvmContractDestination
	}

	if obj.TransactionMessageSignEIP191Destination != nil {
		return obj.TransactionMessageSignEIP191Destination
	}

	if obj.TransactionMessageSignEIP712Destination != nil {
		return obj.TransactionMessageSignEIP712Destination
	}

	if obj.TransactionRawMessageSignDestination != nil {
		return obj.TransactionRawMessageSignDestination
	}

	if obj.TransactionSolContractDestination != nil {
		return obj.TransactionSolContractDestination
	}

	if obj.TransactionTransferToAddressDestination != nil {
		return obj.TransactionTransferToAddressDestination
	}

	if obj.TransactionTransferToWalletDestination != nil {
		return obj.TransactionTransferToWalletDestination
	}

	// all schemas are nil
	return nil
}

type NullableTransactionDestination struct {
	value *TransactionDestination
	isSet bool
}

func (v NullableTransactionDestination) Get() *TransactionDestination {
	return v.value
}

func (v *NullableTransactionDestination) Set(val *TransactionDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDestination(val *TransactionDestination) *NullableTransactionDestination {
	return &NullableTransactionDestination{value: val, isSet: true}
}

func (v NullableTransactionDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


