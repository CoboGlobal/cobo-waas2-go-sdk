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

// checks if the CreateKeyGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateKeyGroupRequest{}

// CreateKeyGroupRequest struct for CreateKeyGroupRequest
type CreateKeyGroupRequest struct {
	GroupType KeyGroupType `json:"group_type"`
	// The number of key share holders in this key share group.  **Notes:** 1. Currently, the available [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) are 2-2, 2-3, and 3-3 schemes (`threshold`-`node_count`), so you can only set `node_count` to 2 or 3.   2. `threshold` must be less than or equal to `node_count`. 
	NodeCount int32 `json:"node_count"`
	// The number of key share holders required to sign an operation.  **Notes:** 1. Currently, the available [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) are 2-2, 2-3, and 3-3 schemes (`threshold`-`node_count`), so you can only set `threshold` to 2 or 3.   2. `threshold` must be less than or equal to `node_count`. 
	Threshold int32 `json:"threshold"`
	KeyHolders []CreateKeyGroupRequestKeyHoldersInner `json:"key_holders"`
}

type _CreateKeyGroupRequest CreateKeyGroupRequest

// NewCreateKeyGroupRequest instantiates a new CreateKeyGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeyGroupRequest(groupType KeyGroupType, nodeCount int32, threshold int32, keyHolders []CreateKeyGroupRequestKeyHoldersInner) *CreateKeyGroupRequest {
	this := CreateKeyGroupRequest{}
	this.GroupType = groupType
	this.NodeCount = nodeCount
	this.Threshold = threshold
	this.KeyHolders = keyHolders
	return &this
}

// NewCreateKeyGroupRequestWithDefaults instantiates a new CreateKeyGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyGroupRequestWithDefaults() *CreateKeyGroupRequest {
	this := CreateKeyGroupRequest{}
	return &this
}

// GetGroupType returns the GroupType field value
func (o *CreateKeyGroupRequest) GetGroupType() KeyGroupType {
	if o == nil {
		var ret KeyGroupType
		return ret
	}

	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value
// and a boolean to check if the value has been set.
func (o *CreateKeyGroupRequest) GetGroupTypeOk() (*KeyGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupType, true
}

// SetGroupType sets field value
func (o *CreateKeyGroupRequest) SetGroupType(v KeyGroupType) {
	o.GroupType = v
}

// GetNodeCount returns the NodeCount field value
func (o *CreateKeyGroupRequest) GetNodeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value
// and a boolean to check if the value has been set.
func (o *CreateKeyGroupRequest) GetNodeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeCount, true
}

// SetNodeCount sets field value
func (o *CreateKeyGroupRequest) SetNodeCount(v int32) {
	o.NodeCount = v
}

// GetThreshold returns the Threshold field value
func (o *CreateKeyGroupRequest) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *CreateKeyGroupRequest) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *CreateKeyGroupRequest) SetThreshold(v int32) {
	o.Threshold = v
}

// GetKeyHolders returns the KeyHolders field value
func (o *CreateKeyGroupRequest) GetKeyHolders() []CreateKeyGroupRequestKeyHoldersInner {
	if o == nil {
		var ret []CreateKeyGroupRequestKeyHoldersInner
		return ret
	}

	return o.KeyHolders
}

// GetKeyHoldersOk returns a tuple with the KeyHolders field value
// and a boolean to check if the value has been set.
func (o *CreateKeyGroupRequest) GetKeyHoldersOk() ([]CreateKeyGroupRequestKeyHoldersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyHolders, true
}

// SetKeyHolders sets field value
func (o *CreateKeyGroupRequest) SetKeyHolders(v []CreateKeyGroupRequestKeyHoldersInner) {
	o.KeyHolders = v
}

func (o CreateKeyGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateKeyGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_type"] = o.GroupType
	toSerialize["node_count"] = o.NodeCount
	toSerialize["threshold"] = o.Threshold
	toSerialize["key_holders"] = o.KeyHolders
	return toSerialize, nil
}

func (o *CreateKeyGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group_type",
		"node_count",
		"threshold",
		"key_holders",
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

	varCreateKeyGroupRequest := _CreateKeyGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateKeyGroupRequest)

	if err != nil {
		return err
	}

	*o = CreateKeyGroupRequest(varCreateKeyGroupRequest)

	return err
}

type NullableCreateKeyGroupRequest struct {
	value *CreateKeyGroupRequest
	isSet bool
}

func (v NullableCreateKeyGroupRequest) Get() *CreateKeyGroupRequest {
	return v.value
}

func (v *NullableCreateKeyGroupRequest) Set(val *CreateKeyGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeyGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeyGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeyGroupRequest(val *CreateKeyGroupRequest) *NullableCreateKeyGroupRequest {
	return &NullableCreateKeyGroupRequest{value: val, isSet: true}
}

func (v NullableCreateKeyGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeyGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


