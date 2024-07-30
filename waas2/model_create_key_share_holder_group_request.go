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

// checks if the CreateKeyShareHolderGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateKeyShareHolderGroupRequest{}

// CreateKeyShareHolderGroupRequest struct for CreateKeyShareHolderGroupRequest
type CreateKeyShareHolderGroupRequest struct {
	KeyShareHolderGroupType KeyShareHolderGroupType `json:"key_share_holder_group_type"`
	// The number of key share holders in this key share holder group.  **Notes:** 1. Currently, the available [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) are 2-2, 2-3, and 3-3 schemes (in the \"threshold - node count\" format), so you can only set `node_count` to 2 or 3.   2. Threshold must be less than or equal to node count. 
	NodeCount int32 `json:"node_count"`
	// The number of key share holders required to sign an operation.  **Notes:** 1. Currently, the available [Threshold Signature Schemes (TSS)](https://manuals.cobo.com/en/portal/mpc-wallets/introduction#threshold-signature-scheme-tss) are 2-2, 2-3, and 3-3 schemes (in the \"threshold - node count\" format), so you can only set `threshold` to 2 or 3.   2. Threshold must be less than or equal to node count. 
	Threshold int32 `json:"threshold"`
	KeyShareHolders []CreateKeyShareHolder `json:"key_share_holders"`
}

type _CreateKeyShareHolderGroupRequest CreateKeyShareHolderGroupRequest

// NewCreateKeyShareHolderGroupRequest instantiates a new CreateKeyShareHolderGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateKeyShareHolderGroupRequest(keyShareHolderGroupType KeyShareHolderGroupType, nodeCount int32, threshold int32, keyShareHolders []CreateKeyShareHolder) *CreateKeyShareHolderGroupRequest {
	this := CreateKeyShareHolderGroupRequest{}
	this.KeyShareHolderGroupType = keyShareHolderGroupType
	this.NodeCount = nodeCount
	this.Threshold = threshold
	this.KeyShareHolders = keyShareHolders
	return &this
}

// NewCreateKeyShareHolderGroupRequestWithDefaults instantiates a new CreateKeyShareHolderGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateKeyShareHolderGroupRequestWithDefaults() *CreateKeyShareHolderGroupRequest {
	this := CreateKeyShareHolderGroupRequest{}
	return &this
}

// GetKeyShareHolderGroupType returns the KeyShareHolderGroupType field value
func (o *CreateKeyShareHolderGroupRequest) GetKeyShareHolderGroupType() KeyShareHolderGroupType {
	if o == nil {
		var ret KeyShareHolderGroupType
		return ret
	}

	return o.KeyShareHolderGroupType
}

// GetKeyShareHolderGroupTypeOk returns a tuple with the KeyShareHolderGroupType field value
// and a boolean to check if the value has been set.
func (o *CreateKeyShareHolderGroupRequest) GetKeyShareHolderGroupTypeOk() (*KeyShareHolderGroupType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyShareHolderGroupType, true
}

// SetKeyShareHolderGroupType sets field value
func (o *CreateKeyShareHolderGroupRequest) SetKeyShareHolderGroupType(v KeyShareHolderGroupType) {
	o.KeyShareHolderGroupType = v
}

// GetNodeCount returns the NodeCount field value
func (o *CreateKeyShareHolderGroupRequest) GetNodeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value
// and a boolean to check if the value has been set.
func (o *CreateKeyShareHolderGroupRequest) GetNodeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeCount, true
}

// SetNodeCount sets field value
func (o *CreateKeyShareHolderGroupRequest) SetNodeCount(v int32) {
	o.NodeCount = v
}

// GetThreshold returns the Threshold field value
func (o *CreateKeyShareHolderGroupRequest) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *CreateKeyShareHolderGroupRequest) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *CreateKeyShareHolderGroupRequest) SetThreshold(v int32) {
	o.Threshold = v
}

// GetKeyShareHolders returns the KeyShareHolders field value
func (o *CreateKeyShareHolderGroupRequest) GetKeyShareHolders() []CreateKeyShareHolder {
	if o == nil {
		var ret []CreateKeyShareHolder
		return ret
	}

	return o.KeyShareHolders
}

// GetKeyShareHoldersOk returns a tuple with the KeyShareHolders field value
// and a boolean to check if the value has been set.
func (o *CreateKeyShareHolderGroupRequest) GetKeyShareHoldersOk() ([]CreateKeyShareHolder, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyShareHolders, true
}

// SetKeyShareHolders sets field value
func (o *CreateKeyShareHolderGroupRequest) SetKeyShareHolders(v []CreateKeyShareHolder) {
	o.KeyShareHolders = v
}

func (o CreateKeyShareHolderGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateKeyShareHolderGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key_share_holder_group_type"] = o.KeyShareHolderGroupType
	toSerialize["node_count"] = o.NodeCount
	toSerialize["threshold"] = o.Threshold
	toSerialize["key_share_holders"] = o.KeyShareHolders
	return toSerialize, nil
}

func (o *CreateKeyShareHolderGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key_share_holder_group_type",
		"node_count",
		"threshold",
		"key_share_holders",
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

	varCreateKeyShareHolderGroupRequest := _CreateKeyShareHolderGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateKeyShareHolderGroupRequest)

	if err != nil {
		return err
	}

	*o = CreateKeyShareHolderGroupRequest(varCreateKeyShareHolderGroupRequest)

	return err
}

type NullableCreateKeyShareHolderGroupRequest struct {
	value *CreateKeyShareHolderGroupRequest
	isSet bool
}

func (v NullableCreateKeyShareHolderGroupRequest) Get() *CreateKeyShareHolderGroupRequest {
	return v.value
}

func (v *NullableCreateKeyShareHolderGroupRequest) Set(val *CreateKeyShareHolderGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateKeyShareHolderGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateKeyShareHolderGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateKeyShareHolderGroupRequest(val *CreateKeyShareHolderGroupRequest) *NullableCreateKeyShareHolderGroupRequest {
	return &NullableCreateKeyShareHolderGroupRequest{value: val, isSet: true}
}

func (v NullableCreateKeyShareHolderGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateKeyShareHolderGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

