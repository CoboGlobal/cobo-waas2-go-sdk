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

// checks if the BaseCreateWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseCreateWallet{}

// BaseCreateWallet struct for BaseCreateWallet
type BaseCreateWallet struct {
	Name string `json:"name"`
}

type _BaseCreateWallet BaseCreateWallet

// NewBaseCreateWallet instantiates a new BaseCreateWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCreateWallet(name string) *BaseCreateWallet {
	this := BaseCreateWallet{}
	this.Name = name
	return &this
}

// NewBaseCreateWalletWithDefaults instantiates a new BaseCreateWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseCreateWalletWithDefaults() *BaseCreateWallet {
	this := BaseCreateWallet{}
	return &this
}

// GetName returns the Name field value
func (o *BaseCreateWallet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseCreateWallet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseCreateWallet) SetName(v string) {
	o.Name = v
}

func (o BaseCreateWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseCreateWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *BaseCreateWallet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varBaseCreateWallet := _BaseCreateWallet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseCreateWallet)

	if err != nil {
		return err
	}

	*o = BaseCreateWallet(varBaseCreateWallet)

	return err
}

type NullableBaseCreateWallet struct {
	value *BaseCreateWallet
	isSet bool
}

func (v NullableBaseCreateWallet) Get() *BaseCreateWallet {
	return v.value
}

func (v *NullableBaseCreateWallet) Set(val *BaseCreateWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCreateWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCreateWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCreateWallet(val *BaseCreateWallet) *NullableBaseCreateWallet {
	return &NullableBaseCreateWallet{value: val, isSet: true}
}

func (v NullableBaseCreateWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCreateWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


