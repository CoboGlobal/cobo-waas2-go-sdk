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

// checks if the EvmEip1559FeeRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmEip1559FeeRate{}

// EvmEip1559FeeRate The transaction fee rate based on the EIP-1559 fee model.
type EvmEip1559FeeRate struct {
	FeeType FeeType `json:"fee_type"`
	// The token ID of the transaction fee.
	TokenId string `json:"token_id"`
	Slow *EvmEip1559FeeBasePrice `json:"slow,omitempty"`
	Recommended EvmEip1559FeeBasePrice `json:"recommended"`
	Fast *EvmEip1559FeeBasePrice `json:"fast,omitempty"`
}

type _EvmEip1559FeeRate EvmEip1559FeeRate

// NewEvmEip1559FeeRate instantiates a new EvmEip1559FeeRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmEip1559FeeRate(feeType FeeType, tokenId string, recommended EvmEip1559FeeBasePrice) *EvmEip1559FeeRate {
	this := EvmEip1559FeeRate{}
	this.FeeType = feeType
	this.TokenId = tokenId
	this.Recommended = recommended
	return &this
}

// NewEvmEip1559FeeRateWithDefaults instantiates a new EvmEip1559FeeRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmEip1559FeeRateWithDefaults() *EvmEip1559FeeRate {
	this := EvmEip1559FeeRate{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetFeeType returns the FeeType field value
func (o *EvmEip1559FeeRate) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeRate) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *EvmEip1559FeeRate) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value
func (o *EvmEip1559FeeRate) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeRate) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *EvmEip1559FeeRate) SetTokenId(v string) {
	o.TokenId = v
}

// GetSlow returns the Slow field value if set, zero value otherwise.
func (o *EvmEip1559FeeRate) GetSlow() EvmEip1559FeeBasePrice {
	if o == nil || IsNil(o.Slow) {
		var ret EvmEip1559FeeBasePrice
		return ret
	}
	return *o.Slow
}

// GetSlowOk returns a tuple with the Slow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeRate) GetSlowOk() (*EvmEip1559FeeBasePrice, bool) {
	if o == nil || IsNil(o.Slow) {
		return nil, false
	}
	return o.Slow, true
}

// HasSlow returns a boolean if a field has been set.
func (o *EvmEip1559FeeRate) HasSlow() bool {
	if o != nil && !IsNil(o.Slow) {
		return true
	}

	return false
}

// SetSlow gets a reference to the given EvmEip1559FeeBasePrice and assigns it to the Slow field.
func (o *EvmEip1559FeeRate) SetSlow(v EvmEip1559FeeBasePrice) {
	o.Slow = &v
}

// GetRecommended returns the Recommended field value
func (o *EvmEip1559FeeRate) GetRecommended() EvmEip1559FeeBasePrice {
	if o == nil {
		var ret EvmEip1559FeeBasePrice
		return ret
	}

	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeRate) GetRecommendedOk() (*EvmEip1559FeeBasePrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value
func (o *EvmEip1559FeeRate) SetRecommended(v EvmEip1559FeeBasePrice) {
	o.Recommended = v
}

// GetFast returns the Fast field value if set, zero value otherwise.
func (o *EvmEip1559FeeRate) GetFast() EvmEip1559FeeBasePrice {
	if o == nil || IsNil(o.Fast) {
		var ret EvmEip1559FeeBasePrice
		return ret
	}
	return *o.Fast
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeeRate) GetFastOk() (*EvmEip1559FeeBasePrice, bool) {
	if o == nil || IsNil(o.Fast) {
		return nil, false
	}
	return o.Fast, true
}

// HasFast returns a boolean if a field has been set.
func (o *EvmEip1559FeeRate) HasFast() bool {
	if o != nil && !IsNil(o.Fast) {
		return true
	}

	return false
}

// SetFast gets a reference to the given EvmEip1559FeeBasePrice and assigns it to the Fast field.
func (o *EvmEip1559FeeRate) SetFast(v EvmEip1559FeeBasePrice) {
	o.Fast = &v
}

func (o EvmEip1559FeeRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmEip1559FeeRate) ToMap() (map[string]interface{}, error) {
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

func (o *EvmEip1559FeeRate) UnmarshalJSON(data []byte) (err error) {
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

	varEvmEip1559FeeRate := _EvmEip1559FeeRate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvmEip1559FeeRate)

	if err != nil {
		return err
	}

	*o = EvmEip1559FeeRate(varEvmEip1559FeeRate)

	return err
}

type NullableEvmEip1559FeeRate struct {
	value *EvmEip1559FeeRate
	isSet bool
}

func (v NullableEvmEip1559FeeRate) Get() *EvmEip1559FeeRate {
	return v.value
}

func (v *NullableEvmEip1559FeeRate) Set(val *EvmEip1559FeeRate) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmEip1559FeeRate) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmEip1559FeeRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmEip1559FeeRate(val *EvmEip1559FeeRate) *NullableEvmEip1559FeeRate {
	return &NullableEvmEip1559FeeRate{value: val, isSet: true}
}

func (v NullableEvmEip1559FeeRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmEip1559FeeRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


