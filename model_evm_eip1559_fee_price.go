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

// checks if the EvmEip1559FeePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmEip1559FeePrice{}

// EvmEip1559FeePrice The transaction fee price based on the EIP-1559 fee model.
type EvmEip1559FeePrice struct {
	FeeType FeeType `json:"fee_type"`
	// The token ID of the transaction fee.
	TokenId *string `json:"token_id,omitempty"`
	Slow *EvmEip1559FeeBasePrice `json:"slow,omitempty"`
	Recommended EvmEip1559FeeBasePrice `json:"recommended"`
	Fast *EvmEip1559FeeBasePrice `json:"fast,omitempty"`
}

type _EvmEip1559FeePrice EvmEip1559FeePrice

// NewEvmEip1559FeePrice instantiates a new EvmEip1559FeePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmEip1559FeePrice(feeType FeeType, recommended EvmEip1559FeeBasePrice) *EvmEip1559FeePrice {
	this := EvmEip1559FeePrice{}
	this.FeeType = feeType
	this.Recommended = recommended
	return &this
}

// NewEvmEip1559FeePriceWithDefaults instantiates a new EvmEip1559FeePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmEip1559FeePriceWithDefaults() *EvmEip1559FeePrice {
	this := EvmEip1559FeePrice{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = feeType
	return &this
}

// GetFeeType returns the FeeType field value
func (o *EvmEip1559FeePrice) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *EvmEip1559FeePrice) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *EvmEip1559FeePrice) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *EvmEip1559FeePrice) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *EvmEip1559FeePrice) SetTokenId(v string) {
	o.TokenId = &v
}

// GetSlow returns the Slow field value if set, zero value otherwise.
func (o *EvmEip1559FeePrice) GetSlow() EvmEip1559FeeBasePrice {
	if o == nil || IsNil(o.Slow) {
		var ret EvmEip1559FeeBasePrice
		return ret
	}
	return *o.Slow
}

// GetSlowOk returns a tuple with the Slow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetSlowOk() (*EvmEip1559FeeBasePrice, bool) {
	if o == nil || IsNil(o.Slow) {
		return nil, false
	}
	return o.Slow, true
}

// HasSlow returns a boolean if a field has been set.
func (o *EvmEip1559FeePrice) HasSlow() bool {
	if o != nil && !IsNil(o.Slow) {
		return true
	}

	return false
}

// SetSlow gets a reference to the given EvmEip1559FeeBasePrice and assigns it to the Slow field.
func (o *EvmEip1559FeePrice) SetSlow(v EvmEip1559FeeBasePrice) {
	o.Slow = &v
}

// GetRecommended returns the Recommended field value
func (o *EvmEip1559FeePrice) GetRecommended() EvmEip1559FeeBasePrice {
	if o == nil {
		var ret EvmEip1559FeeBasePrice
		return ret
	}

	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetRecommendedOk() (*EvmEip1559FeeBasePrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value
func (o *EvmEip1559FeePrice) SetRecommended(v EvmEip1559FeeBasePrice) {
	o.Recommended = v
}

// GetFast returns the Fast field value if set, zero value otherwise.
func (o *EvmEip1559FeePrice) GetFast() EvmEip1559FeeBasePrice {
	if o == nil || IsNil(o.Fast) {
		var ret EvmEip1559FeeBasePrice
		return ret
	}
	return *o.Fast
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvmEip1559FeePrice) GetFastOk() (*EvmEip1559FeeBasePrice, bool) {
	if o == nil || IsNil(o.Fast) {
		return nil, false
	}
	return o.Fast, true
}

// HasFast returns a boolean if a field has been set.
func (o *EvmEip1559FeePrice) HasFast() bool {
	if o != nil && !IsNil(o.Fast) {
		return true
	}

	return false
}

// SetFast gets a reference to the given EvmEip1559FeeBasePrice and assigns it to the Fast field.
func (o *EvmEip1559FeePrice) SetFast(v EvmEip1559FeeBasePrice) {
	o.Fast = &v
}

func (o EvmEip1559FeePrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmEip1559FeePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fee_type"] = o.FeeType
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.Slow) {
		toSerialize["slow"] = o.Slow
	}
	toSerialize["recommended"] = o.Recommended
	if !IsNil(o.Fast) {
		toSerialize["fast"] = o.Fast
	}
	return toSerialize, nil
}

func (o *EvmEip1559FeePrice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_type",
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

	varEvmEip1559FeePrice := _EvmEip1559FeePrice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEvmEip1559FeePrice)

	if err != nil {
		return err
	}

	*o = EvmEip1559FeePrice(varEvmEip1559FeePrice)

	return err
}

type NullableEvmEip1559FeePrice struct {
	value *EvmEip1559FeePrice
	isSet bool
}

func (v NullableEvmEip1559FeePrice) Get() *EvmEip1559FeePrice {
	return v.value
}

func (v *NullableEvmEip1559FeePrice) Set(val *EvmEip1559FeePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmEip1559FeePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmEip1559FeePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmEip1559FeePrice(val *EvmEip1559FeePrice) *NullableEvmEip1559FeePrice {
	return &NullableEvmEip1559FeePrice{value: val, isSet: true}
}

func (v NullableEvmEip1559FeePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmEip1559FeePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


