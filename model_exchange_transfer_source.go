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

// checks if the ExchangeTransferSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeTransferSource{}

// ExchangeTransferSource struct for ExchangeTransferSource
type ExchangeTransferSource struct {
	SourceType WalletSubtype `json:"source_type"`
	// Unique id of the wallet to transfer from.
	WalletId string `json:"wallet_id"`
	// Exchange trading account or any sub wallet info for transfer.
	SubWalletId string `json:"sub_wallet_id"`
}

type _ExchangeTransferSource ExchangeTransferSource

// NewExchangeTransferSource instantiates a new ExchangeTransferSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeTransferSource(sourceType WalletSubtype, walletId string, subWalletId string) *ExchangeTransferSource {
	this := ExchangeTransferSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.SubWalletId = subWalletId
	return &this
}

// NewExchangeTransferSourceWithDefaults instantiates a new ExchangeTransferSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeTransferSourceWithDefaults() *ExchangeTransferSource {
	this := ExchangeTransferSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *ExchangeTransferSource) GetSourceType() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ExchangeTransferSource) GetSourceTypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *ExchangeTransferSource) SetSourceType(v WalletSubtype) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *ExchangeTransferSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *ExchangeTransferSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *ExchangeTransferSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetSubWalletId returns the SubWalletId field value
func (o *ExchangeTransferSource) GetSubWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubWalletId
}

// GetSubWalletIdOk returns a tuple with the SubWalletId field value
// and a boolean to check if the value has been set.
func (o *ExchangeTransferSource) GetSubWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubWalletId, true
}

// SetSubWalletId sets field value
func (o *ExchangeTransferSource) SetSubWalletId(v string) {
	o.SubWalletId = v
}

func (o ExchangeTransferSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeTransferSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["sub_wallet_id"] = o.SubWalletId
	return toSerialize, nil
}

func (o *ExchangeTransferSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
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

	varExchangeTransferSource := _ExchangeTransferSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeTransferSource)

	if err != nil {
		return err
	}

	*o = ExchangeTransferSource(varExchangeTransferSource)

	return err
}

type NullableExchangeTransferSource struct {
	value *ExchangeTransferSource
	isSet bool
}

func (v NullableExchangeTransferSource) Get() *ExchangeTransferSource {
	return v.value
}

func (v *NullableExchangeTransferSource) Set(val *ExchangeTransferSource) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeTransferSource) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeTransferSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeTransferSource(val *ExchangeTransferSource) *NullableExchangeTransferSource {
	return &NullableExchangeTransferSource{value: val, isSet: true}
}

func (v NullableExchangeTransferSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeTransferSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


