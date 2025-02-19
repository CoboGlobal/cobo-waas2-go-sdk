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

// checks if the EthStakingActivityDetailExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthStakingActivityDetailExtra{}

// EthStakingActivityDetailExtra struct for EthStakingActivityDetailExtra
type EthStakingActivityDetailExtra struct {
	PoolType StakingPoolType `json:"pool_type"`
	// The name of the provider.
	ProviderName *string `json:"provider_name,omitempty"`
	// A list of public keys associated with the Ethereum validators for this unstaking operation.
	ValidatorPubkeys []string `json:"validator_pubkeys,omitempty"`
}

type _EthStakingActivityDetailExtra EthStakingActivityDetailExtra

// NewEthStakingActivityDetailExtra instantiates a new EthStakingActivityDetailExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthStakingActivityDetailExtra(poolType StakingPoolType) *EthStakingActivityDetailExtra {
	this := EthStakingActivityDetailExtra{}
	this.PoolType = poolType
	return &this
}

// NewEthStakingActivityDetailExtraWithDefaults instantiates a new EthStakingActivityDetailExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthStakingActivityDetailExtraWithDefaults() *EthStakingActivityDetailExtra {
	this := EthStakingActivityDetailExtra{}
	return &this
}

// GetPoolType returns the PoolType field value
func (o *EthStakingActivityDetailExtra) GetPoolType() StakingPoolType {
	if o == nil {
		var ret StakingPoolType
		return ret
	}

	return o.PoolType
}

// GetPoolTypeOk returns a tuple with the PoolType field value
// and a boolean to check if the value has been set.
func (o *EthStakingActivityDetailExtra) GetPoolTypeOk() (*StakingPoolType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolType, true
}

// SetPoolType sets field value
func (o *EthStakingActivityDetailExtra) SetPoolType(v StakingPoolType) {
	o.PoolType = v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *EthStakingActivityDetailExtra) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingActivityDetailExtra) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *EthStakingActivityDetailExtra) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *EthStakingActivityDetailExtra) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetValidatorPubkeys returns the ValidatorPubkeys field value if set, zero value otherwise.
func (o *EthStakingActivityDetailExtra) GetValidatorPubkeys() []string {
	if o == nil || IsNil(o.ValidatorPubkeys) {
		var ret []string
		return ret
	}
	return o.ValidatorPubkeys
}

// GetValidatorPubkeysOk returns a tuple with the ValidatorPubkeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EthStakingActivityDetailExtra) GetValidatorPubkeysOk() ([]string, bool) {
	if o == nil || IsNil(o.ValidatorPubkeys) {
		return nil, false
	}
	return o.ValidatorPubkeys, true
}

// HasValidatorPubkeys returns a boolean if a field has been set.
func (o *EthStakingActivityDetailExtra) HasValidatorPubkeys() bool {
	if o != nil && !IsNil(o.ValidatorPubkeys) {
		return true
	}

	return false
}

// SetValidatorPubkeys gets a reference to the given []string and assigns it to the ValidatorPubkeys field.
func (o *EthStakingActivityDetailExtra) SetValidatorPubkeys(v []string) {
	o.ValidatorPubkeys = v
}

func (o EthStakingActivityDetailExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthStakingActivityDetailExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pool_type"] = o.PoolType
	if !IsNil(o.ProviderName) {
		toSerialize["provider_name"] = o.ProviderName
	}
	if !IsNil(o.ValidatorPubkeys) {
		toSerialize["validator_pubkeys"] = o.ValidatorPubkeys
	}
	return toSerialize, nil
}

func (o *EthStakingActivityDetailExtra) UnmarshalJSON(data []byte) (err error) {
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

	varEthStakingActivityDetailExtra := _EthStakingActivityDetailExtra{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEthStakingActivityDetailExtra)

	if err != nil {
		return err
	}

	*o = EthStakingActivityDetailExtra(varEthStakingActivityDetailExtra)

	return err
}

type NullableEthStakingActivityDetailExtra struct {
	value *EthStakingActivityDetailExtra
	isSet bool
}

func (v NullableEthStakingActivityDetailExtra) Get() *EthStakingActivityDetailExtra {
	return v.value
}

func (v *NullableEthStakingActivityDetailExtra) Set(val *EthStakingActivityDetailExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableEthStakingActivityDetailExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableEthStakingActivityDetailExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthStakingActivityDetailExtra(val *EthStakingActivityDetailExtra) *NullableEthStakingActivityDetailExtra {
	return &NullableEthStakingActivityDetailExtra{value: val, isSet: true}
}

func (v NullableEthStakingActivityDetailExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthStakingActivityDetailExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


