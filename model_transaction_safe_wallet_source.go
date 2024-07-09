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

// checks if the TransactionSafeWalletSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionSafeWalletSource{}

// TransactionSafeWalletSource The information about the transaction source.
type TransactionSafeWalletSource struct {
	SourceType TransactionSourceType `json:"source_type"`
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	Delegate TransactionSafeWalletSourceDelegate `json:"delegate"`
}

type _TransactionSafeWalletSource TransactionSafeWalletSource

// NewTransactionSafeWalletSource instantiates a new TransactionSafeWalletSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionSafeWalletSource(sourceType TransactionSourceType, walletId string, delegate TransactionSafeWalletSourceDelegate) *TransactionSafeWalletSource {
	this := TransactionSafeWalletSource{}
	this.SourceType = sourceType
	this.WalletId = walletId
	this.Delegate = delegate
	return &this
}

// NewTransactionSafeWalletSourceWithDefaults instantiates a new TransactionSafeWalletSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionSafeWalletSourceWithDefaults() *TransactionSafeWalletSource {
	this := TransactionSafeWalletSource{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *TransactionSafeWalletSource) GetSourceType() TransactionSourceType {
	if o == nil {
		var ret TransactionSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *TransactionSafeWalletSource) GetSourceTypeOk() (*TransactionSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *TransactionSafeWalletSource) SetSourceType(v TransactionSourceType) {
	o.SourceType = v
}

// GetWalletId returns the WalletId field value
func (o *TransactionSafeWalletSource) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *TransactionSafeWalletSource) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *TransactionSafeWalletSource) SetWalletId(v string) {
	o.WalletId = v
}

// GetDelegate returns the Delegate field value
func (o *TransactionSafeWalletSource) GetDelegate() TransactionSafeWalletSourceDelegate {
	if o == nil {
		var ret TransactionSafeWalletSourceDelegate
		return ret
	}

	return o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value
// and a boolean to check if the value has been set.
func (o *TransactionSafeWalletSource) GetDelegateOk() (*TransactionSafeWalletSourceDelegate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delegate, true
}

// SetDelegate sets field value
func (o *TransactionSafeWalletSource) SetDelegate(v TransactionSafeWalletSourceDelegate) {
	o.Delegate = v
}

func (o TransactionSafeWalletSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionSafeWalletSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["delegate"] = o.Delegate
	return toSerialize, nil
}

func (o *TransactionSafeWalletSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"wallet_id",
		"delegate",
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

	varTransactionSafeWalletSource := _TransactionSafeWalletSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionSafeWalletSource)

	if err != nil {
		return err
	}

	*o = TransactionSafeWalletSource(varTransactionSafeWalletSource)

	return err
}

type NullableTransactionSafeWalletSource struct {
	value *TransactionSafeWalletSource
	isSet bool
}

func (v NullableTransactionSafeWalletSource) Get() *TransactionSafeWalletSource {
	return v.value
}

func (v *NullableTransactionSafeWalletSource) Set(val *TransactionSafeWalletSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionSafeWalletSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionSafeWalletSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionSafeWalletSource(val *TransactionSafeWalletSource) *NullableTransactionSafeWalletSource {
	return &NullableTransactionSafeWalletSource{value: val, isSet: true}
}

func (v NullableTransactionSafeWalletSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionSafeWalletSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

