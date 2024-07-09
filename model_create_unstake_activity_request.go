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

// checks if the CreateUnstakeActivityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUnstakeActivityRequest{}

// CreateUnstakeActivityRequest struct for CreateUnstakeActivityRequest
type CreateUnstakeActivityRequest struct {
	// The id of the related staking.
	StakingId string `json:"staking_id"`
	// The amount to stake
	Amount *string `json:"amount,omitempty"`
	Fee TransactionTransferFee `json:"fee"`
	// The initiator of the staking activity.
	Initiator *string `json:"initiator,omitempty"`
}

type _CreateUnstakeActivityRequest CreateUnstakeActivityRequest

// NewCreateUnstakeActivityRequest instantiates a new CreateUnstakeActivityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUnstakeActivityRequest(stakingId string, fee TransactionTransferFee) *CreateUnstakeActivityRequest {
	this := CreateUnstakeActivityRequest{}
	this.StakingId = stakingId
	this.Fee = fee
	return &this
}

// NewCreateUnstakeActivityRequestWithDefaults instantiates a new CreateUnstakeActivityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUnstakeActivityRequestWithDefaults() *CreateUnstakeActivityRequest {
	this := CreateUnstakeActivityRequest{}
	return &this
}

// GetStakingId returns the StakingId field value
func (o *CreateUnstakeActivityRequest) GetStakingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakingId
}

// GetStakingIdOk returns a tuple with the StakingId field value
// and a boolean to check if the value has been set.
func (o *CreateUnstakeActivityRequest) GetStakingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingId, true
}

// SetStakingId sets field value
func (o *CreateUnstakeActivityRequest) SetStakingId(v string) {
	o.StakingId = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreateUnstakeActivityRequest) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUnstakeActivityRequest) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreateUnstakeActivityRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CreateUnstakeActivityRequest) SetAmount(v string) {
	o.Amount = &v
}

// GetFee returns the Fee field value
func (o *CreateUnstakeActivityRequest) GetFee() TransactionTransferFee {
	if o == nil {
		var ret TransactionTransferFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *CreateUnstakeActivityRequest) GetFeeOk() (*TransactionTransferFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *CreateUnstakeActivityRequest) SetFee(v TransactionTransferFee) {
	o.Fee = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *CreateUnstakeActivityRequest) GetInitiator() string {
	if o == nil || IsNil(o.Initiator) {
		var ret string
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUnstakeActivityRequest) GetInitiatorOk() (*string, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *CreateUnstakeActivityRequest) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given string and assigns it to the Initiator field.
func (o *CreateUnstakeActivityRequest) SetInitiator(v string) {
	o.Initiator = &v
}

func (o CreateUnstakeActivityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUnstakeActivityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["staking_id"] = o.StakingId
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["fee"] = o.Fee
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

func (o *CreateUnstakeActivityRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCreateUnstakeActivityRequest := _CreateUnstakeActivityRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUnstakeActivityRequest)

	if err != nil {
		return err
	}

	*o = CreateUnstakeActivityRequest(varCreateUnstakeActivityRequest)

	return err
}

type NullableCreateUnstakeActivityRequest struct {
	value *CreateUnstakeActivityRequest
	isSet bool
}

func (v NullableCreateUnstakeActivityRequest) Get() *CreateUnstakeActivityRequest {
	return v.value
}

func (v *NullableCreateUnstakeActivityRequest) Set(val *CreateUnstakeActivityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUnstakeActivityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUnstakeActivityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUnstakeActivityRequest(val *CreateUnstakeActivityRequest) *NullableCreateUnstakeActivityRequest {
	return &NullableCreateUnstakeActivityRequest{value: val, isSet: true}
}

func (v NullableCreateUnstakeActivityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUnstakeActivityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


