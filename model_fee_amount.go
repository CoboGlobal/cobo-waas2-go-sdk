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

// checks if the FeeAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeAmount{}

// FeeAmount The estimated fee amount in fee_coin.
type FeeAmount struct {
	// The estimated fee amount in fee_coin.
	FeeAmount *string `json:"fee_amount,omitempty"`
}

// NewFeeAmount instantiates a new FeeAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeAmount() *FeeAmount {
	this := FeeAmount{}
	return &this
}

// NewFeeAmountWithDefaults instantiates a new FeeAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeAmountWithDefaults() *FeeAmount {
	this := FeeAmount{}
	return &this
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise.
func (o *FeeAmount) GetFeeAmount() string {
	if o == nil || IsNil(o.FeeAmount) {
		var ret string
		return ret
	}
	return *o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeAmount) GetFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FeeAmount) {
		return nil, false
	}
	return o.FeeAmount, true
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *FeeAmount) HasFeeAmount() bool {
	if o != nil && !IsNil(o.FeeAmount) {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given string and assigns it to the FeeAmount field.
func (o *FeeAmount) SetFeeAmount(v string) {
	o.FeeAmount = &v
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
	if !IsNil(o.FeeAmount) {
		toSerialize["fee_amount"] = o.FeeAmount
	}
	return toSerialize, nil
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


