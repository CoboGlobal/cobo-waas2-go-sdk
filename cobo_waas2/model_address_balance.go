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

// checks if the AddressBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressBalance{}

// AddressBalance The token balance for a specific wallet address.
type AddressBalance struct {
	// The wallet address.
	Address string `json:"address"`
	Balance Balance `json:"balance"`
}

type _AddressBalance AddressBalance

// NewAddressBalance instantiates a new AddressBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressBalance(address string, balance Balance) *AddressBalance {
	this := AddressBalance{}
	this.Address = address
	this.Balance = balance
	return &this
}

// NewAddressBalanceWithDefaults instantiates a new AddressBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressBalanceWithDefaults() *AddressBalance {
	this := AddressBalance{}
	return &this
}

// GetAddress returns the Address field value
func (o *AddressBalance) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressBalance) SetAddress(v string) {
	o.Address = v
}

// GetBalance returns the Balance field value
func (o *AddressBalance) GetBalance() Balance {
	if o == nil {
		var ret Balance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetBalanceOk() (*Balance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *AddressBalance) SetBalance(v Balance) {
	o.Balance = v
}

func (o AddressBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["balance"] = o.Balance
	return toSerialize, nil
}

func (o *AddressBalance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"balance",
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

	varAddressBalance := _AddressBalance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressBalance)

	if err != nil {
		return err
	}

	*o = AddressBalance(varAddressBalance)

	return err
}

type NullableAddressBalance struct {
	value *AddressBalance
	isSet bool
}

func (v NullableAddressBalance) Get() *AddressBalance {
	return v.value
}

func (v *NullableAddressBalance) Set(val *AddressBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBalance(val *AddressBalance) *NullableAddressBalance {
	return &NullableAddressBalance{value: val, isSet: true}
}

func (v NullableAddressBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


