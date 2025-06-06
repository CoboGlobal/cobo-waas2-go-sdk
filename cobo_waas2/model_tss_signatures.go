/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the TSSSignatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TSSSignatures{}

// TSSSignatures struct for TSSSignatures
type TSSSignatures struct {
	Signatures []TSSSignature `json:"signatures,omitempty"`
	SignatureType *TSSSignatureType `json:"signature_type,omitempty"`
	TssProtocol *TSSProtocol `json:"tss_protocol,omitempty"`
}

// NewTSSSignatures instantiates a new TSSSignatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTSSSignatures() *TSSSignatures {
	this := TSSSignatures{}
	return &this
}

// NewTSSSignaturesWithDefaults instantiates a new TSSSignatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTSSSignaturesWithDefaults() *TSSSignatures {
	this := TSSSignatures{}
	return &this
}

// GetSignatures returns the Signatures field value if set, zero value otherwise.
func (o *TSSSignatures) GetSignatures() []TSSSignature {
	if o == nil || IsNil(o.Signatures) {
		var ret []TSSSignature
		return ret
	}
	return o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSSignatures) GetSignaturesOk() ([]TSSSignature, bool) {
	if o == nil || IsNil(o.Signatures) {
		return nil, false
	}
	return o.Signatures, true
}

// HasSignatures returns a boolean if a field has been set.
func (o *TSSSignatures) HasSignatures() bool {
	if o != nil && !IsNil(o.Signatures) {
		return true
	}

	return false
}

// SetSignatures gets a reference to the given []TSSSignature and assigns it to the Signatures field.
func (o *TSSSignatures) SetSignatures(v []TSSSignature) {
	o.Signatures = v
}

// GetSignatureType returns the SignatureType field value if set, zero value otherwise.
func (o *TSSSignatures) GetSignatureType() TSSSignatureType {
	if o == nil || IsNil(o.SignatureType) {
		var ret TSSSignatureType
		return ret
	}
	return *o.SignatureType
}

// GetSignatureTypeOk returns a tuple with the SignatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSSignatures) GetSignatureTypeOk() (*TSSSignatureType, bool) {
	if o == nil || IsNil(o.SignatureType) {
		return nil, false
	}
	return o.SignatureType, true
}

// HasSignatureType returns a boolean if a field has been set.
func (o *TSSSignatures) HasSignatureType() bool {
	if o != nil && !IsNil(o.SignatureType) {
		return true
	}

	return false
}

// SetSignatureType gets a reference to the given TSSSignatureType and assigns it to the SignatureType field.
func (o *TSSSignatures) SetSignatureType(v TSSSignatureType) {
	o.SignatureType = &v
}

// GetTssProtocol returns the TssProtocol field value if set, zero value otherwise.
func (o *TSSSignatures) GetTssProtocol() TSSProtocol {
	if o == nil || IsNil(o.TssProtocol) {
		var ret TSSProtocol
		return ret
	}
	return *o.TssProtocol
}

// GetTssProtocolOk returns a tuple with the TssProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSSignatures) GetTssProtocolOk() (*TSSProtocol, bool) {
	if o == nil || IsNil(o.TssProtocol) {
		return nil, false
	}
	return o.TssProtocol, true
}

// HasTssProtocol returns a boolean if a field has been set.
func (o *TSSSignatures) HasTssProtocol() bool {
	if o != nil && !IsNil(o.TssProtocol) {
		return true
	}

	return false
}

// SetTssProtocol gets a reference to the given TSSProtocol and assigns it to the TssProtocol field.
func (o *TSSSignatures) SetTssProtocol(v TSSProtocol) {
	o.TssProtocol = &v
}

func (o TSSSignatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TSSSignatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signatures) {
		toSerialize["signatures"] = o.Signatures
	}
	if !IsNil(o.SignatureType) {
		toSerialize["signature_type"] = o.SignatureType
	}
	if !IsNil(o.TssProtocol) {
		toSerialize["tss_protocol"] = o.TssProtocol
	}
	return toSerialize, nil
}

type NullableTSSSignatures struct {
	value *TSSSignatures
	isSet bool
}

func (v NullableTSSSignatures) Get() *TSSSignatures {
	return v.value
}

func (v *NullableTSSSignatures) Set(val *TSSSignatures) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSSignatures) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSSignatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSSignatures(val *TSSSignatures) *NullableTSSSignatures {
	return &NullableTSSSignatures{value: val, isSet: true}
}

func (v NullableTSSSignatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSSignatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


