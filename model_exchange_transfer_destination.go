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

// checks if the ExchangeTransferDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeTransferDestination{}

// ExchangeTransferDestination The data for exchange destination.
type ExchangeTransferDestination struct {
	DestinationType TransferDestinationType `json:"destination_type"`
	// Unique id of the wallet to transfer to.
	WalletId string `json:"wallet_id"`
	// Exchange trading account or any sub wallet info for transfer.
	SubWalletId string `json:"sub_wallet_id"`
	// Transaction value (Note that this is an absolute value. If you trade 1.5 ETH, then the value is 1.5) 
	Amount *string `json:"amount,omitempty"`
}

type _ExchangeTransferDestination ExchangeTransferDestination

// NewExchangeTransferDestination instantiates a new ExchangeTransferDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeTransferDestination(destinationType TransferDestinationType, walletId string, subWalletId string) *ExchangeTransferDestination {
	this := ExchangeTransferDestination{}
	this.DestinationType = destinationType
	this.WalletId = walletId
	this.SubWalletId = subWalletId
	return &this
}

// NewExchangeTransferDestinationWithDefaults instantiates a new ExchangeTransferDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeTransferDestinationWithDefaults() *ExchangeTransferDestination {
	this := ExchangeTransferDestination{}
	return &this
}

// GetDestinationType returns the DestinationType field value
func (o *ExchangeTransferDestination) GetDestinationType() TransferDestinationType {
	if o == nil {
		var ret TransferDestinationType
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *ExchangeTransferDestination) GetDestinationTypeOk() (*TransferDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *ExchangeTransferDestination) SetDestinationType(v TransferDestinationType) {
	o.DestinationType = v
}

// GetWalletId returns the WalletId field value
func (o *ExchangeTransferDestination) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *ExchangeTransferDestination) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *ExchangeTransferDestination) SetWalletId(v string) {
	o.WalletId = v
}

// GetSubWalletId returns the SubWalletId field value
func (o *ExchangeTransferDestination) GetSubWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubWalletId
}

// GetSubWalletIdOk returns a tuple with the SubWalletId field value
// and a boolean to check if the value has been set.
func (o *ExchangeTransferDestination) GetSubWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubWalletId, true
}

// SetSubWalletId sets field value
func (o *ExchangeTransferDestination) SetSubWalletId(v string) {
	o.SubWalletId = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ExchangeTransferDestination) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeTransferDestination) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ExchangeTransferDestination) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ExchangeTransferDestination) SetAmount(v string) {
	o.Amount = &v
}

func (o ExchangeTransferDestination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeTransferDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination_type"] = o.DestinationType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["sub_wallet_id"] = o.SubWalletId
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

func (o *ExchangeTransferDestination) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination_type",
		"wallet_id",
		"sub_wallet_id",
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

	varExchangeTransferDestination := _ExchangeTransferDestination{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeTransferDestination)

	if err != nil {
		return err
	}

	*o = ExchangeTransferDestination(varExchangeTransferDestination)

	return err
}

type NullableExchangeTransferDestination struct {
	value *ExchangeTransferDestination
	isSet bool
}

func (v NullableExchangeTransferDestination) Get() *ExchangeTransferDestination {
	return v.value
}

func (v *NullableExchangeTransferDestination) Set(val *ExchangeTransferDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeTransferDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeTransferDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeTransferDestination(val *ExchangeTransferDestination) *NullableExchangeTransferDestination {
	return &NullableExchangeTransferDestination{value: val, isSet: true}
}

func (v NullableExchangeTransferDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeTransferDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


