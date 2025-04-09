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

// checks if the BankAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankAccount{}

// BankAccount struct for BankAccount
type BankAccount struct {
	// The bank account ID.
	BankAccountId string `json:"bank_account_id"`
	// JSON-formatted bank account details.
	Info map[string]interface{} `json:"info"`
}

type _BankAccount BankAccount

// NewBankAccount instantiates a new BankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccount(bankAccountId string, info map[string]interface{}) *BankAccount {
	this := BankAccount{}
	this.BankAccountId = bankAccountId
	this.Info = info
	return &this
}

// NewBankAccountWithDefaults instantiates a new BankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountWithDefaults() *BankAccount {
	this := BankAccount{}
	return &this
}

// GetBankAccountId returns the BankAccountId field value
func (o *BankAccount) GetBankAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankAccountId
}

// GetBankAccountIdOk returns a tuple with the BankAccountId field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccountId, true
}

// SetBankAccountId sets field value
func (o *BankAccount) SetBankAccountId(v string) {
	o.BankAccountId = v
}

// GetInfo returns the Info field value
func (o *BankAccount) GetInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetInfoOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Info, true
}

// SetInfo sets field value
func (o *BankAccount) SetInfo(v map[string]interface{}) {
	o.Info = v
}

func (o BankAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bank_account_id"] = o.BankAccountId
	toSerialize["info"] = o.Info
	return toSerialize, nil
}

func (o *BankAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bank_account_id",
		"info",
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

	varBankAccount := _BankAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBankAccount)

	if err != nil {
		return err
	}

	*o = BankAccount(varBankAccount)

	return err
}

type NullableBankAccount struct {
	value *BankAccount
	isSet bool
}

func (v NullableBankAccount) Get() *BankAccount {
	return v.value
}

func (v *NullableBankAccount) Set(val *BankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccount(val *BankAccount) *NullableBankAccount {
	return &NullableBankAccount{value: val, isSet: true}
}

func (v NullableBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


