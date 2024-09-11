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

// checks if the CreateStakeActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStakeActivity{}

// CreateStakeActivity struct for CreateStakeActivity
type CreateStakeActivity struct {
	// The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization.
	RequestId *string `json:"request_id,omitempty"`
	Source *StakingSource `json:"source,omitempty"`
	// The ID of the staking pool.
	PoolId string `json:"pool_id"`
	// The amount to stake.
	Amount string `json:"amount"`
	Fee TransactionRequestFee `json:"fee"`
	Extra CreateStakeActivityExtra `json:"extra"`
}

type _CreateStakeActivity CreateStakeActivity

// NewCreateStakeActivity instantiates a new CreateStakeActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStakeActivity(poolId string, amount string, fee TransactionRequestFee, extra CreateStakeActivityExtra) *CreateStakeActivity {
	this := CreateStakeActivity{}
	this.PoolId = poolId
	this.Amount = amount
	this.Fee = fee
	this.Extra = extra
	return &this
}

// NewCreateStakeActivityWithDefaults instantiates a new CreateStakeActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStakeActivityWithDefaults() *CreateStakeActivity {
	this := CreateStakeActivity{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateStakeActivity) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStakeActivity) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateStakeActivity) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateStakeActivity) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CreateStakeActivity) GetSource() StakingSource {
	if o == nil || IsNil(o.Source) {
		var ret StakingSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStakeActivity) GetSourceOk() (*StakingSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CreateStakeActivity) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given StakingSource and assigns it to the Source field.
func (o *CreateStakeActivity) SetSource(v StakingSource) {
	o.Source = &v
}

// GetPoolId returns the PoolId field value
func (o *CreateStakeActivity) GetPoolId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *CreateStakeActivity) GetPoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *CreateStakeActivity) SetPoolId(v string) {
	o.PoolId = v
}

// GetAmount returns the Amount field value
func (o *CreateStakeActivity) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateStakeActivity) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateStakeActivity) SetAmount(v string) {
	o.Amount = v
}

// GetFee returns the Fee field value
func (o *CreateStakeActivity) GetFee() TransactionRequestFee {
	if o == nil {
		var ret TransactionRequestFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *CreateStakeActivity) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *CreateStakeActivity) SetFee(v TransactionRequestFee) {
	o.Fee = v
}

// GetExtra returns the Extra field value
func (o *CreateStakeActivity) GetExtra() CreateStakeActivityExtra {
	if o == nil {
		var ret CreateStakeActivityExtra
		return ret
	}

	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value
// and a boolean to check if the value has been set.
func (o *CreateStakeActivity) GetExtraOk() (*CreateStakeActivityExtra, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extra, true
}

// SetExtra sets field value
func (o *CreateStakeActivity) SetExtra(v CreateStakeActivityExtra) {
	o.Extra = v
}

func (o CreateStakeActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStakeActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["pool_id"] = o.PoolId
	toSerialize["amount"] = o.Amount
	toSerialize["fee"] = o.Fee
	toSerialize["extra"] = o.Extra
	return toSerialize, nil
}

func (o *CreateStakeActivity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pool_id",
		"amount",
		"fee",
		"extra",
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

	varCreateStakeActivity := _CreateStakeActivity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateStakeActivity)

	if err != nil {
		return err
	}

	*o = CreateStakeActivity(varCreateStakeActivity)

	return err
}

type NullableCreateStakeActivity struct {
	value *CreateStakeActivity
	isSet bool
}

func (v NullableCreateStakeActivity) Get() *CreateStakeActivity {
	return v.value
}

func (v *NullableCreateStakeActivity) Set(val *CreateStakeActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStakeActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStakeActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStakeActivity(val *CreateStakeActivity) *NullableCreateStakeActivity {
	return &NullableCreateStakeActivity{value: val, isSet: true}
}

func (v NullableCreateStakeActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStakeActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


