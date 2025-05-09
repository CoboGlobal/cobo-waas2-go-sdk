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

// TravelRuleDepositRequestTravelRuleInfo - struct for TravelRuleDepositRequestTravelRuleInfo
type TravelRuleDepositRequestTravelRuleInfo struct {
	SelfCustodyWallet *SelfCustodyWallet
	TravelRuleDepositExchangesOrVASP *TravelRuleDepositExchangesOrVASP
}

// SelfCustodyWalletAsTravelRuleDepositRequestTravelRuleInfo is a convenience function that returns SelfCustodyWallet wrapped in TravelRuleDepositRequestTravelRuleInfo
func SelfCustodyWalletAsTravelRuleDepositRequestTravelRuleInfo(v *SelfCustodyWallet) TravelRuleDepositRequestTravelRuleInfo {
	return TravelRuleDepositRequestTravelRuleInfo{
		SelfCustodyWallet: v,
	}
}

// TravelRuleDepositExchangesOrVASPAsTravelRuleDepositRequestTravelRuleInfo is a convenience function that returns TravelRuleDepositExchangesOrVASP wrapped in TravelRuleDepositRequestTravelRuleInfo
func TravelRuleDepositExchangesOrVASPAsTravelRuleDepositRequestTravelRuleInfo(v *TravelRuleDepositExchangesOrVASP) TravelRuleDepositRequestTravelRuleInfo {
	return TravelRuleDepositRequestTravelRuleInfo{
		TravelRuleDepositExchangesOrVASP: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TravelRuleDepositRequestTravelRuleInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EXCHANGES_OR_VASP'
	if jsonDict["destination_wallet_type"] == "EXCHANGES_OR_VASP" {
		// try to unmarshal JSON data into TravelRuleDepositExchangesOrVASP
		err = json.Unmarshal(data, &dst.TravelRuleDepositExchangesOrVASP)
		if err == nil {
			return nil // data stored in dst.TravelRuleDepositExchangesOrVASP, return on the first match
		} else {
			dst.TravelRuleDepositExchangesOrVASP = nil
			return fmt.Errorf("failed to unmarshal TravelRuleDepositRequestTravelRuleInfo as TravelRuleDepositExchangesOrVASP: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SELF_CUSTODY_WALLET'
	if jsonDict["destination_wallet_type"] == "SELF_CUSTODY_WALLET" {
		// try to unmarshal JSON data into SelfCustodyWallet
		err = json.Unmarshal(data, &dst.SelfCustodyWallet)
		if err == nil {
			return nil // data stored in dst.SelfCustodyWallet, return on the first match
		} else {
			dst.SelfCustodyWallet = nil
			return fmt.Errorf("failed to unmarshal TravelRuleDepositRequestTravelRuleInfo as SelfCustodyWallet: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SelfCustodyWallet'
	if jsonDict["destination_wallet_type"] == "SelfCustodyWallet" {
		// try to unmarshal JSON data into SelfCustodyWallet
		err = json.Unmarshal(data, &dst.SelfCustodyWallet)
		if err == nil {
			return nil // data stored in dst.SelfCustodyWallet, return on the first match
		} else {
			dst.SelfCustodyWallet = nil
			return fmt.Errorf("failed to unmarshal TravelRuleDepositRequestTravelRuleInfo as SelfCustodyWallet: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TravelRuleDepositExchangesOrVASP'
	if jsonDict["destination_wallet_type"] == "TravelRuleDepositExchangesOrVASP" {
		// try to unmarshal JSON data into TravelRuleDepositExchangesOrVASP
		err = json.Unmarshal(data, &dst.TravelRuleDepositExchangesOrVASP)
		if err == nil {
			return nil // data stored in dst.TravelRuleDepositExchangesOrVASP, return on the first match
		} else {
			dst.TravelRuleDepositExchangesOrVASP = nil
			return fmt.Errorf("failed to unmarshal TravelRuleDepositRequestTravelRuleInfo as TravelRuleDepositExchangesOrVASP: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TravelRuleDepositRequestTravelRuleInfo) MarshalJSON() ([]byte, error) {
	if src.SelfCustodyWallet != nil {
		return json.Marshal(&src.SelfCustodyWallet)
	}

	if src.TravelRuleDepositExchangesOrVASP != nil {
		return json.Marshal(&src.TravelRuleDepositExchangesOrVASP)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TravelRuleDepositRequestTravelRuleInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SelfCustodyWallet != nil {
		return obj.SelfCustodyWallet
	}

	if obj.TravelRuleDepositExchangesOrVASP != nil {
		return obj.TravelRuleDepositExchangesOrVASP
	}

	// all schemas are nil
	return nil
}

type NullableTravelRuleDepositRequestTravelRuleInfo struct {
	value *TravelRuleDepositRequestTravelRuleInfo
	isSet bool
}

func (v NullableTravelRuleDepositRequestTravelRuleInfo) Get() *TravelRuleDepositRequestTravelRuleInfo {
	return v.value
}

func (v *NullableTravelRuleDepositRequestTravelRuleInfo) Set(val *TravelRuleDepositRequestTravelRuleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTravelRuleDepositRequestTravelRuleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTravelRuleDepositRequestTravelRuleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTravelRuleDepositRequestTravelRuleInfo(val *TravelRuleDepositRequestTravelRuleInfo) *NullableTravelRuleDepositRequestTravelRuleInfo {
	return &NullableTravelRuleDepositRequestTravelRuleInfo{value: val, isSet: true}
}

func (v NullableTravelRuleDepositRequestTravelRuleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTravelRuleDepositRequestTravelRuleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


