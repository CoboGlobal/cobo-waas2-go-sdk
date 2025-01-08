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

// checks if the CreateClaimActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClaimActivity{}

// CreateClaimActivity struct for CreateClaimActivity
type CreateClaimActivity struct {
	// The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization.
	RequestId *string `json:"request_id,omitempty"`
	// The ID of the staking position. You can retrieve a list of staking positions by calling [List staking positions](https://www.cobo.com/developers/v2/api-references/stakings/list-staking-positions).
	StakingId string `json:"staking_id"`
	Fee *TransactionRequestFee `json:"fee,omitempty"`
}

type _CreateClaimActivity CreateClaimActivity

// NewCreateClaimActivity instantiates a new CreateClaimActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClaimActivity(stakingId string) *CreateClaimActivity {
	this := CreateClaimActivity{}
	this.StakingId = stakingId
	return &this
}

// NewCreateClaimActivityWithDefaults instantiates a new CreateClaimActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClaimActivityWithDefaults() *CreateClaimActivity {
	this := CreateClaimActivity{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateClaimActivity) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClaimActivity) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateClaimActivity) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateClaimActivity) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStakingId returns the StakingId field value
func (o *CreateClaimActivity) GetStakingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakingId
}

// GetStakingIdOk returns a tuple with the StakingId field value
// and a boolean to check if the value has been set.
func (o *CreateClaimActivity) GetStakingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingId, true
}

// SetStakingId sets field value
func (o *CreateClaimActivity) SetStakingId(v string) {
	o.StakingId = v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *CreateClaimActivity) GetFee() TransactionRequestFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionRequestFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClaimActivity) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *CreateClaimActivity) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionRequestFee and assigns it to the Fee field.
func (o *CreateClaimActivity) SetFee(v TransactionRequestFee) {
	o.Fee = &v
}

func (o CreateClaimActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClaimActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	toSerialize["staking_id"] = o.StakingId
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	return toSerialize, nil
}

func (o *CreateClaimActivity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"staking_id",
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

	varCreateClaimActivity := _CreateClaimActivity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateClaimActivity)

	if err != nil {
		return err
	}

	*o = CreateClaimActivity(varCreateClaimActivity)

	return err
}

type NullableCreateClaimActivity struct {
	value *CreateClaimActivity
	isSet bool
}

func (v NullableCreateClaimActivity) Get() *CreateClaimActivity {
	return v.value
}

func (v *NullableCreateClaimActivity) Set(val *CreateClaimActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClaimActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClaimActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClaimActivity(val *CreateClaimActivity) *NullableCreateClaimActivity {
	return &NullableCreateClaimActivity{value: val, isSet: true}
}

func (v NullableCreateClaimActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClaimActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


