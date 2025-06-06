/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the TSSSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TSSSignature{}

// TSSSignature struct for TSSSignature
type TSSSignature struct {
	// The BIP32 path.
	Bip32Path *string `json:"bip32_path,omitempty"`
	// The message hash.
	MsgHash *string `json:"msg_hash,omitempty"`
	// The tweak.
	Tweak *string `json:"tweak,omitempty"`
	// The signature.
	Signature *string `json:"signature,omitempty"`
	// The signature recovery.
	SignatureRecovery *string `json:"signature_recovery,omitempty"`
}

// NewTSSSignature instantiates a new TSSSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTSSSignature() *TSSSignature {
	this := TSSSignature{}
	return &this
}

// NewTSSSignatureWithDefaults instantiates a new TSSSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTSSSignatureWithDefaults() *TSSSignature {
	this := TSSSignature{}
	return &this
}

// GetBip32Path returns the Bip32Path field value if set, zero value otherwise.
func (o *TSSSignature) GetBip32Path() string {
	if o == nil || IsNil(o.Bip32Path) {
		var ret string
		return ret
	}
	return *o.Bip32Path
}

// GetBip32PathOk returns a tuple with the Bip32Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSSignature) GetBip32PathOk() (*string, bool) {
	if o == nil || IsNil(o.Bip32Path) {
		return nil, false
	}
	return o.Bip32Path, true
}

// HasBip32Path returns a boolean if a field has been set.
func (o *TSSSignature) HasBip32Path() bool {
	if o != nil && !IsNil(o.Bip32Path) {
		return true
	}

	return false
}

// SetBip32Path gets a reference to the given string and assigns it to the Bip32Path field.
func (o *TSSSignature) SetBip32Path(v string) {
	o.Bip32Path = &v
}

// GetMsgHash returns the MsgHash field value if set, zero value otherwise.
func (o *TSSSignature) GetMsgHash() string {
	if o == nil || IsNil(o.MsgHash) {
		var ret string
		return ret
	}
	return *o.MsgHash
}

// GetMsgHashOk returns a tuple with the MsgHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSSignature) GetMsgHashOk() (*string, bool) {
	if o == nil || IsNil(o.MsgHash) {
		return nil, false
	}
	return o.MsgHash, true
}

// HasMsgHash returns a boolean if a field has been set.
func (o *TSSSignature) HasMsgHash() bool {
	if o != nil && !IsNil(o.MsgHash) {
		return true
	}

	return false
}

// SetMsgHash gets a reference to the given string and assigns it to the MsgHash field.
func (o *TSSSignature) SetMsgHash(v string) {
	o.MsgHash = &v
}

// GetTweak returns the Tweak field value if set, zero value otherwise.
func (o *TSSSignature) GetTweak() string {
	if o == nil || IsNil(o.Tweak) {
		var ret string
		return ret
	}
	return *o.Tweak
}

// GetTweakOk returns a tuple with the Tweak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSSignature) GetTweakOk() (*string, bool) {
	if o == nil || IsNil(o.Tweak) {
		return nil, false
	}
	return o.Tweak, true
}

// HasTweak returns a boolean if a field has been set.
func (o *TSSSignature) HasTweak() bool {
	if o != nil && !IsNil(o.Tweak) {
		return true
	}

	return false
}

// SetTweak gets a reference to the given string and assigns it to the Tweak field.
func (o *TSSSignature) SetTweak(v string) {
	o.Tweak = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *TSSSignature) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSSignature) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *TSSSignature) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *TSSSignature) SetSignature(v string) {
	o.Signature = &v
}

// GetSignatureRecovery returns the SignatureRecovery field value if set, zero value otherwise.
func (o *TSSSignature) GetSignatureRecovery() string {
	if o == nil || IsNil(o.SignatureRecovery) {
		var ret string
		return ret
	}
	return *o.SignatureRecovery
}

// GetSignatureRecoveryOk returns a tuple with the SignatureRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TSSSignature) GetSignatureRecoveryOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureRecovery) {
		return nil, false
	}
	return o.SignatureRecovery, true
}

// HasSignatureRecovery returns a boolean if a field has been set.
func (o *TSSSignature) HasSignatureRecovery() bool {
	if o != nil && !IsNil(o.SignatureRecovery) {
		return true
	}

	return false
}

// SetSignatureRecovery gets a reference to the given string and assigns it to the SignatureRecovery field.
func (o *TSSSignature) SetSignatureRecovery(v string) {
	o.SignatureRecovery = &v
}

func (o TSSSignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TSSSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bip32Path) {
		toSerialize["bip32_path"] = o.Bip32Path
	}
	if !IsNil(o.MsgHash) {
		toSerialize["msg_hash"] = o.MsgHash
	}
	if !IsNil(o.Tweak) {
		toSerialize["tweak"] = o.Tweak
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.SignatureRecovery) {
		toSerialize["signature_recovery"] = o.SignatureRecovery
	}
	return toSerialize, nil
}

type NullableTSSSignature struct {
	value *TSSSignature
	isSet bool
}

func (v NullableTSSSignature) Get() *TSSSignature {
	return v.value
}

func (v *NullableTSSSignature) Set(val *TSSSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableTSSSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableTSSSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTSSSignature(val *TSSSignature) *NullableTSSSignature {
	return &NullableTSSSignature{value: val, isSet: true}
}

func (v NullableTSSSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTSSSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


