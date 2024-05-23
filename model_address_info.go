/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AddressInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressInfo{}

// AddressInfo The data for address information.
type AddressInfo struct {
	AddressId string `json:"address_id"`
	// Then blockchain address
	AddressStr string `json:"address_str"`
	// ID of the token. Unique in all chains scope.
	TokenId string `json:"token_id"`
	// From address memo
	Memo *string `json:"memo,omitempty"`
	// Derivation path of address pubkey，required for MPC wallet
	Path *string `json:"path,omitempty"`
	Encoding *AddressEncoding `json:"encoding,omitempty"`
	// Address pubkey，required for MPC wallet
	Pubkey *string `json:"pubkey,omitempty"`
}

type _AddressInfo AddressInfo

// NewAddressInfo instantiates a new AddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressInfo(addressId string, addressStr string, tokenId string) *AddressInfo {
	this := AddressInfo{}
	this.AddressId = addressId
	this.AddressStr = addressStr
	this.TokenId = tokenId
	return &this
}

// NewAddressInfoWithDefaults instantiates a new AddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressInfoWithDefaults() *AddressInfo {
	this := AddressInfo{}
	return &this
}

// GetAddressId returns the AddressId field value
func (o *AddressInfo) GetAddressId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetAddressIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *AddressInfo) SetAddressId(v string) {
	o.AddressId = v
}

// GetAddressStr returns the AddressStr field value
func (o *AddressInfo) GetAddressStr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressStr
}

// GetAddressStrOk returns a tuple with the AddressStr field value
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetAddressStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressStr, true
}

// SetAddressStr sets field value
func (o *AddressInfo) SetAddressStr(v string) {
	o.AddressStr = v
}

// GetTokenId returns the TokenId field value
func (o *AddressInfo) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *AddressInfo) SetTokenId(v string) {
	o.TokenId = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *AddressInfo) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *AddressInfo) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *AddressInfo) SetMemo(v string) {
	o.Memo = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *AddressInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *AddressInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *AddressInfo) SetPath(v string) {
	o.Path = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *AddressInfo) GetEncoding() AddressEncoding {
	if o == nil || IsNil(o.Encoding) {
		var ret AddressEncoding
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetEncodingOk() (*AddressEncoding, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *AddressInfo) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given AddressEncoding and assigns it to the Encoding field.
func (o *AddressInfo) SetEncoding(v AddressEncoding) {
	o.Encoding = &v
}

// GetPubkey returns the Pubkey field value if set, zero value otherwise.
func (o *AddressInfo) GetPubkey() string {
	if o == nil || IsNil(o.Pubkey) {
		var ret string
		return ret
	}
	return *o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetPubkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Pubkey) {
		return nil, false
	}
	return o.Pubkey, true
}

// HasPubkey returns a boolean if a field has been set.
func (o *AddressInfo) HasPubkey() bool {
	if o != nil && !IsNil(o.Pubkey) {
		return true
	}

	return false
}

// SetPubkey gets a reference to the given string and assigns it to the Pubkey field.
func (o *AddressInfo) SetPubkey(v string) {
	o.Pubkey = &v
}

func (o AddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address_id"] = o.AddressId
	toSerialize["address_str"] = o.AddressStr
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	if !IsNil(o.Pubkey) {
		toSerialize["pubkey"] = o.Pubkey
	}
	return toSerialize, nil
}

func (o *AddressInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address_id",
		"address_str",
		"token_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddressInfo := _AddressInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressInfo)

	if err != nil {
		return err
	}

	*o = AddressInfo(varAddressInfo)

	return err
}

type NullableAddressInfo struct {
	value *AddressInfo
	isSet bool
}

func (v NullableAddressInfo) Get() *AddressInfo {
	return v.value
}

func (v *NullableAddressInfo) Set(val *AddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressInfo(val *AddressInfo) *NullableAddressInfo {
	return &NullableAddressInfo{value: val, isSet: true}
}

func (v NullableAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


