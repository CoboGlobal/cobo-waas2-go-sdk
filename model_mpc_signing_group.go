/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"encoding/json"
)

// checks if the MpcSigningGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MpcSigningGroup{}

// MpcSigningGroup The information about the Signing Group of an MPC Wallet.
type MpcSigningGroup struct {
	// The ID of the Signing Group.
	UsedKeyGroupId *string `json:"used_key_group_id,omitempty"`
	// The ID of the TSS Nodes that are required to participate in the signature.
	UsedNodeIds []string `json:"used_node_ids,omitempty"`
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

// GetUsedKeyGroupId returns the UsedKeyGroupId field value if set, zero value otherwise.
func (o *MpcSigningGroup) GetUsedKeyGroupId() string {
	if o == nil || IsNil(o.UsedKeyGroupId) {
		var ret string
		return ret
	}
	return *o.UsedKeyGroupId
}

// GetUsedKeyGroupIdOk returns a tuple with the UsedKeyGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcSigningGroup) GetUsedKeyGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.UsedKeyGroupId) {
		return nil, false
	}
	return o.UsedKeyGroupId, true
}

// HasUsedKeyGroupId returns a boolean if a field has been set.
func (o *MpcSigningGroup) HasUsedKeyGroupId() bool {
	if o != nil && !IsNil(o.UsedKeyGroupId) {
		return true
	}

	return false
}

// SetUsedKeyGroupId gets a reference to the given string and assigns it to the UsedKeyGroupId field.
func (o *MpcSigningGroup) SetUsedKeyGroupId(v string) {
	o.UsedKeyGroupId = &v
}

// GetUsedNodeIds returns the UsedNodeIds field value if set, zero value otherwise.
func (o *MpcSigningGroup) GetUsedNodeIds() []string {
	if o == nil || IsNil(o.UsedNodeIds) {
		var ret []string
		return ret
	}
	return o.UsedNodeIds
}

// GetUsedNodeIdsOk returns a tuple with the UsedNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MpcSigningGroup) GetUsedNodeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UsedNodeIds) {
		return nil, false
	}
	return o.UsedNodeIds, true
}

// HasUsedNodeIds returns a boolean if a field has been set.
func (o *MpcSigningGroup) HasUsedNodeIds() bool {
	if o != nil && !IsNil(o.UsedNodeIds) {
		return true
	}

	return false
}

// SetUsedNodeIds gets a reference to the given []string and assigns it to the UsedNodeIds field.
func (o *MpcSigningGroup) SetUsedNodeIds(v []string) {
	o.UsedNodeIds = v
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
	if !IsNil(o.UsedKeyGroupId) {
		toSerialize["used_key_group_id"] = o.UsedKeyGroupId
	}
	if !IsNil(o.UsedNodeIds) {
		toSerialize["used_node_ids"] = o.UsedNodeIds
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


