/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SmartContractInitiator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartContractInitiator{}

// SmartContractInitiator The information about the initiator.
type SmartContractInitiator struct {
	// The initiator's wallet ID.
	WalletId string `json:"wallet_id"`
	// The initiator's wallet address. 
	Address string `json:"address"`
}

type _SmartContractInitiator SmartContractInitiator

// NewSmartContractInitiator instantiates a new SmartContractInitiator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartContractInitiator(walletId string, address string) *SmartContractInitiator {
	this := SmartContractInitiator{}
	this.WalletId = walletId
	this.Address = address
	return &this
}

// NewSmartContractInitiatorWithDefaults instantiates a new SmartContractInitiator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartContractInitiatorWithDefaults() *SmartContractInitiator {
	this := SmartContractInitiator{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *SmartContractInitiator) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *SmartContractInitiator) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *SmartContractInitiator) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddress returns the Address field value
func (o *SmartContractInitiator) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *SmartContractInitiator) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *SmartContractInitiator) SetAddress(v string) {
	o.Address = v
}

func (o SmartContractInitiator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartContractInitiator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

func (o *SmartContractInitiator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_id",
		"address",
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

	varSmartContractInitiator := _SmartContractInitiator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartContractInitiator)

	if err != nil {
		return err
	}

	*o = SmartContractInitiator(varSmartContractInitiator)

	return err
}

type NullableSmartContractInitiator struct {
	value *SmartContractInitiator
	isSet bool
}

func (v NullableSmartContractInitiator) Get() *SmartContractInitiator {
	return v.value
}

func (v *NullableSmartContractInitiator) Set(val *SmartContractInitiator) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartContractInitiator) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartContractInitiator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartContractInitiator(val *SmartContractInitiator) *NullableSmartContractInitiator {
	return &NullableSmartContractInitiator{value: val, isSet: true}
}

func (v NullableSmartContractInitiator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartContractInitiator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


