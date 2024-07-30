/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
)

// checks if the MpcSigningGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MpcSigningGroup{}

// MpcSigningGroup The information about the Signing Group of an MPC Wallet.
type MpcSigningGroup struct {
	// The ID of the Signing Group.
	UsedKeyShareHolderGroupId *string `json:"used_key_share_holder_group_id,omitempty"`
	// The ID of the TSS Nodes that are required to participate in the signature.
	UsedTssNodeIds []string `json:"used_tss_node_ids,omitempty"`
}

// NewMpcSigningGroup instantiates a new MpcSigningGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMpcSigningGroup() *MpcSigningGroup {
	this := MpcSigningGroup{}
	return &this
}

// NewMpcSigningGroupWithDefaults instantiates a new MpcSigningGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMpcSigningGroupWithDefaults() *MpcSigningGroup {
	this := MpcSigningGroup{}
	return &this
}

// GetUsedKeyShareHolderGroupId returns the UsedKeyShareHolderGroupId field value if set, zero value otherwise.
func (o *MpcSigningGroup) GetUsedKeyShareHolderGroupId() string {
	if o == nil || IsNil(o.UsedKeyShareHolderGroupId) {
		var ret string
		return ret
	}
	return *o.UsedKeyShareHolderGroupId
}

// GetUsedKeyShareHolderGroupIdOk returns a tuple with the UsedKeyShareHolderGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcSigningGroup) GetUsedKeyShareHolderGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.UsedKeyShareHolderGroupId) {
		return nil, false
	}
	return o.UsedKeyShareHolderGroupId, true
}

// HasUsedKeyShareHolderGroupId returns a boolean if a field has been set.
func (o *MpcSigningGroup) HasUsedKeyShareHolderGroupId() bool {
	if o != nil && !IsNil(o.UsedKeyShareHolderGroupId) {
		return true
	}

	return false
}

// SetUsedKeyShareHolderGroupId gets a reference to the given string and assigns it to the UsedKeyShareHolderGroupId field.
func (o *MpcSigningGroup) SetUsedKeyShareHolderGroupId(v string) {
	o.UsedKeyShareHolderGroupId = &v
}

// GetUsedTssNodeIds returns the UsedTssNodeIds field value if set, zero value otherwise.
func (o *MpcSigningGroup) GetUsedTssNodeIds() []string {
	if o == nil || IsNil(o.UsedTssNodeIds) {
		var ret []string
		return ret
	}
	return o.UsedTssNodeIds
}

// GetUsedTssNodeIdsOk returns a tuple with the UsedTssNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcSigningGroup) GetUsedTssNodeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UsedTssNodeIds) {
		return nil, false
	}
	return o.UsedTssNodeIds, true
}

// HasUsedTssNodeIds returns a boolean if a field has been set.
func (o *MpcSigningGroup) HasUsedTssNodeIds() bool {
	if o != nil && !IsNil(o.UsedTssNodeIds) {
		return true
	}

	return false
}

// SetUsedTssNodeIds gets a reference to the given []string and assigns it to the UsedTssNodeIds field.
func (o *MpcSigningGroup) SetUsedTssNodeIds(v []string) {
	o.UsedTssNodeIds = v
}

func (o MpcSigningGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MpcSigningGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UsedKeyShareHolderGroupId) {
		toSerialize["used_key_share_holder_group_id"] = o.UsedKeyShareHolderGroupId
	}
	if !IsNil(o.UsedTssNodeIds) {
		toSerialize["used_tss_node_ids"] = o.UsedTssNodeIds
	}
	return toSerialize, nil
}

type NullableMpcSigningGroup struct {
	value *MpcSigningGroup
	isSet bool
}

func (v NullableMpcSigningGroup) Get() *MpcSigningGroup {
	return v.value
}

func (v *NullableMpcSigningGroup) Set(val *MpcSigningGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMpcSigningGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableMpcSigningGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMpcSigningGroup(val *MpcSigningGroup) *NullableMpcSigningGroup {
	return &NullableMpcSigningGroup{value: val, isSet: true}
}

func (v NullableMpcSigningGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMpcSigningGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


