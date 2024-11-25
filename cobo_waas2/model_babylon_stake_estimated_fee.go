/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the BabylonStakeEstimatedFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BabylonStakeEstimatedFee{}

// BabylonStakeEstimatedFee struct for BabylonStakeEstimatedFee
type BabylonStakeEstimatedFee struct {
	PoolType *StakingPoolType `json:"pool_type,omitempty"`
	FeeType *FeeType `json:"fee_type,omitempty"`
	// The amount of the estimated fee.
	FeeAmount *string `json:"fee_amount,omitempty"`
	// The token ID of the staking fee.
	TokenId *string `json:"token_id,omitempty"`
}

// NewBabylonStakeEstimatedFee instantiates a new BabylonStakeEstimatedFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakeEstimatedFee() *BabylonStakeEstimatedFee {
	this := BabylonStakeEstimatedFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = &feeType
	return &this
}

// NewBabylonStakeEstimatedFeeWithDefaults instantiates a new BabylonStakeEstimatedFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakeEstimatedFeeWithDefaults() *BabylonStakeEstimatedFee {
	this := BabylonStakeEstimatedFee{}
	var feeType FeeType = FEETYPE_EVM_EIP_1559
	this.FeeType = &feeType
	return &this
}

// GetPoolType returns the PoolType field value if set, zero value otherwise.
func (o *BabylonStakeEstimatedFee) GetPoolType() StakingPoolType {
	if o == nil || IsNil(o.PoolType) {
		var ret StakingPoolType
		return ret
	}
	return *o.PoolType
}

// GetPoolTypeOk returns a tuple with the PoolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakeEstimatedFee) GetPoolTypeOk() (*StakingPoolType, bool) {
	if o == nil || IsNil(o.PoolType) {
		return nil, false
	}
	return o.PoolType, true
}

// HasPoolType returns a boolean if a field has been set.
func (o *BabylonStakeEstimatedFee) HasPoolType() bool {
	if o != nil && !IsNil(o.PoolType) {
		return true
	}

	return false
}

// SetPoolType gets a reference to the given StakingPoolType and assigns it to the PoolType field.
func (o *BabylonStakeEstimatedFee) SetPoolType(v StakingPoolType) {
	o.PoolType = &v
}

// GetFeeType returns the FeeType field value if set, zero value otherwise.
func (o *BabylonStakeEstimatedFee) GetFeeType() FeeType {
	if o == nil || IsNil(o.FeeType) {
		var ret FeeType
		return ret
	}
	return *o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakeEstimatedFee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil || IsNil(o.FeeType) {
		return nil, false
	}
	return o.FeeType, true
}

// HasFeeType returns a boolean if a field has been set.
func (o *BabylonStakeEstimatedFee) HasFeeType() bool {
	if o != nil && !IsNil(o.FeeType) {
		return true
	}

	return false
}

// SetFeeType gets a reference to the given FeeType and assigns it to the FeeType field.
func (o *BabylonStakeEstimatedFee) SetFeeType(v FeeType) {
	o.FeeType = &v
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise.
func (o *BabylonStakeEstimatedFee) GetFeeAmount() string {
	if o == nil || IsNil(o.FeeAmount) {
		var ret string
		return ret
	}
	return *o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakeEstimatedFee) GetFeeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FeeAmount) {
		return nil, false
	}
	return o.FeeAmount, true
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *BabylonStakeEstimatedFee) HasFeeAmount() bool {
	if o != nil && !IsNil(o.FeeAmount) {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given string and assigns it to the FeeAmount field.
func (o *BabylonStakeEstimatedFee) SetFeeAmount(v string) {
	o.FeeAmount = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *BabylonStakeEstimatedFee) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakeEstimatedFee) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *BabylonStakeEstimatedFee) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *BabylonStakeEstimatedFee) SetTokenId(v string) {
	o.TokenId = &v
}

func (o BabylonStakeEstimatedFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BabylonStakeEstimatedFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PoolType) {
		toSerialize["pool_type"] = o.PoolType
	}
	if !IsNil(o.FeeType) {
		toSerialize["fee_type"] = o.FeeType
	}
	if !IsNil(o.FeeAmount) {
		toSerialize["fee_amount"] = o.FeeAmount
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	return toSerialize, nil
}

type NullableBabylonStakeEstimatedFee struct {
	value *BabylonStakeEstimatedFee
	isSet bool
}

func (v NullableBabylonStakeEstimatedFee) Get() *BabylonStakeEstimatedFee {
	return v.value
}

func (v *NullableBabylonStakeEstimatedFee) Set(val *BabylonStakeEstimatedFee) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakeEstimatedFee) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakeEstimatedFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakeEstimatedFee(val *BabylonStakeEstimatedFee) *NullableBabylonStakeEstimatedFee {
	return &NullableBabylonStakeEstimatedFee{value: val, isSet: true}
}

func (v NullableBabylonStakeEstimatedFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakeEstimatedFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

