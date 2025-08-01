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

// checks if the EstimatedFILFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimatedFILFee{}

// EstimatedFILFee The estimated transaction fee based on the fil fee model.
type EstimatedFILFee struct {
	FeeType FeeType `json:"fee_type"`
	// The token used to pay the transaction fee.
	TokenId string `json:"token_id"`
	Slow *EstimatedFILFeeSlow `json:"slow,omitempty"`
	Recommended EstimatedFILFeeSlow `json:"recommended"`
	Fast *EstimatedFILFeeSlow `json:"fast,omitempty"`
}

type _EstimatedFILFee EstimatedFILFee

// NewEstimatedFILFee instantiates a new EstimatedFILFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimatedFILFee(feeType FeeType, tokenId string, recommended EstimatedFILFeeSlow) *EstimatedFILFee {
	this := EstimatedFILFee{}
	this.FeeType = feeType
	this.TokenId = tokenId
	this.Recommended = recommended
	return &this
}

// NewEstimatedFILFeeWithDefaults instantiates a new EstimatedFILFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimatedFILFeeWithDefaults() *EstimatedFILFee {
	this := EstimatedFILFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetFeeType returns the FeeType field value
func (o *EstimatedFILFee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *EstimatedFILFee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *EstimatedFILFee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value
func (o *EstimatedFILFee) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *EstimatedFILFee) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *EstimatedFILFee) SetTokenId(v string) {
	o.TokenId = v
}

// GetSlow returns the Slow field value if set, zero value otherwise.
func (o *EstimatedFILFee) GetSlow() EstimatedFILFeeSlow {
	if o == nil || IsNil(o.Slow) {
		var ret EstimatedFILFeeSlow
		return ret
	}
	return *o.Slow
}

// GetSlowOk returns a tuple with the Slow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedFILFee) GetSlowOk() (*EstimatedFILFeeSlow, bool) {
	if o == nil || IsNil(o.Slow) {
		return nil, false
	}
	return o.Slow, true
}

// HasSlow returns a boolean if a field has been set.
func (o *EstimatedFILFee) HasSlow() bool {
	if o != nil && !IsNil(o.Slow) {
		return true
	}

	return false
}

// SetSlow gets a reference to the given EstimatedFILFeeSlow and assigns it to the Slow field.
func (o *EstimatedFILFee) SetSlow(v EstimatedFILFeeSlow) {
	o.Slow = &v
}

// GetRecommended returns the Recommended field value
func (o *EstimatedFILFee) GetRecommended() EstimatedFILFeeSlow {
	if o == nil {
		var ret EstimatedFILFeeSlow
		return ret
	}

	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *EstimatedFILFee) GetRecommendedOk() (*EstimatedFILFeeSlow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value
func (o *EstimatedFILFee) SetRecommended(v EstimatedFILFeeSlow) {
	o.Recommended = v
}

// GetFast returns the Fast field value if set, zero value otherwise.
func (o *EstimatedFILFee) GetFast() EstimatedFILFeeSlow {
	if o == nil || IsNil(o.Fast) {
		var ret EstimatedFILFeeSlow
		return ret
	}
	return *o.Fast
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedFILFee) GetFastOk() (*EstimatedFILFeeSlow, bool) {
	if o == nil || IsNil(o.Fast) {
		return nil, false
	}
	return o.Fast, true
}

// HasFast returns a boolean if a field has been set.
func (o *EstimatedFILFee) HasFast() bool {
	if o != nil && !IsNil(o.Fast) {
		return true
	}

	return false
}

// SetFast gets a reference to the given EstimatedFILFeeSlow and assigns it to the Fast field.
func (o *EstimatedFILFee) SetFast(v EstimatedFILFeeSlow) {
	o.Fast = &v
}

func (o EstimatedFILFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimatedFILFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fee_type"] = o.FeeType
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.Slow) {
		toSerialize["slow"] = o.Slow
	}
	toSerialize["recommended"] = o.Recommended
	if !IsNil(o.Fast) {
		toSerialize["fast"] = o.Fast
	}
	return toSerialize, nil
}

func (o *EstimatedFILFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_type",
		"token_id",
		"recommended",
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

	varEstimatedFILFee := _EstimatedFILFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEstimatedFILFee)

	if err != nil {
		return err
	}

	*o = EstimatedFILFee(varEstimatedFILFee)

	return err
}

type NullableEstimatedFILFee struct {
	value *EstimatedFILFee
	isSet bool
}

func (v NullableEstimatedFILFee) Get() *EstimatedFILFee {
	return v.value
}

func (v *NullableEstimatedFILFee) Set(val *EstimatedFILFee) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedFILFee) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedFILFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedFILFee(val *EstimatedFILFee) *NullableEstimatedFILFee {
	return &NullableEstimatedFILFee{value: val, isSet: true}
}

func (v NullableEstimatedFILFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedFILFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


