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

// checks if the BabylonStakingActivityDetailExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BabylonStakingActivityDetailExtra{}

// BabylonStakingActivityDetailExtra struct for BabylonStakingActivityDetailExtra
type BabylonStakingActivityDetailExtra struct {
	PoolType StakingPoolType `json:"pool_type"`
	// The public key of the finality provider.
	FinalityProviderPublicKey *string `json:"finality_provider_public_key,omitempty"`
	// The number of blocks that need to be processed before the locked tokens are unlocked and become accessible.
	StakeBlockTime *int64 `json:"stake_block_time,omitempty"`
	// Whether to automatically broadcast the transaction.  - `true`: Automatically broadcast the transaction. - `false`: The transaction will not be submitted to the blockchain automatically. You can call [Broadcast signed transactions](/v2/api-references/transactions/broadcast-signed-transactions) to broadcast the transaction to the blockchain, or retrieve the signed raw transaction data `raw_tx` by calling [Get transaction information](/v2/api-references/transactions/get-transaction-information) and broadcast it yourself. 
	AutoBroadcast *bool `json:"auto_broadcast,omitempty"`
	// The version of babylon global parameters.
	ParamVersion *int64 `json:"param_version,omitempty"`
	WithdrawFromType *ActivityType `json:"withdraw_from_type,omitempty"`
}

type _BabylonStakingActivityDetailExtra BabylonStakingActivityDetailExtra

// NewBabylonStakingActivityDetailExtra instantiates a new BabylonStakingActivityDetailExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakingActivityDetailExtra(poolType StakingPoolType) *BabylonStakingActivityDetailExtra {
	this := BabylonStakingActivityDetailExtra{}
	this.PoolType = poolType
	return &this
}

// NewBabylonStakingActivityDetailExtraWithDefaults instantiates a new BabylonStakingActivityDetailExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakingActivityDetailExtraWithDefaults() *BabylonStakingActivityDetailExtra {
	this := BabylonStakingActivityDetailExtra{}
	return &this
}

// GetPoolType returns the PoolType field value
func (o *BabylonStakingActivityDetailExtra) GetPoolType() StakingPoolType {
	if o == nil {
		var ret StakingPoolType
		return ret
	}

	return o.PoolType
}

// GetPoolTypeOk returns a tuple with the PoolType field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingActivityDetailExtra) GetPoolTypeOk() (*StakingPoolType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolType, true
}

// SetPoolType sets field value
func (o *BabylonStakingActivityDetailExtra) SetPoolType(v StakingPoolType) {
	o.PoolType = v
}

// GetFinalityProviderPublicKey returns the FinalityProviderPublicKey field value if set, zero value otherwise.
func (o *BabylonStakingActivityDetailExtra) GetFinalityProviderPublicKey() string {
	if o == nil || IsNil(o.FinalityProviderPublicKey) {
		var ret string
		return ret
	}
	return *o.FinalityProviderPublicKey
}

// GetFinalityProviderPublicKeyOk returns a tuple with the FinalityProviderPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingActivityDetailExtra) GetFinalityProviderPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.FinalityProviderPublicKey) {
		return nil, false
	}
	return o.FinalityProviderPublicKey, true
}

// HasFinalityProviderPublicKey returns a boolean if a field has been set.
func (o *BabylonStakingActivityDetailExtra) HasFinalityProviderPublicKey() bool {
	if o != nil && !IsNil(o.FinalityProviderPublicKey) {
		return true
	}

	return false
}

// SetFinalityProviderPublicKey gets a reference to the given string and assigns it to the FinalityProviderPublicKey field.
func (o *BabylonStakingActivityDetailExtra) SetFinalityProviderPublicKey(v string) {
	o.FinalityProviderPublicKey = &v
}

// GetStakeBlockTime returns the StakeBlockTime field value if set, zero value otherwise.
func (o *BabylonStakingActivityDetailExtra) GetStakeBlockTime() int64 {
	if o == nil || IsNil(o.StakeBlockTime) {
		var ret int64
		return ret
	}
	return *o.StakeBlockTime
}

// GetStakeBlockTimeOk returns a tuple with the StakeBlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingActivityDetailExtra) GetStakeBlockTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StakeBlockTime) {
		return nil, false
	}
	return o.StakeBlockTime, true
}

// HasStakeBlockTime returns a boolean if a field has been set.
func (o *BabylonStakingActivityDetailExtra) HasStakeBlockTime() bool {
	if o != nil && !IsNil(o.StakeBlockTime) {
		return true
	}

	return false
}

// SetStakeBlockTime gets a reference to the given int64 and assigns it to the StakeBlockTime field.
func (o *BabylonStakingActivityDetailExtra) SetStakeBlockTime(v int64) {
	o.StakeBlockTime = &v
}

