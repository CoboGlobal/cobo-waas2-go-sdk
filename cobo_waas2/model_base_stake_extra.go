/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BaseStakeExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseStakeExtra{}

// BaseStakeExtra Base stake extra.
type BaseStakeExtra struct {
	PoolType StakingPoolType `json:"pool_type"`
}

type _BaseStakeExtra BaseStakeExtra

// NewBaseStakeExtra instantiates a new BaseStakeExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseStakeExtra(poolType StakingPoolType) *BaseStakeExtra {
	this := BaseStakeExtra{}
	this.PoolType = poolType
	return &this
}

// NewBaseStakeExtraWithDefaults instantiates a new BaseStakeExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseStakeExtraWithDefaults() *BaseStakeExtra {
	this := BaseStakeExtra{}
	return &this
}

// GetPoolType returns the PoolType field value
func (o *BaseStakeExtra) GetPoolType() StakingPoolType {
	if o == nil {
		var ret StakingPoolType
		return ret
	}

	return o.PoolType
}

// GetPoolTypeOk returns a tuple with the PoolType field value
// and a boolean to check if the value has been set.
func (o *BaseStakeExtra) GetPoolTypeOk() (*StakingPoolType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolType, true
}

// SetPoolType sets field value
func (o *BaseStakeExtra) SetPoolType(v StakingPoolType) {
	o.PoolType = v
}

func (o BaseStakeExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseStakeExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pool_type"] = o.PoolType
	return toSerialize, nil
}

func (o *BaseStakeExtra) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pool_type",
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

	varBaseStakeExtra := _BaseStakeExtra{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseStakeExtra)

	if err != nil {
		return err
	}

	*o = BaseStakeExtra(varBaseStakeExtra)

	return err
}

type NullableBaseStakeExtra struct {
	value *BaseStakeExtra
	isSet bool
}

func (v NullableBaseStakeExtra) Get() *BaseStakeExtra {
	return v.value
}

func (v *NullableBaseStakeExtra) Set(val *BaseStakeExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseStakeExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseStakeExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseStakeExtra(val *BaseStakeExtra) *NullableBaseStakeExtra {
	return &NullableBaseStakeExtra{value: val, isSet: true}
}

func (v NullableBaseStakeExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseStakeExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

