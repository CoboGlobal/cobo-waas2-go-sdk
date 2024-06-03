/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the SignMessageSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignMessageSource{}

// SignMessageSource struct for SignMessageSource
type SignMessageSource struct {
	// Unique id of the wallet to sign message.
	WalletId *string `json:"wallet_id,omitempty"`
	// From address
	AddressStr *string `json:"address_str,omitempty"`
	MpcUsedKeyGroup *MpcSigningGroup `json:"mpc_used_key_group,omitempty"`
}

// NewSignMessageSource instantiates a new SignMessageSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignMessageSource() *SignMessageSource {
	this := SignMessageSource{}
	return &this
}

// NewSignMessageSourceWithDefaults instantiates a new SignMessageSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignMessageSourceWithDefaults() *SignMessageSource {
	this := SignMessageSource{}
	return &this
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *SignMessageSource) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignMessageSource) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *SignMessageSource) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *SignMessageSource) SetWalletId(v string) {
	o.WalletId = &v
}

// GetAddressStr returns the AddressStr field value if set, zero value otherwise.
func (o *SignMessageSource) GetAddressStr() string {
	if o == nil || IsNil(o.AddressStr) {
		var ret string
		return ret
	}
	return *o.AddressStr
}

// GetAddressStrOk returns a tuple with the AddressStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignMessageSource) GetAddressStrOk() (*string, bool) {
	if o == nil || IsNil(o.AddressStr) {
		return nil, false
	}
	return o.AddressStr, true
}

// HasAddressStr returns a boolean if a field has been set.
func (o *SignMessageSource) HasAddressStr() bool {
	if o != nil && !IsNil(o.AddressStr) {
		return true
	}

	return false
}

// SetAddressStr gets a reference to the given string and assigns it to the AddressStr field.
func (o *SignMessageSource) SetAddressStr(v string) {
	o.AddressStr = &v
}

// GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field value if set, zero value otherwise.
func (o *SignMessageSource) GetMpcUsedKeyGroup() MpcSigningGroup {
	if o == nil || IsNil(o.MpcUsedKeyGroup) {
		var ret MpcSigningGroup
		return ret
	}
	return *o.MpcUsedKeyGroup
}

// GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignMessageSource) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool) {
	if o == nil || IsNil(o.MpcUsedKeyGroup) {
		return nil, false
	}
	return o.MpcUsedKeyGroup, true
}

// HasMpcUsedKeyGroup returns a boolean if a field has been set.
func (o *SignMessageSource) HasMpcUsedKeyGroup() bool {
	if o != nil && !IsNil(o.MpcUsedKeyGroup) {
		return true
	}

	return false
}

// SetMpcUsedKeyGroup gets a reference to the given MpcSigningGroup and assigns it to the MpcUsedKeyGroup field.
func (o *SignMessageSource) SetMpcUsedKeyGroup(v MpcSigningGroup) {
	o.MpcUsedKeyGroup = &v
}

func (o SignMessageSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignMessageSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	if !IsNil(o.AddressStr) {
		toSerialize["address_str"] = o.AddressStr
	}
	if !IsNil(o.MpcUsedKeyGroup) {
		toSerialize["mpc_used_key_group"] = o.MpcUsedKeyGroup
	}
	return toSerialize, nil
}

type NullableSignMessageSource struct {
	value *SignMessageSource
	isSet bool
}

func (v NullableSignMessageSource) Get() *SignMessageSource {
	return v.value
}

func (v *NullableSignMessageSource) Set(val *SignMessageSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSignMessageSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSignMessageSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignMessageSource(val *SignMessageSource) *NullableSignMessageSource {
	return &NullableSignMessageSource{value: val, isSet: true}
}

func (v NullableSignMessageSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignMessageSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

