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

// checks if the TransactionFeeStationWalletSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionFeeStationWalletSource{}

// TransactionFeeStationWalletSource The information about the transaction source.
type TransactionFeeStationWalletSource struct {
	SourceType TransactionSourceType `json:"source_type"`
	// The Wallet ID.
	WalletId string `json:"wallet_id"`
}

type _TransactionFeeStationWalletSource TransactionFeeStationWalletSource

// NewTransactionFeeStationWalletSource instantiates a new TransactionFeeStationWalletSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionFeeStationWalletSource(sourceType TransactionSourceType, walletId string) *TransactionFeeStationWalletSource {
	this := TransactionFeeStationWalletSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	return &this
}

// NewTransactionFeeStationWalletSourceWithDefaults instantiates a new TransactionFeeStationWalletSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionFeeStationWalletSourceWithDefaults() *TransactionFeeStationWalletSource {
	this := TransactionFeeStationWalletSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *TransactionFeeStationWalletSource) GetSourceType() TransactionSourceType {
	if o == nil {
		var ret TransactionSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeStationWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *TransactionFeeStationWalletSource) SetSourceType(v TransactionSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionFeeStationWalletSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionFeeStationWalletSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionFeeStationWalletSource) SetWalletId(v string) {
	o.WalletId = v
}

func (o TransactionFeeStationWalletSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionFeeStationWalletSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	return toSerialize, nil
}

func (o *TransactionFeeStationWalletSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"wallet_id",
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

	varTransactionFeeStationWalletSource := _TransactionFeeStationWalletSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionFeeStationWalletSource)

	if err != nil {
		return err
	}

	*o = TransactionFeeStationWalletSource(varTransactionFeeStationWalletSource)

	return err
}

type NullableTransactionFeeStationWalletSource struct {
	value *TransactionFeeStationWalletSource
	isSet bool
}

func (v NullableTransactionFeeStationWalletSource) Get() *TransactionFeeStationWalletSource {
	return v.value
}

func (v *NullableTransactionFeeStationWalletSource) Set(val *TransactionFeeStationWalletSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionFeeStationWalletSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionFeeStationWalletSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionFeeStationWalletSource(val *TransactionFeeStationWalletSource) *NullableTransactionFeeStationWalletSource {
	return &NullableTransactionFeeStationWalletSource{value: val, isSet: true}
}

func (v NullableTransactionFeeStationWalletSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionFeeStationWalletSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


