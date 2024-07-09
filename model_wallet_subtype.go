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

// WalletSubtype The wallet type. Possible values include: - `Asset`: Custodial Wallets (Asset Wallets). - `Web3`: Custodial Wallets (Web3  Wallets). - `Main`: Exchange Wallets (Main Account). - `Sub`: Exchange Wallets (Sub Account). - `Org-Controlled`: MPC Wallets (Organization-Controlled Wallets). - `User-Controlled`: MPC Wallets (User-Controlled Wallets). - `Safe{Wallet}`: Smart Contract Wallets (Safe{Wallet}). 
type WalletSubtype string

// List of WalletSubtype
const (
	WALLETSUBTYPE_ASSET WalletSubtype = "Asset"
	WALLETSUBTYPE_WEB3 WalletSubtype = "Web3"
	WALLETSUBTYPE_MAIN WalletSubtype = "Main"
	WALLETSUBTYPE_SUB WalletSubtype = "Sub"
	WALLETSUBTYPE_ORG_CONTROLLED WalletSubtype = "Org-Controlled"
	WALLETSUBTYPE_USER_CONTROLLED WalletSubtype = "User-Controlled"
	WALLETSUBTYPE_SAFEWALLET WalletSubtype = "Safe{Wallet}"
)

// All allowed values of WalletSubtype enum
var AllowedWalletSubtypeEnumValues = []WalletSubtype{
	"Asset",
	"Web3",
	"Main",
	"Sub",
	"Org-Controlled",
	"User-Controlled",
	"Safe{Wallet}",
}

func (v *WalletSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WalletSubtype(value)
	return nil
}

// NewWalletSubtypeFromValue returns a pointer to a valid WalletSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWalletSubtypeFromValue(v string) (*WalletSubtype, error) {
	ev := WalletSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WalletSubtype: valid values are %v", v, AllowedWalletSubtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WalletSubtype) IsValid() bool {
	for _, existing := range AllowedWalletSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WalletSubtype value
func (v WalletSubtype) Ptr() *WalletSubtype {
	return &v
}

type NullableWalletSubtype struct {
	value *WalletSubtype
	isSet bool
}

func (v NullableWalletSubtype) Get() *WalletSubtype {
	return v.value
}

func (v *NullableWalletSubtype) Set(val *WalletSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletSubtype(val *WalletSubtype) *NullableWalletSubtype {
	return &NullableWalletSubtype{value: val, isSet: true}
}

func (v NullableWalletSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