// GetAutoBroadcast returns the AutoBroadcast field value if set, zero value otherwise.
func (o *BabylonStakingActivityDetailExtra) GetAutoBroadcast() bool {
	if o == nil || IsNil(o.AutoBroadcast) {
		var ret bool
		return ret
	}
	return *o.AutoBroadcast
}

// GetAutoBroadcastOk returns a tuple with the AutoBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingActivityDetailExtra) GetAutoBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoBroadcast) {
		return nil, false
	}
	return o.AutoBroadcast, true
}

// HasAutoBroadcast returns a boolean if a field has been set.
func (o *BabylonStakingActivityDetailExtra) HasAutoBroadcast() bool {
	if o != nil && !IsNil(o.AutoBroadcast) {
		return true
	}

	return false
}

// SetAutoBroadcast gets a reference to the given bool and assigns it to the AutoBroadcast field.
func (o *BabylonStakingActivityDetailExtra) SetAutoBroadcast(v bool) {
	o.AutoBroadcast = &v
}

// GetParamVersion returns the ParamVersion field value if set, zero value otherwise.
func (o *BabylonStakingActivityDetailExtra) GetParamVersion() int64 {
	if o == nil || IsNil(o.ParamVersion) {
		var ret int64
		return ret
	}
	return *o.ParamVersion
}

// GetParamVersionOk returns a tuple with the ParamVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingActivityDetailExtra) GetParamVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.ParamVersion) {
		return nil, false
	}
	return o.ParamVersion, true
}

// HasParamVersion returns a boolean if a field has been set.
func (o *BabylonStakingActivityDetailExtra) HasParamVersion() bool {
	if o != nil && !IsNil(o.ParamVersion) {
		return true
	}

	return false
}

// SetParamVersion gets a reference to the given int64 and assigns it to the ParamVersion field.
func (o *BabylonStakingActivityDetailExtra) SetParamVersion(v int64) {
	o.ParamVersion = &v
}

// GetWithdrawFromType returns the WithdrawFromType field value if set, zero value otherwise.
func (o *BabylonStakingActivityDetailExtra) GetWithdrawFromType() ActivityType {
	if o == nil || IsNil(o.WithdrawFromType) {
		var ret ActivityType
		return ret
	}
	return *o.WithdrawFromType
}

// GetWithdrawFromTypeOk returns a tuple with the WithdrawFromType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BabylonStakingActivityDetailExtra) GetWithdrawFromTypeOk() (*ActivityType, bool) {
	if o == nil || IsNil(o.WithdrawFromType) {
		return nil, false
	}
	return o.WithdrawFromType, true
}

// HasWithdrawFromType returns a boolean if a field has been set.
func (o *BabylonStakingActivityDetailExtra) HasWithdrawFromType() bool {
	if o != nil && !IsNil(o.WithdrawFromType) {
		return true
	}

	return false
}

// SetWithdrawFromType gets a reference to the given ActivityType and assigns it to the WithdrawFromType field.
func (o *BabylonStakingActivityDetailExtra) SetWithdrawFromType(v ActivityType) {
	o.WithdrawFromType = &v
}

func (o BabylonStakingActivityDetailExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BabylonStakingActivityDetailExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pool_type"] = o.PoolType
	if !IsNil(o.FinalityProviderPublicKey) {
		toSerialize["finality_provider_public_key"] = o.FinalityProviderPublicKey
	}
	if !IsNil(o.StakeBlockTime) {
		toSerialize["stake_block_time"] = o.StakeBlockTime
	}
	if !IsNil(o.AutoBroadcast) {
		toSerialize["auto_broadcast"] = o.AutoBroadcast
	}
	if !IsNil(o.ParamVersion) {
		toSerialize["param_version"] = o.ParamVersion
	}
	if !IsNil(o.WithdrawFromType) {
		toSerialize["withdraw_from_type"] = o.WithdrawFromType
	}
	return toSerialize, nil
}

func (o *BabylonStakingActivityDetailExtra) UnmarshalJSON(data []byte) (err error) {
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

	varBabylonStakingActivityDetailExtra := _BabylonStakingActivityDetailExtra{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBabylonStakingActivityDetailExtra)

	if err != nil {
		return err
	}

	*o = BabylonStakingActivityDetailExtra(varBabylonStakingActivityDetailExtra)

	return err
}

type NullableBabylonStakingActivityDetailExtra struct {
	value *BabylonStakingActivityDetailExtra
	isSet bool
}

func (v NullableBabylonStakingActivityDetailExtra) Get() *BabylonStakingActivityDetailExtra {
	return v.value
}

func (v *NullableBabylonStakingActivityDetailExtra) Set(val *BabylonStakingActivityDetailExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakingActivityDetailExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakingActivityDetailExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakingActivityDetailExtra(val *BabylonStakingActivityDetailExtra) *NullableBabylonStakingActivityDetailExtra {
	return &NullableBabylonStakingActivityDetailExtra{value: val, isSet: true}
}

func (v NullableBabylonStakingActivityDetailExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakingActivityDetailExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


