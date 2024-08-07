/*
Cobo Wallet as a Service 2.0

Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ActivityTimeline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityTimeline{}

// ActivityTimeline The timeline of the staking activity.
type ActivityTimeline struct {
	Action ActivityAction `json:"action"`
	Status *string `json:"status,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	// The transaction ID.
	TransactionId *string `json:"transaction_id,omitempty"`
}

type _ActivityTimeline ActivityTimeline

// NewActivityTimeline instantiates a new ActivityTimeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityTimeline(action ActivityAction) *ActivityTimeline {
	this := ActivityTimeline{}
	this.Action = action
	return &this
}

// NewActivityTimelineWithDefaults instantiates a new ActivityTimeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityTimelineWithDefaults() *ActivityTimeline {
	this := ActivityTimeline{}
	return &this
}

// GetAction returns the Action field value
func (o *ActivityTimeline) GetAction() ActivityAction {
	if o == nil {
		var ret ActivityAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ActivityTimeline) GetActionOk() (*ActivityAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ActivityTimeline) SetAction(v ActivityAction) {
	o.Action = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ActivityTimeline) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityTimeline) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ActivityTimeline) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ActivityTimeline) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ActivityTimeline) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityTimeline) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ActivityTimeline) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ActivityTimeline) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *ActivityTimeline) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityTimeline) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *ActivityTimeline) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *ActivityTimeline) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o ActivityTimeline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityTimeline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return toSerialize, nil
}

func (o *ActivityTimeline) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varActivityTimeline := _ActivityTimeline{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivityTimeline)

	if err != nil {
		return err
	}

	*o = ActivityTimeline(varActivityTimeline)

	return err
}

type NullableActivityTimeline struct {
	value *ActivityTimeline
	isSet bool
}

func (v NullableActivityTimeline) Get() *ActivityTimeline {
	return v.value
}

func (v *NullableActivityTimeline) Set(val *ActivityTimeline) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityTimeline) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityTimeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityTimeline(val *ActivityTimeline) *NullableActivityTimeline {
	return &NullableActivityTimeline{value: val, isSet: true}
}

func (v NullableActivityTimeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityTimeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


