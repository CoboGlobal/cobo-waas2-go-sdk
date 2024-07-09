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

// checks if the FeeAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeAmount{}

// FeeAmount The maximum transaction fee.
type FeeAmount struct {
	// The maximum fee that you are willing to pay for the transaction. The transaction will fail if the transaction fee exceeds the maximum fee.
	MaxFeeAmount string `json:"max_fee_amount"`
}

type _FeeAmount FeeAmount

// NewFeeAmount instantiates a new FeeAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeAmount(maxFeeAmount string) *FeeAmount {
	this := FeeAmount{}
	this.MaxFeeAmount = maxFeeAmount
	return &this
}

// NewFeeAmountWithDefaults instantiates a new FeeAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeAmountWithDefaults() *FeeAmount {
	this := FeeAmount{}
	return &this
}

// GetMaxFeeAmount returns the MaxFeeAmount field value
func (o *FeeAmount) GetMaxFeeAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxFeeAmount
}

// GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field value
// and a boolean to check if the value has been set.
func (o *FeeAmount) GetMaxFeeAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFeeAmount, true
}

// SetMaxFeeAmount sets field value
func (o *FeeAmount) SetMaxFeeAmount(v string) {
	o.MaxFeeAmount = v
}

func (o FeeAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max_fee_amount"] = o.MaxFeeAmount
	return toSerialize, nil
}

func (o *FeeAmount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"max_fee_amount",
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

	varFeeAmount := _FeeAmount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFeeAmount)

	if err != nil {
		return err
	}

	*o = FeeAmount(varFeeAmount)

	return err
}

type NullableFeeAmount struct {
	value *FeeAmount
	isSet bool
}

func (v NullableFeeAmount) Get() *FeeAmount {
	return v.value
}

func (v *NullableFeeAmount) Set(val *FeeAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeAmount(val *FeeAmount) *NullableFeeAmount {
	return &NullableFeeAmount{value: val, isSet: true}
}

func (v NullableFeeAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


