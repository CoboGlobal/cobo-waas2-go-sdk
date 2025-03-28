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

// TravelRuleWithdrawRequestTravelRuleInfo - struct for TravelRuleWithdrawRequestTravelRuleInfo
type TravelRuleWithdrawRequestTravelRuleInfo struct {
	SelfCustodyWallet *SelfCustodyWallet
	TravelRuleWithdrawExchangesOrVASP *TravelRuleWithdrawExchangesOrVASP
}

// SelfCustodyWalletAsTravelRuleWithdrawRequestTravelRuleInfo is a convenience function that returns SelfCustodyWallet wrapped in TravelRuleWithdrawRequestTravelRuleInfo
func SelfCustodyWalletAsTravelRuleWithdrawRequestTravelRuleInfo(v *SelfCustodyWallet) TravelRuleWithdrawRequestTravelRuleInfo {
	return TravelRuleWithdrawRequestTravelRuleInfo{
		SelfCustodyWallet: v,
	}
}

// TravelRuleWithdrawExchangesOrVASPAsTravelRuleWithdrawRequestTravelRuleInfo is a convenience function that returns TravelRuleWithdrawExchangesOrVASP wrapped in TravelRuleWithdrawRequestTravelRuleInfo
func TravelRuleWithdrawExchangesOrVASPAsTravelRuleWithdrawRequestTravelRuleInfo(v *TravelRuleWithdrawExchangesOrVASP) TravelRuleWithdrawRequestTravelRuleInfo {
	return TravelRuleWithdrawRequestTravelRuleInfo{
		TravelRuleWithdrawExchangesOrVASP: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TravelRuleWithdrawRequestTravelRuleInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'EXCHANGES_OR_VASP'
	if jsonDict["destination_wallet_type"] == "EXCHANGES_OR_VASP" {
		// try to unmarshal JSON data into TravelRuleWithdrawExchangesOrVASP
		err = json.Unmarshal(data, &dst.TravelRuleWithdrawExchangesOrVASP)
		if err == nil {
			return nil // data stored in dst.TravelRuleWithdrawExchangesOrVASP, return on the first match
		} else {
			dst.TravelRuleWithdrawExchangesOrVASP = nil
			return fmt.Errorf("failed to unmarshal TravelRuleWithdrawRequestTravelRuleInfo as TravelRuleWithdrawExchangesOrVASP: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TravelRuleWithdrawRequestTravelRuleInfo as SelfCustodyWallet: %s", err.Error())
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
			return fmt.Errorf("failed to unmarshal TravelRuleWithdrawRequestTravelRuleInfo as SelfCustodyWallet: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TravelRuleWithdrawExchangesOrVASP'
	if jsonDict["destination_wallet_type"] == "TravelRuleWithdrawExchangesOrVASP" {
		// try to unmarshal JSON data into TravelRuleWithdrawExchangesOrVASP
		err = json.Unmarshal(data, &dst.TravelRuleWithdrawExchangesOrVASP)
		if err == nil {
			return nil // data stored in dst.TravelRuleWithdrawExchangesOrVASP, return on the first match
		} else {
			dst.TravelRuleWithdrawExchangesOrVASP = nil
			return fmt.Errorf("failed to unmarshal TravelRuleWithdrawRequestTravelRuleInfo as TravelRuleWithdrawExchangesOrVASP: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TravelRuleWithdrawRequestTravelRuleInfo) MarshalJSON() ([]byte, error) {
	if src.SelfCustodyWallet != nil {
		return json.Marshal(&src.SelfCustodyWallet)
	}

	if src.TravelRuleWithdrawExchangesOrVASP != nil {
		return json.Marshal(&src.TravelRuleWithdrawExchangesOrVASP)
	}

	return []byte(`{}`), nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TravelRuleWithdrawRequestTravelRuleInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SelfCustodyWallet != nil {
		return obj.SelfCustodyWallet
	}

	if obj.TravelRuleWithdrawExchangesOrVASP != nil {
		return obj.TravelRuleWithdrawExchangesOrVASP
	}

	// all schemas are nil
	return nil
}

type NullableTravelRuleWithdrawRequestTravelRuleInfo struct {
	value *TravelRuleWithdrawRequestTravelRuleInfo
	isSet bool
}

func (v NullableTravelRuleWithdrawRequestTravelRuleInfo) Get() *TravelRuleWithdrawRequestTravelRuleInfo {
	return v.value
}

func (v *NullableTravelRuleWithdrawRequestTravelRuleInfo) Set(val *TravelRuleWithdrawRequestTravelRuleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTravelRuleWithdrawRequestTravelRuleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTravelRuleWithdrawRequestTravelRuleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTravelRuleWithdrawRequestTravelRuleInfo(val *TravelRuleWithdrawRequestTravelRuleInfo) *NullableTravelRuleWithdrawRequestTravelRuleInfo {
	return &NullableTravelRuleWithdrawRequestTravelRuleInfo{value: val, isSet: true}
}

func (v NullableTravelRuleWithdrawRequestTravelRuleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTravelRuleWithdrawRequestTravelRuleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


