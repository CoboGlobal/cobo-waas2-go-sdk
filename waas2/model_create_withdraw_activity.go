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

// checks if the CreateWithdrawActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWithdrawActivity{}

// CreateWithdrawActivity struct for CreateWithdrawActivity
type CreateWithdrawActivity struct {
	// The id of the related staking.
	StakingId string `json:"staking_id"`
	// The amount to stake
	Amount *string `json:"amount,omitempty"`
	// The withdraw to address.
	Address *string `json:"address,omitempty"`
	Fee TransactionRequestFee `json:"fee"`
}

type _CreateWithdrawActivity CreateWithdrawActivity

// NewCreateWithdrawActivity instantiates a new CreateWithdrawActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWithdrawActivity(stakingId string, fee TransactionRequestFee) *CreateWithdrawActivity {
	this := CreateWithdrawActivity{}
	this.StakingId = stakingId
	this.Fee = fee
	return &this
}

// NewCreateWithdrawActivityWithDefaults instantiates a new CreateWithdrawActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWithdrawActivityWithDefaults() *CreateWithdrawActivity {
	this := CreateWithdrawActivity{}
	return &this
}

// GetStakingId returns the StakingId field value
func (o *CreateWithdrawActivity) GetStakingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakingId
}

// GetStakingIdOk returns a tuple with the StakingId field value
// and a boolean to check if the value has been set.
func (o *CreateWithdrawActivity) GetStakingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingId, true
}

// SetStakingId sets field value
func (o *CreateWithdrawActivity) SetStakingId(v string) {
	o.StakingId = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreateWithdrawActivity) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawActivity) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreateWithdrawActivity) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CreateWithdrawActivity) SetAmount(v string) {
	o.Amount = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CreateWithdrawActivity) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWithdrawActivity) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateWithdrawActivity) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CreateWithdrawActivity) SetAddress(v string) {
	o.Address = &v
}

// GetFee returns the Fee field value
func (o *CreateWithdrawActivity) GetFee() TransactionRequestFee {
	if o == nil {
		var ret TransactionRequestFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *CreateWithdrawActivity) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *CreateWithdrawActivity) SetFee(v TransactionRequestFee) {
	o.Fee = v
}

func (o CreateWithdrawActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWithdrawActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["staking_id"] = o.StakingId
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["fee"] = o.Fee
	return toSerialize, nil
}

func (o *CreateWithdrawActivity) UnmarshalJSON(data []byte) (err error) {
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

	varCreateWithdrawActivity := _CreateWithdrawActivity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWithdrawActivity)

	if err != nil {
		return err
	}

	*o = CreateWithdrawActivity(varCreateWithdrawActivity)

	return err
}

type NullableCreateWithdrawActivity struct {
	value *CreateWithdrawActivity
	isSet bool
}

func (v NullableCreateWithdrawActivity) Get() *CreateWithdrawActivity {
	return v.value
}

func (v *NullableCreateWithdrawActivity) Set(val *CreateWithdrawActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWithdrawActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWithdrawActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWithdrawActivity(val *CreateWithdrawActivity) *NullableCreateWithdrawActivity {
	return &NullableCreateWithdrawActivity{value: val, isSet: true}
}

func (v NullableCreateWithdrawActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWithdrawActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


