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

// checks if the UpdateKeyGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateKeyGroupRequest{}

// UpdateKeyGroupRequest struct for UpdateKeyGroupRequest
type UpdateKeyGroupRequest struct {
	// The available action of key group update.
	UpdateKeyShareGroupAction string `json:"update_key_share_group_action"`
}

type _UpdateKeyGroupRequest UpdateKeyGroupRequest

// NewUpdateKeyGroupRequest instantiates a new UpdateKeyGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateKeyGroupRequest(updateKeyShareGroupAction string) *UpdateKeyGroupRequest {
	this := UpdateKeyGroupRequest{}
	this.UpdateKeyShareGroupAction = updateKeyShareGroupAction
	return &this
}

// NewUpdateKeyGroupRequestWithDefaults instantiates a new UpdateKeyGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateKeyGroupRequestWithDefaults() *UpdateKeyGroupRequest {
	this := UpdateKeyGroupRequest{}
	return &this
}

// GetUpdateKeyShareGroupAction returns the UpdateKeyShareGroupAction field value
func (o *UpdateKeyGroupRequest) GetUpdateKeyShareGroupAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateKeyShareGroupAction
}

// GetUpdateKeyShareGroupActionOk returns a tuple with the UpdateKeyShareGroupAction field value
// and a boolean to check if the value has been set.
func (o *UpdateKeyGroupRequest) GetUpdateKeyShareGroupActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateKeyShareGroupAction, true
}

// SetUpdateKeyShareGroupAction sets field value
func (o *UpdateKeyGroupRequest) SetUpdateKeyShareGroupAction(v string) {
	o.UpdateKeyShareGroupAction = v
}

func (o UpdateKeyGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateKeyGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["update_key_share_group_action"] = o.UpdateKeyShareGroupAction
	return toSerialize, nil
}

func (o *UpdateKeyGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"update_key_share_group_action",
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

	varUpdateKeyGroupRequest := _UpdateKeyGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateKeyGroupRequest)

	if err != nil {
		return err
	}

	*o = UpdateKeyGroupRequest(varUpdateKeyGroupRequest)

	return err
}

type NullableUpdateKeyGroupRequest struct {
	value *UpdateKeyGroupRequest
	isSet bool
}

func (v NullableUpdateKeyGroupRequest) Get() *UpdateKeyGroupRequest {
	return v.value
}

func (v *NullableUpdateKeyGroupRequest) Set(val *UpdateKeyGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateKeyGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateKeyGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateKeyGroupRequest(val *UpdateKeyGroupRequest) *NullableUpdateKeyGroupRequest {
	return &NullableUpdateKeyGroupRequest{value: val, isSet: true}
}

func (v NullableUpdateKeyGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateKeyGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


