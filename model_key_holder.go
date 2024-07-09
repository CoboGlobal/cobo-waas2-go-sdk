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

// checks if the KeyHolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyHolder{}

// KeyHolder The data for MPC Wallets' key share holder information.
type KeyHolder struct {
	// The key share holder name.
	Name *string `json:"name,omitempty"`
	Type *KeyHolderType `json:"type,omitempty"`
	// Key share holder's TSS Node ID.
	TssNodeId *string `json:"tss_node_id,omitempty"`
	// Whether the key share holder's TSS Node is online. - `true`: The TSS Node is online.  - `false`: The TSS Node is offline. 
	Online *bool `json:"online,omitempty"`
	Status *KeyHolderStatus `json:"status,omitempty"`
}

// NewKeyHolder instantiates a new KeyHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyHolder() *KeyHolder {
	this := KeyHolder{}
	return &this
}

// NewKeyHolderWithDefaults instantiates a new KeyHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyHolderWithDefaults() *KeyHolder {
	this := KeyHolder{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeyHolder) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyHolder) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeyHolder) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeyHolder) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KeyHolder) GetType() KeyHolderType {
	if o == nil || IsNil(o.Type) {
		var ret KeyHolderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyHolder) GetTypeOk() (*KeyHolderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KeyHolder) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given KeyHolderType and assigns it to the Type field.
func (o *KeyHolder) SetType(v KeyHolderType) {
	o.Type = &v
}

// GetTssNodeId returns the TssNodeId field value if set, zero value otherwise.
func (o *KeyHolder) GetTssNodeId() string {
	if o == nil || IsNil(o.TssNodeId) {
		var ret string
		return ret
	}
	return *o.TssNodeId
}

// GetTssNodeIdOk returns a tuple with the TssNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyHolder) GetTssNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TssNodeId) {
		return nil, false
	}
	return o.TssNodeId, true
}

// HasTssNodeId returns a boolean if a field has been set.
func (o *KeyHolder) HasTssNodeId() bool {
	if o != nil && !IsNil(o.TssNodeId) {
		return true
	}

	return false
}

// SetTssNodeId gets a reference to the given string and assigns it to the TssNodeId field.
func (o *KeyHolder) SetTssNodeId(v string) {
	o.TssNodeId = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *KeyHolder) GetOnline() bool {
	if o == nil || IsNil(o.Online) {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyHolder) GetOnlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Online) {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *KeyHolder) HasOnline() bool {
	if o != nil && !IsNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *KeyHolder) SetOnline(v bool) {
	o.Online = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KeyHolder) GetStatus() KeyHolderStatus {
	if o == nil || IsNil(o.Status) {
		var ret KeyHolderStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyHolder) GetStatusOk() (*KeyHolderStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KeyHolder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given KeyHolderStatus and assigns it to the Status field.
func (o *KeyHolder) SetStatus(v KeyHolderStatus) {
	o.Status = &v
}

func (o KeyHolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyHolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TssNodeId) {
		toSerialize["tss_node_id"] = o.TssNodeId
	}
	if !IsNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableKeyHolder struct {
	value *KeyHolder
	isSet bool
}

func (v NullableKeyHolder) Get() *KeyHolder {
	return v.value
}

func (v *NullableKeyHolder) Set(val *KeyHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyHolder(val *KeyHolder) *NullableKeyHolder {
	return &NullableKeyHolder{value: val, isSet: true}
}

func (v NullableKeyHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


