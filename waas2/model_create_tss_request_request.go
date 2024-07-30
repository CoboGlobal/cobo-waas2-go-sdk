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

// checks if the CreateTssRequestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTssRequestRequest{}

// CreateTssRequestRequest struct for CreateTssRequestRequest
type CreateTssRequestRequest struct {
	Type TSSRequestType `json:"type"`
	// The target key share holder group ID.
	TargetKeyShareHolderGroupId string `json:"target_key_share_holder_group_id"`
	SourceKeyShareHolderGroup *SourceGroup `json:"source_key_share_holder_group,omitempty"`
}

type _CreateTssRequestRequest CreateTssRequestRequest

// NewCreateTssRequestRequest instantiates a new CreateTssRequestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTssRequestRequest(type_ TSSRequestType, targetKeyShareHolderGroupId string) *CreateTssRequestRequest {
	this := CreateTssRequestRequest{}
	this.Type = type_
	this.TargetKeyShareHolderGroupId = targetKeyShareHolderGroupId
	return &this
}

// NewCreateTssRequestRequestWithDefaults instantiates a new CreateTssRequestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTssRequestRequestWithDefaults() *CreateTssRequestRequest {
	this := CreateTssRequestRequest{}
	return &this
}

// GetType returns the Type field value
func (o *CreateTssRequestRequest) GetType() TSSRequestType {
	if o == nil {
		var ret TSSRequestType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateTssRequestRequest) GetTypeOk() (*TSSRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateTssRequestRequest) SetType(v TSSRequestType) {
	o.Type = v
}

// GetTargetKeyShareHolderGroupId returns the TargetKeyShareHolderGroupId field value
func (o *CreateTssRequestRequest) GetTargetKeyShareHolderGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetKeyShareHolderGroupId
}

// GetTargetKeyShareHolderGroupIdOk returns a tuple with the TargetKeyShareHolderGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateTssRequestRequest) GetTargetKeyShareHolderGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetKeyShareHolderGroupId, true
}

// SetTargetKeyShareHolderGroupId sets field value
func (o *CreateTssRequestRequest) SetTargetKeyShareHolderGroupId(v string) {
	o.TargetKeyShareHolderGroupId = v
}

// GetSourceKeyShareHolderGroup returns the SourceKeyShareHolderGroup field value if set, zero value otherwise.
func (o *CreateTssRequestRequest) GetSourceKeyShareHolderGroup() SourceGroup {
	if o == nil || IsNil(o.SourceKeyShareHolderGroup) {
		var ret SourceGroup
		return ret
	}
	return *o.SourceKeyShareHolderGroup
}

// GetSourceKeyShareHolderGroupOk returns a tuple with the SourceKeyShareHolderGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTssRequestRequest) GetSourceKeyShareHolderGroupOk() (*SourceGroup, bool) {
	if o == nil || IsNil(o.SourceKeyShareHolderGroup) {
		return nil, false
	}
	return o.SourceKeyShareHolderGroup, true
}

// HasSourceKeyShareHolderGroup returns a boolean if a field has been set.
func (o *CreateTssRequestRequest) HasSourceKeyShareHolderGroup() bool {
	if o != nil && !IsNil(o.SourceKeyShareHolderGroup) {
		return true
	}

	return false
}

// SetSourceKeyShareHolderGroup gets a reference to the given SourceGroup and assigns it to the SourceKeyShareHolderGroup field.
func (o *CreateTssRequestRequest) SetSourceKeyShareHolderGroup(v SourceGroup) {
	o.SourceKeyShareHolderGroup = &v
}

func (o CreateTssRequestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTssRequestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["target_key_share_holder_group_id"] = o.TargetKeyShareHolderGroupId
	if !IsNil(o.SourceKeyShareHolderGroup) {
		toSerialize["source_key_share_holder_group"] = o.SourceKeyShareHolderGroup
	}
	return toSerialize, nil
}

func (o *CreateTssRequestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"target_key_share_holder_group_id",
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

	varCreateTssRequestRequest := _CreateTssRequestRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTssRequestRequest)

	if err != nil {
		return err
	}

	*o = CreateTssRequestRequest(varCreateTssRequestRequest)

	return err
}

type NullableCreateTssRequestRequest struct {
	value *CreateTssRequestRequest
	isSet bool
}

func (v NullableCreateTssRequestRequest) Get() *CreateTssRequestRequest {
	return v.value
}

func (v *NullableCreateTssRequestRequest) Set(val *CreateTssRequestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTssRequestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTssRequestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTssRequestRequest(val *CreateTssRequestRequest) *NullableCreateTssRequestRequest {
	return &NullableCreateTssRequestRequest{value: val, isSet: true}
}

func (v NullableCreateTssRequestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTssRequestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


