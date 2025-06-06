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

// checks if the FeeStationDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeStationDestination{}

// FeeStationDestination struct for FeeStationDestination
type FeeStationDestination struct {
	// The destination address.
	Address *string `json:"address,omitempty"`
	// The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON.
	Memo *string `json:"memo,omitempty"`
	// The transfer amount. For example, if you trade 1.5 BTC, then the value is `1.5`. 
	Amount string `json:"amount"`
}

type _FeeStationDestination FeeStationDestination

// NewFeeStationDestination instantiates a new FeeStationDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeStationDestination(amount string) *FeeStationDestination {
	this := FeeStationDestination{}
	this.Amount = amount
	return &this
}

// NewFeeStationDestinationWithDefaults instantiates a new FeeStationDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeStationDestinationWithDefaults() *FeeStationDestination {
	this := FeeStationDestination{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *FeeStationDestination) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeStationDestination) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *FeeStationDestination) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *FeeStationDestination) SetAddress(v string) {
	o.Address = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *FeeStationDestination) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeStationDestination) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *FeeStationDestination) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *FeeStationDestination) SetMemo(v string) {
	o.Memo = &v
}

// GetAmount returns the Amount field value
func (o *FeeStationDestination) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FeeStationDestination) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FeeStationDestination) SetAmount(v string) {
	o.Amount = v
}

func (o FeeStationDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeStationDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

func (o *FeeStationDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
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

	varFeeStationDestination := _FeeStationDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFeeStationDestination)

	if err != nil {
		return err
	}

	*o = FeeStationDestination(varFeeStationDestination)

	return err
}

type NullableFeeStationDestination struct {
	value *FeeStationDestination
	isSet bool
}

func (v NullableFeeStationDestination) Get() *FeeStationDestination {
	return v.value
}

func (v *NullableFeeStationDestination) Set(val *FeeStationDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeStationDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeStationDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeStationDestination(val *FeeStationDestination) *NullableFeeStationDestination {
	return &NullableFeeStationDestination{value: val, isSet: true}
}

func (v NullableFeeStationDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeStationDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


