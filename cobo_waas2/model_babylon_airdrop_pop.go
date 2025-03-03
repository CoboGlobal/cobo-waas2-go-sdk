/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BabylonAirdropPop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BabylonAirdropPop{}

// BabylonAirdropPop Proof of Participation (PoP) details used for airdrop registration.
type BabylonAirdropPop struct {
	// The Babylon (BABY) address used to receive BABY rewards.
	BabyAddress string `json:"baby_address"`
	// The Bitcoin (BTC) address used for staking.
	BtcAddress string `json:"btc_address"`
	// The public key corresponding to the `btc_address`, represented in hex format.
	BtcPublicKey string `json:"btc_public_key"`
	// A BTC signature that signs the `baby_address`.
	BtcSignBaby string `json:"btc_sign_baby"`
	// A BABY signature that signs the `btc_address`.
	BabySignBtc string `json:"baby_sign_btc"`
	// The public key corresponding to the `baby_address`, represented in base64 format.
	BabyPublicKey string `json:"baby_public_key"`
}

type _BabylonAirdropPop BabylonAirdropPop

// NewBabylonAirdropPop instantiates a new BabylonAirdropPop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonAirdropPop(babyAddress string, btcAddress string, btcPublicKey string, btcSignBaby string, babySignBtc string, babyPublicKey string) *BabylonAirdropPop {
	this := BabylonAirdropPop{}
	this.BabyAddress = babyAddress
	this.BtcAddress = btcAddress
	this.BtcPublicKey = btcPublicKey
	this.BtcSignBaby = btcSignBaby
	this.BabySignBtc = babySignBtc
	this.BabyPublicKey = babyPublicKey
	return &this
}

// NewBabylonAirdropPopWithDefaults instantiates a new BabylonAirdropPop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonAirdropPopWithDefaults() *BabylonAirdropPop {
	this := BabylonAirdropPop{}
	return &this
}

// GetBabyAddress returns the BabyAddress field value
func (o *BabylonAirdropPop) GetBabyAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BabyAddress
}

// GetBabyAddressOk returns a tuple with the BabyAddress field value
// and a boolean to check if the value has been set.
func (o *BabylonAirdropPop) GetBabyAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BabyAddress, true
}

// SetBabyAddress sets field value
func (o *BabylonAirdropPop) SetBabyAddress(v string) {
	o.BabyAddress = v
}

// GetBtcAddress returns the BtcAddress field value
func (o *BabylonAirdropPop) GetBtcAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtcAddress
}

// GetBtcAddressOk returns a tuple with the BtcAddress field value
// and a boolean to check if the value has been set.
func (o *BabylonAirdropPop) GetBtcAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtcAddress, true
}

// SetBtcAddress sets field value
func (o *BabylonAirdropPop) SetBtcAddress(v string) {
	o.BtcAddress = v
}

// GetBtcPublicKey returns the BtcPublicKey field value
func (o *BabylonAirdropPop) GetBtcPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtcPublicKey
}

// GetBtcPublicKeyOk returns a tuple with the BtcPublicKey field value
// and a boolean to check if the value has been set.
func (o *BabylonAirdropPop) GetBtcPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtcPublicKey, true
}

// SetBtcPublicKey sets field value
func (o *BabylonAirdropPop) SetBtcPublicKey(v string) {
	o.BtcPublicKey = v
}

// GetBtcSignBaby returns the BtcSignBaby field value
func (o *BabylonAirdropPop) GetBtcSignBaby() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BtcSignBaby
}

// GetBtcSignBabyOk returns a tuple with the BtcSignBaby field value
// and a boolean to check if the value has been set.
func (o *BabylonAirdropPop) GetBtcSignBabyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BtcSignBaby, true
}

// SetBtcSignBaby sets field value
func (o *BabylonAirdropPop) SetBtcSignBaby(v string) {
	o.BtcSignBaby = v
}

// GetBabySignBtc returns the BabySignBtc field value
func (o *BabylonAirdropPop) GetBabySignBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BabySignBtc
}

// GetBabySignBtcOk returns a tuple with the BabySignBtc field value
// and a boolean to check if the value has been set.
func (o *BabylonAirdropPop) GetBabySignBtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BabySignBtc, true
}

// SetBabySignBtc sets field value
func (o *BabylonAirdropPop) SetBabySignBtc(v string) {
	o.BabySignBtc = v
}

// GetBabyPublicKey returns the BabyPublicKey field value
func (o *BabylonAirdropPop) GetBabyPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BabyPublicKey
}

// GetBabyPublicKeyOk returns a tuple with the BabyPublicKey field value
// and a boolean to check if the value has been set.
func (o *BabylonAirdropPop) GetBabyPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BabyPublicKey, true
}

// SetBabyPublicKey sets field value
func (o *BabylonAirdropPop) SetBabyPublicKey(v string) {
	o.BabyPublicKey = v
}

func (o BabylonAirdropPop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BabylonAirdropPop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["baby_address"] = o.BabyAddress
	toSerialize["btc_address"] = o.BtcAddress
	toSerialize["btc_public_key"] = o.BtcPublicKey
	toSerialize["btc_sign_baby"] = o.BtcSignBaby
	toSerialize["baby_sign_btc"] = o.BabySignBtc
	toSerialize["baby_public_key"] = o.BabyPublicKey
	return toSerialize, nil
}

func (o *BabylonAirdropPop) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"baby_address",
		"btc_address",
		"btc_public_key",
		"btc_sign_baby",
		"baby_sign_btc",
		"baby_public_key",
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

	varBabylonAirdropPop := _BabylonAirdropPop{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBabylonAirdropPop)

	if err != nil {
		return err
	}

	*o = BabylonAirdropPop(varBabylonAirdropPop)

	return err
}

type NullableBabylonAirdropPop struct {
	value *BabylonAirdropPop
	isSet bool
}

func (v NullableBabylonAirdropPop) Get() *BabylonAirdropPop {
	return v.value
}

func (v *NullableBabylonAirdropPop) Set(val *BabylonAirdropPop) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonAirdropPop) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonAirdropPop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonAirdropPop(val *BabylonAirdropPop) *NullableBabylonAirdropPop {
	return &NullableBabylonAirdropPop{value: val, isSet: true}
}

func (v NullableBabylonAirdropPop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonAirdropPop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


