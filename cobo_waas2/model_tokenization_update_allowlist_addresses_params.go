/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TokenizationUpdateAllowlistAddressesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizationUpdateAllowlistAddressesParams{}

// TokenizationUpdateAllowlistAddressesParams struct for TokenizationUpdateAllowlistAddressesParams
type TokenizationUpdateAllowlistAddressesParams struct {
	Action TokenizationUpdateAddressAction `json:"action"`
	Source TokenizationTokenOperationSource `json:"source"`
	// A list of addresses to manage. For 'add' operations, notes can be provided. For 'remove' operations, notes are ignored.
	Addresses []TokenizationUpdateAllowlistAddressesParamsAddressesInner `json:"addresses"`
}

type _TokenizationUpdateAllowlistAddressesParams TokenizationUpdateAllowlistAddressesParams

// NewTokenizationUpdateAllowlistAddressesParams instantiates a new TokenizationUpdateAllowlistAddressesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizationUpdateAllowlistAddressesParams(action TokenizationUpdateAddressAction, source TokenizationTokenOperationSource, addresses []TokenizationUpdateAllowlistAddressesParamsAddressesInner) *TokenizationUpdateAllowlistAddressesParams {
	this := TokenizationUpdateAllowlistAddressesParams{}
	this.Action = action
	this.Source = source
	this.Addresses = addresses
	return &this
}

// NewTokenizationUpdateAllowlistAddressesParamsWithDefaults instantiates a new TokenizationUpdateAllowlistAddressesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizationUpdateAllowlistAddressesParamsWithDefaults() *TokenizationUpdateAllowlistAddressesParams {
	this := TokenizationUpdateAllowlistAddressesParams{}
	return &this
}

// GetAction returns the Action field value
func (o *TokenizationUpdateAllowlistAddressesParams) GetAction() TokenizationUpdateAddressAction {
	if o == nil {
		var ret TokenizationUpdateAddressAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *TokenizationUpdateAllowlistAddressesParams) GetActionOk() (*TokenizationUpdateAddressAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *TokenizationUpdateAllowlistAddressesParams) SetAction(v TokenizationUpdateAddressAction) {
	o.Action = v
}

// GetSource returns the Source field value
func (o *TokenizationUpdateAllowlistAddressesParams) GetSource() TokenizationTokenOperationSource {
	if o == nil {
		var ret TokenizationTokenOperationSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TokenizationUpdateAllowlistAddressesParams) GetSourceOk() (*TokenizationTokenOperationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TokenizationUpdateAllowlistAddressesParams) SetSource(v TokenizationTokenOperationSource) {
	o.Source = v
}

// GetAddresses returns the Addresses field value
func (o *TokenizationUpdateAllowlistAddressesParams) GetAddresses() []TokenizationUpdateAllowlistAddressesParamsAddressesInner {
	if o == nil {
		var ret []TokenizationUpdateAllowlistAddressesParamsAddressesInner
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *TokenizationUpdateAllowlistAddressesParams) GetAddressesOk() ([]TokenizationUpdateAllowlistAddressesParamsAddressesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *TokenizationUpdateAllowlistAddressesParams) SetAddresses(v []TokenizationUpdateAllowlistAddressesParamsAddressesInner) {
	o.Addresses = v
}

func (o TokenizationUpdateAllowlistAddressesParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizationUpdateAllowlistAddressesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["source"] = o.Source
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

func (o *TokenizationUpdateAllowlistAddressesParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"source",
		"addresses",
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

	varTokenizationUpdateAllowlistAddressesParams := _TokenizationUpdateAllowlistAddressesParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenizationUpdateAllowlistAddressesParams)

	if err != nil {
		return err
	}

	*o = TokenizationUpdateAllowlistAddressesParams(varTokenizationUpdateAllowlistAddressesParams)

	return err
}

type NullableTokenizationUpdateAllowlistAddressesParams struct {
	value *TokenizationUpdateAllowlistAddressesParams
	isSet bool
}

func (v NullableTokenizationUpdateAllowlistAddressesParams) Get() *TokenizationUpdateAllowlistAddressesParams {
	return v.value
}

func (v *NullableTokenizationUpdateAllowlistAddressesParams) Set(val *TokenizationUpdateAllowlistAddressesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizationUpdateAllowlistAddressesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizationUpdateAllowlistAddressesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizationUpdateAllowlistAddressesParams(val *TokenizationUpdateAllowlistAddressesParams) *NullableTokenizationUpdateAllowlistAddressesParams {
	return &NullableTokenizationUpdateAllowlistAddressesParams{value: val, isSet: true}
}

func (v NullableTokenizationUpdateAllowlistAddressesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizationUpdateAllowlistAddressesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


