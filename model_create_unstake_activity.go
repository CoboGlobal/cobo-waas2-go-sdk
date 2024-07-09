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

// checks if the CreateUnstakeActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUnstakeActivity{}

// CreateUnstakeActivity struct for CreateUnstakeActivity
type CreateUnstakeActivity struct {
	// The id of the related staking.
	StakingId string `json:"staking_id"`
	// The amount to stake
	Amount *string `json:"amount,omitempty"`
	Fee TransactionTransferFee `json:"fee"`
}

type _CreateUnstakeActivity CreateUnstakeActivity

// NewCreateUnstakeActivity instantiates a new CreateUnstakeActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUnstakeActivity(stakingId string, fee TransactionTransferFee) *CreateUnstakeActivity {
	this := CreateUnstakeActivity{}
	this.StakingId = stakingId
	this.Fee = fee
	return &this
}

// NewCreateUnstakeActivityWithDefaults instantiates a new CreateUnstakeActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUnstakeActivityWithDefaults() *CreateUnstakeActivity {
	this := CreateUnstakeActivity{}
	return &this
}

// GetStakingId returns the StakingId field value
func (o *CreateUnstakeActivity) GetStakingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakingId
}

// GetStakingIdOk returns a tuple with the StakingId field value
// and a boolean to check if the value has been set.
func (o *CreateUnstakeActivity) GetStakingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingId, true
}

// SetStakingId sets field value
func (o *CreateUnstakeActivity) SetStakingId(v string) {
	o.StakingId = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreateUnstakeActivity) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUnstakeActivity) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreateUnstakeActivity) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CreateUnstakeActivity) SetAmount(v string) {
	o.Amount = &v
}

// GetFee returns the Fee field value
func (o *CreateUnstakeActivity) GetFee() TransactionTransferFee {
	if o == nil {
		var ret TransactionTransferFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *CreateUnstakeActivity) GetFeeOk() (*TransactionTransferFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *CreateUnstakeActivity) SetFee(v TransactionTransferFee) {
	o.Fee = v
}

func (o CreateUnstakeActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUnstakeActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["staking_id"] = o.StakingId
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["fee"] = o.Fee
	return toSerialize, nil
}

func (o *CreateUnstakeActivity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"staking_id",
		"fee",
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

	varCreateUnstakeActivity := _CreateUnstakeActivity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUnstakeActivity)

	if err != nil {
		return err
	}

	*o = CreateUnstakeActivity(varCreateUnstakeActivity)

	return err
}

type NullableCreateUnstakeActivity struct {
	value *CreateUnstakeActivity
	isSet bool
}

func (v NullableCreateUnstakeActivity) Get() *CreateUnstakeActivity {
	return v.value
}

func (v *NullableCreateUnstakeActivity) Set(val *CreateUnstakeActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUnstakeActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUnstakeActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUnstakeActivity(val *CreateUnstakeActivity) *NullableCreateUnstakeActivity {
	return &NullableCreateUnstakeActivity{value: val, isSet: true}
}

func (v NullableCreateUnstakeActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUnstakeActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


