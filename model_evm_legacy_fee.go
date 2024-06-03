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

// checks if the EvmLegacyFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmLegacyFee{}

// EvmLegacyFee The estimated transaction fee when using the legacy method.
type EvmLegacyFee struct {
	FeeType FeeType `json:"fee_type"`
	Slow *EvmLegacyFeeSlow `json:"slow,omitempty"`
	Standard EvmLegacyFeeSlow `json:"standard"`
	Fast *EvmLegacyFeeSlow `json:"fast,omitempty"`
}

type _EvmLegacyFee EvmLegacyFee

// NewEvmLegacyFee instantiates a new EvmLegacyFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmLegacyFee(feeType FeeType, standard EvmLegacyFeeSlow) *EvmLegacyFee {
	this := EvmLegacyFee{}
	this.FeeType = feeType
	this.Standard = standard
	return &this
}

// NewEvmLegacyFeeWithDefaults instantiates a new EvmLegacyFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmLegacyFeeWithDefaults() *EvmLegacyFee {
	this := EvmLegacyFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetFeeType returns the FeeType field value
func (o *EvmLegacyFee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *EvmLegacyFee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *EvmLegacyFee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetSlow returns the Slow field value if set, zero value otherwise.
func (o *EvmLegacyFee) GetSlow() EvmLegacyFeeSlow {
	if o == nil || IsNil(o.Slow) {
		var ret EvmLegacyFeeSlow
		return ret
	}
	return *o.Slow
}

// GetSlowOk returns a tuple with the Slow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmLegacyFee) GetSlowOk() (*EvmLegacyFeeSlow, bool) {
	if o == nil || IsNil(o.Slow) {
		return nil, false
	}
	return o.Slow, true
}

// HasSlow returns a boolean if a field has been set.
func (o *EvmLegacyFee) HasSlow() bool {
	if o != nil && !IsNil(o.Slow) {
		return true
	}

	return false
}

// SetSlow gets a reference to the given EvmLegacyFeeSlow and assigns it to the Slow field.
func (o *EvmLegacyFee) SetSlow(v EvmLegacyFeeSlow) {
	o.Slow = &v
}

// GetStandard returns the Standard field value
func (o *EvmLegacyFee) GetStandard() EvmLegacyFeeSlow {
	if o == nil {
		var ret EvmLegacyFeeSlow
		return ret
	}

	return o.Standard
}

// GetStandardOk returns a tuple with the Standard field value
// and a boolean to check if the value has been set.
func (o *EvmLegacyFee) GetStandardOk() (*EvmLegacyFeeSlow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Standard, true
}

// SetStandard sets field value
func (o *EvmLegacyFee) SetStandard(v EvmLegacyFeeSlow) {
	o.Standard = v
}

// GetFast returns the Fast field value if set, zero value otherwise.
func (o *EvmLegacyFee) GetFast() EvmLegacyFeeSlow {
	if o == nil || IsNil(o.Fast) {
		var ret EvmLegacyFeeSlow
		return ret
	}
	return *o.Fast
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmLegacyFee) GetFastOk() (*EvmLegacyFeeSlow, bool) {
	if o == nil || IsNil(o.Fast) {
		return nil, false
	}
	return o.Fast, true
}

// HasFast returns a boolean if a field has been set.
func (o *EvmLegacyFee) HasFast() bool {
	if o != nil && !IsNil(o.Fast) {
		return true
	}

	return false
}

// SetFast gets a reference to the given EvmLegacyFeeSlow and assigns it to the Fast field.
func (o *EvmLegacyFee) SetFast(v EvmLegacyFeeSlow) {
	o.Fast = &v
}

func (o EvmLegacyFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmLegacyFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fee_type"] = o.FeeType
	if !IsNil(o.Slow) {
		toSerialize["slow"] = o.Slow
	}
	toSerialize["standard"] = o.Standard
	if !IsNil(o.Fast) {
		toSerialize["fast"] = o.Fast
	}
	return toSerialize, nil
}

func (o *EvmLegacyFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_type",
		"standard",
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

	varEvmLegacyFee := _EvmLegacyFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvmLegacyFee)

	if err != nil {
		return err
	}

	*o = EvmLegacyFee(varEvmLegacyFee)

	return err
}

type NullableEvmLegacyFee struct {
	value *EvmLegacyFee
	isSet bool
}

func (v NullableEvmLegacyFee) Get() *EvmLegacyFee {
	return v.value
}

func (v *NullableEvmLegacyFee) Set(val *EvmLegacyFee) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmLegacyFee) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmLegacyFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmLegacyFee(val *EvmLegacyFee) *NullableEvmLegacyFee {
	return &NullableEvmLegacyFee{value: val, isSet: true}
}

func (v NullableEvmLegacyFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmLegacyFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


