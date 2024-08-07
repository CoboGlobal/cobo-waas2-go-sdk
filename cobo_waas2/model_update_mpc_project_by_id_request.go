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

// checks if the UpdateMpcProjectByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMpcProjectByIdRequest{}

// UpdateMpcProjectByIdRequest struct for UpdateMpcProjectByIdRequest
type UpdateMpcProjectByIdRequest struct {
	// The project's new name.
	Name string `json:"name"`
}

type _UpdateMpcProjectByIdRequest UpdateMpcProjectByIdRequest

// NewUpdateMpcProjectByIdRequest instantiates a new UpdateMpcProjectByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMpcProjectByIdRequest(name string) *UpdateMpcProjectByIdRequest {
	this := UpdateMpcProjectByIdRequest{}
	this.Name = name
	return &this
}

// NewUpdateMpcProjectByIdRequestWithDefaults instantiates a new UpdateMpcProjectByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMpcProjectByIdRequestWithDefaults() *UpdateMpcProjectByIdRequest {
	this := UpdateMpcProjectByIdRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateMpcProjectByIdRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateMpcProjectByIdRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateMpcProjectByIdRequest) SetName(v string) {
	o.Name = v
}

func (o UpdateMpcProjectByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMpcProjectByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UpdateMpcProjectByIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varUpdateMpcProjectByIdRequest := _UpdateMpcProjectByIdRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMpcProjectByIdRequest)

	if err != nil {
		return err
	}

	*o = UpdateMpcProjectByIdRequest(varUpdateMpcProjectByIdRequest)

	return err
}

type NullableUpdateMpcProjectByIdRequest struct {
	value *UpdateMpcProjectByIdRequest
	isSet bool
}

func (v NullableUpdateMpcProjectByIdRequest) Get() *UpdateMpcProjectByIdRequest {
	return v.value
}

func (v *NullableUpdateMpcProjectByIdRequest) Set(val *UpdateMpcProjectByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMpcProjectByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMpcProjectByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMpcProjectByIdRequest(val *UpdateMpcProjectByIdRequest) *NullableUpdateMpcProjectByIdRequest {
	return &NullableUpdateMpcProjectByIdRequest{value: val, isSet: true}
}

func (v NullableUpdateMpcProjectByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMpcProjectByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


