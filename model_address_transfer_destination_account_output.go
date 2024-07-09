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

// checks if the AddressTransferDestinationAccountOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressTransferDestinationAccountOutput{}

// AddressTransferDestinationAccountOutput struct for AddressTransferDestinationAccountOutput
type AddressTransferDestinationAccountOutput struct {
	// The destination address.
	Address *string `json:"address,omitempty"`
	// The memo that identifies a transaction in order to credit the correct account. For transfers out of Cobo Portal, it is highly recommended to include a memo for the chains such as XRP, EOS, XLM, IOST, BNB_BNB, ATOM, LUNA, and TON.
	Memo *string `json:"memo,omitempty"`
	// The quantity of the token in the transaction. For example, if you trade 1.5 ETH, then the value is `1.5`. 
	Amount *string `json:"amount,omitempty"`
}

// NewAddressTransferDestinationAccountOutput instantiates a new AddressTransferDestinationAccountOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressTransferDestinationAccountOutput() *AddressTransferDestinationAccountOutput {
	this := AddressTransferDestinationAccountOutput{}
	return &this
}

// NewAddressTransferDestinationAccountOutputWithDefaults instantiates a new AddressTransferDestinationAccountOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressTransferDestinationAccountOutputWithDefaults() *AddressTransferDestinationAccountOutput {
	this := AddressTransferDestinationAccountOutput{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AddressTransferDestinationAccountOutput) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestinationAccountOutput) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AddressTransferDestinationAccountOutput) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AddressTransferDestinationAccountOutput) SetAddress(v string) {
	o.Address = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *AddressTransferDestinationAccountOutput) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestinationAccountOutput) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *AddressTransferDestinationAccountOutput) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *AddressTransferDestinationAccountOutput) SetMemo(v string) {
	o.Memo = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AddressTransferDestinationAccountOutput) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressTransferDestinationAccountOutput) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AddressTransferDestinationAccountOutput) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AddressTransferDestinationAccountOutput) SetAmount(v string) {
	o.Amount = &v
}

func (o AddressTransferDestinationAccountOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressTransferDestinationAccountOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableAddressTransferDestinationAccountOutput struct {
	value *AddressTransferDestinationAccountOutput
	isSet bool
}

func (v NullableAddressTransferDestinationAccountOutput) Get() *AddressTransferDestinationAccountOutput {
	return v.value
}

func (v *NullableAddressTransferDestinationAccountOutput) Set(val *AddressTransferDestinationAccountOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressTransferDestinationAccountOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressTransferDestinationAccountOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressTransferDestinationAccountOutput(val *AddressTransferDestinationAccountOutput) *NullableAddressTransferDestinationAccountOutput {
	return &NullableAddressTransferDestinationAccountOutput{value: val, isSet: true}
}

func (v NullableAddressTransferDestinationAccountOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressTransferDestinationAccountOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


