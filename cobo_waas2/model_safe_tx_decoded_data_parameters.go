/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the SafeTxDecodedDataParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SafeTxDecodedDataParameters{}

// SafeTxDecodedDataParameters The information about the decoded parameters of the transaction.
type SafeTxDecodedDataParameters struct {
	// The name of the parameter.
	Name *string `json:"name,omitempty"`
	// The data type of the parameter.
	Type *string `json:"type,omitempty"`
	// The value of the parameter.
	Value *string `json:"value,omitempty"`
	// The decoded value of the parameter (if applicable).
	ValueDecoded []SafeTxSubTransaction `json:"value_decoded,omitempty"`
}

// NewSafeTxDecodedDataParameters instantiates a new SafeTxDecodedDataParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSafeTxDecodedDataParameters() *SafeTxDecodedDataParameters {
	this := SafeTxDecodedDataParameters{}
	return &this
}

// NewSafeTxDecodedDataParametersWithDefaults instantiates a new SafeTxDecodedDataParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSafeTxDecodedDataParametersWithDefaults() *SafeTxDecodedDataParameters {
	this := SafeTxDecodedDataParameters{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SafeTxDecodedDataParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeTxDecodedDataParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SafeTxDecodedDataParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SafeTxDecodedDataParameters) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SafeTxDecodedDataParameters) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeTxDecodedDataParameters) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SafeTxDecodedDataParameters) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SafeTxDecodedDataParameters) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SafeTxDecodedDataParameters) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeTxDecodedDataParameters) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SafeTxDecodedDataParameters) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SafeTxDecodedDataParameters) SetValue(v string) {
	o.Value = &v
}

// GetValueDecoded returns the ValueDecoded field value if set, zero value otherwise.
func (o *SafeTxDecodedDataParameters) GetValueDecoded() []SafeTxSubTransaction {
	if o == nil || IsNil(o.ValueDecoded) {
		var ret []SafeTxSubTransaction
		return ret
	}
	return o.ValueDecoded
}

// GetValueDecodedOk returns a tuple with the ValueDecoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafeTxDecodedDataParameters) GetValueDecodedOk() ([]SafeTxSubTransaction, bool) {
	if o == nil || IsNil(o.ValueDecoded) {
		return nil, false
	}
	return o.ValueDecoded, true
}

// HasValueDecoded returns a boolean if a field has been set.
func (o *SafeTxDecodedDataParameters) HasValueDecoded() bool {
	if o != nil && !IsNil(o.ValueDecoded) {
		return true
	}

	return false
}

// SetValueDecoded gets a reference to the given []SafeTxSubTransaction and assigns it to the ValueDecoded field.
func (o *SafeTxDecodedDataParameters) SetValueDecoded(v []SafeTxSubTransaction) {
	o.ValueDecoded = v
}

func (o SafeTxDecodedDataParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SafeTxDecodedDataParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueDecoded) {
		toSerialize["value_decoded"] = o.ValueDecoded
	}
	return toSerialize, nil
}

type NullableSafeTxDecodedDataParameters struct {
	value *SafeTxDecodedDataParameters
	isSet bool
}

func (v NullableSafeTxDecodedDataParameters) Get() *SafeTxDecodedDataParameters {
	return v.value
}

func (v *NullableSafeTxDecodedDataParameters) Set(val *SafeTxDecodedDataParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSafeTxDecodedDataParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSafeTxDecodedDataParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSafeTxDecodedDataParameters(val *SafeTxDecodedDataParameters) *NullableSafeTxDecodedDataParameters {
	return &NullableSafeTxDecodedDataParameters{value: val, isSet: true}
}

func (v NullableSafeTxDecodedDataParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSafeTxDecodedDataParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


